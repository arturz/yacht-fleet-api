package models

import (
	"errors"
	"sort"
)

// migracja jachtu z jednego portu do drugiego
type Migration struct {
	ID             int `json:"id"`
	SourceMarinaID int `json:"source_marina_id"`
	MarinaID       int `json:"marina_id"`
	YachtID        int `json:"yacht_id"`
}

var Migrations = map[int]Migration{}
var migrations_id_counter = 0

func NewMigration(yacht_id int, marina_id int) (*Migration, error) {
	yacht, ok := Yachts[yacht_id]

	if !ok {
		return nil, errors.New("Yacht not found")
	}

	source_marina_id := yacht.MarinaID

	yacht.MarinaID = marina_id
	Yachts[yacht_id] = yacht

	migration := &Migration{
		ID:             migrations_id_counter,
		SourceMarinaID: source_marina_id,
		MarinaID:       marina_id,
		YachtID:        yacht_id,
	}

	migrations_id_counter++

	Migrations[migration.ID] = *migration

	return migration, nil
}

func ClearMigrations() {
	Migrations = map[int]Migration{}
}

func GetMigrations() []Migration {
	migrations := make([]Migration, 0, len(Migrations))
	for _, migration := range Migrations {
		migrations = append(migrations, migration)
	}

	sort.Slice(migrations, func(i, j int) bool {
		return migrations[i].ID < migrations[j].ID
	})

	return migrations
}
