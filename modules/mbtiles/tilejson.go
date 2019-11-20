package mbtiles

import (
	"fmt"

	"github.com/pkg/errors"
)

//From:
// - https://github.com/mapbox/tilejson-spec/tree/master/2.2.0
// - https://github.com/mapbox/mbtiles-spec/blob/master/1.3/spec.md

//GetTileJSON generate a TileJSON from database
func (c *Client) GetTileJSON() ([]byte, error) {
	metas, err := c.GetMetadataList()
	if err != nil {
		return nil, errors.Wrap(err, "metadata read failed")
	}

	//minzoom: defaut 0
	if _, ok := metas["minzoom"]; !ok {
		metas["minzoom"] = "0"
	}
	//maxzoom: defaut 22
	if _, ok := metas["maxzoom"]; !ok {
		metas["maxzoom"] = "22"
	}
	//bounds: default
	if _, ok := metas["bounds"]; !ok {
		metas["bounds"] = "-180.0,-85,180,85"
	}
	//center: default 0,0,0
	if _, ok := metas["center"]; !ok {
		metas["center"] = "0,0,0"
	}
	//TODO vector_layers json

	json := fmt.Sprintf(`{
		"tilejson": "2.2.0",
		"name": "%s",
		"description": "%s",
		"minzoom": %s,
		"maxzoom": %s,
		"bounds": [%s],
		"center": [%s],
		"version": "%s",
		"attribution": "%s",
		"scheme": "xyz",
		"tiles": ["/api/v1/tiles/{z}/{x}/{y}.%s"]
	}`, metas["name"], metas["description"], metas["minzoom"], metas["maxzoom"], metas["bounds"], metas["center"], metas["version"], metas["attribution"], metas["format"])
	return []byte(json), nil
}
