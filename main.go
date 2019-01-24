package main

import (
	_ "gin_jwt_demo/model"
	"gin_jwt_demo/routes"
)

func main() {
	routers := routes.InitRouter()
	routers.Run(":8080")
}
