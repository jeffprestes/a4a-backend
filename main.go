package main

import (
	"gopkg.in/macaron.v1"
)

func main() {
	server := macaron.Classic()
	server.Use(macaron.Renderer())
	server.Use(macaron.Static("static"))
	// server.Get("/", handlerMain)
	server.Get("/nft", handleQueryNFTItem)
	server.Run()
}
