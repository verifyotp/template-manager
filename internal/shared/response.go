package shared

import (
	"template-manager/internal/entity"
)

type LoginResponse struct {
	Account *entity.Account `json:"account"`
	Session *entity.Session `json:"session"`
}
