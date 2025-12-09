package main

import "math"
import "fmt"

func checkPrime(num int) bool {
	return checkPrimeHelper(num,2);
}

func checkPrimeHelper(num, i int) bool{
    if(num == 1){
        return false
    }
    
    if(i > int(math.Sqrt(float64(num)))){
      return true
    }

    if(num%i == 0){
        return false
    }

    return checkPrimeHelper(num,i+1)
}

func main(){
	fmt.Println(checkPrime(1))
}