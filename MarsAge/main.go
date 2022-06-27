package main
import "fmt"

func main() {
    var age int
    fmt.Scanln("Enter the age",&age)

    mars := mars_age(age)
    fmt.Println(mars)
}

func mars_age(a int) int { //function with return type and arguments
    return a*365/687
} 