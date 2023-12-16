package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iofabela/meli_bootcamp_go_w5-5/internal/domain"
	product_batch "github.com/iofabela/meli_bootcamp_go_w5-5/internal/product_batches"
	"github.com/iofabela/meli_bootcamp_go_w5-5/pkg/web"
)

var (
	EMPTY_FIELD = "field %s cant be empty"
	ZERO_FIELD  = "field %s cant be empty or zero"
)

type ProductBatch struct {
	product_batch_service product_batch.Service
}

func NewProductBatch(pb product_batch.Service) *ProductBatch {
	return &ProductBatch{
		product_batch_service: pb,
	}
}

type requestBatches struct {
	BatchNumber        *int
	CurrentQuantity    *int
	CurrentTemperature *int
	DueDate            *string
	InitialQuantity    *int
	ManufacturingDate  *string
	ManufacturingHour  *string
	MinumumTemperature *int
	ProductId          *int
	SectionId          *int
}

func (rq *requestBatches) Parse() domain.ProductBatches {
	return domain.ProductBatches{
		BatchNumber:        *rq.BatchNumber,
		CurrentQuantity:    *rq.CurrentQuantity,
		DueDate:            *rq.DueDate,
		InitialQuantity:    *rq.InitialQuantity,
		ManufacturingDate:  *rq.ManufacturingDate,
		ManufacturingHour:  *rq.ManufacturingHour,
		MinumumTemperature: *rq.MinumumTemperature,
		ProductId:          *rq.ProductId,
		SectionId:          *rq.SectionId,
	}
}

func ValidateBatch(validated requestBatches) error {
	if validated.BatchNumber == nil {
		return fmt.Errorf(ZERO_FIELD, "batch_number")
	}
	if validated.CurrentQuantity == nil {
		return fmt.Errorf(ZERO_FIELD, "current_quantity")
	}
	if validated.DueDate == nil {
		return fmt.Errorf(EMPTY_FIELD, "due_date")
	}
	if validated.InitialQuantity == nil {
		return fmt.Errorf(ZERO_FIELD, "initial_quantity")
	}
	if validated.ManufacturingHour == nil {
		return fmt.Errorf(ZERO_FIELD, "manufacturing_hour")
	}
	if validated.MinumumTemperature == nil {
		return fmt.Errorf(ZERO_FIELD, "minimun_temperature")
	}
	if validated.ProductId == nil {
		return fmt.Errorf(ZERO_FIELD, "product_id")
	}
	if validated.SectionId == nil {
		return fmt.Errorf(ZERO_FIELD, "section_id")
	}
	return nil
}

// CreateProductBatch godoc
// @Summary Create product batch
// @Tags ProductBatch
// @Description create a new product batch
// @Accept json
// @Produce  json
// @Param Batch body requestBatches true "New Batch"
// @Success 201 {object} web.response
// @Success 409 {object} web.errorResponse
// @Success 409 {object} web.errorResponse
// @Router /productBatches/ [post]
func (s *ProductBatch) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var product_batch requestBatches
		err := c.ShouldBindJSON(&product_batch)
		if err != nil {
			web.Error(c, http.StatusBadRequest, err.Error())
			return
		}
		err = ValidateBatch(product_batch)
		if err != nil {
			web.Error(c, http.StatusUnprocessableEntity, err.Error())
			return
		}
		new_batch := product_batch.Parse()
		createdInt, err := s.product_batch_service.Save(c, new_batch)
		new_batch.Id = createdInt
		if err != nil {
			web.Error(c, http.StatusInternalServerError, CantSave)
			return
		}
		web.Success(c, http.StatusCreated, new_batch)
	}
}
