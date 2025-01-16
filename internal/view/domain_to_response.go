package view

import (
	"server-go/internal/domain"
	"server-go/internal/schemas/response"
	"time"
)

func DomainToResponse(userDomain domain.UserDomainInterface) response.UserResponse {
	return response.UserResponse{
		ID:        "1",
		FirstName: userDomain.GetFirstName(),
		LastName:  userDomain.GetLastName(),
		Email:     userDomain.GetEmail(),
		Phone:     userDomain.GetPhone(),
		Role:      userDomain.GetRole(),
		IsActive:  userDomain.GetIsActive(),
		CreatedAt: time.Now().String(),
		UpdatedAt: time.Now().String(),
	}
}
