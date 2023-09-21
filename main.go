package main

import (
	"github.com/FranMT-S/chi-zinc-server/src/constants"
	myServer "github.com/FranMT-S/chi-zinc-server/src/server"
)

func main() {
	constants.InitializeVarEnviroment()

	myServer.Server().Start()

}
