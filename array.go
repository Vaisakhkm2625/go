package main
import ("fmt")

func main() {

    var arr = [10] int { 1, 2, 3, 4, 5, 6, 7, 8, 9, 10 }
    var arr1 = [...]int{ 1, 2, 3, 4, 5, 6, 7, 8, 9, 10 }

    cars := [...]string {"volvo", "BMW", "Ford","Mazda"}

    fmt.Println(arr)
    fmt.Println(arr1)
    fmt.Println(cars)

    cars[2] = "FordDeez"

    fmt.Println(cars)

    //initialize speific elemnets
    arrsp := [5]int{1:10,2:23}
    fmt.Println(arrsp)

    //length of a array
    fmt.Println(len(arrsp))


}

