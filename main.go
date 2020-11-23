package main

import (
	"net/http"

	"github.com/marcodeltorob/golang-bootcamp-2020/config"
	"github.com/marcodeltorob/golang-bootcamp-2020/controller"
	"github.com/marcodeltorob/golang-bootcamp-2020/router"
	"github.com/marcodeltorob/golang-bootcamp-2020/usecase"
	"github.com/unrolled/render"
)

func main() {
	configuration := config.GetConfig()

	dataService := NewDatabase(configuration.DbUser, configuration.DbPassword, configuration.DbHostname, configuration.DbName)

	locationUseCase := usecase.New(dataService)

	locationController := controller.New(render.New(), locationUseCase)

	r := router.Setup(locationController)

	http.ListenAndServe("localhost:9090", r)
}
