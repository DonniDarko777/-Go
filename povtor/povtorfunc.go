package main

import "fmt"

/*func main() {
	jojo := email("Евгений", 21)
	printjojo(jojo)
}
func printjojo(jojo string) {
	fmt.Println(jojo)
}
func email(name string, age int) string {
	result := fmt.Sprintf("Привет, %s, тебе сейчас %d , О ужас ", name, age)
	return result
}
*/
// Задача вывести цифры и буквы в одной строке
func main() {
	a := c("Жека", 41)
	b(a)
}
func b(a string) {
	fmt.Println(a)
}
func c(name string, age int) string {
	result := fmt.Sprintf("Привет %s тебе %d", name, age)
	return result
}
