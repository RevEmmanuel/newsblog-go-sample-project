package models

type UpdateUserDTO struct {
	FirstName   string       `json:"firstname" validate:"omitempty,max=30" `
	LastName    string       `json:"lastname" validate:"omitempty,max=30"`
	Password    string       `json:"password" validate:"omitempty,min=5"`
	PhoneNumber uint         `json:"phonenumber" validate:"omitempty"`
	Preferences []Preference `json:"preferences" validate:"omitempty,validPreferences"`
}
