package server

import (
	"github.com/larship/barbershop/models"
	"net/http"
)

func (s *Server) MakeRoutes() {
	s.router.HandleFunc("/", mainHandler)
	s.router.HandleFunc("/haircut-types", getHaircutTypesHandler)
	s.router.HandleFunc("/barbershops", getBarbershopsHandler)
	s.router.HandleFunc("/hairdressers", getHairdressersHandler)
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	ResponseError(w, http.StatusBadRequest, "")
}

func getHaircutTypesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		ResponseError(w, http.StatusBadRequest, "")
		return
	}

	barbershopUuid := r.URL.Query().Get("barbershop")
	if barbershopUuid == "" {
		ResponseError(w, http.StatusBadRequest, "Не указана парикмахерская")
		return
	}

	haircutTypeList := models.GetHaircutTypes(barbershopUuid)
	ResponseSuccess(w, http.StatusOK, haircutTypeList)
}

func getBarbershopsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		ResponseError(w, http.StatusBadRequest, "")
		return
	}

	city := r.URL.Query().Get("city")
	if city == "" {
		ResponseError(w, http.StatusBadRequest, "Не указан город")
		return
	}

	barbershopList := models.GetBarbershops(city)
	ResponseSuccess(w, http.StatusOK, barbershopList)
}

func getHairdressersHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		ResponseError(w, http.StatusBadRequest, "")
		return
	}

	city := r.URL.Query().Get("barbershop")
	if city == "" {
		ResponseError(w, http.StatusBadRequest, "Не указана парикмахерская")
		return
	}

	hairdressersList := models.GetHairdressers(city)
	ResponseSuccess(w, http.StatusOK, hairdressersList)
}
