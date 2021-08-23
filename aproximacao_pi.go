package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Digite o número de iterações para a aproximação de pi:")
	var numero float64
	fmt.Scanln(&numero)
	if numero == 0 {
		aprox_pi_inf()
	} else {
		fmt.Printf("Aproximação com %.0f iterações -> %.9f", numero, aprox_pi(numero))
	}
}

func aprox_pi(n float64) float64 {
	var i, retorno float64
	retorno = 0
	for i = 0; i < n; i++ {
		retorno += math.Pow(-1, i) / ((2 * i) + 1)
	}
	return retorno * 4
}

func aprox_pi_inf() float64 {
	fmt.Println("Aproximação infinita! Aperte Ctrl + C para parar!")
	var i float64
	var retorno float64
	retorno = 0
	for i = 0; ; i++ {
		retorno += math.Pow(-1, i) / ((2 * i) + 1)
		if int(i)%1000000 == 0 {
			fmt.Printf("%.20f\n", retorno*4)
		}
	}
}
