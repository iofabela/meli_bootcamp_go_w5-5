package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/iofabela/meli_bootcamp_go_w5-5/internal/domain"
	"github.com/iofabela/meli_bootcamp_go_w5-5/internal/seller"
	"github.com/iofabela/meli_bootcamp_go_w5-5/tests"
	"github.com/iofabela/meli_bootcamp_go_w5-5/tests/mocks"
	"github.com/stretchr/testify/assert"
)

// Create Server
func createServerSeller(sellerService seller.Service) *gin.Engine {
	r := gin.Default()

	seller := NewSeller(sellerService)

	sellersGroup := r.Group("/sellers")
	{
		sellersGroup.GET("/", seller.GetAll())
		sellersGroup.GET("/:id", seller.Get())
		sellersGroup.POST("/", seller.Create())
		sellersGroup.PATCH("/:id", seller.Update())
		sellersGroup.DELETE("/:id", seller.Delete())
	}
	return r
}

func createBadJsonRequestSeller(method string, url string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte("{dasd")))
	req.Header.Add("Content-Type", "application/json")

	return req, httptest.NewRecorder()
}

// Test Create Ok
func TestCreateOkSellerHandler(t *testing.T) {
	//arrange
	r := createServerSeller(&mocks.MockServiceSeller{
		MockRepo: mocks.MockSellerRepo{
			MockSeller: mocks.MockListSellers,
		},
	})
	resp := struct {
		Data domain.Seller `json:"data"`
	}{}
	req, rr := tests.CreateRequestTest(http.MethodPost, "/sellers/", mocks.MockNewSeller)
	r.ServeHTTP(rr, req)
	//response code
	assert.Equal(t, 201, rr.Code)

	err := json.Unmarshal(rr.Body.Bytes(), &resp)
	assert.Nil(t, err)

	mocks.MockNewSeller.ID = 3
	//correct answer on date
	assert.Equal(t, mocks.MockNewSeller, resp.Data)
}

// Create Bad Request
func TestCreateBadRequestSellerHandler(t *testing.T) {
	//arrange
	r := createServerSeller(&mocks.MockServiceSeller{
		MockRepo: mocks.MockSellerRepo{
			MockSeller: mocks.MockListSellers,
		},
	})

	objRes := struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	}{}

	req, rr := createBadJsonRequestSeller(http.MethodPost, "/sellers/")
	r.ServeHTTP(rr, req)
	//response code
	assert.Equal(t, 400, rr.Code)

	err := json.Unmarshal(rr.Body.Bytes(), &objRes)
	assert.Nil(t, err)

}

// Create Fail
func TestCreateFaildSellerHandler(t *testing.T) {
	//arrange
	r := createServerSeller(&mocks.MockServiceSeller{
		MockRepo: mocks.MockSellerRepo{
			MockSeller: mocks.MockListSellers,
		},
	})

	objRes := struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	}{}

	req, rr := tests.CreateRequestTest(http.MethodPost, "/sellers/", mocks.MockMissingCid)
	r.ServeHTTP(rr, req)
	//response code
	assert.Equal(t, 422, rr.Code)

	err := json.Unmarshal(rr.Body.Bytes(), &objRes)
	assert.Nil(t, err)
	//correct answer on date
	assert.Equal(t, "unprocessable_entity", objRes.Code)
}

func TestCreateMissingFielSellerHandler(t *testing.T) {
	//arrange
	r := createServerSeller(&mocks.MockServiceSeller{
		MockRepo: mocks.MockSellerRepo{
			MockSeller: mocks.MockListSellers,
		},
	})

	objRes := struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	}{}

	req, rr := tests.CreateRequestTest(http.MethodPost, "/sellers/", mocks.MockMissingCid)
	r.ServeHTTP(rr, req)

	//Test de código de respuesta válido
	assert.Equal(t, 422, rr.Code)

	// Test cuerpo de respuesta válido
	err := json.Unmarshal(rr.Body.Bytes(), &objRes)
	assert.Nil(t, err)

	// Test codigo y mensaje correctos en la respuesta
	assert.Equal(t, "unprocessable_entity", objRes.Code)
}

// Create Conflict
func TestCreateWithConflictSellerHandler(t *testing.T) {
	//arrange
	r := createServerSeller(&mocks.MockServiceSeller{
		MockRepo: mocks.MockSellerRepo{
			MockSeller: mocks.MockListSellers,
		},
	})

	objRes := struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	}{}

	req, rr := tests.CreateRequestTest(http.MethodPost, "/sellers/", mocks.MockNewSellerWithConflict)
	r.ServeHTTP(rr, req)
	//response code
	assert.Equal(t, 409, rr.Code)

	err := json.Unmarshal(rr.Body.Bytes(), &objRes)
	assert.Nil(t, err)
	//correct answer on date
	assert.Equal(t, "conflict", objRes.Code)
	assert.Equal(t, "there is already a seller with that cid", objRes.Message)
}

// Find All
func TestFindAllSellerHandler(t *testing.T) {
	//arrange
	r := createServerSeller(&mocks.MockServiceSeller{
		MockRepo: mocks.MockSellerRepo{
			MockSeller: mocks.MockListSellers,
		},
	})

	objRes := struct {
		Data []domain.Seller `json:"data"`
	}{}

	req, rr := tests.CreateRequestTest(http.MethodGet, "/sellers/", nil)
	r.ServeHTTP(rr, req)
	//response code
	assert.Equal(t, 200, rr.Code)

	err := json.Unmarshal(rr.Body.Bytes(), &objRes)
	assert.Nil(t, err)
	//correct answer on date
	assert.Equal(t, mocks.MockListSellers, objRes.Data)
}

// Find By Id No Existente
func TestFindByIdNonExistentSellerHandler(t *testing.T) {
	//arrange
	r := createServerSeller(&mocks.MockServiceSeller{
		MockRepo: mocks.MockSellerRepo{
			MockSeller: mocks.MockListSellers,
		},
	})

	objRes := struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	}{}

	req, rr := tests.CreateRequestTest(http.MethodGet, "/sellers/3", nil)
	r.ServeHTTP(rr, req)
	//response  code
	assert.Equal(t, 404, rr.Code)

	err := json.Unmarshal(rr.Body.Bytes(), &objRes)
	assert.Nil(t, err)
	//correct answer on date
	assert.Equal(t, "not_found", objRes.Code)
	assert.Equal(t, "no seller with the id was found 3", objRes.Message)
}

// Find By id Existent
func TestFindByIdExistentSellerHandler(t *testing.T) {
	//arrange
	r := createServerSeller(&mocks.MockServiceSeller{
		MockRepo: mocks.MockSellerRepo{
			MockSeller: mocks.MockListSellers,
		},
	})

	sellerExistent := mocks.MockListSellers[0]

	objRes := struct {
		Data domain.Seller `json:"data"`
	}{}

	req, rr := tests.CreateRequestTest(http.MethodGet, "/sellers/1", nil)
	r.ServeHTTP(rr, req)
	//response code
	assert.Equal(t, 200, rr.Code)

	err := json.Unmarshal(rr.Body.Bytes(), &objRes)
	assert.Nil(t, err)
	//correct answer on date
	assert.Equal(t, sellerExistent, objRes.Data)
}

// Update Ok
func TestUpdateOkSellerHandler(t *testing.T) {
	//arrange
	r := createServerSeller(&mocks.MockServiceSeller{
		MockRepo: mocks.MockSellerRepo{
			MockSeller: mocks.MockListSellers,
		},
	})

	objRes := struct {
		Data domain.Seller `json:"data"`
	}{}

	req, rr := tests.CreateRequestTest(http.MethodPatch, "/sellers/2", mocks.MockUpdateSeller)
	r.ServeHTTP(rr, req)

	mocks.MockUpdateSeller.ID = 2
	//response code
	assert.Equal(t, 200, rr.Code)

	err := json.Unmarshal(rr.Body.Bytes(), &objRes)
	assert.Nil(t, err)
	//correct answer on date

	assert.Equal(t, mocks.MockUpdateSeller, objRes.Data)
}

// Update No Existente
func TestUpdateNonExistentSellerHandler(t *testing.T) {
	//arrange
	r := createServerSeller(&mocks.MockServiceSeller{
		MockRepo: mocks.MockSellerRepo{
			MockSeller: mocks.MockListSellers,
		},
	})

	fieldsToUpdate := domain.Seller{
		CompanyName: "MELI",
		Address:     "Argentina",
	}

	objRes := struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	}{}

	req, rr := tests.CreateRequestTest(http.MethodPatch, "/sellers/3", fieldsToUpdate)
	r.ServeHTTP(rr, req)
	//response code
	assert.Equal(t, 404, rr.Code)

	err := json.Unmarshal(rr.Body.Bytes(), &objRes)
	assert.Nil(t, err)
	//correct answer on date
	assert.Equal(t, "not_found", objRes.Code)
	assert.Equal(t, "no seller with the id was found 3", objRes.Message)
}
func TestUpdateConflictSellerHandler(t *testing.T) {
	//arrange
	r := createServerSeller(&mocks.MockServiceSeller{
		MockRepo: mocks.MockSellerRepo{
			MockSeller: mocks.MockListSellers,
		},
	})

	fieldsToUpdate := domain.Seller{
		CID:         3,
		CompanyName: "MELI",
		Address:     "Argentina",
		Telephone:   "123",
	}

	objRes := struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	}{}

	req, rr := tests.CreateRequestTest(http.MethodPatch, "/sellers/1", fieldsToUpdate)
	r.ServeHTTP(rr, req)
	//Test de código de respuesta válido
	assert.Equal(t, 409, rr.Code)

	// Test cuerpo de respuesta válido
	err := json.Unmarshal(rr.Body.Bytes(), &objRes)
	assert.Nil(t, err)

	// Test codigo y mensaje correctos en la respuesta
	assert.Equal(t, "conflict", objRes.Code)
	assert.Equal(t, "there is already a seller with that cid", objRes.Message)
}

func TestUpdateIDNonIntSellerHandler(t *testing.T) {
	//arrange
	r := createServerSeller(&mocks.MockServiceSeller{
		MockRepo: mocks.MockSellerRepo{
			MockSeller: mocks.MockListSellers,
		},
	})

	req, rr := tests.CreateRequestTest(http.MethodPatch, "/sellers/notInt", nil)
	r.ServeHTTP(rr, req)

	objRes := struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	}{}

	//response code
	assert.Equal(t, 400, rr.Code)

	err := json.Unmarshal(rr.Body.Bytes(), &objRes)
	assert.Nil(t, err)
	assert.Equal(t, "bad_request", objRes.Code)
	assert.Equal(t, "invalid id, must be integer", objRes.Message)
}

// Delete No existente
func TestDeleteNonExistentSellerHandler(t *testing.T) {
	//arrange
	r := createServerSeller(&mocks.MockServiceSeller{
		MockRepo: mocks.MockSellerRepo{
			MockSeller: mocks.MockListSellers,
		},
	})

	objRes := struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	}{}

	req, rr := tests.CreateRequestTest(http.MethodDelete, "/sellers/3", nil)
	r.ServeHTTP(rr, req)
	//response code
	assert.Equal(t, 404, rr.Code)

	err := json.Unmarshal(rr.Body.Bytes(), &objRes)
	assert.Nil(t, err)
	//correct answer on date
	assert.Equal(t, "not_found", objRes.Code)
	assert.Equal(t, "no seller with the id was found 3", objRes.Message)
}

// Delete Ok
func TestDeleteOkSellerHandler(t *testing.T) {
	//arrange
	r := createServerSeller(&mocks.MockServiceSeller{
		MockRepo: mocks.MockSellerRepo{
			MockSeller: mocks.MockListSellers,
		},
	})

	req, rr := tests.CreateRequestTest(http.MethodDelete, "/sellers/1", nil)
	r.ServeHTTP(rr, req)
	//response code
	assert.Equal(t, 204, rr.Code)
	assert.Equal(t, "", rr.Body.String())
}
func TestDeleteIDNonIntSellerHandler(t *testing.T) {
	//arrange
	r := createServerSeller(&mocks.MockServiceSeller{
		MockRepo: mocks.MockSellerRepo{
			MockSeller: mocks.MockListSellers,
		},
	})

	req, rr := tests.CreateRequestTest(http.MethodDelete, "/sellers/notInt", nil)
	r.ServeHTTP(rr, req)

	objRes := struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	}{}

	//response code
	assert.Equal(t, 400, rr.Code)

	err := json.Unmarshal(rr.Body.Bytes(), &objRes)
	assert.Nil(t, err)

	assert.Equal(t, "bad_request", objRes.Code)
	assert.Equal(t, "invalid id, must be integer", objRes.Message)
}

//////////

func TestUpdateBadJsonSellerHandler(t *testing.T) {
	//arrange
	r := createServerSeller(&mocks.MockServiceSeller{
		MockRepo: mocks.MockSellerRepo{
			MockSeller: mocks.MockListSellers,
		},
	})

	objRes := struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	}{}

	req, rr := createBadJsonRequestSeller(http.MethodPatch, "/sellers/1")
	r.ServeHTTP(rr, req)

	//Test de código de respuesta válido
	assert.Equal(t, 404, rr.Code)

	// Test cuerpo de respuesta válido
	err := json.Unmarshal(rr.Body.Bytes(), &objRes)
	assert.Nil(t, err)
}

func TestFindIDNonIntSellerHandler(t *testing.T) {
	//arrange
	r := createServerSeller(&mocks.MockServiceSeller{
		MockRepo: mocks.MockSellerRepo{
			MockSeller: mocks.MockListSellers,
		},
	})

	req, rr := tests.CreateRequestTest(http.MethodGet, "/sellers/notInt", nil)
	r.ServeHTTP(rr, req)

	objRes := struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	}{}

	//response code
	assert.Equal(t, 400, rr.Code)

	err := json.Unmarshal(rr.Body.Bytes(), &objRes)
	assert.Nil(t, err)
	assert.Equal(t, "bad_request", objRes.Code)
	assert.Equal(t, "invalid id, must be integer", objRes.Message)
}
