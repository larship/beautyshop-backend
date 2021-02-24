package server

import (
	"github.com/larship/beautyshop/auth"
	"github.com/larship/beautyshop/models"
	"net/http"
)

func (s *Server) MakeRoutes() {
	s.router.HandleFunc("/", mainHandler)
	s.router.HandleFunc("/service-types", getServiceTypesHandler)
	s.router.HandleFunc("/beautyshops", getBeautyshopsHandler)
	s.router.HandleFunc("/beautyshop", getBeautyshopHandler)
	s.router.HandleFunc("/workers", getWorkersHandler)
	s.router.HandleFunc("/workers/add", addWorkerHandler)
	s.router.HandleFunc("/schedule", getScheduleHandler)
	s.router.HandleFunc("/client/auth", authClientHandler)
	s.router.HandleFunc("/client/new", newClientHandler)
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	ResponseError(w, r, http.StatusBadRequest, "")
}

func getServiceTypesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		ResponseError(w, r, http.StatusBadRequest, "")
		return
	}

	beautyshopUuid := r.URL.Query().Get("beautyshopUuid")
	if beautyshopUuid == "" {
		ResponseError(w, r, http.StatusBadRequest, "Не указан салон красоты")
		return
	}

	serviceTypeList := models.GetServiceTypes(beautyshopUuid)
	ResponseSuccess(w, http.StatusOK, serviceTypeList)
}

func getBeautyshopsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		ResponseError(w, r, http.StatusBadRequest, "")
		return
	}

	city := r.URL.Query().Get("city")
	if city == "" {
		ResponseError(w, r, http.StatusBadRequest, "Не указан город")
		return
	}

	beautyshopList := models.GetBeautyshops(city)
	ResponseSuccess(w, http.StatusOK, beautyshopList)
}

func getBeautyshopHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		ResponseError(w, r, http.StatusBadRequest, "")
		return
	}

	beautyshopUuid := r.URL.Query().Get("uuid")
	if beautyshopUuid == "" {
		ResponseError(w, r, http.StatusBadRequest, "Не указан салон красоты")
		return
	}

	beautyshop := models.GetBeautyshopByUuid(beautyshopUuid)
	ResponseSuccess(w, http.StatusOK, beautyshop)
}

func getWorkersHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		ResponseError(w, r, http.StatusBadRequest, "")
		return
	}

	beautyshopUuid := r.URL.Query().Get("beautyshopUuid")
	if beautyshopUuid == "" {
		ResponseError(w, r, http.StatusBadRequest, "Не указан салон красоты")
		return
	}

	workersList := models.GetWorkers(beautyshopUuid)
	ResponseSuccess(w, http.StatusOK, workersList)
}

func addWorkerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		ResponseError(w, r, http.StatusBadRequest, "")
		return
	}

	beautyshopUuid := r.FormValue("beautyshopUuid")
	fullName := r.FormValue("fullName")
	description := r.FormValue("description")

	if beautyshopUuid == "" || fullName == "" {
		ResponseError(w, r, http.StatusBadRequest, "Не указан UUID салона красоты или Fullname мастера")
		return
	}

	if models.AddWorker(beautyshopUuid, fullName, description) {
		ResponseSuccess(w, http.StatusOK, "")
	} else {
		ResponseError(w, r, http.StatusBadRequest, "Ошибка при добавлении")
	}
}

func getScheduleHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		ResponseError(w, r, http.StatusBadRequest, "")
		return
	}

	beautyshopUuid := r.FormValue("beautyshop")
	if beautyshopUuid == "" {
		ResponseError(w, r, http.StatusBadRequest, "Не указан салон красоты")
		return
	}

	schedule := models.GetScheduleItems(beautyshopUuid, "", "")
	ResponseSuccess(w, http.StatusOK, schedule)
}

func authClientHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		ResponseError(w, r, http.StatusBadRequest, "")
		return
	}

	clientUuid := r.FormValue("clientUuid")
	sessionId := r.FormValue("sessionId")
	salt := r.FormValue("salt")

	if clientUuid == "" || sessionId == "" || salt == "" {
		ResponseError(w, r, http.StatusBadRequest, "Недостаточно данных для аутентификации")
		return
	}

	client := auth.CheckAuth(clientUuid, sessionId, salt)

	if client != nil {
		ResponseSuccess(w, http.StatusOK, client)
	} else {
		ResponseError(w, r, http.StatusBadRequest, "Ошибка при авторизации")
	}
}

func newClientHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		ResponseError(w, r, http.StatusBadRequest, "")
		return
	}

	fullName := r.FormValue("fullName")
	phone := r.FormValue("phone")

	if fullName == "" || phone == "" {
		ResponseError(w, r, http.StatusBadRequest, "Не указано имя клиента или его телефон")
		return
	}

	user := auth.CreateUser(fullName, phone)

	if user != nil {
		ResponseSuccess(w, http.StatusOK, user)
	} else {
		ResponseError(w, r, http.StatusBadRequest, "Ошибка при добавлении клиента")
	}
}
