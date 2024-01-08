package main

import (
    "fmt"
)

type student struct {
    Rollno int
    name string 
}

var (

    v1 = student{10,"Anirudh"}
    v2 = student{Rollno:10}
    v3 = student{}
    p = &student{11,"Asrith"}
) 
func main() {
    var a = [10]int{1,2,3,4,5,6,7,8,9,10}
    fmt.Println(a)
    b := [10]string{"I","Love","Go"}
    fmt.Println(b)
    var c [2] bool
    c[1] = true
    c[0] = false
    fmt.Println(c)
    var as[] int = a[1:4]
    var bs[] string = b[2:10]
    var cs[] bool = c[0:2]
    fmt.Println(as,bs,cs)
    as[0] = 10
    bs[0] = "Is this worth it?"
    cs[0] = true
    fmt.Println(as,bs,cs)
    fmt.Println(a,b,c)
    d := []int{1,1,1,1,1,1,1,1}
    fmt.Println(d)
    a1 := a[:]
    a2 := a[1:]
    a3 := a[4:7]
    a4 := a[0:10]
    fmt.Println(a1,a2,a3,a4)
    fmt.Println(len(a1),cap(a1))
    fmt.Println(len(a2),cap(a2))
    fmt.Println(len(a3),cap(a3))
    fmt.Println(len(a4),cap(a4))

}