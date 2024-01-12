package service

import (
	"awesomeProject/handler"
	"awesomeProject/model"
	"errors"
	"strconv"
)

type service struct {
	store Account
}

// New - is a factory function for delivery layer
// nolint:golint // service should not be used without proper initialisation with required dependency
func New(s Account) handler.Account { 
	return service{store: s}
}

func (s service) Create(a []model.Account) error {
	return s.store.Create(a)
}

func (s service) Get() ([]model.Account, error) {
	return s.store.Get()
}

func (s service) Transfer(senderAcct, receiverAcct string, amount float64) error {
	var sender model.Account

	resp := s.store.GetByAcct(senderAcct, receiverAcct)
	if len(resp) != 2 {
		return errors.New("data not found")
	}

	if resp[0].Id == senderAcct {
		sender = resp[0]
	} else {
		sender = resp[1]
	}

	amt, _ := strconv.ParseFloat(sender.Balance, 64)
	if amount > amt {
		return errors.New("balance is low")
	}

	s.store.Update(senderAcct, receiverAcct, amount)

	return nil
}
