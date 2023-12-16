package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/iofabela/meli_bootcamp_go_w5-5/internal/domain"
	"github.com/iofabela/meli_bootcamp_go_w5-5/internal/locality"
	"github.com/iofabela/meli_bootcamp_go_w5-5/pkg/web"
)

type RequestLocalityPost struct {
	ID           int    `json:"id" binding:"required"`
	LocalityName string `json:"locality_name" binding:"required"`
	ProvinceName string `json:"province_name" binding:"required"`
	CountryName  string `json:"country_name" binding:"required"`
}

type RequestSellerPost struct {
	CID         int    `json:"cid" binding:"required"`
	CompanyName string `json:"company_name" binding:"required"`
	Address     string `json:"address" binding:"required"`
	Telephone   string `json:"telephone" binding:"required"`
	LocalityId  int    `json:"locality_id" binding:"required"`
}

type Locality struct {
	localityService locality.Service
}

func NewLocality(l locality.Service) *Locality {
	return &Locality{
		localityService: l,
	}
}

// CreateLocality godoc
// @Summary Create Locality
// @Tags Localities
// @Description Create Locality
// @Accept json
// @Produce json
// @Param Locality body RequestLocalityPost true "Locality to store"
// @Success 201 {object} web.response
// @Failure 409 {object} web.errorResponse
// @Failure 400 {object} web.errorResponse
// @Failure 422 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /localities [post]
func (l *Locality) CreateLocality() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req RequestLocalityPost

		if err := ctx.ShouldBindJSON(&req); err != nil {
			if strings.Contains(err.Error(), "'required' tag") {
				web.Error(ctx, 422, err.Error())
				return
			}
			web.Error(ctx, 400, err.Error())
			return
		}

		locality := domain.Locality{
			ID:           req.ID,
			LocalityName: req.LocalityName,
			ProvinceName: req.ProvinceName,
			CountryName:  req.CountryName,
		}
		//save locality in DB
		id, err := l.localityService.SaveLocality(ctx, locality)

		//return error
		if err != nil {

			if strings.Contains(err.Error(), "exists") {
				web.Error(ctx, 409, err.Error())
				return
			}

			web.Error(ctx, 500, "internal server error")
			return
		}
		//return and add id
		locality.ID = id
		web.Success(ctx, 201, locality)

	}
}

// ReportSellers godoc
// @Summary Get Report of Sellers by Locality
// @Tags Localities
// @Description Get Report of Sellers by Locality
// @Produce json
// @Param id query int false "locality id"
// @Success 200 {object} web.response
// @Failure 400 {object} web.errorResponse
// @Failure 404 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /localities/reportSellers [get]
func (l *Locality) ReportSellers() gin.HandlerFunc {
	return func(c *gin.Context) {

		//I get the query ID and verify that it exists
		stringId, containsId := c.GetQuery("id")

		if containsId {
			id, err := strconv.Atoi(stringId)
			if err != nil {
				web.Error(c, 400, "%s", "id must be integer")
				return
			}
			lc, err := l.localityService.SellerReport(c, id)
			if err != nil {
				if err.Error() == "seller not found" {
					web.Error(c, 404, err.Error())
					return
				}
				web.Error(c, 500, "internal server error")
				return
			}
			web.Success(c, 200, lc)
			return
		}

		lcs, err := l.localityService.GetAllSellerReports(c)

		if err != nil {
			web.Error(c, 500, "internal Server Error")
			return
		}

		//return all reports
		web.Success(c, 200, lcs)
	}
}

// ReportCarries godoc
// @Summary Get Report of Carries by Locality
// @Tags Localities
// @Description Get Report of Carries by Locality
// @Produce json
// @Param id query int false "locality id"
// @Success 200 {object} web.response
// @Failure 400 {object} web.errorResponse
// @Failure 404 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /localities/reportCarries [get]
func (l *Locality) ReportCarries() gin.HandlerFunc {
	return func(c *gin.Context) {
		//Obtengo el query ID y verifico que exista
		stringId, containsId := c.GetQuery("id")

		if containsId {
			// Lo convierto a entero y si hay error retorno un 400.
			id, err := strconv.Atoi(stringId)
			if err != nil {
				web.Error(c, http.StatusBadRequest, "%s", "id must be integer")
				return
			}
			// Busco el reporte
			lc, err := l.localityService.GetCarryReport(c, id)
			if err != nil {
				// Si hay un error "not found" retorno un 404
				if err.Error() == "locality not found" {
					web.Error(c, http.StatusNotFound, err.Error())
					return
				}
				// Otro error de BBDD retorno un 500 y muestro por consola el error
				web.Error(c, http.StatusInternalServerError, "internal server error")
				fmt.Printf("[SERVER INFO] error in localityService.GetCarryReport: %v\n", err.Error())
				return
			}
			// Retorno el reporte encontrado en caso de exito
			web.Success(c, http.StatusOK, lc)
			return
		}

		// Si el id no existe obtengo todos los reportes
		lcs, err := l.localityService.GetAllCarryReports(c)

		// En caso de error de BBDD retorno un 500 y muestro por consola el error
		if err != nil {
			web.Error(c, http.StatusInternalServerError, "internal Server Error")
			fmt.Printf("[SERVER INFO] error in localityService.GetAllCarryReports: %v\n", err.Error())
			return
		}
		// Retorno todos los reportes encontrados
		web.Success(c, http.StatusOK, lcs)
	}
}
