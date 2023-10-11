package main
import "math"
import "fmt"

func main(){
var e1 , e2 float64
fmt.Print("età persona 1 = ")
fmt.Scan(&e1)
fmt.Print("età persona 2 = ")
fmt.Scan(&e2)
sum := e1 + e2
media := (e1 + e2) /2
sum10 := e1 + e2 + 20
media10 := (e1 + e2 +20) /2
var mediaArrotondataDifetto float64 = math.Floor(media)
var mediaArrotondataEccesso float64 = math.Ceil(media)

fmt.Println("Somma età = ",sum, "\nMedia età = ",media,"\nMedia età arrotondata per difetto all'intero inferiore = ",mediaArrotondataDifetto)
fmt.Println("Media età arrotondata per eccesso all'intero superiore =" ,mediaArrotondataEccesso , "\nSomma età a 10 anni = ",sum10,"\nMedia età a 10 anni = ", media10)
}
