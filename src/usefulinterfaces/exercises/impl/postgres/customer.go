package postgres

import (
	"usefulinterfaces/exercises/impl"
)

type CustomerService struct {
	ID   int
	Name string
}

func (c *CustomerService) Get(id int) (*impl.Customer, error) {
	panic("not implemented")
}
