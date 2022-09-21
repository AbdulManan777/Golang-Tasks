package main

import "fmt"


type employee struct {
	name string
	salary int
	position string
	
	}


	type company struct {
		companyName string
		employees []employee
		
		}

func main() {
			
		emp1:= employee{
				name: "Hamza",
				salary: 8000,
				position: "Front-end developer",

			}


			emp2:= employee{
				name: "Saad",
				salary: 9000,
				position: "Back-end developer",

			}

			emp3:= employee{
				name: "Ali",
				salary: 12000,
				position: "Full-Stack developer",


			}


			emplo := []employee{emp1, emp2, emp3}

			company:= company{"Tetra", emplo}


			fmt.Printf("Company Details are", company)




		}