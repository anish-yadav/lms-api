package handler

type StudentResponse struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	Email          string `json:"email"`
	RegistrationNo string `json:"registration_no"`
	Role           string `json:"string"`
}
