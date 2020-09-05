package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {

	resp, err := http.Get("http://www.google.com")
	lw := logWriter{}
	if err != nil {
		fmt.Println("error occured", err)
		os.Exit(1)
	}
	//fmt.Println(resp)
	// bs := make([]byte, 99999)
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))

	//io.Copy(os.Stdout, resp.Body) //same work done by io.copy

	//custom writer
	io.Copy(lw, resp.Body)

}

func (logWriter) Write(bs []byte) (int, error) {

	fmt.Println(string(bs))
	fmt.Println("just wrote this much bytes", len(bs))

	return len(bs), nil

}
