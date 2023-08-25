package usecase
import "github.com/ArianNexux/challenge-golang-route-api/internal/route/entity"

type ListRouteInputDTO struct {
	Id string
	Name string
	Destination entity.Coord 
	Source entity.Coord 
}


type ListRouteOutputDTO struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Destination entity.Coord `json:"destination"`
	Source entity.Coord `json:"source"`
}

type ListRouteUseCase struct {
	Repository entity.RouteRepository
}

func NewListRouteUseCase(repository entity.RouteRepository) *ListRouteUseCase {
	return &ListRouteUseCase{
        Repository: repository,
    }
}

func (u *ListRouteUseCase) Execute() (error, []ListRouteOutputDTO) {
	route, err := u.Repository.List()
    if err!= nil {
        return err, nil
    }
	var result []ListRouteOutputDTO

	for _, r := range route {
		
		result = append(result, ListRouteOutputDTO{
            Id: r.Id,
            Name: r.Name,
            Destination: r.Destination,
            Source: r.Source,
        })
	}
    return nil, result
}