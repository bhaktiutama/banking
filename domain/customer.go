package domain

import (
	"github.com/bhaktiutama/banking/dto"
	"github.com/bhaktiutama/banking/errs"
)

// MODELS - ENTITIES

// domain object
type Customer struct {
	Id          string `db:"customer_id"`
	Name        string
	City        string
	ZipCode     string
	DateofBirth string `db:"date_of_birth"`
	Status      string
}

func (c Customer) statusAsText() string {
	statusAsText := "active"
	if c.Status == "0" {
		statusAsText = "inactive"
	}
	return statusAsText
}

func (c Customer) ToDto() dto.CustomerResponse {

	return dto.CustomerResponse{
		Id:          c.Id,
		Name:        c.Name,
		City:        c.City,
		ZipCode:     c.ZipCode,
		DateofBirth: c.DateofBirth,
		Status:      c.statusAsText(),
	}
}

// secondary port
type CustomerRepository interface {
	// status == 1 status == 0 status = ""
	FindAll(status string) ([]Customer, *errs.AppError)
	ById(string) (*Customer, *errs.AppError)
}
