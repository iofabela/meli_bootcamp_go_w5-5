package mocks

import (
	"context"
	"errors"
	"fmt"

	"github.com/iofabela/meli_bootcamp_go_w5-5/internal/domain"
)

type MockBuyerRepository struct {
	DataMock []domain.Buyer
}

var MockDataBuyers []domain.Buyer = []domain.Buyer{
	{ID: 1, CardNumberID: "ABC1234", FirstName: "JUAN", LastName: "GONZALEZ"},
	{ID: 2, CardNumberID: "XYZ1234", FirstName: "MIA", LastName: "RODRIGUEZ"},
}

func (m *MockBuyerRepository) GetAll(ctx context.Context) ([]domain.Buyer, error) {
	return m.DataMock, nil
}

func (m *MockBuyerRepository) Get(ctx context.Context, id int) (domain.Buyer, error) {

	var buyerObtained domain.Buyer
	for i, buyer := range m.DataMock {
		if id == buyer.ID {
			buyerObtained = m.DataMock[i]
			return buyerObtained, nil
		}
	}
	err := fmt.Sprintf("error: buyer with id:%v not found", id)
	return domain.Buyer{}, errors.New(err)
}

func (m *MockBuyerRepository) Exists(ctx context.Context, cardNumberID string) bool {

	for _, buyer := range m.DataMock {
		if cardNumberID == buyer.CardNumberID {
			return true
		}
	}
	return false
}

func (m *MockBuyerRepository) Save(ctx context.Context, b domain.Buyer) (int, error) {

	if b.CardNumberID == "" || b.FirstName == "" || b.LastName == "" {
		return 0, errors.New("error: JSON keys required are not included")
	}

	if m.Exists(ctx, b.CardNumberID) {
		return 0, errors.New("error: buyer with this card_number_id already exist")
	}

	b.ID = m.DataMock[len(m.DataMock)-1].ID + 1
	m.DataMock = append(m.DataMock, b)

	return b.ID, nil
}

func (m *MockBuyerRepository) Update(ctx context.Context, b domain.Buyer) error {

	_, err := m.Get(ctx, b.ID)
	if err != nil {
		return nil
	}

	if b.FirstName == "" && b.LastName == "" {
		return errors.New("error: both keys {first_name, last_name} are empty. At least one of them need be included.")
	}

	for i, buyer := range m.DataMock {
		if buyer.ID == b.ID {
			m.DataMock[i] = b
		}
	}

	return nil
}

func (m *MockBuyerRepository) Delete(ctx context.Context, id int) error {

	_, err := m.Get(ctx, id)
	if err != nil {
		return nil
	}

	var index int
	for i, buyer := range m.DataMock {
		if buyer.ID == id {
			index = i
		}
	}

	m.DataMock = append(m.DataMock[:index], m.DataMock[index+1:]...)

	return nil
}

func (m *MockBuyerRepository) GetPurchaseOrders(ctx context.Context, id int) ([]domain.BuyerOrders, error) {
	return []domain.BuyerOrders{}, nil
}
