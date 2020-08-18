/*
Copyright 2020 Kohl's Department Stores, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/KohlsTechnology/blackbox-helloworld-responder/pkg/version"
)

func main() {
	log.SetOutput(os.Stdout)
	log.Println(version.Get())

	var portHTTP int = 8080
	var portTCP int = 8081
	var err error

	envPort, ok := os.LookupEnv("HELLO_WORLD_HTTP_PORT")
	if ok {
		portHTTP, err = strconv.Atoi(envPort)
		if err != nil {
			log.Fatal(err)
		}
	}
	envPort, ok = os.LookupEnv("HELLO_WORLD_TCP_PORT")
	if ok {
		portTCP, err = strconv.Atoi(envPort)
		if err != nil {
			log.Fatal(err)
		}
	}

	// Start the TCP Responder
	go tcpHelloServer(portTCP)

	// Start the HTTP Responder
	httpHelloServer(portHTTP)
}

// httpHelloServer starts a simple HelloWorld Web Server
func httpHelloServer(port int) {
	log.Printf("Starting 'Hello World!' HTTP Server on %d\n", port)
	mux := http.NewServeMux()
	httpsrv := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  10 * time.Second,
		Handler:      mux,
	}
	mux.HandleFunc("/", httpHelloHandler)
	httpsrv.ListenAndServe()
}

// httpHelloServer starts a simple HelloWorld TCP Server
func tcpHelloServer(port int) {
	log.Printf("Starting 'Hello World!' TCP Server on %d\n", port)
	listener, err := net.Listen("tcp4", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	for {
		connection, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("TCP %s\n", connection.RemoteAddr())
		connection.Write([]byte(`Hello World!`))
		connection.Close()
	}
}

// httpHelloServer responds with a simple static "Hello World!" text
func httpHelloHandler(res http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()

	log.Printf("HTTP %s", req.RemoteAddr)
	header := res.Header()

	header.Set("Content-Type", "text/plain")
	header.Set("Cache-Control", "no-store")
	res.WriteHeader(http.StatusOK)

	fmt.Fprint(res, `Hello World!`)
}
