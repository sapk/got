package main

import (
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"github.com/rs/zerolog"
	"xorm.io/core"
	"xorm.io/xorm"
)

//TODO rename to TileReader

//Client client ot parse and access mbtile file
type Client struct {
	DB *xorm.Engine
}

//TABLE metadata (name text, value text);

//Metadata hold OSM metadata https://github.com/mapbox/mbtiles-spec/blob/master/1.3/spec.md#metadata
type Metadata struct {
	Name  string `xorm:"text 'name'"`
	Value string `xorm:"text 'value'"`
}

//TableName table name in db
func (m *Metadata) TableName() string {
	return "metadata"
}

//TABLE tiles (zoom_level integer, tile_column integer, tile_row integer, tile_data blob);

//Tiles hold OSM rendered tiles https://github.com/mapbox/mbtiles-spec/blob/master/1.3/spec.md#tiles
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

func tileClient(debug bool, fPath string) *Client {
	dbLogger := setupLogger(debug, "database")
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
	return &Client{DB: db}
}

//GetTile retrive a tile from database
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
	appEngine.SetLogger(generateSQLLogger(logger))
	appEngine.ShowSQL(true)
	return appEngine, nil
}

//TODO better
func generateSQLLogger(l *zerolog.Logger) *logger {
	return &logger{l}
}

type logger struct {
	logger *zerolog.Logger
}

func (l *logger) Debug(v ...interface{}) {
	l.logger.Print(v...) //TODO debug level
}

func (l *logger) Debugf(format string, v ...interface{}) {
	l.logger.Printf(format, v...) //TODO debug level
}

func (l *logger) Error(v ...interface{}) {
	l.logger.Print(v...) //TODO Error level
}

func (l *logger) Errorf(format string, v ...interface{}) {
	l.logger.Printf(format, v...) //TODO Error level
}

func (l *logger) Info(v ...interface{}) {
	l.logger.Print(v...) //TODO Info level
}

func (l *logger) Infof(format string, v ...interface{}) {
	l.logger.Printf(format, v...) //TODO Info level
}

func (l *logger) Warn(v ...interface{}) {
	l.logger.Print(v...) //TODO Warn level
}

func (l *logger) Warnf(format string, v ...interface{}) {
	log.Printf(format, v...) //TODO Warn level
}

func (l *logger) Level() core.LogLevel {
	return core.LOG_DEBUG //TODO
}

func (l *logger) SetLevel(lvl core.LogLevel) {
	l.logger.Debug().Msgf("xorm.log.SetLevel %d", lvl)
	//zerolog.SetGlobalLevel(zerolog.DebugLevel)
}
func (l *logger) ShowSQL(show ...bool) {
	//TODO
}

func (l *logger) IsShowSQL() bool {
	return true
}
