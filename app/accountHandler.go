package app

import (
	"encoding/json"
	"net/http"

	"github.com/bhaktiutama/banking/dto"
	"github.com/bhaktiutama/banking/service"
	"github.com/gorilla/mux"
)

type AccountHandler struct {
	service service.AccountService
}

func (h AccountHandler) NewAccount(w http.ResponseWriter, r *http.Request) {

	// get from url
	vars := mux.Vars(r)
	customerId := vars["customer_id"]

	// decode from request body
	var request dto.NewAccountRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	} else {
		request.CustomerId = customerId
		account, appError := h.service.NewAccount(request)
		if appError != nil {
			writeResponse(w, appError.Code, appError.Message)
		} else {
			writeResponse(w, http.StatusCreated, account)
		}
	}
}

func (h AccountHandler) MakeTransaction(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	customerId := vars["customer_id"]
	accountId := vars["account_id"]

	var request dto.TransactionRequest

	// log.Println(json.NewDecoder(r.Body))

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	} else {

		request.AccountId = accountId
		request.CustomerId = customerId

		// fmt.Printf("test: %f", request.Amount)

		transaction, appError := h.service.MakeTransaction(request)
		if appError != nil {
			writeResponse(w, appError.Code, appError.Message)
		} else {
			writeResponse(w, http.StatusCreated, transaction)
		}

	}

}
