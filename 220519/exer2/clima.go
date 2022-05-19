package exer2

import "fmt"

var (
	temperatura int8    = 23
	umidade     int8    = 80
	atm         float32 = 1.0033
)

func Clima() {
	fmt.Printf("Hoje teremos %v graus, "+
		"com umidade relativa do ar em %v porcento e "+
		"%vkgf/cm² de pressão atmosférica",
		temperatura, umidade, atm)
}
