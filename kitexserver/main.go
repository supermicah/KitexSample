package main

import (
	api "kitexserver/kitex_gen/api/sample"
	"log"
)

func main() {
	svr := api.NewServer(new(SampleImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
