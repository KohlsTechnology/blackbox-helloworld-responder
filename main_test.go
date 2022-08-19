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
	"bufio"
	"io"
	"net"
	"net/http"
	"testing"
	"time"
)

var strHelloWorld = "Hello World!"

func TestHTTP(t *testing.T) {
	go httpHelloServer(8080)
	time.Sleep(time.Second)

	resp, err := http.Get("http://127.0.0.1:8080")
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	if http.StatusOK != resp.StatusCode {
		t.Fatalf("http test returned status: %d, wanted: %d", resp.StatusCode, http.StatusOK)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	if strHelloWorld != string(body) {
		t.Errorf("returned body: %s, wanted: %s", string(body), strHelloWorld)
	}
}

func TestTCP(t *testing.T) {
	go tcpHelloServer(8081)
	time.Sleep(time.Second)

	connection, err := net.Dial("tcp4", "127.0.0.1:8081")
	if err != nil {
		t.Fatal(err)
	}

	defer connection.Close()

	reply, _ := bufio.NewReader(connection).ReadString('\n')

	if strHelloWorld != reply {
		t.Errorf("returned body: %s, wanted: %s", reply, strHelloWorld)
	}

}
