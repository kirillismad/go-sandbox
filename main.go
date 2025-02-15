package main

import (
	"os"
	"sandbox/http/echo_example"
)

func main() {
	os.Setenv("ECHO_EXAMPLE_FILE_PATH", "./http/echo_example/testdata/websocket.png")
	e := echo_example.BuildServer()
	e.Logger.Fatal(e.Start(":8080"))
}
