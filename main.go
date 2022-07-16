package main

import "fmt"

func main() {
	record := Record{
		Patient{
			FirstName: "Mike",
			LastName:  "Hull",
		},
		Docter{
			FirstName: "Steven",
			LastName:  "Berg",
		},
		Medicine{
			MedicineName:     "Tylenol",
			MedicineQuantity: "300mg As Needed",
		},
	}
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
