package handler

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/extmatperez/meli_bootcamp_go_w5-5/internal/carry"
	"github.com/extmatperez/meli_bootcamp_go_w5-5/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w5-5/pkg/web"
	"github.com/gin-gonic/gin"
)

type Carry struct {
	carryService carry.Service
}

func NewCarry(c carry.Service) *Carry {
	return &Carry{
		carryService: c,
	}
}

type postRequestCarry struct {
	CID         string `json:"cid" binding:"required"`
	CompanyName string `json:"company_name"`
	Address     string `json:"address"`
	Telephone   string `json:"telephone"`
	LocalityID  int    `json:"locality_id" binding:"required"`
}

// CreateCarry godoc
// @Summary Create Carry
// @Tags Carries
// @Description create Carries
// @Accept json
// @Produce json
// @Param carry body postRequestCarry true "Carry to store"
// @Success 201 {object} web.response
// @Failure 400 {object} web.errorResponse
// @Failure 409 {object} web.errorResponse
// @Failure 422 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /carries [post]
func (c *Carry) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req postRequestCarry

		// Obtengo el Request del body y si hay error lo retorno
		if err := ctx.ShouldBindJSON(&req); err != nil {
			// Si falta un campo requerido retorno un 422
			if strings.Contains(err.Error(), "'required' tag") {
				web.Error(ctx, http.StatusUnprocessableEntity, err.Error())
				return
			}
			// Cualquier otro error en el body retorno un 400
			web.Error(ctx, http.StatusBadRequest, err.Error())
			return
		}

		// Seteo el Carry con el request
		carry := domain.Carry{
			CID:         req.CID,
			CompanyName: req.CompanyName,
			Address:     req.Address,
			Telephone:   req.Telephone,
			LocalityID:  req.LocalityID,
		}

		// Guardo el Carry en la BBDD
		id, err := c.carryService.Save(ctx, carry)

		// Retorno si hay error
		if err != nil {
			// Si es un error de existencia en BBDD retorno un 409
			if strings.Contains(err.Error(), "exists") {
				web.Error(ctx, http.StatusConflict, err.Error())
				return
			}
			// Otro error de BBDD retorno un 500 y muestro por consola el error.
			web.Error(ctx, http.StatusInternalServerError, "internal server error")
			fmt.Printf("[SERVER INFO] error in carryService.Save: %v\n", err.Error())
			return
		}

		// Agrego el id generado y lo retorno
		carry.ID = id
		web.Success(ctx, http.StatusCreated, carry)
	}
}
