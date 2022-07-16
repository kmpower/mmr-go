package main

import "fmt"

func main() {

	mmr := Mmr2()

	fmt.Println(mmr.PatientName())
	fmt.Println(mmr.Mmr1())
	fmt.Println(mmr)
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

func (r *Record) Mmr1() (string, string, string, string, string, string) {
	return r.Patient.FirstName,
		r.Patient.LastName,
		r.Docter.FirstName,
		r.Docter.LastName,
		r.Medicine.MedicineName,
		r.Medicine.MedicineQuantity
}

func (r *Record) PatientName() (string, string) {
	return r.Patient.FirstName,
		r.Patient.LastName
}

func Mmr2() Record {
	return Record{
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
}
