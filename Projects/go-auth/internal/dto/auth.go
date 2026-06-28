package dto

type RegisterRequest struct {
	Name  		string 		`json:"name"`
	Email 		string		`json:"email"`
	Password 	string 		`json:"password"`
}	

type LoginRequest	struct {
	Email 		string	`json:"name"`
	Password 	string  `json:"password"`
}

type UserResponse struct {
	ID 		int64 	`json:"id"`
	Name 	string 	`json:"name"`
	Email 	string 	`json:"email"`
}

type AuthResponse struct {
	User 			UserResponse 	`json:"user"`
	AccessToken 	string 			`json:"accessToken"`
	ExpiresIn		int				`json:"expiresIn"`
}

