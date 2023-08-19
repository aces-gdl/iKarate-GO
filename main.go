package main

import (
	"iKarate-GO/initializers"
	"iKarate-GO/server"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectTODB()
	initializers.SyncDatabase()

}

func main() {

	server.Init()
}
