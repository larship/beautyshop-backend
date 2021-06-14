package server

import (
	"encoding/json"
	"github.com/larship/beautyshop/auth"
	"github.com/larship/beautyshop/models"
	"log"
	"net/http"
	"regexp"
)

type createCheckInParams struct {
	BeautyshopUuid  string `json:"beautyshopUuid"`
	ClientUuid      string `json:"clientUuid"`
	WorkerUuid      string `json:"workerUuid"`
	ServiceTypeUuid string `json:"serviceTypeUuid"`
	StartDate       int64  `json:"startDate"`
}

type cancelCheckInParams struct {
	Uuid string `json:"uuid"`
}

func (s *Server) MakeRoutes() {
	s.router.HandleFunc("/", mainHandler)
	s.router.HandleFunc("/beautyshops", authMiddleware(getBeautyshopsHandler)) // TODO /beautyshop/list
	s.router.HandleFunc("/beautyshop", authMiddleware(getBeautyshopHandler))
	s.router.HandleFunc("/beautyshop/list-for-admin", authMiddleware(getBeautyshopListByAdminHandler))
	s.router.HandleFunc("/beautyshop/service-types", authMiddleware(getBeautyshopServiceTypesHandler))
	s.router.HandleFunc("/workers", authMiddleware(getWorkersHandler))
	s.router.HandleFunc("/workers/add", authMiddleware(addWorkerHandler))
	s.router.HandleFunc("/check-in/list-for-client", authMiddleware(getClientCheckInList))
	s.router.HandleFunc("/check-in/list-for-beautyshop", authMiddleware(getBeautyshopCheckInList))
	s.router.HandleFunc("/check-in/create", authMiddleware(createCheckInHandler))
	s.router.HandleFunc("/check-in/cancel", authMiddleware(cancelCheckInHandler))
	s.router.HandleFunc("/client/auth", authClientHandler)
	s.router.HandleFunc("/client/new", newClientHandler)
	s.router.HandleFunc("/admin/auth", authAdminHandler)
	s.router.HandleFunc("/admin/send-security-code", sendSecurityCodeHandler)
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	ResponseError(w, r, http.StatusBadRequest, "")
}

func authMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		clientUuid := r.Header.Get("Auth-Client-Uuid")
		sessionId := r.Header.Get("Auth-Session-Id")
		salt := r.Header.Get("Auth-Salt")

		log.Printf("AuthMiddleware: clientUuid = %s, sessionId = %s, salt = %s", clientUuid, sessionId, salt)

		if clientUuid == "" || sessionId == "" || salt == "" {
			ResponseError(w, r, http.StatusForbidden, "Ошибка аутентификации")
			return
		}

		client := auth.CheckAuth(clientUuid, sessionId, salt)

		if client == nil {
			ResponseError(w, r, http.StatusForbidden, "Ошибка аутентификации")
			return
		}

		log.Printf("AuthMiddleware: успешная авторизация")

		next(w, r)
	}
}

func getBeautyshopServiceTypesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		ResponseError(w, r, http.StatusBadRequest, "")
		return
	}

	beautyshopUuid := r.URL.Query().Get("beautyshopUuid")
	if beautyshopUuid == "" {
		ResponseError(w, r, http.StatusBadRequest, "Не указан салон красоты")
		return
	}

	serviceTypeList := models.GetBeautyshopServiceTypes(beautyshopUuid)
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

func getBeautyshopListByAdminHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		ResponseError(w, r, http.StatusBadRequest, "")
		return
	}

	adminUuid := r.URL.Query().Get("adminUuid")
	if adminUuid == "" {
		ResponseError(w, r, http.StatusBadRequest, "Не указан идентификатор администратора")
		return
	}

	beautyshop := models.GetBeautyshopListByAdmin(adminUuid)
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

func getClientCheckInList(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		ResponseError(w, r, http.StatusBadRequest, "")
		return
	}

	clientUuid := r.FormValue("uuid")
	if clientUuid == "" {
		ResponseError(w, r, http.StatusBadRequest, "Не указан идентификатор клиента")
		return
	}

	checkInList := models.GetClientCheckInList(clientUuid)
	ResponseSuccess(w, http.StatusOK, checkInList)
}

func getBeautyshopCheckInList(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		ResponseError(w, r, http.StatusBadRequest, "")
		return
	}

	beautyshopUuid := r.FormValue("uuid")
	dateFrom := r.FormValue("dateFrom")
	dateTo := r.FormValue("dateTo")
	if beautyshopUuid == "" || dateFrom == "" || dateTo == "" {
		ResponseError(w, r, http.StatusBadRequest, "Недостаточно данных")
		return
	}

	checkInList := models.GetBeautyshopCheckInList(beautyshopUuid, dateFrom, dateTo)
	ResponseSuccess(w, http.StatusOK, checkInList)
}

func createCheckInHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		ResponseError(w, r, http.StatusBadRequest, "")
		return
	}

	var params createCheckInParams
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&params); err != nil {
		ResponseError(w, r, http.StatusBadRequest, "Произошла ошибка при парсинге запроса")
		return
	}

	if params.BeautyshopUuid == "" || params.ClientUuid == "" || params.WorkerUuid == "" ||
		params.ServiceTypeUuid == "" || params.StartDate == 0 {
		ResponseError(w, r, http.StatusBadRequest, "Недостаточно данных")
		return
	}

	checkInItem := models.CreateCheckIn(params.BeautyshopUuid, params.ClientUuid, params.WorkerUuid, params.ServiceTypeUuid, params.StartDate)

	if checkInItem != nil {
		ResponseSuccess(w, http.StatusOK, checkInItem)
	} else {
		ResponseError(w, r, http.StatusBadRequest, "Ошибка при добавлении записи")
	}
}

func cancelCheckInHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		ResponseError(w, r, http.StatusBadRequest, "")
		return
	}

	var params cancelCheckInParams
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&params); err != nil {
		ResponseError(w, r, http.StatusBadRequest, "Произошла ошибка при парсинге запроса")
		return
	}

	if params.Uuid == "" {
		ResponseError(w, r, http.StatusBadRequest, "Недостаточно данных")
		return
	}

	status := models.CancelCheckIn(params.Uuid)
	ResponseSuccess(w, http.StatusOK, status)
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

func authAdminHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		ResponseError(w, r, http.StatusBadRequest, "")
		return
	}

	re := regexp.MustCompile(`\D`)

	phone := string(re.ReplaceAll([]byte(r.FormValue("phone")), []byte("")))
	code := string(re.ReplaceAll([]byte(r.FormValue("code")), []byte("")))

	if phone == "" || code == "" {
		ResponseError(w, r, http.StatusBadRequest, "Недостаточно данных для аутентификации администратора")
		return
	}

	client := auth.CheckAdminAuth(phone, code)

	if client != nil {
		ResponseSuccess(w, http.StatusOK, client)
	} else {
		ResponseError(w, r, http.StatusBadRequest, "Ошибка при авторизации")
	}
}

func sendSecurityCodeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		ResponseError(w, r, http.StatusBadRequest, "")
		return
	}

	re := regexp.MustCompile(`\D`)

	phone := string(re.ReplaceAll([]byte(r.FormValue("phone")), []byte("")))

	if phone == "" {
		ResponseError(w, r, http.StatusBadRequest, "Не указан телефон для отправки кода подтверждения")
		return
	}

	status := auth.SendSecurityCode(phone)

	if !status {
		ResponseError(w, r, http.StatusBadRequest, "Ошибка отправки кода")
		return
	}

	ResponseSuccess(w, http.StatusOK, status)
}
