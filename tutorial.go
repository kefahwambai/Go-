package main
import "fmt"

type gasEngine struct{
	mpg uint8
	gallons uint8
}

func (e gasEngine) milesLeft() uint8{
	return e.gallons * e.mpg
}

func canMakeIt(e gasEngine, miles uint8){
	if miles <=e.milesLeft(){
		fmt.Println("You can make it!")
	} else {
		fmt.Println("You're risking!")
	}

}

type engine interface{
	milesLeft() uint8
}

func main(){

	myEngine := gasEngine{mpg: 5, gallons: 5}
	canMakeIt(myEngine, 9)
}