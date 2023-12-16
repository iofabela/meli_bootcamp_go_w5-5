package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/extmatperez/meli_bootcamp_go_w5-5/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w5-5/internal/warehouse"
	"github.com/extmatperez/meli_bootcamp_go_w5-5/tests"
	"github.com/extmatperez/meli_bootcamp_go_w5-5/tests/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func createWarehouseServer(warehouseService warehouse.Service) *gin.Engine {
	r := gin.New()
	gin.SetMode(gin.ReleaseMode)

	handler := NewWarehouse(warehouseService)

	whRoutes := r.Group("/warehouses")
	{
		whRoutes.GET("/", handler.GetAll())
		whRoutes.POST("/", handler.Create())
		whRoutes.GET("/:id", handler.Get())
		whRoutes.PATCH("/:id", handler.Update())
		whRoutes.DELETE("/:id", handler.Delete())
	}

	return r
}

func createBadJsonRequest(method string, url string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte("{dasd")))
	req.Header.Add("Content-Type", "application/json")

	return req, httptest.NewRecorder()
}

func TestCreateOkWarehouse(t *testing.T) {
	r := createWarehouseServer(&mocks.MockWarehouseService{
		MockRepository: mocks.MockWarehouseRepository{
			MockData: mocks.MockDataWarehouse,
		},
	})
	warehouseOk := domain.Warehouse{
		Address:            "Calle Falsa 123",
		Telephone:          "+54 9 11 5487-5421",
		WarehouseCode:      "CTX-458",
		MinimumCapacity:    5,
		MinimumTemperature: 7,
	}

	objRes := struct {
		Data domain.Warehouse `json:"data"`
	}{}

	req, rr := tests.CreateRequestTest(http.MethodPost, "/warehouses/", warehouseOk)
	r.ServeHTTP(rr, req)

	//Test de código de respuesta válido
	assert.Equal(t, 201, rr.Code)

	// Test cuerpo de respuesta válido
	err := json.Unmarshal(rr.Body.Bytes(), &objRes)
	assert.Nil(t, err)

	warehouseOk.ID = 3

	// Test de data correcta en respuesta
	assert.Equal(t, warehouseOk, objRes.Data)
}

func TestCreateConflictWarehouse(t *testing.T) {
	r := createWarehouseServer(&mocks.MockWarehouseService{
		MockRepository: mocks.MockWarehouseRepository{
			MockData: mocks.MockDataWarehouse,
		},
	})
	warehouseConflict := domain.Warehouse{
		Address:            "Calle Falsa 123",
		Telephone:          "+54 9 11 5487-5421",
		WarehouseCode:      "CTX-555",
		MinimumCapacity:    5,
		MinimumTemperature: 7,
	}

	objRes := struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	}{}

	req, rr := tests.CreateRequestTest(http.MethodPost, "/warehouses/", warehouseConflict)
	r.ServeHTTP(rr, req)

	//Test de código de respuesta válido
	assert.Equal(t, 409, rr.Code)

	// Test cuerpo de respuesta válido
	err := json.Unmarshal(rr.Body.Bytes(), &objRes)
	assert.Nil(t, err)

	// Test codigo y mensaje correctos en la respuesta
	assert.Equal(t, "conflict", objRes.Code)
	assert.Equal(t, "warehouse with code CTX-555 already exists", objRes.Message)
}

func TestCreateMissingFieldWarehouse(t *testing.T) {
	r := createWarehouseServer(&mocks.MockWarehouseService{
		MockRepository: mocks.MockWarehouseRepository{
			MockData: mocks.MockDataWarehouse,
		},
	})

	warehouseMissingField := domain.Warehouse{
		Address:            "Calle Falsa 123",
		Telephone:          "+54 9 11 5487-5421",
		MinimumCapacity:    5,
		MinimumTemperature: 7,
	}

	objRes := struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	}{}

	req, rr := tests.CreateRequestTest(http.MethodPost, "/warehouses/", warehouseMissingField)
	r.ServeHTTP(rr, req)

	//Test de código de respuesta válido
	assert.Equal(t, 422, rr.Code)

	// Test cuerpo de respuesta válido
	err := json.Unmarshal(rr.Body.Bytes(), &objRes)
	assert.Nil(t, err)

	// Test codigo y mensaje correctos en la respuesta
	assert.Equal(t, "unprocessable_entity", objRes.Code)
	assert.Equal(t, "field warehouse_code is required", objRes.Message)
}

func TestCreateBadJsonWarehouse(t *testing.T) {
	r := createWarehouseServer(&mocks.MockWarehouseService{
		MockRepository: mocks.MockWarehouseRepository{
			MockData: mocks.MockDataWarehouse,
		},
	})

	objRes := struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	}{}

	req, rr := createBadJsonRequest(http.MethodPost, "/warehouses/")
	r.ServeHTTP(rr, req)

	//Test de código de respuesta válido
	assert.Equal(t, 422, rr.Code)

	// Test cuerpo de respuesta válido
	err := json.Unmarshal(rr.Body.Bytes(), &objRes)
	assert.Nil(t, err)

	// Test codigo correcto en la respuesta
	assert.Equal(t, "unprocessable_entity", objRes.Code)
}

func TestCreateServerErrorWarehouse(t *testing.T) {
	r := createWarehouseServer(&mocks.MockWarehouseServiceError{
		MockRepository: mocks.MockWarehouseRepository{
			MockData: mocks.MockDataWarehouse,
		},
	})

	warehouseOk := domain.Warehouse{
		Address:            "Calle Falsa 123",
		Telephone:          "+54 9 11 5487-5421",
		WarehouseCode:      "CTX-458",
		MinimumCapacity:    5,
		MinimumTemperature: 7,
	}

	objRes := struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	}{}

	req, rr := tests.CreateRequestTest(http.MethodPost, "/warehouses/", warehouseOk)
	r.ServeHTTP(rr, req)

	//Test de código de respuesta válido
	assert.Equal(t, 500, rr.Code)

	// Test cuerpo de respuesta válido
	err := json.Unmarshal(rr.Body.Bytes(), &objRes)
	assert.Nil(t, err)

	// Test codigo y mensaje correctos en la respuesta
	assert.Equal(t, "internal_server_error", objRes.Code)
	assert.Equal(t, "internal server error", objRes.Message)
}

func TestFindAllWarehouse(t *testing.T) {
	r := createWarehouseServer(&mocks.MockWarehouseService{
		MockRepository: mocks.MockWarehouseRepository{
			MockData: mocks.MockDataWarehouse,
		},
	})

	objRes := struct {
		Data []domain.Warehouse `json:"data"`
	}{}

	req, rr := tests.CreateRequestTest(http.MethodGet, "/warehouses/", nil)
	r.ServeHTTP(rr, req)

	//Test de código de respuesta válido
	assert.Equal(t, 200, rr.Code)

	// Test cuerpo de respuesta válido
	err := json.Unmarshal(rr.Body.Bytes(), &objRes)
	assert.Nil(t, err)

	// Test de data correcta en respuesta
	assert.Equal(t, mocks.MockDataWarehouse, objRes.Data)
}

func TestFindAllServerErrorWarehouse(t *testing.T) {
	r := createWarehouseServer(&mocks.MockWarehouseServiceError{
		MockRepository: mocks.MockWarehouseRepository{
			MockData: mocks.MockDataWarehouse,
		},
	})

	objRes := struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	}{}

	req, rr := tests.CreateRequestTest(http.MethodGet, "/warehouses/", nil)
	r.ServeHTTP(rr, req)
	//Test de código de respuesta válido
	assert.Equal(t, 500, rr.Code)

	// Test cuerpo de respuesta válido
	err := json.Unmarshal(rr.Body.Bytes(), &objRes)
	assert.Nil(t, err)

	// Test codigo y mensaje correctos en la respuesta
	assert.Equal(t, "internal_server_error", objRes.Code)
	assert.Equal(t, "internal server error", objRes.Message)
}

func TestFindAllEmptyListWarehouse(t *testing.T) {
	r := createWarehouseServer(&mocks.MockWarehouseService{
		MockRepository: mocks.MockWarehouseRepository{
			MockData: mocks.MockEmptyDataWarehouse,
		},
	})

	objRes := struct {
		Data []domain.Warehouse `json:"data"`
	}{}

	req, rr := tests.CreateRequestTest(http.MethodGet, "/warehouses/", nil)
	r.ServeHTTP(rr, req)

	//Test de código de respuesta válido
	assert.Equal(t, 200, rr.Code)

	// Test cuerpo de respuesta válido
	err := json.Unmarshal(rr.Body.Bytes(), &objRes)
	assert.Nil(t, err)

	// Test de data correcta en respuesta
	assert.Equal(t, mocks.MockEmptyDataWarehouse, objRes.Data)
}

func TestFindByIdNonExistentWarehouse(t *testing.T) {
	r := createWarehouseServer(&mocks.MockWarehouseService{
		MockRepository: mocks.MockWarehouseRepository{
			MockData: mocks.MockDataWarehouse,
		},
	})

	objRes := struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	}{}

	req, rr := tests.CreateRequestTest(http.MethodGet, "/warehouses/3", nil)
	r.ServeHTTP(rr, req)
	//Test de código de respuesta válido
	assert.Equal(t, 404, rr.Code)

	// Test cuerpo de respuesta válido
	err := json.Unmarshal(rr.Body.Bytes(), &objRes)
	assert.Nil(t, err)

	// Test codigo y mensaje correctos en la respuesta
	assert.Equal(t, "not_found", objRes.Code)
	assert.Equal(t, "warehouse not found", objRes.Message)
}

func TestFindByIdExistentWarehouse(t *testing.T) {
	r := createWarehouseServer(&mocks.MockWarehouseService{
		MockRepository: mocks.MockWarehouseRepository{
			MockData: mocks.MockDataWarehouse,
		},
	})

	warehouseExistent := mocks.MockDataWarehouse[0]

	objRes := struct {
		Data domain.Warehouse `json:"data"`
	}{}

	req, rr := tests.CreateRequestTest(http.MethodGet, "/warehouses/1", nil)
	r.ServeHTTP(rr, req)

	//Test de código de respuesta válido
	assert.Equal(t, 200, rr.Code)

	// Test cuerpo de respuesta válido
	err := json.Unmarshal(rr.Body.Bytes(), &objRes)
	assert.Nil(t, err)

	// Test de data correcta en respuesta
	assert.Equal(t, warehouseExistent, objRes.Data)
}

func TestUpdateOkWarehouse(t *testing.T) {
	r := createWarehouseServer(&mocks.MockWarehouseService{
		MockRepository: mocks.MockWarehouseRepository{
			MockData: mocks.MockDataWarehouse,
		},
	})

	warehouseUpdated := domain.Warehouse{
		MinimumCapacity:    10,
		MinimumTemperature: 17,
		WarehouseCode:      "CTX-558",
		Address:            "domicilio 123",
		Telephone:          "+54 9 11-5654-4613",
	}

	objRes := struct {
		Data domain.Warehouse `json:"data"`
	}{}

	req, rr := tests.CreateRequestTest(http.MethodPatch, "/warehouses/1", warehouseUpdated)
	r.ServeHTTP(rr, req)

	warehouseUpdated.ID = 1

	//Test de código de respuesta válido
	assert.Equal(t, 200, rr.Code)

	// Test cuerpo de respuesta válido
	err := json.Unmarshal(rr.Body.Bytes(), &objRes)
	assert.Nil(t, err)

	// Test de data correcta en respuesta
	assert.Equal(t, warehouseUpdated, objRes.Data)
}

func TestUpdateConflictWarehouse(t *testing.T) {
	r := createWarehouseServer(&mocks.MockWarehouseService{
		MockRepository: mocks.MockWarehouseRepository{
			MockData: mocks.MockDataWarehouse,
		},
	})

	fieldsToUpdate := domain.Warehouse{
		MinimumCapacity:    10,
		MinimumTemperature: 17,
		WarehouseCode:      "CTX-576",
	}

	objRes := struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	}{}

	req, rr := tests.CreateRequestTest(http.MethodPatch, "/warehouses/1", fieldsToUpdate)
	r.ServeHTTP(rr, req)
	//Test de código de respuesta válido
	assert.Equal(t, 409, rr.Code)

	// Test cuerpo de respuesta válido
	err := json.Unmarshal(rr.Body.Bytes(), &objRes)
	assert.Nil(t, err)

	// Test codigo y mensaje correctos en la respuesta
	assert.Equal(t, "conflict", objRes.Code)
	assert.Equal(t, "warehouse with code CTX-576 already exists", objRes.Message)
}
func TestUpdateNonExistentWarehouse(t *testing.T) {
	r := createWarehouseServer(&mocks.MockWarehouseService{
		MockRepository: mocks.MockWarehouseRepository{
			MockData: mocks.MockDataWarehouse,
		},
	})

	fieldsToUpdate := domain.Warehouse{
		MinimumCapacity:    10,
		MinimumTemperature: 17,
	}

	objRes := struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	}{}

	req, rr := tests.CreateRequestTest(http.MethodPatch, "/warehouses/3", fieldsToUpdate)
	r.ServeHTTP(rr, req)
	//Test de código de respuesta válido
	assert.Equal(t, 404, rr.Code)

	// Test cuerpo de respuesta válido
	err := json.Unmarshal(rr.Body.Bytes(), &objRes)
	assert.Nil(t, err)

	// Test codigo y mensaje correctos en la respuesta
	assert.Equal(t, "not_found", objRes.Code)
	assert.Equal(t, "warehouse not found", objRes.Message)
}

func TestUpdateBadJsonWarehouse(t *testing.T) {
	r := createWarehouseServer(&mocks.MockWarehouseService{
		MockRepository: mocks.MockWarehouseRepository{
			MockData: mocks.MockDataWarehouse,
		},
	})

	objRes := struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	}{}

	req, rr := createBadJsonRequest(http.MethodPatch, "/warehouses/1")
	r.ServeHTTP(rr, req)

	//Test de código de respuesta válido
	assert.Equal(t, 422, rr.Code)

	// Test cuerpo de respuesta válido
	err := json.Unmarshal(rr.Body.Bytes(), &objRes)
	assert.Nil(t, err)

	// Test codigo correcto en la respuesta
	assert.Equal(t, "unprocessable_entity", objRes.Code)
}

func TestUpdateServerErrorWarehouse(t *testing.T) {
	r := createWarehouseServer(&mocks.MockWarehouseServiceError{
		MockRepository: mocks.MockWarehouseRepository{
			MockData: mocks.MockDataWarehouse,
		},
	})

	warehouseUpdated := domain.Warehouse{
		MinimumCapacity: 10,
	}

	objRes := struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	}{}

	req, rr := tests.CreateRequestTest(http.MethodPatch, "/warehouses/1", warehouseUpdated)
	r.ServeHTTP(rr, req)

	//Test de código de respuesta válido
	assert.Equal(t, 500, rr.Code)

	// Test cuerpo de respuesta válido
	err := json.Unmarshal(rr.Body.Bytes(), &objRes)
	assert.Nil(t, err)

	// Test codigo y mensaje correctos en la respuesta
	assert.Equal(t, "internal_server_error", objRes.Code)
	assert.Equal(t, "internal server error", objRes.Message)
}

func TestDeleteNonExistentWarehouse(t *testing.T) {
	r := createWarehouseServer(&mocks.MockWarehouseService{
		MockRepository: mocks.MockWarehouseRepository{
			MockData: mocks.MockDataWarehouse,
		},
	})

	objRes := struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	}{}

	req, rr := tests.CreateRequestTest(http.MethodDelete, "/warehouses/3", nil)
	r.ServeHTTP(rr, req)
	//Test de código de respuesta válido
	assert.Equal(t, 404, rr.Code)

	// Test cuerpo de respuesta válido
	err := json.Unmarshal(rr.Body.Bytes(), &objRes)
	assert.Nil(t, err)

	// Test codigo y mensaje correctos en la respuesta
	assert.Equal(t, "not_found", objRes.Code)
	assert.Equal(t, "warehouse not found", objRes.Message)
}

func TestDeleteOkWarehouse(t *testing.T) {
	r := createWarehouseServer(&mocks.MockWarehouseService{
		MockRepository: mocks.MockWarehouseRepository{
			MockData: mocks.MockDataWarehouse,
		},
	})

	req, rr := tests.CreateRequestTest(http.MethodDelete, "/warehouses/1", nil)
	r.ServeHTTP(rr, req)

	//Test de código de respuesta válido
	assert.Equal(t, 204, rr.Code)

	// Test de respuesta vacia
	assert.Equal(t, "", rr.Body.String())
}

func TestUpdateIDNonIntWarehouse(t *testing.T) {
	r := createWarehouseServer(&mocks.MockWarehouseService{
		MockRepository: mocks.MockWarehouseRepository{
			MockData: mocks.MockDataWarehouse,
		},
	})

	req, rr := tests.CreateRequestTest(http.MethodPatch, "/warehouses/notInt", nil)
	r.ServeHTTP(rr, req)

	objRes := struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	}{}

	//Test de código de respuesta válido
	assert.Equal(t, 400, rr.Code)

	// Test cuerpo de respuesta válido
	err := json.Unmarshal(rr.Body.Bytes(), &objRes)
	assert.Nil(t, err)

	// Test de data correcta en respuesta
	assert.Equal(t, "bad_request", objRes.Code)
	assert.Equal(t, "id must be integer", objRes.Message)
}

func TestDeleteIDNonIntWarehouse(t *testing.T) {
	r := createWarehouseServer(&mocks.MockWarehouseService{
		MockRepository: mocks.MockWarehouseRepository{
			MockData: mocks.MockDataWarehouse,
		},
	})

	req, rr := tests.CreateRequestTest(http.MethodDelete, "/warehouses/notInt", nil)
	r.ServeHTTP(rr, req)

	objRes := struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	}{}

	//Test de código de respuesta válido
	assert.Equal(t, 400, rr.Code)

	// Test cuerpo de respuesta válido
	err := json.Unmarshal(rr.Body.Bytes(), &objRes)
	assert.Nil(t, err)

	// Test de data correcta en respuesta
	assert.Equal(t, "bad_request", objRes.Code)
	assert.Equal(t, "id must be integer", objRes.Message)
}
