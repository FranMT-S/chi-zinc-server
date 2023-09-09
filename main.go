package main

import (
	"fmt"
	"time"

	"github.com/FranMT-S/Challenge-Go/src/constants"
	myServer "github.com/FranMT-S/Challenge-Go/src/server"
	// "github.com/FranMT-S/Challenge-Go/src/core"
	// "github.com/FranMT-S/Challenge-Go/src/core/bulker"
	// "github.com/FranMT-S/Challenge-Go/src/core/parser"
)

func main() {
	constants.InitializeVarEnviroment()

	startTime := time.Now() // Registra el tiempo de inicio
	myServer.Server().Start()

	endTime := time.Now()
	duration := endTime.Sub(startTime)
	seconds := duration.Seconds()

	fmt.Printf("El código se ejecutó en %.2f segundos\n", seconds)
}
