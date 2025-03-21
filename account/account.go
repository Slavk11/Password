package account

import (
	"errors"
	"github.com/fatih/color"
	"math/rand/v2"
	"net/url"
	"time"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890-*!")

type Account struct {
	login    string
	password string
	url      string
}

type AccountWithTimeStamp struct {
	createdAt time.Time
	updatedAt time.Time
	Account
}

func (acc *Account) OutputPassword() {
	color.Cyan(acc.login)
}

func (acc *Account) generatePassword(n int) {
	res := make([]rune, n)
	for i := range res {
		res[i] = letterRunes[rand.IntN(len(letterRunes))]
	}
	acc.password = string(res)
}

//func newAccount(login, password, urlString string) (*Account, error) {
//	if login == "" {
//		return nil, errors.New("INVALID_LOGIN")
//	}
//	_, err := url.ParseRequestURI(urlString)
//	if err != nil {
//		return nil, errors.New("INVALID_URL")
//	}
//	newAcc := &Account{
//		login,
//		password,
//		urlString,
//	}
//	if password == "" {
//		newAcc.generatePassword(12)
//	}
//	return newAcc, nil
//
//}

func NewAccountWithTimeStamp(login, password, urlString string) (*AccountWithTimeStamp, error) {
	if login == "" {
		return nil, errors.New("INVALID_LOGIN")
	}
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("INVALID_URL")
	}
	newAcc := &AccountWithTimeStamp{
		time.Now(),
		time.Now(),
		Account{
			login,
			password,
			urlString,
		},
	}
	if password == "" {
		newAcc.generatePassword(12)
	}
	return newAcc, nil

}
