package main

import "fmt"

// Create functions to view data for patient, doctor, medicine, record

func (r *Patient) PatientDetail() string {
	s := fmt.Sprint(r.FirstName + " " + r.LastName)
	return s
}

func (r *Doctor) DoctorDetail() string {
	s := fmt.Sprint(r.FirstName + " " + r.LastName)
	return s
}

func (r *Medicine) MedicineDetail() string {
	s := fmt.Sprint(r.MedicineName + " " + r.MedicineQuantity)
	return s
}

// Unknown function
func NewPatient() *Patient {
	p := new(Patient)
	return p
}
