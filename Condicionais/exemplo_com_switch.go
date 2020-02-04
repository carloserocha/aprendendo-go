package main

import (
	"fmt"
	"time"
)

func main() {
	print := fmt.Println
	_, e := time.Now().ISOWeek() // underline ignora a saída do parâmetro selecionado (no caso o primeiro)

	isoWeek := e - 6 // 7-1=6 dias na semana (pois o array começa a contar do zero)

	// não se esqueça que está ao contrário
	switch isoWeek {
	case 0:
		print("Monday again? No problem")
		break
	case 1:
		print("Tuesday again? No problem")
		break
	case 2:
		print("Wednesday again? No problem")
		break
	case 3:
		print("Thursday again? No problem")
		break
	case 4:
		print("Friday again? HAPPY HOUR")
		break
	case 5:
		print("Saturday again? HAPPY SATURDAY")
		break
	case 6:
		print("Sunday again? HAPPY SUNDAY")
	}

	print("Happy day!")
}

/*
	time.Now().ISOWeek():
		Segunda: 6
		Terça: 5
		Quarta: 4
		Quinta: 3
		Sexta: 2
		Sabado: 1
		Domingo: 0
*/
