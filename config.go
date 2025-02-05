package main

import (
	"os"
)

var Mode string = os.Getenv("mode")
var FiberDataUrl string = os.Getenv("FiberDataUrl")
var FiberDataApiKey string = os.Getenv("FiberDataApiKey")

func GetAllowedOrigins() []string {
	if Mode == "prod" {
		return []string{"https://*", "http://*"}
	} else {
		return []string{"https://portal.eeplus.com"}
	}
}
