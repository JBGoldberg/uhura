package main

import (
	"github.com/JBGoldberg/uhura/cmd"

	log "github.com/sirupsen/logrus"
)

func main() {

	log.Printf("uhura - v0.0.1")

	cmd.Execute()

}
