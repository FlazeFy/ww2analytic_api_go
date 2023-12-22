package main

import (
	"ww2analytic/packages/database"
	"ww2analytic/routes"
)

func main() {
	database.Init()
	e := routes.InitV1()

	e.Logger.Fatal(e.Start(":1323"))
}
