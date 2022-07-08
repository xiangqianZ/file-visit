package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

var Success = true

func main() {

	pwd,_ := os.Getwd()
	fmt.Println("current path: ", pwd)

	argArray := os.Args

	port := "8080"
	if len(argArray) > 1 {
		port = argArray[1]
	}

	fmt.Println("open port:", port)

	go printSuccessLog()

	http.Handle("/", http.FileServer(http.Dir(pwd)))
	e := http.ListenAndServe(":" + port, nil)

	if e != nil {
		fmt.Println("open error", e)
		Success = false
	}
}

func printSuccessLog()  {
	time.Sleep(time.Duration(2) * time.Second)

	if !Success {
		return
	}

	fmt.Println("open Success")
}
