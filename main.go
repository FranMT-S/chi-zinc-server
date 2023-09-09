package main

import (
	"fmt"
	"time"

	"github.com/FranMT-S/chi-zinc-server/src/constants"
	myServer "github.com/FranMT-S/chi-zinc-server/src/server"
	// "github.com/FranMT-S/chi-zinc-server/src/core"
	// "github.com/FranMT-S/chi-zinc-server/src/core/bulker"
	// "github.com/FranMT-S/chi-zinc-server/src/core/parser"
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
