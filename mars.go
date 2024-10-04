package main

import (
	"fmt"
	"log"
)

type Planeta struct {
	nombre   string
	gravity float64
}

func getPeso() float64 {
	var peso float64

	fmt.Print("Ingresa tu peso en la Tierra (kg): ")
	_, err := fmt.Scanln(&peso)

	// Controlar los datos
	if err != nil {
		log.Fatalf("Error al ingresar el peso: %v", err)
	}

	return peso
}

func calculatePeso(peso float64, gravity float64) float64 {
	return peso * gravity
}

func getResults(peso float64, planetas []Planeta)  {

	for _, planet := range planetas {
		pesoPlanet := calculatePeso(peso, planet.gravity)
		fmt.Printf("Tu peso en %s sería: %.2f kg\n", planet.nombre, pesoPlanet)
	}
}

func main()  {
	
	fmt.Println("Vamos a calcular cual sería tu peso en distintos planetas!")	

	pesoTierra := getPeso()
	
	planetas := []Planeta{
		{"Marte", 0.38},
		{"Júpiter", 2.34},
		{"Venus", 0.91},
	}

	getResults(pesoTierra, planetas)
}