package main

import "fmt"

func main() {
	fmt.Println(enter(60))
}
func enter(age int) (string, bool) {
	if age >= 18 && age < 45 {
		return "Входи мужик", true
	} else if age >= 45 && age < 65 {
		return "вам завтра на работу точно нужно сюда ? ", true
	} else if age >= 65 {
		return "Ты слишком стары для клуба ", true
	}

	return "Пошел вон сопляк", false // такой метод являеться более правильным в данной ситуации
}
