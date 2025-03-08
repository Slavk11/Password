package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
	"time"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890-*!")

type account struct {
	login    string
	password string
	url      string
}

type accountWithTimeStamp struct {
	createdAt time.Time
	updatedAt time.Time
	acc       account
}

func (acc *account) outputPassword() {
	fmt.Println(acc.login, acc.password, acc.url)
}

func (acc *account) generatePassword(n int) {
	res := make([]rune, n)
	for i := range res {
		res[i] = letterRunes[rand.IntN(len(letterRunes))]
	}
	acc.password = string(res)
}

//func newAccount(login, password, urlString string) (*account, error) {
//	if login == "" {
//		return nil, errors.New("INVALID_LOGIN")
//	}
//	_, err := url.ParseRequestURI(urlString)
//	if err != nil {
//		return nil, errors.New("INVALID_URL")
//	}
//	newAcc := &account{
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

func newAccountWithTimeStamp(login, password, urlString string) (*accountWithTimeStamp, error) {
	if login == "" {
		return nil, errors.New("INVALID_LOGIN")
	}
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("INVALID_URL")
	}
	newAcc := &accountWithTimeStamp{
		time.Now(),
		time.Now(),
		account{
			login,
			password,
			urlString,
		},
	}
	if password == "" {
		newAcc.acc.generatePassword(12)
	}
	return newAcc, nil

}
