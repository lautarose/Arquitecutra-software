package app

import (
	"Arquitecutra-software/mvc-auth/auth/controller"
)

func mapUrls() {
	//POST
	router.POST("/login", controller.InsertUser)
}
