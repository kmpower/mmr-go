package main

func PatientData() Patient {
	return Patient{
		FirstName: "Mike",
		LastName:  "Hull",
	}
}

func DoctorData() Doctor {
	return Doctor{
		FirstName: "Steven",
		LastName:  "Berg",
	}
}

func MedicineData() Medicine {
	return Medicine{
		MedicineName:     "Tylenol",
		MedicineQuantity: "300mg As Needed",
	}
}

func Records() Record {
	return Record{
		Patient:  PatientData(),
		Doctor:   DoctorData(),
		Medicine: MedicineData(),
	}
}
