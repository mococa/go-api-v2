package main

import (
	"github.com/mococa/go-api-v2/pkg/cmd/server"
	"github.com/mococa/go-api-v2/pkg/db"
)

func main() {
	const port string = ":3333"

	db := db.Init()

	server := server.NewServer(port, db)
	server.Start()
}
