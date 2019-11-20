package mbtiles

import (
	"fmt"

	"github.com/rs/zerolog"
	"xorm.io/xorm"

	//Needed for opening mbtiles
	_ "github.com/mattn/go-sqlite3"

	"github.com/sapk/got/modules/log"
)

//Client client ot parse and access mbtile file
type Client struct {
	DB *xorm.Engine
}

//Metadata hold OSM metadata https://github.com/mapbox/mbtiles-spec/blob/master/1.3/spec.md#metadata
//TABLE metadata (name text, value text);
type Metadata struct {
	Name  string `xorm:"text 'name'"`
	Value string `xorm:"text 'value'"`
}

//TableName table name in db
func (m *Metadata) TableName() string {
	return "metadata"
}

//Tiles hold OSM rendered tiles https://github.com/mapbox/mbtiles-spec/blob/master/1.3/spec.md#tiles
//TABLE tiles (zoom_level integer, tile_column integer, tile_row integer, tile_data blob);
type Tiles struct {
	Zoom   int     `xorm:"integer 'zoom_level'"`
	Column int     `xorm:"integer 'tile_column'"`
	Row    int     `xorm:"integer 'tile_row'"`
	Data   []uint8 `xorm:"blob 'tile_data'"`
}

//TableName table name in db
func (m *Tiles) TableName() string {
	return "tiles"
}

//NewClient coonect to mbtiles database
func NewClient(debug bool, fPath string) *Client {
	dbLogger := log.NewLogger(debug, "database")
	db, err := initDB(dbLogger, fPath)
	if err != nil {
		dbLogger.Fatal().Err(err).Msg("Fail to read mbtiles file")
	}
	if debug {
		m := new(Metadata)
		total, err := db.Count(m)
		if err != nil {
			dbLogger.Fatal().Err(err).Msg("Fail to read mbtiles file")
		} else {
			dbLogger.Debug().Int64("count", total).Msg("Metadata")
		}
		t := new(Tiles)
		total, err = db.Count(t)
		if err != nil {
			dbLogger.Fatal().Err(err).Msg("Fail to read mbtiles file")
		} else {
			dbLogger.Debug().Int64("count", total).Msg("Tiles")
		}
	}

	//Simple cache for metadata
	cacher := xorm.NewLRUCacher(xorm.NewMemoryStore(), 100)
	err = db.MapCacher(&Metadata{}, cacher)
	if err != nil {
		dbLogger.Fatal().Err(err).Msg("Fail to setup metadata mbtiles cache")
	}

	return &Client{DB: db}
}

//GetMetadata retrieve a metadata from database
func (c *Client) GetMetadata(name string) (string, error) {
	var m = Metadata{Name: name}
	has, err := c.DB.Get(&m)
	if err != nil {
		return "", err
	}
	if !has {
		return "", fmt.Errorf("metadata not found")
	}
	return m.Value, nil
}

//GetMetadataList retrieve all the metadata from database
func (c *Client) GetMetadataList() (map[string]string, error) {
	results, err := c.DB.QueryString("select * from metadata")
	if err != nil {
		return nil, err
	}
	out := make(map[string]string)
	for _, line := range results {
		out[line["name"]] = line["value"]
	}
	return out, nil
}

//GetTile retrieve a tile from database
func (c *Client) GetTile(z, x, y int) ([]byte, error) {
	yCorr := (1 << z) - 1 - y
	var t = Tiles{Zoom: z, Column: x, Row: yCorr}
	has, err := c.DB.Get(&t)
	if err != nil {
		return []byte{}, err
	}
	if !has {
		return []byte{}, fmt.Errorf("tile not found")
	}
	return t.Data, nil
}

//initDB start the database connection and settings
func initDB(logger *zerolog.Logger, filepath string) (*xorm.Engine, error) {
	logger.Debug().Msgf("Opening file '%s' ...", filepath)
	appEngine, err := xorm.NewEngine("sqlite3", filepath)
	if err != nil {
		return nil, err
	}
	appEngine.SetLogger(log.NewSQLLogger(logger))
	appEngine.ShowSQL(true)
	return appEngine, nil
}
