package main

import (
	"github.com/hydrophonics/database"
	"github.com/hydrophonics/database/migrate"
	"github.com/hydrophonics/database/seed"
	"github.com/hydrophonics/routers"
)

func main() {

	db := database.Init()
	defer db.Close()

	migrate.Migrate(db)
	seed.Category(db)

	routers.StartServer()
}
