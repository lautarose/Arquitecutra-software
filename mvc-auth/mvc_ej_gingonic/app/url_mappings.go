package app

import "Arquitecutra-software/mvc-auth/mvc_ej_gingonic/controller"

func MapUrls() {

	//POST
	router.POST("/test", controller.InsertBody)

}
