package main

import (
	"github.com/hydrophonics/config"
	"github.com/hydrophonics/database"
	"github.com/hydrophonics/database/migrate"
	"github.com/hydrophonics/database/seed"
	"github.com/hydrophonics/routers"
)

func main() {

	err := config.LoadConfig() // loads config
	if err != nil {
		panic("error reading toml")
	}

	db := database.Init()
	defer db.Close()

	migrate.Migrate(db) // create table, adds constraint
	seed.Category(db)   // seeds db values

	routers.StartServer()
}
