package main

import "fmt"

func main(){
	var b, h float64 
	fmt.Println("Inserisci la base:")
	fmt.Scan(&b)
	fmt.Println("Inserisci l'altezza:")
	fmt.Scan(&h)
	area := b * h
	perimetro := (b + h ) * 2
	fmt.Println("Perimetro =" , perimetro,"\nArea = " ,area)
}
