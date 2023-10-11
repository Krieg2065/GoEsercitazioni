package main
import "fmt"

func main(){
	var s int
	fmt.Print("Secondi? ")
	fmt.Scan(&s)
	h := s/60/60
	m :=s/60 - h*60
	sec := s - h*60*60 - m * 60
	
	stri := fmt.Sprintf("h:m:sec - %d:%d:%d",h,m,sec)
	fmt.Println(stri)
}
