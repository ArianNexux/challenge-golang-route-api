package usecase
import "github.com/ArianNexux/challenge-golang-route-api/internal/route/entity"

type CreateRouteInputDTO struct {
	Name string `json:"name"`
	Destination entity.Coord `json:"destination"`
	Source entity.Coord `json:"source"`
}


type CreateRouteOutputDTO struct {
	Name string `json:"name"`
	Destination entity.Coord `json:"destination"`
	Source entity.Coord `json:"source"`
}

type CreateRouteUseCase struct {
	Repository entity.RouteRepository
}

func NewCreateRouteUseCase(repository entity.RouteRepository) *CreateRouteUseCase {
	return &CreateRouteUseCase{
        Repository: repository,
    }
}

func (uc *CreateRouteUseCase) Execute(input CreateRouteInputDTO) (error, CreateRouteOutputDTO) {
	route := entity.NewRoute(input.Name, input.Source, input.Destination)

    err := uc.Repository.Create(route)
    if err!= nil {
        return err, CreateRouteOutputDTO{}
    }
    return nil, CreateRouteOutputDTO{
        Name: route.Name,
        Destination: route.Destination,
        Source: route.Source,
    }
}