package models

import "sort"

type Charter struct {
	ID      int    `json:"id"`
	YachtID int    `json:"yacht_id"`
	Captain string `json:"captain"`
}

var Charters = map[int]Charter{}
var charters_id_counter = 0

func NewCharter(captain string, yacht_id int) *Charter {
	charter := &Charter{
		Captain: captain,
		ID:      charters_id_counter,
		YachtID: yacht_id,
	}

	charters_id_counter++

	Charters[charter.ID] = *charter

	return charter
}

func GetCharters() []Charter {
	charters := make([]Charter, 0, len(Charters))
	for _, charter := range Charters {
		charters = append(charters, charter)
	}

	sort.Slice(charters, func(i, j int) bool {
		return charters[i].ID < charters[j].ID
	})

	return charters
}
