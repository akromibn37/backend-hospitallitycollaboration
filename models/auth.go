package models

type AuthRequest struct {
	UserName string `json:"username"`
	PassWord string `json:"password"`
}

type AuthResponse struct {
	UserId string `json:"user_id"`
}
