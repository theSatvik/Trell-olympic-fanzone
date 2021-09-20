package server

import (
	"gitlab.com/satvikshrivas26/olympic-fanzone/mappings"
	// "gitlab.com/satvikshrivas26/olympic-fanzone/controllers"
	_ "github.com/go-sql-driver/mysql"
	
)

func Init() {
	// logger.Init()
	// Init()
	// es.Init()
	// redis.Init()
	mappings.CreateUrlMappings()
	// Listen and server on 0.0.0.0:8080
	mappings.Router.Run(":8080")
}
