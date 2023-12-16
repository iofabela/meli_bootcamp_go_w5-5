package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/iofabela/meli_bootcamp_go_w5-5/internal/domain"
	"github.com/iofabela/meli_bootcamp_go_w5-5/tests"
	"github.com/iofabela/meli_bootcamp_go_w5-5/tests/mocks"
	"github.com/stretchr/testify/assert"
)

func createBuyerServer(mockData []domain.Buyer) *gin.Engine {

	gin.SetMode(gin.ReleaseMode)

	b := NewBuyer(&mocks.MockBuyerService{
		DataMock: mockData,
	})

	r := gin.Default()

	br := r.Group("/buyers")
	br.GET("/", b.GetAll())
	br.GET("/:id", b.Get())
	br.POST("/", b.Create())
	br.PATCH("/:id", b.Update())
	br.DELETE("/:id", b.Delete())

	return r
}

func TestCreateOkBuyer(t *testing.T) {
	// arrange
	r := createBuyerServer(mocks.MockDataBuyers)

	objRes := struct {
		Data domain.Buyer `json:"data"`
	}{}

	newBuyer := domain.Buyer{
		CardNumberID: "KDFS1234",
		FirstName:    "ALONSO",
		LastName:     "RODRIGUEZ",
	}

	req, rr := tests.CreateRequestTest(http.MethodPost, "/buyers/", newBuyer)

	// act
	r.ServeHTTP(rr, req)

	// assert

	//Verificación código
	assert.Equal(t, 201, rr.Code)

	//Verificación cuerpo de respuesta válido
	err := json.Unmarshal(rr.Body.Bytes(), &objRes)
	assert.Nil(t, err)

	// Verificación contenido
	newBuyer.ID = 3
	assert.Equal(t, newBuyer, objRes.Data)
}

func TestCreateFailBuyer(t *testing.T) {
	// arrange
	r := createBuyerServer(mocks.MockDataBuyers)

	objRes := struct {
		Code string `json:"code"`
		Msg  string `json:"message"`
	}{}

	newBuyer := domain.Buyer{}

	req, rr := tests.CreateRequestTest(http.MethodPost, "/buyers/", newBuyer)

	// act
	r.ServeHTTP(rr, req)

	// assert

	// Verificación código
	assert.Equal(t, 422, rr.Code)

	//Verificación cuerpo de respuesta válido
	err := json.Unmarshal(rr.Body.Bytes(), &objRes)
	assert.Nil(t, err)

	// Verificación contenido
	msg := "error: JSON keys required are not included."
	assert.ErrorContains(t, errors.New(msg), objRes.Msg)

}

func TestCreateConflictBuyer(t *testing.T) {
	// arrange
	r := createBuyerServer(mocks.MockDataBuyers)

	objRes := struct {
		Code string `json:"code"`
		Msg  string `json:"message"`
	}{}

	newBuyer := domain.Buyer{
		CardNumberID: "ABC1234",
		FirstName:    "ALONSO",
		LastName:     "RODRIGUEZ",
	}

	req, rr := tests.CreateRequestTest(http.MethodPost, "/buyers/", newBuyer)

	// act
	r.ServeHTTP(rr, req)

	// assert

	// Verificación código
	assert.Equal(t, 409, rr.Code)

	//Verificación cuerpo de respuesta válido
	err := json.Unmarshal(rr.Body.Bytes(), &objRes)
	assert.Nil(t, err)

	// Verificación contenido
	msg := fmt.Sprintf("error: buyer with card_number_id:%s already exist", newBuyer.CardNumberID)
	assert.ErrorContains(t, errors.New(msg), objRes.Msg)
}

func TestFindAllBuyer(t *testing.T) {
	// arrange
	r := createBuyerServer(mocks.MockDataBuyers)

	objRes := struct {
		Data []domain.Buyer `json:"data"`
	}{}

	req, rr := tests.CreateRequestTest(http.MethodGet, "/buyers/", nil)

	// act
	r.ServeHTTP(rr, req)

	// assert

	// Verificación código
	assert.Equal(t, 200, rr.Code)

	// Verificación contenido válido
	err := json.Unmarshal(rr.Body.Bytes(), &objRes)
	assert.Nil(t, err)

	// Verificación contenido
	assert.Equal(t, mocks.MockDataBuyers, objRes.Data)
}

func TestFindByIdNonExistentBuyer(t *testing.T) {
	// arrange
	r := createBuyerServer(mocks.MockDataBuyers)

	objRes := struct {
		Code string `json:"code"`
		Msg  string `json:"message"`
	}{}

	req, rr := tests.CreateRequestTest(http.MethodGet, "/buyers/3", nil)

	// act
	r.ServeHTTP(rr, req)

	// assert

	// Verificación código
	assert.Equal(t, 404, rr.Code)

	// Verificación contenido válido
	err := json.Unmarshal(rr.Body.Bytes(), &objRes)
	assert.Nil(t, err)

	// Verificación contenido
	msg := "error: buyer with id:3 not found"
	assert.ErrorContains(t, errors.New(msg), objRes.Msg)
}

func TestFindByIdExistentBuyer(t *testing.T) {
	// arrange
	r := createBuyerServer(mocks.MockDataBuyers)

	objRes := struct {
		Data domain.Buyer `json:"data"`
	}{}

	req, rr := tests.CreateRequestTest(http.MethodGet, "/buyers/1", nil)

	// act
	r.ServeHTTP(rr, req)

	// assert

	// Verificación código
	assert.Equal(t, 200, rr.Code)

	// Verificación contenido válido
	err := json.Unmarshal(rr.Body.Bytes(), &objRes)
	assert.Nil(t, err)

	// Verificación contenido
	assert.Equal(t, mocks.MockDataBuyers[0], objRes.Data)
}

func TestUpdateOkBuyer(t *testing.T) {
	// arrange
	r := createBuyerServer(mocks.MockDataBuyers)

	objRes := struct {
		Data domain.Buyer `json:"data"`
	}{}

	buyerToUpdate := domain.Buyer{
		ID:           1,
		CardNumberID: "ABC1234",
		FirstName:    "ALFREDO",
		LastName:     "LOPEZ",
	}

	req, rr := tests.CreateRequestTest(http.MethodPatch, "/buyers/1", buyerToUpdate)

	// act
	r.ServeHTTP(rr, req)

	// assert

	// Verificación código
	assert.Equal(t, 200, rr.Code)

	// Verificación contenido válido
	err := json.Unmarshal(rr.Body.Bytes(), &objRes)
	assert.Nil(t, err)

	// Verificación contenido
	assert.Equal(t, buyerToUpdate, objRes.Data)
}

func TestUpdateNonExistentBuyer(t *testing.T) {
	// arrange
	r := createBuyerServer(mocks.MockDataBuyers)

	objRes := struct {
		Code string `json:"code"`
		Msg  string `json:"message"`
	}{}

	buyerToUpdate := domain.Buyer{
		FirstName: "ALFREDO",
		LastName:  "LOPEZ",
	}

	req, rr := tests.CreateRequestTest(http.MethodPatch, "/buyers/3", buyerToUpdate)

	// act
	r.ServeHTTP(rr, req)

	// assert

	// Verificación código
	assert.Equal(t, 404, rr.Code)

	// Verificación contenido válido
	err := json.Unmarshal(rr.Body.Bytes(), &objRes)
	assert.Nil(t, err)

	// Verificación contenido
	msg := "error: buyer with id:3 not found"
	assert.ErrorContains(t, errors.New(msg), objRes.Msg)
}

func TestDeleteNonExistentBuyer(t *testing.T) {
	// arrange
	r := createBuyerServer(mocks.MockDataBuyers)

	objRes := struct {
		Code string `json:"code"`
		Msg  string `json:"message"`
	}{}

	req, rr := tests.CreateRequestTest(http.MethodDelete, "/buyers/3", nil)

	// act
	r.ServeHTTP(rr, req)

	// assert

	// Verificación código
	assert.Equal(t, 404, rr.Code)

	// Verificación contenido válido
	err := json.Unmarshal(rr.Body.Bytes(), &objRes)
	assert.Nil(t, err)

	// Verificación contenido
	msg := "error: buyer with id:3 not found"
	assert.ErrorContains(t, errors.New(msg), objRes.Msg)
}

func TestDeleteOkBuyer(t *testing.T) {
	// arrange
	r := createBuyerServer(mocks.MockDataBuyers)

	req, rr := tests.CreateRequestTest(http.MethodDelete, "/buyers/1", nil)

	// act
	r.ServeHTTP(rr, req)

	// assert

	// Verificación código
	assert.Equal(t, 204, rr.Code)

	// Verificación contenido
	assert.Equal(t, "", rr.Body.String())

}

func TestUpdateKeysRequiredBuyer(t *testing.T) {
	// arrange
	r := createBuyerServer(mocks.MockDataBuyers)

	objRes := struct {
		Code string `json:"code"`
		Msg  string `json:"message"`
	}{}

	buyerToUpdate := domain.Buyer{}

	req, rr := tests.CreateRequestTest(http.MethodPatch, "/buyers/2", buyerToUpdate)

	// act
	r.ServeHTTP(rr, req)

	// assert

	// Verificación código
	assert.Equal(t, 400, rr.Code)

	// Verificación contenido válido
	err := json.Unmarshal(rr.Body.Bytes(), &objRes)
	assert.Nil(t, err)

	// Verificación contenido
	msg := "error: both keys {first_name, last_name} are empty. At least one of them need be included."
	assert.ErrorContains(t, errors.New(msg), objRes.Msg)
}

func TestUpdateKeysedBuyer(t *testing.T) {
	// arrange
	r := createBuyerServer(mocks.MockDataBuyers)

	req, rr := tests.CreateRequestTest(http.MethodPatch, "/buyers/2", "")

	// act
	r.ServeHTTP(rr, req)

	// assert

	// Verificación código
	assert.Equal(t, 422, rr.Code)

}

func TestUpdateIdNotCorrectBuyer(t *testing.T) {
	// arrange
	r := createBuyerServer(mocks.MockDataBuyers)

	req, rr := tests.CreateRequestTest(http.MethodPatch, "/buyers/error", "")

	// act
	r.ServeHTTP(rr, req)

	// assert

	// Verificación código
	assert.Equal(t, 400, rr.Code)

}

func TestGetIdNotCorrectBuyer(t *testing.T) {
	// arrange
	r := createBuyerServer(mocks.MockDataBuyers)

	req, rr := tests.CreateRequestTest(http.MethodGet, "/buyers/error", "")

	// act
	r.ServeHTTP(rr, req)

	// assert

	// Verificación código
	assert.Equal(t, 400, rr.Code)

}

func TestDeleteIdNotCorrectBuyer(t *testing.T) {
	// arrange
	r := createBuyerServer(mocks.MockDataBuyers)

	req, rr := tests.CreateRequestTest(http.MethodDelete, "/buyers/error", "")

	// act
	r.ServeHTTP(rr, req)

	// assert

	// Verificación código
	assert.Equal(t, 400, rr.Code)

}

func TestFindAllDbEmptyBuyer(t *testing.T) {
	// arrange
	mockEmpty := []domain.Buyer{}
	r := createBuyerServer(mockEmpty)

	objRes := struct {
		Data []domain.Buyer `json:"data"`
	}{}

	req, rr := tests.CreateRequestTest(http.MethodGet, "/buyers/", nil)

	// act
	r.ServeHTTP(rr, req)

	// assert

	// Verificación código
	assert.Equal(t, 200, rr.Code)

	// Verificación contenido válido
	err := json.Unmarshal(rr.Body.Bytes(), &objRes)
	assert.Nil(t, err)

	// Verificación contenido
	assert.Equal(t, mockEmpty, objRes.Data)
}
