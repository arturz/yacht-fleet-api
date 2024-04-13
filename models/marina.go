package models

import "sort"

type Marina struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var Marinas = map[int]Marina{}
var marinas_id_counter = 0

func NewMarina(name string) *Marina {
	marina := &Marina{
		Name: name,
		ID:   marinas_id_counter,
	}

	marinas_id_counter++

	Marinas[marina.ID] = *marina

	return marina
}

func ClearMarinas() {
	Marinas = map[int]Marina{}
}
func GetMarinas() []Marina {
	marinas := make([]Marina, 0, len(Marinas))
	for _, marina := range Marinas {
		marinas = append(marinas, marina)
	}

	sort.Slice(marinas, func(i, j int) bool {
		return marinas[i].ID < marinas[j].ID
	})

	return marinas
}
