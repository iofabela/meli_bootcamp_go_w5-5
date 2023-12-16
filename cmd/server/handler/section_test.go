package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/extmatperez/meli_bootcamp_go_w5-5/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w5-5/tests"
	"github.com/extmatperez/meli_bootcamp_go_w5-5/tests/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type mockSectionService struct {
	mockData []domain.Section
}

func (mk *mockSectionService) GetAll(ctx context.Context) ([]domain.Section, error) {
	return mk.mockData, nil
}

func (mk *mockSectionService) Get(ctx context.Context, id int) (domain.Section, error) {
	for _, section := range mk.mockData {
		if section.ID == id {
			return section, nil
		}
	}
	return domain.Section{}, fmt.Errorf("no se encontro seccion: %d", id)
}

func (mk *mockSectionService) Save(ctx context.Context, s domain.Section) (int, error) {
	newId := len(mk.mockData) + 1
	s.ID = newId
	mk.mockData = append(mk.mockData, s)
	return newId, nil
}

func (mk *mockSectionService) Update(ctx context.Context, s domain.Section) (domain.Section, error) {
	for _, section := range mk.mockData {
		if section.ID == s.ID {
			section = s
			return section, nil
		}
	}
	return domain.Section{}, fmt.Errorf("no se encontro seccion: %d", s.ID)
}

func (mk *mockSectionService) Delete(ctx context.Context, id int) error {
	for i, section := range mk.mockData {
		if section.ID == id {
			mk.mockData = append(mk.mockData[:i], mk.mockData[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("no se encontro seccion: %d", id)
}

func (mk *mockSectionService) ReportProductsGetAll(ctx context.Context) ([]domain.ProductReport, error) {
	return []domain.ProductReport{}, nil
}
func (mk *mockSectionService) ReportProductsGet(ctx context.Context, id int) (domain.ProductReport, error) {
	return domain.ProductReport{}, nil
}

type mockSectionErrorService struct {
	mockData []domain.Section
}

func (mk *mockSectionErrorService) GetAll(ctx context.Context) ([]domain.Section, error) {
	return []domain.Section{}, fmt.Errorf("no se pueden obtener las secciones")
}

func (mk *mockSectionErrorService) Get(ctx context.Context, id int) (domain.Section, error) {
	for _, section := range mk.mockData {
		if section.ID == id {
			return section, nil
		}
	}
	return domain.Section{}, fmt.Errorf("no se encontro seccion: %d", id)
}

func (mk *mockSectionErrorService) Save(ctx context.Context, s domain.Section) (int, error) {
	return 0, fmt.Errorf("no se pudo crear seccion")
}

func (mk *mockSectionErrorService) Update(ctx context.Context, s domain.Section) (domain.Section, error) {
	return domain.Section{}, fmt.Errorf("no se actualizo seccion: %d", s.ID)
}

func (mk *mockSectionErrorService) Delete(ctx context.Context, id int) error {
	return fmt.Errorf("no se encontro elimino: %d", id)
}

func (mk *mockSectionErrorService) ReportProductsGetAll(ctx context.Context) ([]domain.ProductReport, error) {
	return []domain.ProductReport{}, fmt.Errorf("Err")
}
func (mk *mockSectionErrorService) ReportProductsGet(ctx context.Context, id int) (domain.ProductReport, error) {
	return domain.ProductReport{}, fmt.Errorf("Err")
}

func createSectionServer() *gin.Engine {
	mService := mockSectionService{mocks.MockListaSections}
	s := NewSection(&mService)

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	productos := router.Group("/section")
	{
		productos.GET("/", s.GetAll())
		productos.GET("/:id", s.Get())
		productos.POST("/", s.Create())
		productos.PATCH("/:id", s.Update())
		productos.DELETE("/:id", s.Delete())
	}
	return router
}

func createSectionEmptyServer() *gin.Engine {
	mService := mockSectionService{[]domain.Section{}}
	s := NewSection(&mService)

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	productos := router.Group("/section")
	{
		productos.GET("/", s.GetAll())
		productos.GET("/:id", s.Get())
		productos.POST("/", s.Create())
		productos.PATCH("/:id", s.Update())
		productos.DELETE("/:id", s.Delete())
	}
	return router
}

func createSectionErrorServer() *gin.Engine {
	mService := mockSectionErrorService{mocks.MockListaSections}
	s := NewSection(&mService)

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	productos := router.Group("/section")
	{
		productos.GET("/", s.GetAll())
		productos.GET("/:id", s.Get())
		productos.POST("/", s.Create())
		productos.PATCH("/:id", s.Update())
		productos.DELETE("/:id", s.Delete())
	}
	return router
}

func TestCreateOkHandler(t *testing.T) {
	expectedData := mocks.MockNuevaSection
	expectedData.ID = len(mocks.MockListaSections) + 1
	r := createSectionServer()
	req, res := tests.CreateRequestTest(http.MethodPost, "/section/", mocks.MockNuevaSectionRequest)
	r.ServeHTTP(res, req)
	assert.Equal(t, http.StatusCreated, res.Code)
	resData := map[string]domain.Section{}
	jsonErr := json.Unmarshal(res.Body.Bytes(), &resData)
	assert.Nil(t, jsonErr)
	assert.Equal(t, expectedData, resData["data"])
}

func TestCreateConflictHandler(t *testing.T) {
	expectedMessage := fmt.Sprintf(EmptyField, "section_number")
	r := createSectionServer()
	req, res := tests.CreateRequestTest(http.MethodPost, "/section/", mocks.MockNuevaSectionRequestConflict)
	r.ServeHTTP(res, req)
	assert.Equal(t, http.StatusUnprocessableEntity, res.Code)
	resData := map[string]string{}
	jsonErr := json.Unmarshal(res.Body.Bytes(), &resData)
	assert.Nil(t, jsonErr)
	responseErr := fmt.Errorf(resData["message"])
	assert.EqualError(t, responseErr, expectedMessage)
}

func TestCreateInvalidJSONHandler(t *testing.T) {
	r := createSectionServer()
	req, res := tests.CreateRequestTest(http.MethodPost, "/section/", []byte("{invalid}"))
	r.ServeHTTP(res, req)
	assert.Equal(t, http.StatusBadRequest, res.Code)
	resData := map[string]string{}
	jsonErr := json.Unmarshal(res.Body.Bytes(), &resData)
	assert.Nil(t, jsonErr)
}

func TestCreateErrorHandler(t *testing.T) {
	r := createSectionErrorServer()
	req, res := tests.CreateRequestTest(http.MethodPost, "/section/", mocks.MockNuevaSectionRequest)
	r.ServeHTTP(res, req)
	assert.Equal(t, http.StatusInternalServerError, res.Code)
	resData := map[string]string{}
	jsonErr := json.Unmarshal(res.Body.Bytes(), &resData)
	assert.Nil(t, jsonErr)
}

func TestFindAllHandler(t *testing.T) {
	expectedSections := mocks.MockListaSections
	r := createSectionServer()
	req, res := tests.CreateRequestTest(http.MethodGet, "/section/", nil)
	r.ServeHTTP(res, req)
	assert.Equal(t, http.StatusOK, res.Code)
	resData := map[string][]domain.Section{}
	jsonErr := json.Unmarshal(res.Body.Bytes(), &resData)
	assert.Nil(t, jsonErr)
	assert.Equal(t, expectedSections, resData["data"])
}

func TestFindAllEmptyHandler(t *testing.T) {
	expectedSections := []domain.Section{}
	r := createSectionEmptyServer()
	req, res := tests.CreateRequestTest(http.MethodGet, "/section/", nil)
	r.ServeHTTP(res, req)
	assert.Equal(t, http.StatusOK, res.Code)
	resData := map[string][]domain.Section{}
	jsonErr := json.Unmarshal(res.Body.Bytes(), &resData)
	assert.Nil(t, jsonErr)
	assert.Equal(t, expectedSections, resData["data"])
}

func TestFindAllErrorEmptyHandler(t *testing.T) {
	r := createSectionErrorServer()
	req, res := tests.CreateRequestTest(http.MethodGet, "/section/", nil)
	r.ServeHTTP(res, req)
	assert.Equal(t, http.StatusInternalServerError, res.Code)
	resData := map[string]string{}
	jsonErr := json.Unmarshal(res.Body.Bytes(), &resData)
	assert.Nil(t, jsonErr)
}

func TestFindByIdNonExistantHandler(t *testing.T) {
	searchedId := 10
	r := createSectionServer()
	req, res := tests.CreateRequestTest(http.MethodGet, fmt.Sprintf("/section/%d", searchedId), nil)
	r.ServeHTTP(res, req)
	assert.Equal(t, http.StatusNotFound, res.Code)
	resData := map[string]string{}
	jsonErr := json.Unmarshal(res.Body.Bytes(), &resData)
	assert.Nil(t, jsonErr)
}

func TestFindByIdInvalidHandler(t *testing.T) {
	searchedId := "d"
	r := createSectionServer()
	req, res := tests.CreateRequestTest(http.MethodGet, fmt.Sprintf("/section/%s", searchedId), mocks.MockNuevaSectionRequest)
	r.ServeHTTP(res, req)
	assert.Equal(t, http.StatusBadRequest, res.Code)
	resData := map[string]string{}
	jsonErr := json.Unmarshal(res.Body.Bytes(), &resData)
	assert.Nil(t, jsonErr)
}

func TestFindByIdHandler(t *testing.T) {
	expectedData := mocks.MockBuscarSection
	searchedId := 2
	r := createSectionServer()
	req, res := tests.CreateRequestTest(http.MethodGet, fmt.Sprintf("/section/%d", searchedId), mocks.MockNuevaSectionRequest)
	r.ServeHTTP(res, req)
	assert.Equal(t, http.StatusOK, res.Code)
	resData := map[string]domain.Section{}
	jsonErr := json.Unmarshal(res.Body.Bytes(), &resData)
	assert.Nil(t, jsonErr)
	assert.Equal(t, expectedData, resData["data"])
}

func TestUpdateExistantHandler(t *testing.T) {
	expectedData := mocks.MockActualizarSection
	searchedId := 2
	r := createSectionServer()
	req, res := tests.CreateRequestTest(http.MethodPatch, fmt.Sprintf("/section/%d", searchedId), mocks.MockActualizarSectionRequest)
	r.ServeHTTP(res, req)
	assert.Equal(t, http.StatusOK, res.Code)
	resData := map[string]domain.Section{}
	jsonErr := json.Unmarshal(res.Body.Bytes(), &resData)
	assert.Nil(t, jsonErr)
	assert.Equal(t, expectedData, resData["data"])
}

func TestUpdateInvalidHandler(t *testing.T) {
	searchedId := "d"
	r := createSectionServer()
	req, res := tests.CreateRequestTest(http.MethodPatch, fmt.Sprintf("/section/%s", searchedId), mocks.MockActualizarSectionRequest)
	r.ServeHTTP(res, req)
	assert.Equal(t, http.StatusBadRequest, res.Code)
	resData := map[string]string{}
	jsonErr := json.Unmarshal(res.Body.Bytes(), &resData)
	assert.Nil(t, jsonErr)
}

func TestUpdateNonExistantHandler(t *testing.T) {
	searchedId := 10
	r := createSectionServer()
	req, res := tests.CreateRequestTest(http.MethodPatch, fmt.Sprintf("/section/%d", searchedId), mocks.MockActualizarSectionRequest)
	r.ServeHTTP(res, req)
	assert.Equal(t, http.StatusNotFound, res.Code)
	resData := map[string]string{}
	jsonErr := json.Unmarshal(res.Body.Bytes(), &resData)
	assert.Nil(t, jsonErr)
}

func TestUpdateInvalidJSONHandler(t *testing.T) {
	searchedId := 2
	r := createSectionServer()
	req, res := tests.CreateRequestTest(http.MethodPatch, fmt.Sprintf("/section/%d", searchedId), []byte("{invalid}"))
	r.ServeHTTP(res, req)
	assert.Equal(t, http.StatusBadRequest, res.Code)
	resData := map[string]string{}
	jsonErr := json.Unmarshal(res.Body.Bytes(), &resData)
	assert.Nil(t, jsonErr)
}

func TestUpdateErrorHandler(t *testing.T) {
	searchedId := 2
	r := createSectionErrorServer()
	req, res := tests.CreateRequestTest(http.MethodPatch, fmt.Sprintf("/section/%d", searchedId), mocks.MockActualizarSectionRequest)
	r.ServeHTTP(res, req)
	assert.Equal(t, http.StatusInternalServerError, res.Code)
	resData := map[string]string{}
	jsonErr := json.Unmarshal(res.Body.Bytes(), &resData)
	assert.Nil(t, jsonErr)
}

func TestDeleteNonExistantHandler(t *testing.T) {
	searchedId := 10
	r := createSectionServer()
	req, res := tests.CreateRequestTest(http.MethodDelete, fmt.Sprintf("/section/%d", searchedId), mocks.MockActualizarSectionRequest)
	r.ServeHTTP(res, req)
	assert.Equal(t, http.StatusNotFound, res.Code)
	resData := map[string]string{}
	jsonErr := json.Unmarshal(res.Body.Bytes(), &resData)
	assert.Nil(t, jsonErr)
}

func TestDeleteInvalidHandler(t *testing.T) {
	searchedId := "d"
	r := createSectionServer()
	req, res := tests.CreateRequestTest(http.MethodDelete, fmt.Sprintf("/section/%s", searchedId), mocks.MockActualizarSectionRequest)
	r.ServeHTTP(res, req)
	assert.Equal(t, http.StatusBadRequest, res.Code)
	resData := map[string]string{}
	jsonErr := json.Unmarshal(res.Body.Bytes(), &resData)
	assert.Nil(t, jsonErr)
}

func TestDeleteErrorHandler(t *testing.T) {
	searchedId := 2
	r := createSectionErrorServer()
	req, res := tests.CreateRequestTest(http.MethodDelete, fmt.Sprintf("/section/%d", searchedId), mocks.MockActualizarSectionRequest)
	r.ServeHTTP(res, req)
	assert.Equal(t, http.StatusInternalServerError, res.Code)
	resData := map[string]string{}
	jsonErr := json.Unmarshal(res.Body.Bytes(), &resData)
	assert.Nil(t, jsonErr)
}

func TestDeleteHandler(t *testing.T) {
	searchedId := 2
	r := createSectionServer()
	req, res := tests.CreateRequestTest(http.MethodDelete, fmt.Sprintf("/section/%d", searchedId), mocks.MockActualizarSectionRequest)
	r.ServeHTTP(res, req)
	assert.Equal(t, http.StatusNoContent, res.Code)
}
