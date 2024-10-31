package main

import (
	"encoding/json"
	"os"
	"path"
)

type TilemapLayerJSON struct {
	Data   []int  `json:"data"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
	Name   string `json:"name"`
}

type TileMapJSON struct {
	Layers   []TilemapLayerJSON `json:"layers"`
	Tilesets []map[string]any   `json:"tilesets"`
}

func (t *TileMapJSON) GenTilesets() ([]Tileset, error) {
	tilesets := make([]Tileset, 0)

	for _, tilesetData := range t.Tilesets {
		tilesetPath := path.Join("assets/maps/", tilesetData["source"].(string))
		tileset, err := NewTileset(tilesetPath, int(tilesetData["firstgid"].(float64)))
		if err != nil {
			return nil, err
		}

		tilesets = append(tilesets, tileset)
	}

	return tilesets, nil
}

func NewTilemapJSON(filepath string) (*TileMapJSON, error) {
	contents, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	var tilemapJSON TileMapJSON
	err = json.Unmarshal(contents, &tilemapJSON)
	if err != nil {
		return nil, err
	}

	return &tilemapJSON, nil
}
