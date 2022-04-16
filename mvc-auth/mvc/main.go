package main

import (
	"mvc-auth/mvc/app"
	"mvc-auth/mvc/db"
)

func main() {
	db.StartDbEngine()
	app.StartRoute()
}
