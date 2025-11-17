package main

import (
	"Flag"
	"fmt"
	"strings"

	
)

func main() {

	template := `ДОБРО ПОЖАЛОВАТЬ, {username}! Ваш код доступа: {code}`
	
	flag.Parse()

	username := flag.Arg(0)
	code := flag.Arg(1)

	message := strings.ReplaceAll(template, "{username}", username)
	message = strings.ReplaceAll(message, "{code}", code)


	fmt.Println(message)

}
