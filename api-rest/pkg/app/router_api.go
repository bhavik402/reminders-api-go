package app

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/bhavik402/remidners-api-go/api-rest/internal/data/reminders"
	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
	"go.uber.org/zap"
)

func (a *App) WebApiRouter() *httprouter.Router {
	router := httprouter.New()
	router.HandleOPTIONS = true
	router.GlobalOPTIONS = http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			addCorsHeaders(&w)
		},
	)
	router.HandlerFunc(http.MethodGet, "/v1/reminders/", a.GetAllReminders)
	router.HandlerFunc(http.MethodPost, "/v1/reminders/", a.CreateAReminder)

	router.HandlerFunc(http.MethodGet, "/v1/reminders/:id", a.GetAReminder)
	router.HandlerFunc(http.MethodDelete, "/v1/reminders/:id", a.RemoveAReminder)

	router.HandlerFunc(http.MethodPut, "/v1/reminders/status/:id", a.FlipStatusReminder)
	router.HandlerFunc(http.MethodPut, "/v1/reminders/flag/:id", a.FlipFlagReminder)

	return router
}

func addCorsHeaders(wr *http.ResponseWriter) {
	w := *wr
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
}

func (a *App) LogRequest(msg string, r *http.Request) {
	a.Logger.Info(
		msg,
		zap.String("method", r.Method),
		zap.String("url", r.URL.String()),
	)
}

func (a *App) GetAllReminders(w http.ResponseWriter, r *http.Request) {
	a.LogRequest("GET ALL REMINDERS", r)
	data, err := a.Models.Reminders.Storage.ReadAll()
	if err != nil {
		fmt.Println(fmt.Errorf("failed to read data: %w", err))
	}

	js, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return
	}

	addCorsHeaders(&w)
	w.WriteHeader(http.StatusOK)
	w.Write(js)
}

func (a *App) GetAReminder(w http.ResponseWriter, r *http.Request) {
	a.LogRequest("GET A REMINDER", r)
	params := strings.Split(r.URL.Path, "/")
	fmt.Println(params[3])

	data, err := a.Models.Reminders.Storage.Read(params[3])
	if err != nil {
		fmt.Println(fmt.Errorf("failed to read data: %w", err))
		w.WriteHeader(http.StatusNotFound)
	}

	js, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return
	}

	addCorsHeaders(&w)
	w.WriteHeader(http.StatusOK)
	w.Write(js)
}

func (a *App) CreateAReminder(w http.ResponseWriter, r *http.Request) {
	a.LogRequest("CREATE A REMINDER", r)
	body := ""
	// _, err := r.Body.Read([]byte(body))
	// fmt.Println(body)
	// if err != nil {
	// 	fmt.Println(fmt.Println("failed to read request body: %w", err))
	// }

	fmt.Println(body)

	var rm reminders.Reminder
	err := json.NewDecoder(r.Body).Decode(&rm)
	// err = json.Unmarshal([]byte(body), &rm)
	if err != nil {
		fmt.Println(fmt.Errorf("failed to unmarshal request body: %w", err))
	}
	rm.Id = uuid.NewString()

	err = a.Models.Reminders.Storage.Save(&rm)

	if err != nil {
		// fmt.Println(err.Error())
		//todo: these currently do the same thing but need to differentiate and clean
		switch {
		case errors.Is(err, reminders.ErrFailedToOpenDB):
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		case errors.Is(err, reminders.ErrFailedToOpenDB):
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		case errors.Is(err, reminders.ErrFailedToOpenDB):
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		default:
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
	} else {
		addCorsHeaders(&w)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Successfully Created"))
	}
}

func (a *App) FlipStatusReminder(w http.ResponseWriter, r *http.Request) {
	a.LogRequest("UPDATE A REMINDER", r)
	params := strings.Split(r.URL.Path, "/")
	fmt.Println(params[0])
	fmt.Println(params[1])
	fmt.Println(params[2])
	fmt.Println(params[2])
	fmt.Println(params[3])
	fmt.Println(params[4])

	_, err := a.Models.Reminders.Storage.FlipStatus(params[4])
	if err != nil {
		fmt.Println(fmt.Errorf("failed to update data: %w", err))
	}

	addCorsHeaders(&w)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Updated A Reminder"))
}

func (a *App) FlipFlagReminder(w http.ResponseWriter, r *http.Request) {
	a.LogRequest("UPDATE FLAG for a REMINDER", r)
	params := strings.Split(r.URL.Path, "/")
	fmt.Println(params[0])
	fmt.Println(params[1])
	fmt.Println(params[2])
	fmt.Println(params[2])
	fmt.Println(params[3])
	fmt.Println(params[4])

	_, err := a.Models.Reminders.Storage.FlipFlag(params[4])
	if err != nil {
		fmt.Println(fmt.Errorf("failed to update data: %w", err))
	}

	addCorsHeaders(&w)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Updated Flag on a Reminder"))
}

func (a *App) RemoveAReminder(w http.ResponseWriter, r *http.Request) {
	a.LogRequest("DELETE A REMINDER", r)
	params := strings.Split(r.URL.Path, "/")
	fmt.Println(params[0])
	fmt.Println(params[1])
	fmt.Println(params[2])
	fmt.Println(params[2])
	fmt.Println(params[3])

	err := a.Models.Reminders.Storage.Remove(params[3])
	if err != nil {
		fmt.Println(fmt.Errorf("failed to read data: %w", err))
	}

	addCorsHeaders(&w)
	w.WriteHeader(http.StatusOK)
}
