package sdproject

import (
	"fmt"
	"log"
)

func NewTracer(logType string, categorie string, message string) {
	if logType == "error" {
		log.Println(fmt.Sprintf("ERROR :: (%s) ~> %s", categorie, message))
	}

	if logType == "info" {
		log.Println(fmt.Sprintf("INFO :: (%s) ~> %s", categorie, message))
	}
}
