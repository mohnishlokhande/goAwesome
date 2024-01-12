package store

import (
	"awesomeProject/model"
	"awesomeProject/service"
	"errors"
	"fmt"
	"strconv"
)

var acct []model.Account

type store struct{}

func New() service.Account { 
	return &store{}
}

func (s store) Create(a []model.Account) error {
	if a == nil {
		return errors.New("data not found")
	}

	acct = a

	return nil
}

func (s store) Get() ([]model.Account, error) {
	if acct == nil {
		return nil, errors.New("data not found")
	}

	return acct, nil
}

func (s store) GetByAcct(senderAcct, receiverAcct string) []model.Account {
	var data []model.Account
	for i, tc := range acct {
		if tc.Id == senderAcct || tc.Id == receiverAcct {
			data = append(data, acct[i])
		}
	}

	return data
}

func (s store) Update(senderAcct, receiverAcct string, amount float64) {
	for i, tc := range acct {
		if tc.Id == senderAcct {
			senderBal, _ := strconv.ParseFloat(tc.Balance, 64)
			senderBal -= amount
			acct[i].Balance = fmt.Sprintf("%.2f", senderBal)
		}

		if tc.Id == receiverAcct {
			receiverBal, _ := strconv.ParseFloat(tc.Balance, 64)
			receiverBal += amount
			acct[i].Balance = fmt.Sprintf("%.2f", receiverBal)
		}
	}
}
