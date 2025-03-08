package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
	"time"
)

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

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890-*!")

func main() {

	login := promptData("Введите логин")
	password := promptData("Введите пароль")
	url := promptData("Введите URL")

	myAccount, err := newAccountWithTimeStamp(login, password, url)
	if err != nil {
		fmt.Println("Неверный формат URL или Логин")
		return
	}
	myAccount.acc.outputPassword()
	fmt.Println(myAccount)
}

func promptData(prompt string) string {
	fmt.Print(prompt + ": ")
	var res string
	fmt.Scanln(&res)
	return res
}
