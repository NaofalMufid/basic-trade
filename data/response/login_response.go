package response

import "github.com/google/uuid"

type LoginResponse struct {
	ID    int       `json:"id"`
	UUID  uuid.UUID `json:"uuid"`
	Name  string    `json:"name"`
	Email string    `json:"email"`
	Token string    `json:"token"`
}
