package main

import "fmt"

func main() {
	//mas, entered := enter(15) //можно так ебаться но  не понял зачем пока еще ...
	//fmt.Println(mas)
	//fmt.Println(entered)
	fmt.Println(enter(20)) // или помно прописать вот так ! Такого хуя ??
}

func enter(age int) (string, bool) {
	if age >= 18 {
		return "Входи мужик", true
	} /*else {
		return "Пошел вон сопляк ", false  // можно убрать else   и сделать проще
	}*/
	return "Пошел вон сопляк", false // один хрен если не правильно вернет противоположное тип мелкий ты ))
}
