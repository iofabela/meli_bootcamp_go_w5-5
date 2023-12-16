package handler

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/extmatperez/meli_bootcamp_go_w5-5/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w5-5/internal/employee"
	"github.com/extmatperez/meli_bootcamp_go_w5-5/tests"
	"github.com/extmatperez/meli_bootcamp_go_w5-5/tests/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func createServerEmployee(employeeService employee.Service) *gin.Engine {
	r := gin.Default()

	handler := NewEmployee(employeeService)

	employeeRoutes := r.Group("/employees")
	{
		employeeRoutes.GET("/", handler.GetAll())
		employeeRoutes.POST("/", handler.Create())
		employeeRoutes.GET("/:id", handler.Get())
		employeeRoutes.PATCH("/:id", handler.Update())
		employeeRoutes.DELETE("/:id", handler.Delete())
	}

	return r
}

func TestCreateOkEmployee(t *testing.T) {
	//Crea el server y define las rutas
	r := createServerEmployee(&mocks.MockEmployeeService{
		MockRepository: mocks.MockEmployeeRepository{
			MockData: mocks.MockEmployees,
		},
	})

	resp := struct {
		Data domain.Employee `json:"data"`
	}{}

	req, rr := tests.CreateRequestTest(http.MethodPost, "/employees/", mocks.MockNewEmployee)
	r.ServeHTTP(rr, req)

	//Respuesta válida
	assert.Equal(t, 201, rr.Code)

	//Cuerpo de respuesta válido
	err := json.Unmarshal(rr.Body.Bytes(), &resp)
	assert.Nil(t, err)

	resp.Data.ID = 4

	//Datos correctos
	assert.Equal(t, mocks.MockNewEmployee, resp.Data)
}

func TestCreateFailEmployee(t *testing.T) {
	//Crea el server y define las rutas
	r := createServerEmployee(&mocks.MockEmployeeService{
		MockRepository: mocks.MockEmployeeRepository{
			MockData: mocks.MockEmployees,
		},
	})

	resp := struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	}{}
	//Crea request de tipo post y response para obtener el resultado
	req, rr := tests.CreateRequestTest(http.MethodPost, "/employees/", "")
	//Indica al servidor que puede atender la solicitud
	r.ServeHTTP(rr, req)

	//Valida el codigo de respuesta
	assert.Equal(t, 422, rr.Code)
	
	//Valida que el cuerpo de la respuesta sea correcto
	err := json.Unmarshal(rr.Body.Bytes(), &resp)
	assert.Nil(t, err)

	//Test que valida los datos correctos en la respuesta recibida
	assert.Equal(t, "unprocessable_entity", resp.Code)
}

func TestCreateConflictEmployee(t *testing.T) {

	r := createServerEmployee(&mocks.MockEmployeeService{
		MockRepository: mocks.MockEmployeeRepository{
			MockData: mocks.MockEmployees,
		},
	})

	resp := struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	}{}

	req, rr := tests.CreateRequestTest(http.MethodPost, "/employees/", mocks.MockEmployees[1])
	r.ServeHTTP(rr, req)

	assert.Equal(t, 409, rr.Code)

	err := json.Unmarshal(rr.Body.Bytes(), &resp)
	assert.Nil(t, err)

	assert.Equal(t, "conflict", resp.Code)
	assert.Equal(t, "El employee ya existe", resp.Message)
}

func TestFindAllEmptyEmployees(t *testing.T) {

	r := createServerEmployee(&mocks.MockEmployeeService{
		MockRepository: mocks.MockEmployeeRepository{
			MockData: mocks.MockEmptyEmployees,
		},
	})

	resp := struct {
		Data []domain.Employee `json:"data"`
	}{}

	req, rr := tests.CreateRequestTest(http.MethodGet, "/employees/", nil)

	r.ServeHTTP(rr, req)

	assert.Equal(t, 200, rr.Code)

	err := json.Unmarshal(rr.Body.Bytes(), &resp)
	assert.Nil(t, err)
	assert.True(t, len(resp.Data) == 0)
}

func TestFindAllEmployee(t *testing.T) {
	//Crea request de tipo post y response para obtener el resultado
	r := createServerEmployee(&mocks.MockEmployeeService{
		MockRepository: mocks.MockEmployeeRepository{
			MockData: mocks.MockEmployees,
		},
	})

	resp := struct {
		Data []domain.Employee `json:"data"`
	}{}

	//Crea request de tipo get y response para obtener el resultado
	req, rr := tests.CreateRequestTest(http.MethodGet, "/employees/", nil)
	//Indica al servidor que puede atender la solicitud
	r.ServeHTTP(rr, req)

	//Valida el codigo de respuesta
	assert.Equal(t, 200, rr.Code)

	//Valida que el cuerpo de la respuesta sea correcto
	err := json.Unmarshal(rr.Body.Bytes(), &resp)
	assert.Nil(t, err)

	//Los datos deben ser iguales del actual a los esperados
	assert.Equal(t, mocks.MockEmployees, resp.Data)
}

func TestFindByIdNonExistentEmployee(t *testing.T) {
	//Crea el server y define las rutas
	r := createServerEmployee(&mocks.MockEmployeeService{
		MockRepository: mocks.MockEmployeeRepository{
			MockData: mocks.MockEmployees,
		},
	})

	resp := struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	}{}

	//Crea request de tipo get y response para obtener el resultado
	req, rr := tests.CreateRequestTest(http.MethodGet, "/employees/8", nil)
	//Indica al servidor que puede atender la solicitud
	r.ServeHTTP(rr, req)

	//Valida el codigo de respuesta
	assert.Equal(t, 404, rr.Code)

	//Valida que el cuerpo de la respuesta sea correcto
	err := json.Unmarshal(rr.Body.Bytes(), &resp)
	assert.Nil(t, err)

	//Los datos deben ser iguales del actual a los esperados
	assert.Equal(t, "not_found", resp.Code)
	assert.Equal(t, "El id 8 no existe", resp.Message)
}

func TestFindByIdExistentEmployee(t *testing.T) {

	r := createServerEmployee(&mocks.MockEmployeeService{
		MockRepository: mocks.MockEmployeeRepository{
			MockData: mocks.MockEmployees,
		},
	})

	employeeExistent := mocks.MockEmployees[0]

	resp := struct {
		Data domain.Employee `json:"data"`
	}{}

	req, rr := tests.CreateRequestTest(http.MethodGet, "/employees/1", nil)
	r.ServeHTTP(rr, req)

	assert.Equal(t, 200, rr.Code)

	err := json.Unmarshal(rr.Body.Bytes(), &resp)
	assert.Nil(t, err)

	assert.Equal(t, employeeExistent, resp.Data)
}

func TestFindIDNonIntEmployee(t *testing.T) {

	r := createServerEmployee(&mocks.MockEmployeeService{
		MockRepository: mocks.MockEmployeeRepository{
			MockData: mocks.MockEmployees,
		},
	})

	resp := struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	}{}

	req, rr := tests.CreateRequestTest(http.MethodGet, "/employees/string", nil)
	r.ServeHTTP(rr, req)

	assert.Equal(t, 400, rr.Code)

	err := json.Unmarshal(rr.Body.Bytes(), &resp)
	assert.Nil(t, err)
	assert.Equal(t, "bad_request", resp.Code)
	assert.Equal(t, "El id es invalido", resp.Message)
}

func TestUpdateOkEmployee(t *testing.T) {

	r := createServerEmployee(&mocks.MockEmployeeService{
		MockRepository: mocks.MockEmployeeRepository{
			MockData: mocks.MockEmployees,
		},
	})

	resp := struct {
		Data domain.Employee `json:"data"`
	}{}

	req, rr := tests.CreateRequestTest(http.MethodPatch, "/employees/2", mocks.MockUpdateEmployee)
	r.ServeHTTP(rr, req)

	mocks.MockUpdateEmployee.ID = 2

	assert.Equal(t, 200, rr.Code)

	err := json.Unmarshal(rr.Body.Bytes(), &resp)
	assert.Nil(t, err)

	assert.Equal(t, mocks.MockUpdateEmployee, resp.Data)
}

func TestUpdateNonExistentEmployee(t *testing.T) {

	r := createServerEmployee(&mocks.MockEmployeeService{
		MockRepository: mocks.MockEmployeeRepository{
			MockData: mocks.MockEmployees,
		},
	})

	resp := struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	}{}

	req, rr := tests.CreateRequestTest(http.MethodPatch, "/employees/8", &mocks.MockUpdateEmployee)
	r.ServeHTTP(rr, req)

	assert.Equal(t, 404, rr.Code)

	err := json.Unmarshal(rr.Body.Bytes(), &resp)
	assert.Nil(t, err)

	assert.Equal(t, "not_found", resp.Code)
	assert.Equal(t, "El id 8 no existe", resp.Message)
}

func TestUpdateIDNonIntEmployee(t *testing.T) {
	
	r := createServerEmployee(&mocks.MockEmployeeService{
		MockRepository: mocks.MockEmployeeRepository{
			MockData: mocks.MockEmployees,
		},
	})

	resp := struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	}{}

	req, rr := tests.CreateRequestTest(http.MethodPatch, "/employees/string", nil)
	r.ServeHTTP(rr, req)

	//Test de código de respuesta válido
	assert.Equal(t, 400, rr.Code)

	// Test cuerpo de respuesta válido
	err := json.Unmarshal(rr.Body.Bytes(), &resp)
	assert.Nil(t, err)

	// Test de data correcta en respuesta
	assert.Equal(t, "bad_request", resp.Code)
	assert.Equal(t, "El id es invalido", resp.Message)
}

func TestDeleteNonExistentEmployee(t *testing.T) {

	r := createServerEmployee(&mocks.MockEmployeeService{
		MockRepository: mocks.MockEmployeeRepository{
			MockData: mocks.MockEmployees,
		},
	})

	resp := struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	}{}

	req, rr := tests.CreateRequestTest(http.MethodDelete, "/employees/8", nil)

	r.ServeHTTP(rr, req)

	assert.Equal(t, 404, rr.Code)

	err := json.Unmarshal(rr.Body.Bytes(), &resp)
	assert.Nil(t, err)

	assert.Equal(t, "not_found", resp.Code)
	assert.Equal(t, "El id 8 no existe", resp.Message)
}

func TestDeleteOkEmployee(t *testing.T) {

	r := createServerEmployee(&mocks.MockEmployeeService{
		MockRepository: mocks.MockEmployeeRepository{
			MockData: mocks.MockEmployees,
		},
	})

	req, rr := tests.CreateRequestTest(http.MethodDelete, "/employees/1", nil)

	r.ServeHTTP(rr, req)

	assert.Equal(t, 204, rr.Code)
}

func TestDeleteIDNonIntEmployee(t *testing.T) {
	
	r := createServerEmployee(&mocks.MockEmployeeService{
		MockRepository: mocks.MockEmployeeRepository{
			MockData: mocks.MockEmployees,
		},
	})

	resp := struct {
		Code string `json:"code"`
		Message string `json:"message"`
	}{}

	req, rr := tests.CreateRequestTest(http.MethodDelete, "/employees/string", nil)
	r.ServeHTTP(rr, req)

	//Test de código de respuesta válido
	assert.Equal(t, 400, rr.Code)

	// Test cuerpo de respuesta válido
	err := json.Unmarshal(rr.Body.Bytes(), &resp)
	assert.Nil(t, err)

	// Test de data correcta en respuesta
	assert.Equal(t, "bad_request", resp.Code)
	assert.Equal(t, "El id es invalido", resp.Message)
}
