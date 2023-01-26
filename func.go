package main

import "fmt"

/*func main() {
	print()
	print()
	print()
}
func print() {
	fmt.Println("go")
} тут 3 раза выведиться go на экран */
func main() {
	massage := sayhello("Евгений", 21)
	printmassage(massage)
	printmassage(massage)
	//printmassage("вызов 1")
	//printmassage("вызов 2")
	//printmassage("вызов 3")

}

//аргумент функции
func printmassage(massage string) {
	fmt.Println(massage)
}

func sayhello(name string, age int) string {
	result := fmt.Sprintf("Привет, %s! тебе %d лет ", name, age) // передача строки %s передача цифрв %d
	//return "Привет , " + name + "!!!" + " " + "Тебе" + age + "лет" //return возвращает , полкчилась Конкатенация (объединение ) но число не выведет
	return result

}
