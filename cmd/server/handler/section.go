package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/extmatperez/meli_bootcamp_go_w5-5/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w5-5/internal/section"
	"github.com/extmatperez/meli_bootcamp_go_w5-5/pkg/web"
	"github.com/gin-gonic/gin"
)

var (
	EmptyField     = "field %s cant be empty "
	CantFind       = "cant find section: %d"
	InvalidId      = "ID given isnt valid"
	CantGet        = "Cant get sections"
	CantGetReports = "Cant get product reports"
	CantSave       = "Cant save section"
	CantUpdate     = "Cant update section: %d"
	CantDelete     = "Cant delete section: %d"
)

type Section struct {
	sectionService section.Service
}

func NewSection(s section.Service) *Section {
	return &Section{
		sectionService: s,
	}
}

type request struct {
	SectionNumber      *int `json:"section_number"`
	CurrentTemperature *int `json:"current_temperature"`
	MinimumTemperature *int `json:"minimum_temperature"`
	CurrentCapacity    *int `json:"current_capacity"`
	MinimumCapacity    *int `json:"minimum_capacity"`
	MaximumCapacity    *int `json:"maximum_capacity"`
	WarehouseID        *int `json:"warehouse_id"`
	ProductTypeID      *int `json:"product_type_id"`
}

// ListSections godoc
// @Summary List sections
// @Tags Sections
// @Description get all registered sections
// @Produce  json
// @Success 200 {object} web.response
// @Router /sections [get]
func (s *Section) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		allSections, err := s.sectionService.GetAll(c)
		if err != nil {
			web.Error(c, http.StatusInternalServerError, CantGet)
			return
		}
		if len(allSections) == 0 {
			web.Success(c, http.StatusOK, []domain.Section{})
			return
		}
		web.Success(c, http.StatusOK, allSections)
	}
}

// GetSection godoc
// @Summary Get section
// @Tags Sections
// @Description search a section given a valid id
// @Produce  json
// @Param id path int true "id"
// @Success 200 {object} web.response
// @Router /sections/{id} [get]
func (s *Section) Get() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			web.Error(c, http.StatusBadRequest, InvalidId)
			return
		}
		section, err := s.sectionService.Get(c, id)
		if err != nil {
			web.Error(c, http.StatusNotFound, CantFind, id)
			return
		}
		web.Success(c, http.StatusOK, section)
	}
}

// CreateSecion godoc
// @Summary Create section
// @Tags Sections
// @Description create a new section
// @Accept json
// @Produce  json
// @Param product body request true "New section"
// @Success 201 {object} web.response
// @Router /sections/ [post]
func (s *Section) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var section request
		err := c.ShouldBindJSON(&section)
		if err != nil {
			web.Error(c, http.StatusBadRequest, err.Error())
			return
		}
		errs := Validate(section)
		if len(errs) > 0 {
			web.Error(c, http.StatusUnprocessableEntity, errs[0].Error())
			return
		}
		newSection := domain.Section{
			SectionNumber:      *section.SectionNumber,
			CurrentTemperature: *section.CurrentTemperature,
			MinimumTemperature: *section.MinimumTemperature,
			CurrentCapacity:    *section.CurrentCapacity,
			MinimumCapacity:    *section.MinimumCapacity,
			MaximumCapacity:    *section.MaximumCapacity,
			WarehouseID:        *section.WarehouseID,
			ProductTypeID:      *section.ProductTypeID,
		}
		createdInt, err := s.sectionService.Save(c, newSection)
		newSection.ID = createdInt
		if err != nil {
			web.Error(c, http.StatusInternalServerError, CantSave)
			return
		}
		web.Success(c, http.StatusCreated, newSection)
	}
}

// UpdateSection godoc
// @Summary Update Section
// @Tags Sections
// @Description search a section given a valid id
// @Produce  json
// @Param id path int true "id"
// @Param product body request true "New data for section"
// @Success 200 {object} web.response
// @Router /sections/{id} [patch]
func (s *Section) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			web.Error(c, http.StatusBadRequest, InvalidId)
			return
		}
		oldSection, err := s.sectionService.Get(c, id)
		if err != nil {
			web.Error(c, http.StatusNotFound, CantFind, id)
			return
		}

		var section request
		err = c.ShouldBindJSON(&section)
		if err != nil {
			web.Error(c, http.StatusBadRequest, err.Error())
			return
		}
		updated := partialUpdate(oldSection, section)
		oldSection.ID = id
		upSection, err := s.sectionService.Update(c, updated)
		if err != nil {
			web.Error(c, http.StatusInternalServerError, CantUpdate, id)
			return
		}
		web.Success(c, http.StatusOK, upSection)
	}
}

// DeleteSection godoc
// @Summary Delete section
// @Tags Sections
// @Description delete a section given a valid id
// @Produce  json
// @Param id path int true "id"
// @Success 204
// @Router /sections/{id} [delete]
func (s *Section) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			web.Error(c, http.StatusBadRequest, InvalidId)
			return
		}
		_, err = s.sectionService.Get(c, id)
		if err != nil {
			web.Error(c, http.StatusNotFound, CantFind, id)
			return
		}
		err = s.sectionService.Delete(c, id)
		if err != nil {
			web.Error(c, http.StatusInternalServerError, CantDelete, id)
			return
		}
		web.Success(c, http.StatusNoContent, nil)
	}
}

// GetSection godoc
// @Summary Get Product Reports
// @Tags Sections
// @Description Get all products Reports
// @Produce  json
// @Success 200 {object} web.response
// @Param id query int false "id"
// @Router /sections/reportProducts [get]
func (s *Section) GetProductReport() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam, ok := c.GetQuery("id")
		if !ok {
			reportParam, err := s.sectionService.ReportProductsGetAll(c)
			if err != nil {
				web.Error(c, http.StatusInternalServerError, CantGetReports)
				return
			}
			web.Success(c, http.StatusOK, reportParam)
			return
		}
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Error(c, http.StatusBadRequest, InvalidId)
			return
		}
		report, err := s.sectionService.ReportProductsGet(c, id)
		if err != nil {
			if strings.Contains(err.Error(), "sql: no rows in result set") {
				web.Error(c, http.StatusNotFound, CantFind, id)
				return
			}
			web.Error(c, http.StatusBadRequest, err.Error())
			return
		}
		web.Success(c, http.StatusOK, report)
	}
}

func Validate(section request) []error {
	errors := []error{}
	if section.SectionNumber == nil {
		errors = append(errors, fmt.Errorf(EmptyField, "section_number"))
	}
	if section.CurrentTemperature == nil {
		errors = append(errors, fmt.Errorf(EmptyField, "current_temperature"))
	}
	if section.MinimumTemperature == nil {
		errors = append(errors, fmt.Errorf(EmptyField, "minimum_temperature"))
	}
	if section.CurrentCapacity == nil {
		errors = append(errors, fmt.Errorf(EmptyField, "current_capacity"))
	}
	if section.MinimumCapacity == nil {
		errors = append(errors, fmt.Errorf(EmptyField, "minimum_capacity"))
	}
	if section.MaximumCapacity == nil {
		errors = append(errors, fmt.Errorf(EmptyField, "maximum_capacity"))
	}
	if section.WarehouseID == nil {
		errors = append(errors, fmt.Errorf(EmptyField, "warehouse_id"))
	}
	if section.ProductTypeID == nil {
		errors = append(errors, fmt.Errorf(EmptyField, "product_type_id"))
	}
	return errors
}

func partialUpdate(oldSection domain.Section, newSection request) domain.Section {
	if newSection.SectionNumber != nil {
		oldSection.SectionNumber = *newSection.SectionNumber
	}
	if newSection.CurrentTemperature != nil {
		oldSection.CurrentTemperature = *newSection.CurrentTemperature
	}
	if newSection.MinimumTemperature != nil {
		oldSection.MinimumTemperature = *newSection.MinimumTemperature
	}
	if newSection.CurrentCapacity != nil {
		oldSection.CurrentCapacity = *newSection.CurrentCapacity
	}
	if newSection.MinimumCapacity != nil {
		oldSection.MinimumCapacity = *newSection.MinimumCapacity
	}
	if newSection.MaximumCapacity != nil {
		oldSection.MaximumCapacity = *newSection.MaximumCapacity
	}
	if newSection.WarehouseID != nil {
		oldSection.WarehouseID = *newSection.WarehouseID
	}
	if newSection.ProductTypeID != nil {
		oldSection.ProductTypeID = *newSection.ProductTypeID
	}
	return oldSection
}
