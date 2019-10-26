package main

import (
	"blog-app/app/server"
	"log"
)

func main() {


	err:= server.Start()


	if err != nil {
		log.Fatal("Error in starting Server", err)
	}
}

