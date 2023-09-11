package constants

import (
	"fmt"
	"os"
)

func InitializeVarEnviroment() {

	// Variables base de datos
	os.Setenv("INDEX", "Test")
	os.Setenv("URL", "http://localhost:4080/api/")
	os.Setenv("DB_USER", "admin")
	os.Setenv("DB_PASSWORD", "Complexpass#123")

	fmt.Println("Variables de entorno establecidas")
}
