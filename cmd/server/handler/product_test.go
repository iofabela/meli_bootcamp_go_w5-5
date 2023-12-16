package handler

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/extmatperez/meli_bootcamp_go_w5-5/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w5-5/tests"
	"github.com/extmatperez/meli_bootcamp_go_w5-5/tests/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// CreateServerProduct ...
func CreateServerProduct(data []domain.Product) *gin.Engine {

	// act
	p := NewProduct(&mocks.MockRepositoryProduct{
		Data: data,
	})

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	pr := r.Group("/products")
	{
		pr.GET("/", p.GetAll())
		pr.GET("/:id", p.Get())
		pr.POST("/", p.Create())
		pr.PATCH("/:id", p.Update())
		pr.DELETE("/:id", p.Delete())
	}

	return r
}

func TestCreateOkProduct(t *testing.T) {
	resp := struct {
		Data domain.Product `json:"data"`
	}{}
	// crear el Server y definir las Rutas
	r := CreateServerProduct(mocks.MockListProducts)
	// crear Request del tipo GET y Response para obtener el resultado
	req, rr := tests.CreateRequestTest(http.MethodPost, "/products/", mocks.MockUpdateProduct)
	// indicar al servidor que pueda atender la solicitud
	r.ServeHTTP(rr, req)

	// Validation
	assert.Equal(t, 201, rr.Code)

	// Test validating
	err := json.Unmarshal(rr.Body.Bytes(), &resp)
	assert.Nil(t, err)

	// Test de data correcta en respuesta
	resp.Data.ID = 2
	assert.Equal(t, mocks.MockUpdateProduct, resp.Data)
}

func TestCreateFailProduct(t *testing.T) {
	resp := struct {
		ErrCode string `json:"code"`
		Message string `json:"message"`
	}{}

	// crear el Server y definir las Rutas
	r := CreateServerProduct(mocks.MockListProducts)
	// crear Request del tipo GET y Response para obtener el resultado
	req, rr := tests.CreateRequestTest(http.MethodPost, "/products/", "")
	// indicar al servidor que pueda atender la solicitud
	r.ServeHTTP(rr, req)

	// Validation
	assert.Equal(t, 422, rr.Code)

	// Test validating
	err := json.Unmarshal(rr.Body.Bytes(), &resp)
	assert.Nil(t, err)

	// Test de data correcta en respuesta
	assert.Equal(t, "unprocessable_entity", resp.ErrCode)
}

func TestCreateConflictProduct(t *testing.T) {

	resp := struct {
		ErrCode string `json:"code"`
		Message string `json:"message"`
	}{}

	// crear el Server y definir las Rutas
	r := CreateServerProduct(mocks.MockListProducts)
	// crear Request del tipo GET y Response para obtener el resultado
	req, rr := tests.CreateRequestTest(http.MethodPost, "/products/", mocks.MockListProducts[1])
	// indicar al servidor que pueda atender la solicitud
	r.ServeHTTP(rr, req)

	// Validation
	assert.Equal(t, 409, rr.Code)

	// Test validating
	err := json.Unmarshal(rr.Body.Bytes(), &resp)
	assert.Nil(t, err)

	// Test de data correcta en respuesta
	assert.Equal(t, "conflict", resp.ErrCode)
	assert.Equal(t, "error. the product with the code: raspberry, already exists", resp.Message)
}

func TestFindAllProducts(t *testing.T) {

	resp := struct {
		Data []domain.Product `json:"data"`
	}{}

	// crear el Server y definir las Rutas
	r := CreateServerProduct(mocks.MockListProducts)
	// crear Request del tipo GET y Response para obtener el resultado
	req, rr := tests.CreateRequestTest(http.MethodGet, "/products/", nil)
	// indicar al servidor que pueda atender la solicitud
	r.ServeHTTP(rr, req)

	// Validation
	assert.Equal(t, 200, rr.Code)
	err := json.Unmarshal(rr.Body.Bytes(), &resp)
	assert.Nil(t, err)
	assert.True(t, len(mocks.MockListProducts) > 0)
	// La data sea igual
	assert.Equal(t, mocks.MockListProducts, resp.Data)
}

func TestFindAllProductsWithEmptyList(t *testing.T) {

	resp := struct {
		Data []domain.Product `json:"data"`
	}{}

	// crear el Server y definir las Rutas
	r := CreateServerProduct(mocks.MockDataEmptyProduct)
	// crear Request del tipo GET y Response para obtener el resultado
	req, rr := tests.CreateRequestTest(http.MethodGet, "/products/", nil)
	// indicar al servidor que pueda atender la solicitud
	r.ServeHTTP(rr, req)

	// Validation
	assert.Equal(t, 200, rr.Code)
	err := json.Unmarshal(rr.Body.Bytes(), &resp)
	assert.Nil(t, err)
	assert.True(t, len(resp.Data) == 0)
}

func TestFindByIdExistentProduct(t *testing.T) {

	resp := struct {
		Data domain.Product `json:"data"`
	}{}

	// crear el Server y definir las Rutas
	r := CreateServerProduct(mocks.MockListProducts)
	// crear Request del tipo GET y Response para obtener el resultado
	req, rr := tests.CreateRequestTest(http.MethodGet, "/products/2", nil)
	// indicar al servidor que pueda atender la solicitud
	r.ServeHTTP(rr, req)

	// Validation code request
	assert.Equal(t, 200, rr.Code)

	// Test validating
	err := json.Unmarshal(rr.Body.Bytes(), &resp)
	assert.Nil(t, err)

	// La data sea igual
	assert.Equal(t, mocks.MockListProducts[1], resp.Data)
}

func TestFindByIdNonExistentProduct(t *testing.T) {

	resp := struct {
		ErrCode string `json:"code"`
		Message string `json:"message"`
	}{}

	// crear el Server y definir las Rutas
	r := CreateServerProduct(mocks.MockListProducts)
	// crear Request del tipo GET y Response para obtener el resultado
	req, rr := tests.CreateRequestTest(http.MethodGet, "/products/5", nil)
	// indicar al servidor que pueda atender la solicitud
	r.ServeHTTP(rr, req)

	// Validation code request
	assert.Equal(t, 404, rr.Code)

	// Test validating
	err := json.Unmarshal(rr.Body.Bytes(), &resp)
	assert.Nil(t, err)

	// La data sea igual
	assert.Equal(t, "not_found", resp.ErrCode)
	assert.Equal(t, "error. No product found with the entered id: 5", resp.Message)
}

func TestGetIdNonExistentProduct(t *testing.T) {
	resp := struct {
		ErrCode string `json:"code"`
		Message string `json:"message"`
	}{}

	// crear el Server y definir las Rutas
	r := CreateServerProduct(mocks.MockListProducts)
	// crear Request del tipo GET y Response para obtener el resultado
	req, rr := tests.CreateRequestTest(http.MethodGet, "/products/string", nil)
	// indicar al servidor que pueda atender la solicitud
	r.ServeHTTP(rr, req)

	// Validation code request
	assert.Equal(t, 400, rr.Code)

	// Test validating
	err := json.Unmarshal(rr.Body.Bytes(), &resp)
	assert.Nil(t, err)

	// La data sea igual
	assert.Equal(t, "bad_request", resp.ErrCode)
	assert.Equal(t, "error. The id [0] entered must be of type *integer*", resp.Message)
}

func TestUpdateOkProduct(t *testing.T) {
	resp := struct {
		Data domain.Product `json:"data"`
	}{}

	beforeUpdated := mocks.MockListProducts[1]
	updateMock := mocks.MockUpdateProduct

	// crear el Server y definir las Rutas
	r := CreateServerProduct(mocks.MockListProducts)
	// crear Request del tipo GET y Response para obtener el resultado
	req, rr := tests.CreateRequestTest(http.MethodPatch, "/products/2", updateMock)
	// indicar al servidor que pueda atender la solicitud
	r.ServeHTTP(rr, req)

	// Validation
	assert.Equal(t, 200, rr.Code)
	err := json.Unmarshal(rr.Body.Bytes(), &resp)
	assert.Nil(t, err)
	// La data sea igual
	assert.NotEqual(t, beforeUpdated, resp.Data)
	assert.Equal(t, updateMock, resp.Data)
}

func TestUpdateNonExistentProduct(t *testing.T) {
	resp := struct {
		ErrCode string `json:"code"`
		Message string `json:"message"`
	}{}

	// crear el Server y definir las Rutas
	r := CreateServerProduct(mocks.MockListProducts)
	// crear Request del tipo GET y Response para obtener el resultado
	req, rr := tests.CreateRequestTest(http.MethodPatch, "/products/5", &mocks.MockUpdateProduct)
	// indicar al servidor que pueda atender la solicitud
	r.ServeHTTP(rr, req)

	// Validation
	assert.Equal(t, 404, rr.Code)
	err := json.Unmarshal(rr.Body.Bytes(), &resp)
	assert.Nil(t, err)
	// The message and the code are the same as expected.
	assert.Equal(t, "not_found", resp.ErrCode)
	assert.Equal(t, "error. No product found with the entered id: 5", resp.Message)
}

func TestUpdateIdConflict(t *testing.T) {
	resp := struct {
		ErrCode string `json:"code"`
		Message string `json:"message"`
	}{}

	// crear el Server y definir las Rutas
	r := CreateServerProduct(mocks.MockListProducts)
	// crear Request del tipo GET y Response para obtener el resultado
	req, rr := tests.CreateRequestTest(http.MethodPatch, "/products/string", &mocks.MockUpdateProduct)
	// indicar al servidor que pueda atender la solicitud
	r.ServeHTTP(rr, req)

	// Validation
	assert.Equal(t, 400, rr.Code)
	err := json.Unmarshal(rr.Body.Bytes(), &resp)
	assert.Nil(t, err)
	// The message and the code are the same as expected.
	assert.Equal(t, "bad_request", resp.ErrCode)
	assert.Equal(t, "error. The entered id must be of type *integer*", resp.Message)
}

func TestUpdateIdFailProduct(t *testing.T) {
	resp := struct {
		ErrCode string `json:"code"`
		Message string `json:"message"`
	}{}

	// crear el Server y definir las Rutas
	r := CreateServerProduct(mocks.MockListProducts)
	// crear Request del tipo GET y Response para obtener el resultado
	req, rr := tests.CreateRequestTest(http.MethodPatch, "/products/2", &mocks.MockUpdateConflictProduct)
	// indicar al servidor que pueda atender la solicitud
	r.ServeHTTP(rr, req)

	// Validation
	assert.Equal(t, 400, rr.Code)
	err := json.Unmarshal(rr.Body.Bytes(), &resp)
	assert.Nil(t, err)
	// The message and the code are the same as expected.
	assert.Equal(t, "bad_request", resp.ErrCode, resp.Message)
}

func TestUpdateConflictProductCode(t *testing.T) {
	resp := struct {
		ErrCode string `json:"code"`
		Message string `json:"message"`
	}{}

	// crear el Server y definir las Rutas
	r := CreateServerProduct(mocks.MockListProducts)
	// crear Request del tipo GET y Response para obtener el resultado
	req, rr := tests.CreateRequestTest(http.MethodPatch, "/products/2", &mocks.MockListProducts[1])
	// indicar al servidor que pueda atender la solicitud
	r.ServeHTTP(rr, req)

	// Validation
	assert.Equal(t, 409, rr.Code)
	err := json.Unmarshal(rr.Body.Bytes(), &resp)
	assert.Nil(t, err)
	// The message and the code are the same as expected.
	assert.Equal(t, "conflict", resp.ErrCode)
	assert.Equal(t, "error. the product with code: raspberry, already exists", resp.Message)
}

func TestUpdateFail(t *testing.T) {
	resp := struct {
		ErrCode string `json:"code"`
		Message string `json:"message"`
	}{}

	// crear el Server y definir las Rutas
	r := CreateServerProduct(mocks.MockListProducts)
	// crear Request del tipo GET y Response para obtener el resultado
	req, rr := tests.CreateRequestTest(http.MethodPatch, "/products/2", "")
	// indicar al servidor que pueda atender la solicitud
	r.ServeHTTP(rr, req)

	// Validation
	assert.Equal(t, 422, rr.Code)
	err := json.Unmarshal(rr.Body.Bytes(), &resp)
	assert.Nil(t, err)
	// The message and the code are the same as expected.
	assert.Equal(t, "unprocessable_entity", resp.ErrCode)
}

func TestDeleteNonExistent(t *testing.T) {
	resp := struct {
		ErrCode string `json:"code"`
		Message string `json:"message"`
	}{}

	// crear el Server y definir las Rutas
	r := CreateServerProduct(mocks.MockListProducts)
	// crear Request del tipo GET y Response para obtener el resultado
	req, rr := tests.CreateRequestTest(http.MethodDelete, "/products/5", nil)
	// indicar al servidor que pueda atender la solicitud
	r.ServeHTTP(rr, req)

	// Validation
	assert.Equal(t, 404, rr.Code)
	err := json.Unmarshal(rr.Body.Bytes(), &resp)
	assert.Nil(t, err)
	// The message and the code are the same as expected.
	assert.Equal(t, "not_found", resp.ErrCode)
	assert.Equal(t, "error. No product found with the entered id: 5", resp.Message)
}

func TestDeleteIdNonExistent(t *testing.T) {
	resp := struct {
		ErrCode string `json:"code"`
		Message string `json:"message"`
	}{}

	// crear el Server y definir las Rutas
	r := CreateServerProduct(mocks.MockListProducts)
	// crear Request del tipo GET y Response para obtener el resultado
	req, rr := tests.CreateRequestTest(http.MethodDelete, "/products/string", nil)
	// indicar al servidor que pueda atender la solicitud
	r.ServeHTTP(rr, req)

	// Validation
	assert.Equal(t, 400, rr.Code)
	err := json.Unmarshal(rr.Body.Bytes(), &resp)
	assert.Nil(t, err)
	// The message and the code are the same as expected.
	assert.Equal(t, "bad_request", resp.ErrCode)
	assert.Equal(t, "error. The entered id must be of type *integer*", resp.Message)
}

func TestDeleteOkProduct(t *testing.T) {
	// crear el Server y definir las Rutas
	r := CreateServerProduct(mocks.MockListProducts)
	// crear Request del tipo GET y Response para obtener el resultado
	req, rr := tests.CreateRequestTest(http.MethodDelete, "/products/1", nil)
	// indicar al servidor que pueda atender la solicitud
	r.ServeHTTP(rr, req)

	// Validation code request
	assert.Equal(t, 204, rr.Code)
}
