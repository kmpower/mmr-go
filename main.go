package main

import (
	"fmt"
	"log"
)

func init() {
	log.SetPrefix("Logs: ")
}

func main() {
	ReadDB()
	fmt.Println(PatientData())
	fmt.Println(DoctorData())
	fmt.Println(MedicineData())
	fmt.Println(Records())
}
