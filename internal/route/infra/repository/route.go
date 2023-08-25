package repository

import ("database/sql"
"encoding/json"
 "github.com/ArianNexux/challenge-golang-route-api/internal/route/entity")
type RouteRepositoryMySQL struct {
	db *sql.DB
}


func NewRouteRepositoryMySQL(db *sql.DB) *RouteRepositoryMySQL {
	return &RouteRepositoryMySQL{
        db: db,
    }
}


func (r *RouteRepositoryMySQL) Create(route *entity.Route)  (error) {
    routeSource, _ := json.Marshal(&route.Source)
    routeDestination, _ := json.Marshal(&route.Destination)
	query := `INSERT INTO routes (name, source, destination) VALUES (?, ?, ?)`
    _, err := r.db.Exec(query,route.Id, route.Name, routeSource, routeDestination)
    if err!= nil {
        return nil
    }

    return nil
}

func (r *RouteRepositoryMySQL) List() ([]*entity.Route, error) { 

    rows, err := r.db.Query("SELECT id, name, source, destination FROM routes")
    if err!= nil {
        return nil, err
    }
    defer rows.Close()

    var routes []*entity.Route
    for rows.Next() {
      
        var route entity.Route
        var routeSource, routeDestination []byte
        err := rows.Scan(&route.Id, &route.Name, &routeSource, &routeDestination)
        if err!= nil {
            return nil, err
        }
        _ = json.Unmarshal(routeSource, &route.Source)
        _ = json.Unmarshal(routeDestination, &route.Destination)
       
        routes = append(routes, &route)
    }

    return routes, nil
}