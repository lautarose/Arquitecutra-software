package app

import (
	pingpong "mvc/controller"
)

func mapUrls() {
	//GET
	router.GET("/ping", pingpong.PingPong)
}
