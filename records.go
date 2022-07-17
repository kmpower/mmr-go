package main

// create struct for all items patient, doctor, medicine, record
type Patient struct {
	FirstName string
	LastName  string
}

type Doctor struct {
	FirstName string
	LastName  string
}

type Medicine struct {
	MedicineName     string
	MedicineQuantity string
}

type Record struct {
	Patient
	Doctor
	Medicine
}
