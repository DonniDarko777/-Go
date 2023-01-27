package main

import "fmt"

/*func main() {
	fmt.Println(enter(20))
}
func enter(age int) (string, bool) {
	if age >= 18 {
		return "Входи мужик", true
	}
	return "Пошел вон сопляк", false
}
*/
func main() {
	fmt.Println(enter(25))

}
func enter(age int) (string, bool) {
	if age >= 20 {
		return "Входи", true
	}
	return "Выходи ", false

}
