package main
import "fmt"

func main() {
    var age int
    fmt.Scanln("Enter the age",&age)

    mars := mars_age(age)
    fmt.Println(mars)
}

func mars_age(a int) int {
    return a*365/687
} 