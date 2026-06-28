package handlers

type AuthHandler struct {
	service *services.AuthService
}

func NewAuthHandler(service *services.AuthService) *AuthHandler{
	return &AuthHandler{
		service: service,
	}
}