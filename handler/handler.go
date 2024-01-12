package handler

import (
	"awesomeProject/model"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type handler struct {
	service Account
}

func New(s Account) handler {
	return handler{service: s}
}

var flag bool

func (h handler) Router(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.Get(w, r)
	case http.MethodPatch:
		h.Transfer(w, r)
	case http.MethodPost:
		h.Create(w, r)
	}
}

func (h handler) Create(w http.ResponseWriter, r *http.Request) {

	var acct []model.Account
	json.NewDecoder(r.Body).Decode(&acct)

	err := h.service.Create(acct)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "cannot read the data")
	} else {
		flag = true
		w.WriteHeader(http.StatusCreated)
		fmt.Fprint(w, "data is inserted")
	}
}

func (h handler) Get(w http.ResponseWriter, r *http.Request) {
	if !flag {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "insert is not done yet")
		return
	}

	e, err := h.service.Get()
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "data not found")
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(e)
	}
}

func (h handler) Transfer(w http.ResponseWriter, r *http.Request) {
	if !flag {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "insert is not done yet")
		return
	}

	amount := r.FormValue("amount")
	amount = strings.TrimSpace(amount)
	sender := r.FormValue("sender")
	sender = strings.TrimSpace(sender)
	receiver := r.FormValue("receiver")
	receiver = strings.TrimSpace(receiver)

	if amount == "" || sender == "" || receiver == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "missing transaction details")
		return
	}

	amt, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "incorrect amount")
		return
	}

	err = h.service.Transfer(sender, receiver, amt)
	if err == errors.New("data not found") {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "data not found")
		return
	}

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "balance is low")
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "balance transferred")
}
