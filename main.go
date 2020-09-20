package main

import (
	"playUp/configs/database"
	"playUp/configs/router"
	"playUp/main/master"
)

func main() {
	db := database.ConnectionDB()
	r := router.CreateRouter()

	master.Init(db, r)
	router.StartServer(r)
}
