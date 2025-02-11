package main

import (
	"fmt"
	"github.com/DiegoJCordeiro/golang-study/activity/client/cfg"
)

func main() {

	env, err := cfg.LoadConfiguration("./cmd/client")

	if err != nil {
		panic(err)
	}

	fmt.Println(env.DBPort)
}
