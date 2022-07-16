package main

import "fmt"

func main() {
	record := Record{}
	fmt.Println(record)
}

type Patient struct {
	FirstName string
	LastName  string
}

type Docter struct {
	FirstName string
	LastName  string
}

type Medicine struct {
	MedicineName     string
	MedicineQuantity string
}

type Record struct {
	Patient
	Docter
	Medicine
}
