package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Appointment struct {
	gorm.Model
	DoctorName string
	PatientName string
	Especiality string
	Complaint string
	Date time.Time
}

type AppoitmentResponse struct {
	ID uint `json:"id"`
	DoctorName string `json:"doctor_name"`
	PatientName string `json:"patient_name"`
	Especiality string `json:"especiality"`
	Complaint string `json:"complaint"`
	Date time.Time `json:"date"`
}
