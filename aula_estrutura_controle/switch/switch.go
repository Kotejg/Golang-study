package main

import "fmt"

func diaDaSemana(numero int8) string {
	switch numero {
	case 1:
		return "domingo"

	case 2:
		return "segunda"

	case 3:
		return "terca"

	case 4:
		return "quarta"

	case 5:
		return "quinta"

	case 6:
		return "sexta"

	case 7:
		return "sabado"

	default:
		return "numero invalido"
	}
}

func diaDaSemana2(numero int8) string {
	var dia string
	switch {
	case numero == 1:
		dia = "domingo"
	case numero == 2:
		dia = "segunda"

	case numero == 3:
		dia = "terca"

	case numero == 4:
		dia = "quarta"

	case numero == 5:
		dia = "quinta"

	case numero == 6:
		dia = "sexta"

	case numero == 7:
		dia = "sabado"

	default:
		dia = "numero invalido"
	}

	return dia
}

func main() {
	fmt.Println("switch")

	fmt.Println(diaDaSemana(2))
	fmt.Println("---")
	fmt.Println(diaDaSemana2(2))
}
