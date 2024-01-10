package main

import "fmt"

func main(){

    // declaration
    var x int = 10
    //assignatio
    x = 1
    fmt.Println("assgined value of x:",x)
    // go wiill infer the type
    y := 1

    fmt.Println("y type infered",y)

    //types
    //bool
    t := true
    f := false

    fmt.Println("boolean types: ",f,t)

    //string
    s := "string"
    fmt.Println("string type :",s)

    // go has complex number types
    var c complex64 = 1.4+3.2i;
    fmt.Println("Complex number: ",c)

    //arithmatic operations
    fmt.Println("20+10=",20+10)
    fmt.Println("20-10=",20-10)
    fmt.Println("20*10=",20*10)
    fmt.Println("20/10=",20/10)


    fmt.Println("5/2",5/2)
    fmt.Println("5/2",float32(5)/float32(2))

    //assignment operators
    a:=1
    fmt.Println(a)
    a+=10
    fmt.Println(a)
    a-=5
    fmt.Println(a)
    a*=5
    fmt.Println(a)
    a/=5
    fmt.Println(a)


}
