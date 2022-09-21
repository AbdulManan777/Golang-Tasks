package main

import "fmt"

type Person struct {
    Name string
    Age  int
    City string
 }


 func structArgum(p Person) string {
    return p.Name
 }
 
 func main() {
    p := Person{
       Name: "Manan",
       Age:  21,
       City: "Islamabad",
    }
 
    name := structArgum(p)

    fmt.Println(name)
 }