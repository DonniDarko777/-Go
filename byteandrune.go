package main

import "fmt"

func main() {
	var a byte = 233      // байт это херня выводимая из таблицы аскаii типа символы разные
	fmt.Printf("%c\n", a) //обязательно выводить инфу так для типа byte
	var b rune = 'f'      // руна  это херня преобразующая симболы в биты
	fmt.Println(b)
}
