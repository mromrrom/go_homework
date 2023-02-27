package hw02unpackstring //hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")

func Main() {

}

func Unpack(str string) (string, error) {

	// что значит _ перед переменной - что мы задекларировали её но не хотим использовать

	/*
		1. берём строку
		2. разбиваем на массив рун (range)
		2.2. н = 0
		3. берём н символ    --проверяем на строку (строка или руна)
		4. если 1 символ - строка - идём дальше. если не строка - возвр. ошибку
		5. берём н+1 символ проверяем можем сконвертировать число или нет. если можем пишем в пемененную х
		5.2. если можем - делаем новую текст переменную m = "символ н повторить x раз", переменную н увеличиваем на 2
		5.3. если не можем - ничего не делаем, переменную н увеличиваем на 1
		6. повторяем цикл пока н < длина(строки)

	*/

	var m rune
	var isNum bool = true
	//var prevNum int = 0
	//strRunes := []rune(str)
	str1 := str + "b"
	var builder strings.Builder

	for _, v := range str1 {

		// условие 1. v - число, m - число или 0. Возвращаем ошибку   if amount, err := strconv.Atoi(string(v)); err == nil
		// условие 2. v - число, m - буква. Добавляем в массив rr, m, v-1 раз
		if amount, error := strconv.Atoi(string(v)); error == nil {

			//fmt.Printf("%d seems like a number! ", amount)

			if isNum == true {
				return "", ErrInvalidString
				// here will be an error
			} else {
				if amount >= 1 {
					builder.WriteString(strings.Repeat(string(m), amount))
				} else {
					//builder.WriteString(string(m))
				}

			}

			isNum = true
			//prevNum = amount
			// условие 3. v - буква, m - число или буква. Добавляем в массив rr
		} else {
			//fmt.Printf("%q seems like a letter! ", v)
			//fmt.Printf("%q seems a previous variable! ", m)
			if isNum == false {
				builder.WriteString(string(m))
			} else {
				//?
			}

			isNum = false
		}

		m = v

	}

	res := builder.String()
	return res, nil

}
