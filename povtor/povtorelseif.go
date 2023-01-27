package main

import "fmt"

/*func main() {
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

*/

func main() {
	fmt.Println(vozrost(80))
}
func vozrost(age int) (string, bool) {
	if age >= 18 && age < 25 {
		return "Входи это то что тебе нужно", true
	} else if age >= 27 && age < 40 {
		return "Может подумаешь о работе и будущем ?", true
	} else if age >= 60 {
		return "Тебе точно пора на кладбище", true
	}
	return "Пошел вон сопляк ", false
}
