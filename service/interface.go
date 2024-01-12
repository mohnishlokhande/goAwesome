package service

import "awesomeProject/model"

type Account interface { 
	Create(a []model.Account) error
	Get() ([]model.Account, error)
	GetByAcct(senderAcct, receiverAcct string) []model.Account
	Update(senderAcct, receiverAcct string, amount float64)
}
