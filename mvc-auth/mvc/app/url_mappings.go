package app

import (
	userController "mvc-auth/mvc/controllers/user"

	log "github.com/sirupsen/logrus"
)

func mapUrls() {

	// Users Mapping
	router.GET("/user", userController.GetUsers)
	router.POST("/user", userController.UserInsert)

	log.Info("Finishing mappings configurations")
}
