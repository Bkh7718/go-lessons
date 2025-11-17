package main

import (
	
	"Flag"
	"fmt"        // Все нужные нам библиотеки 
	"strings"

	
)

func main() {

	//template := `ДОБРО ПОЖАЛОВАТЬ, {username}! Ваш код доступа: {code}`
	
	flag.Parse()

	Number:= flag.Arg(0) // Тут мы получаем аргумент из командной строки

	

	
	
	correctNumber := strings.ReplaceAll(Number, "(" , "")
	correctNumber = strings.ReplaceAll(correctNumber, ")" , "")
	correctNumber = strings.ReplaceAll(correctNumber, "-" , "")
	correctNumber = strings.ReplaceAll(correctNumber, " " , "")	 


	fmt.Println(correctNumber) // Вуаля!

}
