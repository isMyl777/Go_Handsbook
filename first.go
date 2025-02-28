package main

import "fmt"

func main() {
    // 往标准输出打印一行。
    // 用包名fmt限制打印函数。
    fmt.Println("你好世界")
    fmt.Println(42)
    fmt.Printf("%d - %b \n", 42, 2)
    fmt.Printf("%d - %b -%2f \n", 42, 2,10)
    // 循环
    // for i:= 60;i<=122;i++ {
    //     fmt.Printf("%d \t %b \t %x \t %q \n",i,i,i,i)

    // }

    var xx string = "hello"
    var yy int32 = 10
    var zz float32 = 3.1415
    runes :=[]rune(xx)
    for i,s := range runes {
        fmt.Printf("%d,%c \n",i,s)
    }

    fmt.Println(xx)
    fmt.Println(yy)
    fmt.Println(zz)

    if yy>5 {
        fmt.Println("nice to meet you")
    }

    // 调用当前包的另一个函数。
    // beyondHello()
    // xunhuan()
}
func learnMultiple(x,y int) (sum,prod int){
    return x+y,x*y
}

func beyondHello() {
    fmt.Println(" beyond Hello")

    var a int
    a = 10
    y := 20
    // sum,prod    := a+y,a*y
    // fmt.Println(sum,prod)
    // learnTypes()
    sum,prod := learnMultiple(a,y)
    fmt.Println(sum,prod)
}

func learnTypes(){
    var a int = 10
    var b float64 = 3.14
    var c string = "hello"
    var d bool = true
    str := "少说话多读书"  
    var e string = str
    fmt.Println(a,b,c,str,d,e)
}
func xunhuan() {
    numbers := []int{1, 2, 3, 4, 5}

    for index, value := range numbers {
        fmt.Printf("Index: %d, Value: %d\n", index, value)
    }
}
