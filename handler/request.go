package handler

import "fmt"

type CreateAppointmentRequest struct {
	DoctorName string `json:"doctor_name"`
	PatientName string `json:"patient_name"`
	Especiality string `json:"especiality"`
	Complaint string `json:"complaint"`
	Date string `json:"date"`
}

func errParamsIsRequired(name, typ string) error {
	return fmt.Errorf("Param: %s (type: %s) is required", name, typ)
}

func (r *CreateAppointmentRequest) Validate() error {
	if r.DoctorName == "" {
		return errParamsIsRequired("doctor_id", "string")
	}
	if r.PatientName == "" {
		return errParamsIsRequired("patient_id", "string")
	}
	if r.Date == "" {
		return errParamsIsRequired("date", "string")
	}

	return nil
}

type UpdateOpeningRequest struct {
	DoctorID string `json:"doctor_id"`
	Date     string `json:"date"`
}