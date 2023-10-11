package main
import "math"
import "fmt"

func main(){
	
	var r float64
	fmt.Println("inserisci raggio :")
	fmt.Scan(&r)
	circonferenza := 2*r*math.Pi
	area := r*r*math.Pi
	fmt.Println("Circonferenza = ",circonferenza,"\nArea = ",area)
	
}
