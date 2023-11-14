package initializers

import (
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

var CurrentUser uuid.UUID

func LoadEnvVariables() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error cargando archivo .env")
	}

	fmt.Println("...Inicializando")

}
