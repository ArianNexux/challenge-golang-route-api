package main 
import(
	"github.com/ArianNexux/challenge-golang-route-api/internal/route/infra/repository"
	"github.com/ArianNexux/challenge-golang-route-api/internal/route/usecase"
	"database/sql"
	"net/http"
	"github.com/go-chi/chi/v5"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"log"
)

func main() {

 	db, err := sql.Open("mysql", "root:root@tcp(db:3306)/routes")
	defer db.Close()
	repository := repository.NewRouteRepositoryMySQL(db)
	if err != nil {
		panic(err)
	}
	r := chi.NewRouter()
	r.Route("/api/routes", func(r chi.Router){
		r.Get("/", func(rw http.ResponseWriter, r *http.Request){
			uc := usecase.NewListRouteUseCase(repository)
			err, response := uc.Execute()
			if err != nil {
				rw.WriteHeader(400)
				rw.Write([]byte(err.Error()))
                rw.Write([]byte("Erro ao listar a rota"))
                return
			}
			rw.Header().Set("Content-Type", "application/json")
			json.NewEncoder(rw).Encode(response)
		})


		r.Post("/", func(rw http.ResponseWriter, r *http.Request){
			var input usecase.CreateRouteInputDTO
			
			err := json.NewDecoder(r.Body).Decode(&input)
			if err != nil {
				log.Printf("error decoding sakura response: %v", err)
				if e, ok := err.(*json.SyntaxError); ok {
					log.Printf("syntax error at byte offset %d", e.Offset)
				}
				log.Printf("sakura response: %q", r.Body)
			}

			uc := usecase.NewCreateRouteUseCase(repository)

		    err, response := uc.Execute(input)
			if err != nil  {
				rw.Write([]byte(err.Error()))
                rw.Write([]byte("Erro ao cadastrar"))
                return
			}

			rw.WriteHeader(201)
			json.NewEncoder(rw).Encode(response)
			return
		})
	})

	fmt.Println(http.ListenAndServe(":3006", r))
	fmt.Println("Rodando o servidor na porta 30006")
}