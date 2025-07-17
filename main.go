package main

import (
	"fmt"

	"github.com/antchfx/xmlquery"
	_ "github.com/dgrijalva/jwt-go"
	_ "github.com/gogo/protobuf/proto"

	//- "github.com/gogs/gogs"
	_ "github.com/hashicorp/golang-lru"
	//_ "github.com/owncast/owncast/logging"

	_ "github.com/gorilla/websocket"
	_ "github.com/traefik/yaegi"
)

func main() {
	fmt.Println("Hello world")
	wadl, err := xmlquery.LoadURL("https://httpbin.org/get")
	if err != nil {
		panic(err)
	}

	attr := xmlquery.FindOne(wadl, "//application/@xmlns")
	fmt.Println(attr.InnerText())
}
