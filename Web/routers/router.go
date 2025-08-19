package routers

import(
	"github.com/gorilla/mux"
	"web/controllers"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/movies", controllers.GetAllMovies).Methods("GET")
	router.HandleFunc("/api/movies", controllers.CreateMovie).Methods("POST")
	router.HandleFunc("/api/movies/{id}", controllers.MarkMovieAsWatched).Methods("PUT")
	router.HandleFunc("/api/movies/{id}", controllers.DeleteOneMovie).Methods("DELETE")
	router.HandleFunc("/api/movies", controllers.DeleteAllMovies).Methods("DELETE")

	return router;
}