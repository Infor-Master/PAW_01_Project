package services

import (
	"projetoapi/docs"
)

func FormatSwagger() {
	//http://localhost:8081/swagger/index.html
	docs.SwaggerInfo.Title = "API do projecto de PAW"
	docs.SwaggerInfo.Description = "Essa api permite registar em tempo real o número de pessoas numa zona."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8081"
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
}
