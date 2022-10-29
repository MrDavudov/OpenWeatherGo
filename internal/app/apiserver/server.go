package apiserver

import (
	"encoding/json"
	"net/http"
	"syscall/js"

	"github.com/MrDavudov/OpenWeatherGO/internal/app/store"
	"github.com/gorilla/mux"
)

type server struct {
	router *mux.Router
	store  store.Store
}

func newServer(store store.Store) *server {
	s := &server{
		router:       mux.NewRouter(),
		store:        store,
	}

	s.configureRouter()

	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) configureRouter() {
	s.router.HandleFunc("/city", s.handleCityCreate()).Methods("POST")
	s.router.HandleFunc("/", s.handelAllCities()).Methods("GET")
}

func (s *server) handleCityCreate() http.HandlerFunc {
	type request struct {
		Name	string	`json:"name"`
		Country	string	`json:"country"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		w := &model.Weather{
			Name: req.Name,
			Counter: req.Country,
		}
		if err := s.store.City().Create(w); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		s.respond(w, r, http.StatusCreated, w)
	}
}

func (s *server) handelAllCities() http.HandlerFunc {

}

func (s *server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(w, r, code, map[string]string{"error": err.Error()})
}

func (s *server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}