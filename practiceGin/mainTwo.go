package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
)

func main() {

	for {
		l, _ := net.Listen("tcp", ":0")
		var port = l.Addr().String()[5:]
		l.Close()
		fmt.Println(port)
		go func() {
			l1, err := net.Listen("tcp", ":"+port)
			if err != nil {
				log.Fatal("Error while binding port: ", err.Error())
			} else {
				fmt.Println("bound well! ", port)
			}
			http.Serve(l1, nil)
		}()
	}
}
