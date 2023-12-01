package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"goyave.dev/goyave/v4"
)

func Register(router *goyave.Router) {
	router.Get("/api", sleep)
}

// Handler function for the "/api" route
func sleep(response *goyave.Response, request *goyave.Request) {
	goyave.Logger.Printf("Starting to sleep for '%s'...\n", request.Data["id"])
	time.Sleep(5 * time.Second)
	goyave.Logger.Printf("... I'm awake for '%s'", request.Data["id"])
	response.String(http.StatusOK, fmt.Sprintf(`{"now":%d, "id": "%s"}`,
		time.Now().UnixMilli(), request.Data["id"]))
}

func main() {
	goyave.Logger.Println("Launching...")
	if err := goyave.Start(Register); err != nil {
		os.Exit(err.(*goyave.Error).ExitCode)
	}
}
