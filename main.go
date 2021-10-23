package main

import (
	"gopkg.in/macaron.v1"
)

func main() {
	server := macaron.Classic()
	server.Use(macaron.Renderer())
	server.Get("/", handlerMain)
	server.Get("/nft", handleQueryNFTItem)
	server.Run()
}
