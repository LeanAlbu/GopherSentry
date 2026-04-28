package processor

import "fmt"

type Job struct{
	ID int
	Line string
}

type Result struct{
	JobID int
	IsThreat bool
	Message string
}


func main(){
	age := 5
	
	if age < 18 {
		fmt.Println("teste")
	}
}
