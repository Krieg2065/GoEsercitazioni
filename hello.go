package main
import f "fmt"

func main(){
	//date altezzi in cm di due persone, calcolare la media in decimale
	var alt1,alt2 float64
	
	f.Scan(&alt1 , &alt2)
	
	media := (alt1+alt2) /2
	f.Println("altezza media : " , media)
}
