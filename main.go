package main

import (
	"fmt"
	"math/rand/v2"
)

type account struct {
	login    string
	password string
	url      string
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

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890-*!")

func main() {

	login := promptData("Введите логин")
	//password := promptData("Введите пароль")
	url := promptData("Введите URL")

	account1 := account{
		login: login,
		url:   url,
	}
	account1.generatePassword(12)
	account1.outputPassword()
	fmt.Println(account1)
}

func promptData(prompt string) string {
	fmt.Print(prompt + ": ")
	var res string
	fmt.Scan(&res)
	return res
}
