package response

import "github.com/google/uuid"

type AdminResponse struct {
	UUID  uuid.UUID `json:"uuid"`
	Name  string    `json:"name"`
	Email string    `json:"email"`
}
