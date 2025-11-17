package main

import (
	"Flag"
	"fmt"
	"strings"

	
)

func main() {

	template := `ДОБРО ПОЖАЛОВАТЬ, {username}! Ваш код доступа: {code}`
	
	flag.Parse()

	username := flag.Arg(0) // Тут мы получаем ПЕРВЫЙ! аргумент из командной строки
	code := flag.Arg(1) // Тут мы получаем ВТОРОЙ! аргумент из командной строки

	message := strings.ReplaceAll(template, "{username}", username) // Заменяем {username} на значение переменной Которой присвоили первый аргумент
	message = strings.ReplaceAll(message, "{code}", code) // Заменяем {code} на значение переменной Которой присвоили второй аргумент


	fmt.Println(message) // Вуаля!

}
