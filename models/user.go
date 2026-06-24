package models

import "gorm.io/gorm"
//gorm.Model this is called a embedding
//its automatically adds:
// type Model struct {
// 	ID        uint
// 	CreatedAt time.Time
// 	UpdatedAt time.Time
// 	DeletedAt DeletedAt
// }

//the struct becomes a 
// type User struct {
// 	ID        uint
// 	CreatedAt time.Time
// 	UpdatedAt time.Time
// 	DeletedAt gorm.DeletedAt

// 	Name string
// 	Email string
// 	Password string
// }

type User struct { //for this struct stores in the db for the database perpose
	gorm.Model
	Name	string	`json:"name"`
	Email	string	`json:"email" gorm:"unique;not null"`
	Password	string	`json:"-"` //never send password bak in json
}

type SignupInput struct {
	Name	string `json:"name" binding:"required"`
	Email	 string `json:"email" binding:"required,email"`
	Passowrd string `json:"password" binding:"required,min=6"`
}

type LoginInput struct {
	Email	string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}