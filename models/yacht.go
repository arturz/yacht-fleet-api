package models

import (
	"fmt"
	"sort"
)

type Yacht struct {
	ID       int    `json:"id"`
	MarinaID int    `json:"marina_id"`
	Name     string `json:"name"`
}

var Yachts = map[int]Yacht{}
var yachts_id_counter = 0

func NewYacht(name string, marina_id int) *Yacht {
	yacht := &Yacht{
		Name:     name,
		ID:       yachts_id_counter,
		MarinaID: marina_id,
	}

	yachts_id_counter++

	Yachts[yacht.ID] = *yacht

	return yacht
}

func ClearYachts() {
	Yachts = map[int]Yacht{}
}

func GetYachts() []Yacht {
	yachts := make([]Yacht, 0, len(Yachts))
	for _, yacht := range Yachts {
		yachts = append(yachts, yacht)
	}

	sort.Slice(yachts, func(i, j int) bool {
		return yachts[i].ID < yachts[j].ID
	})

	return yachts
}

func HashYacht(yacht_id int) string {
	yacht := Yachts[yacht_id]
	hash := fmt.Sprintf("%d:%d:%s", yacht.ID, yacht.MarinaID, yacht.Name)
	return hash
}
