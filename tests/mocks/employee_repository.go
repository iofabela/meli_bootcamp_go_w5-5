package mocks

import (
	"context"
	"errors"

	"github.com/iofabela/meli_bootcamp_go_w5-5/internal/domain"
)

type MockEmployeeRepository struct {
	MockData []domain.Employee
}

var MockNewEmployee domain.Employee = domain.Employee{
	ID:           4,
	CardNumberID: "0101",
	FirstName:    "Gabriela",
	LastName:     "Bejarano",
	WarehouseID:  2,
}

var MockIncompleteEmployee domain.Employee = domain.Employee{
	CardNumberID: "0111",
	FirstName:    "Laura",
	LastName:     "Toro",
	WarehouseID:  4,
}

var MockUpdateEmployee domain.Employee = domain.Employee{
	ID:           2,
	CardNumberID: "0103",
	FirstName:    "Eliza",
	LastName:     "Candelo",
	WarehouseID:  6,
}

var MockEmployeeNonExistent domain.Employee = domain.Employee{
	ID:           9,
	CardNumberID: "0110",
	FirstName:    "Ana",
	LastName:     "Hurtado",
	WarehouseID:  6,
}

var MockEmptyEmployees []domain.Employee = []domain.Employee{}

var MockEmployees []domain.Employee = []domain.Employee{
	{
		ID:           1,
		CardNumberID: "0102",
		FirstName:    "Maria",
		LastName:     "Satizabal",
		WarehouseID:  7,
	},
	{
		ID:           2,
		CardNumberID: "0103",
		FirstName:    "Juan",
		LastName:     "Pedroza",
		WarehouseID:  9,
	},
	{
		ID:           3,
		CardNumberID: "0104",
		FirstName:    "Victoria",
		LastName:     "Segura",
		WarehouseID:  4,
	},
}

func (r *MockEmployeeRepository) GetAll(ctx context.Context) ([]domain.Employee, error) {
	return r.MockData, nil
}

func (r *MockEmployeeRepository) Get(ctx context.Context, id int) (domain.Employee, error) {
	for _, empTest := range r.MockData {
		if empTest.ID == id {
			return empTest, nil
		}
	}
	return domain.Employee{}, errors.New("El id no existe")
}

func (r *MockEmployeeRepository) Exists(ctx context.Context, cardNumberID string) bool {
	for _, empTest := range r.MockData {
		if empTest.CardNumberID == cardNumberID {
			return true
		}
	}
	return false
}

func (r *MockEmployeeRepository) Save(ctx context.Context, e domain.Employee) (int, error) {
	if e.CardNumberID == "" {
		return 0, errors.New("El CardNumberID es requerido")
	}
	if r.Exists(ctx, e.CardNumberID) {
		return 0, errors.New("El empleado ya existe")
	}
	e.ID = r.MockData[len(r.MockData)-1].ID + 1
	r.MockData = append(r.MockData, e)
	return e.ID, nil
}

func (r *MockEmployeeRepository) Update(ctx context.Context, e domain.Employee) error {
	for i, empTest := range r.MockData {
		if empTest.ID == e.ID {
			if e.CardNumberID != empTest.CardNumberID && r.Exists(ctx, e.CardNumberID) {
				return errors.New("El empleado ya existe")
			}
			r.MockData[i] = e
			return nil
		}
	}
	return errors.New("El id no existe")
}

func (r *MockEmployeeRepository) Delete(ctx context.Context, id int) error {
	for i, empTest := range r.MockData {
		if empTest.ID == id {
			r.MockData = append(r.MockData[:i], r.MockData[i+1:]...)
			return nil
		}
	}
	return errors.New("El id no existe")
}

func (r *MockEmployeeRepository) GetInboundOrders(ctx context.Context, id int) ([]domain.EmployeeOrders, error) {
	return nil, nil
}
