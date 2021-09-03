package student

type Student struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	Email          string `json:"email"`
	RegistrationNo string `json:"registration_no"`
	password       string `json:"password"`
	Role           string `json:"string"`
}

func NewStudent() *Student {
	return &Student{}
}
