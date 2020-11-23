package controller

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/marcodeltorob/golang-bootcamp-2020/model"
	"github.com/unrolled/render"
)

type LocationUseCase interface {
	GetLocation(woeid int) (*model.Location, error)
	GetAllLocations() ([]model.Location, error)
}

type Location struct {
	render  *render.Render
	useCase LocationUseCase
}

func New(r *render.Render, uc LocationUseCase) Location {

	return Location{r, uc}
}

func (l Location) GetAll(w http.ResponseWriter, r *http.Request) {
	locations, err := l.useCase.GetAllLocations()

	if err != nil {
		response := model.Response{
			Status:  http.StatusBadRequest,
			Message: "An Error has occurred: " + err.Error(),
		}
		l.render.JSON(w, http.StatusBadRequest, response)
		return
	}

	l.render.JSON(w, http.StatusOK, locations)
}

func (l Location) GetOne(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	woeid, err := strconv.Atoi(vars["woeid"])

	if err != nil {
		response := model.Response{
			Status:  http.StatusBadRequest,
			Message: "Woeid must be a int value: " + err.Error(),
		}
		l.render.JSON(w, http.StatusBadRequest, response)
		return
	}

	location, err := l.useCase.GetLocation(woeid)

	if err != nil {
		response := model.Response{
			Status:  http.StatusBadRequest,
			Message: "An Error has ocurred: " + err.Error(),
		}
		l.render.JSON(w, http.StatusBadRequest, response)
		return
	}

	l.render.JSON(w, http.StatusOK, location)
}
