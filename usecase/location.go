package usecase

import "github.com/marcodeltorob/golang-bootcamp-2020/model"

type DBService interface {
	GetLocationById(woid int) (*model.Location, error)
	GetAllLocations() ([]model.Location, error)
}
type LocationUseCase struct {
	dbService DBService
}

func New(dbs DBService) LocationUseCase {
	return LocationUseCase{dbs}
}

func (uc LocationUseCase) GetAllLocations() ([]model.Location, error) {

	return uc.dbService.GetAllLocations()
}

func (uc LocationUseCase) GetLocation(woeid int) (*model.Location, error) {
	return uc.dbService.GetLocationById(woeid)
}
