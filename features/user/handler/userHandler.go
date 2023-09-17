package handler

import (
	"bookingoto-try/features/user"
	"bookingoto-try/helpers"
	"encoding/json"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type UserHandler struct {
	userUseCase user.UserUseCase
}

func NewUserHandler(ctm *mux.Router, userUseCase user.UserUseCase) {

	middleUrl := os.Getenv("MIDDLE_URL")
	userHandlers := UserHandler{userUseCase}
	ctm.HandleFunc(middleUrl+"/customer", userHandlers.H_GetAll).Methods("GET")
	ctm.HandleFunc(middleUrl+"/customer/{cst_id}", userHandlers.H_GetOneId).Methods("GET")

}

func (e *UserHandler) H_GetAll(w http.ResponseWriter, r *http.Request) {
	result, err := e.userUseCase.ReadAllUser()
	if err != nil {
		helpers.Logger("error", "In Server: "+err.Error())
		msg := helpers.MsgErr(http.StatusBadRequest, "bad request", err.Error())
		helpers.Response(w, http.StatusBadRequest, msg)
		return
	}

	rMsg := helpers.MsgOk(200, "Successfully")
	rMsg["customer"] = result
	helpers.Logger("info", "get all customer")
	helpers.Response(w, http.StatusOK, rMsg)
}

func (e *UserHandler) H_GetOneId(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	cstId := params["cst_id"]

	result, err := e.userUseCase.ReadSingleId(cstId)
	if err != nil {
		helpers.Logger("error", "In Server: "+err.Error())
		msg := helpers.MsgErr(http.StatusBadRequest, "bad request", err.Error())
		helpers.Response(w, http.StatusBadRequest, msg)
		return
	}

	logger, _ := json.Marshal(result)
	helpers.Logger("info", "get one user: "+string(logger))
	helpers.Response(w, http.StatusOK, result)
}
