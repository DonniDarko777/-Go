package main

import (
	"fmt" //Пакет fmt реализует форматированный ввод-вывод с функциями, аналогичными C printf и scanf.
	"reflect"
	/* reflect типичное использование заключается в том, чтобы принять значение со статическим типом interface{}
	   и извлечь информацию о его динамическом типе, вызвав TypeOf, который возвращает тип*/)

func main() {
	var massage string // строка
	massage = "переменная с указание типа данных"
	var malibu int //число
	malibu = 7
	gogo := "без типа данных строка" //без указания типа
	var nini string                  // нулевая строка в ней ничего не указано выведит пустоту
	var namber int                   // нулевое сзначение для типи данных число выведет 0
	var polozh uint                  // выводит положительные тока числа
	polozh = 100
	var plav float64 //переменная с плавающей точкой 1.67
	plav = 99.7654
	var b bool // истина или лож , булевая херня
	fmt.Println("переменная истина или лож", b)
	fmt.Println("переменная с плавающей точкой", plav)
	fmt.Println("Выводит только положительны числа например:", polozh)
	fmt.Println(namber)
	fmt.Println(nini, "нулевая строка в ней ничего не указано выведит пустоту")
	fmt.Println(gogo)
	fmt.Println(massage, "тип строка")
	fmt.Println("переменная c указания  данных:", malibu, "Тип число")
	fmt.Println(reflect.TypeOf(gogo))                 // выводит тип даных feflrct.Typeof(переменная)
	fmt.Println(reflect.TypeOf(massage))              // выводит тип даных feflrct.Typeof(переменная)
	fmt.Println(reflect.TypeOf(malibu), "тип данных") // выводит тип даных feflrct.Typeof(переменная)
}
