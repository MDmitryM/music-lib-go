package musiclib

type User struct {
	Email    string `json:"email" binding:"required, email"`
	Password string `json:"password" binding:"required"`
	Name     string `json:"name" binding:"required"`
}
