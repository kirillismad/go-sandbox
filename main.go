package main

import (
	swaggergenerated "sandbox/http/swagger_generated"
)

func main() {
	// os.Setenv("ECHO_EXAMPLE_FILE_PATH", "./http/echo_example/testdata/websocket.png")
	// e := echo_example.BuildServer()
	// e.Logger.Fatal(e.Start(":8080"))

	s, shutdown := swaggergenerated.NewServer()
	defer shutdown()
	s.Port = 8080
	s.Serve()
}
