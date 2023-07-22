/**
 * @author Enéas Almeida <eneas.eng@yahoo.com>
 * @description Conversão de tipos
 */

package src

import (
	"fmt"
	"strconv"
)

func ConvertStringToInt(str string) int {
	value, err := strconv.Atoi(str)

	if err != nil {
		panic(err)
	}

	return value
}

func ConvertIntToString(value int) string {
	return strconv.Itoa(value)
}

func ConvertFloatToString(value float64) string {
	str := strconv.FormatFloat(value, 'f', 2, 64)

	fmt.Println(str)

	return str
}

func ConvertStringToFloat(str string) float64 {
	value, err := strconv.ParseFloat(str, 64)

	if err != nil {
		panic(err)
	}

	return value
}
