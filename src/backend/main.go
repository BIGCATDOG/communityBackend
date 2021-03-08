package main

import (

	"fmt"
    "github.com/untils"
	"github.com/valyala/gorpc"
	"log"
)

func main()  {
	fmt.Printf("server is running")
	utils.Add()
	s := &gorpc.Server{
		// Accept clients on this TCP address.
		Addr: "127.0.0.1:12345",

		// Echo handler - just return back the message we received from the client
		Handler: func(clientAddr string, request interface{}) interface{} {
			log.Printf("Obtained request %+v from the client %s\n", request, clientAddr)
			return request
		},
	}
	if err := s.Serve(); err != nil {
		log.Fatalf("Cannot start rpc server: %s", err)
	}
}
