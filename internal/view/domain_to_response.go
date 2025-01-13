package view

import (
	"server-go/internal/domain"
	"server-go/internal/schemas/response"
	"time"
)

func DomainToResponse(ud domain.UserDomainInterface) response.UserResponse {
	return response.UserResponse{
		ID:        "1",
		FirstName: ud.GetFirstName(),
		LastName:  ud.GetLastName(),
		Email:     ud.GetEmail(),
		Phone:     ud.GetPhone(),
		Role:      ud.GetRole(),
		IsActive:  ud.GetIsActive(),
		CreatedAt: time.Now().String(),
		UpdatedAt: time.Now().String(),
	}
}
