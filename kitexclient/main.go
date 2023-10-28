package main

import (
	"context"
	"fmt"
	"log"

	"kitexclient/kitex_gen/api/sample"
	"kitexclient/kitex_gen/dragon"

	"github.com/cloudwego/kitex/client"
)

func main() {
	c, err := sample.NewClient("example-server", client.WithHostPorts("0.0.0.0:8888"))
	if err != nil {
		log.Fatal(err)
	}
	resp, err := c.DragonSay(context.Background(), &dragon.DragonSayRequest{Message: "nihao"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp)
}
