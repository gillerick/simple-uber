package simple_uber

import (
	"log"
	"simple-uber/internal/configs"
	"simple-uber/internal/database"
	"simple-uber/internal/http"
	"simple-uber/internal/repositories"
)

func main() {
	service := Service{}

	if err := service.Run(); err != nil {
		log.Fatalf("unable to start services %s", err)
	}

}

type Service struct {
	HttpServer *http.Server
}

func (s *Service) Run() error {
	var err error

	// Fetch app configurations. Empty paths reads configs from a set default path
	ymlConfig := configs.ReadYaml("")
	config := configs.GetConfig(*ymlConfig)

	//Setup a database connection
	pgDb, err := database.NewConnection(config.DB)
	if err != nil {
		log.Fatal("could not establish connection with the database")
	}

	//Set up database
	dbHandler := repositories.NewDatabaseHandler(pgDb)
}
