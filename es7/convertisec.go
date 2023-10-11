package main
import "fmt"

func main(){
	var s int
	fmt.Print("Secondi? ")
	fmt.Scan(&s)
	h,m := s/60/60,s/60
	stri := fmt.Sprintf("h:m:sec - %d:%d:%d",h,m,s)
	fmt.Println(stri)
}
