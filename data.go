package main

import (
	"database/sql"
	"fmt"
)

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

func ReadDB() {
	db := dbConn()
	defer db.Close()
	// Execute the query
	rows, err := db.Query("SELECT * FROM sampletb")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	// Get column names
	columns, err := rows.Columns()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	// Make a slice for the values
	values := make([]sql.RawBytes, len(columns))

	// rows.Scan wants '[]interface{}' as an argument, so we must copy the
	// references into such a slice
	// See http://code.google.com/p/go-wiki/wiki/InterfaceSlice for details
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	// Fetch rows
	for rows.Next() {
		// get RawBytes from data
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		// Now do something with the data.
		// Here we just print each column as a string.
		var value string
		for i, col := range values {
			// Here we can check if the value is nil (NULL value)
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			fmt.Println(columns[i], ": ", value)
		}
		fmt.Println("-----------------------------------")
	}
	if err = rows.Err(); err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
}
