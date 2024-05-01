package schemas

import "time"

type CreateUserData struct {
	Email     string    `json:"email" binding:"required"`
	Password  string    `json:"password" binding:"required"`
	Name      string    `json:"name" binding:"required"`
	BirthDate time.Time `json:"birthDate"`
}

// TODO: Make birth date required again, fix json problem
