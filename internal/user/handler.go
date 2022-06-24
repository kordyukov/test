package user

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"test/internal/handlers"
)

const (
	usersURL = "/users"
	userURL  = "/users/:uuid"
)

type handler struct {
}

func NewHandler() handlers.Handler {
	return &handler{}
}

func (h *handler) Register(router *httprouter.Router) {
	router.GET(usersURL, h.GetList)
	router.POST(usersURL, h.CreateUser)
	router.GET(userURL, h.GetUserByUUID)
	router.PUT(userURL, h.UpdateUser)
	router.PATCH(userURL, h.PartiallyUpdateUser)
	router.DELETE(userURL, h.DeleteUser)
}

func (h *handler) GetList(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("this is list if users"))
}

func (h *handler) GetUserByUUID(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("this is GetUserByUUID"))
}
func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("this is CreateUser"))
}
func (h *handler) UpdateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("this is UpdateUser"))
}
func (h *handler) PartiallyUpdateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("this is PartiallyUpdateUser"))
}
func (h *handler) DeleteUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("this is DeleteUser"))
}
