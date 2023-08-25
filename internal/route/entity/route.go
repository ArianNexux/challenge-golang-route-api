package entity


type Coord struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type Route struct {
	Id       string
	Name     string
	Source   Coord
	Destination Coord
}

type RouteRepository interface{
	Create(route *Route) (error)
	List() ([]*Route, error)
}

func NewRoute(name string, source Coord, destination Coord) *Route{
	return &Route{
        Name:     name,
        Source:   source,
        Destination: destination,
	}
}