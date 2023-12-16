package handler

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/iofabela/meli_bootcamp_go_w5-5/internal/domain"
	"github.com/iofabela/meli_bootcamp_go_w5-5/internal/seller"
	"github.com/iofabela/meli_bootcamp_go_w5-5/pkg/web"

	"github.com/gin-gonic/gin"
)

type Seller struct {
	sellerService seller.Service
}

type postSeller struct {
	CID         int    `json:"cid" binding:"required"`
	CompanyName string `json:"company_name" binding:"required"`
	Address     string `json:"address" binding:"required"`
	Telephone   string `json:"telephone" binding:"required"`
}

type patchSeller struct {
	CID         *int   `json:"cid"`
	CompanyName string `json:"company_name"`
	Address     string `json:"address"`
	Telephone   string `json:"telephone"`
}

func NewSeller(s seller.Service) *Seller {
	return &Seller{sellerService: s}
}

// GetAllSeller godoc
// @Summary Get all Sellers
// @Tags Sellers
// @Description get all Sellers
// @Produce json
// @Success 200 {object} web.response
// @Failure 404 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /sellers [get]
func (s *Seller) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		p, err := s.sellerService.GetAll(c)

		if err != nil {
			web.Error(c, 500, "internal server error")
			return
		}

		if len(p) == 0 {
			web.Success(c, 200, []domain.Seller{})
			return
		}
		web.Success(c, 200, p)
	}
}

// GetSellers godoc
// @Summary Get sellers by id
// @Tags Sellers
// @Description get sellers by id
// @Produce json
// @Success 200 {object} web.response
// @Failure 400 {object} web.errorResponse
// @Failure 404 {object} web.errorResponse
// @Param id path int true "sellers id"
// @Router /sellers/{id} [get]
func (s *Seller) Get() gin.HandlerFunc {
	return func(c *gin.Context) {
		p, errCode, err := getSellerByParamID(s, c)
		if err != nil {
			web.Error(c, errCode, err.Error())
			return
		}
		web.Success(c, 200, p)
	}
}

// CreateSellers godoc
// @Summary Create Sellers
// @Tags Sellers
// @Description create Sellers
// @Accept json
// @Produce json
// @Param Sellers body postSeller true "Sellers to store"
// @Success 201 {object} web.response
// @Failure 409 {object} web.errorResponse
// @Failure 400 {object} web.errorResponse
// @Router /sellers [post]
func (s *Seller) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req RequestSellerPost

		if err := ctx.ShouldBindJSON(&req); err != nil {
			if strings.Contains(err.Error(), "'required' tag") {
				web.Error(ctx, 422, err.Error())
				return
			}
			web.Error(ctx, 400, err.Error())
			return
		}

		seller := domain.Seller{
			CID:         req.CID,
			CompanyName: req.CompanyName,
			Address:     req.Address,
			Telephone:   req.Telephone,
			LocalityId:  req.LocalityId,
		}
		//save seller in DB
		id, err := s.sellerService.Save(ctx, seller)
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
		seller.ID = id
		web.Success(ctx, 201, seller)
	}
}

// UpadteSellers godoc
// @Summary Update Sellers
// @Tags Sellers
// @Description update Sellers by id
// @Accept json
// @Produce json
// @Param id path int true "sellers id"
// @Param sellers body patchSeller true "sellers to update"
// @Success 200 {object} web.response
// @Failure 400 {object} web.errorResponse
// @Failure 404 {object} web.errorResponse
// @Failure 409 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /sellers/{id} [patch]
func (s *Seller) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req patchSeller

		id, err := strconv.ParseInt(c.Param("id"), 10, 64)

		if err != nil {
			web.Error(c, 400, "invalid id, must be integer")
			return
		}

		se, err := s.sellerService.Get(c, int(id))
		if err != nil {
			web.Error(c, 404, "no seller with the id was found %d", id)
			return
		}

		if err := c.ShouldBindJSON(&req); err != nil {
			web.Error(c, 404, err.Error())
			return
		}
		if req.CID != nil {
			if *req.CID != se.CID {
				if s.sellerService.Exists(c, *req.CID) {
					web.Error(c, 409, "there is already a seller with that cid")
					return
				}
				se.CID = *req.CID
			}
		}

		if req.CompanyName != "" {
			se.CompanyName = req.CompanyName
		}
		if req.Address != "" {
			se.Address = req.Address
		}
		if req.Telephone != "" {
			se.Telephone = req.Telephone
		}

		if err := s.sellerService.Update(c, se); err != nil {
			web.Error(c, 500, "internal server error")
			fmt.Printf("Error: %v\n", err.Error())
			return
		}

		web.Success(c, 200, se)

	}
}

// DeleteSeller godoc
// @Summary Delete Seller by id
// @Tags Sellers
// @Description delete Seller by id
// @Param id path int true "Seller id"
// @Success 204 {object} nil
// @Failure 400 {object} web.errorResponse
// @Failure 404 {object} web.errorResponse
// @Router /sellers/{id} [delete]
func (s *Seller) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {

		id, err := strconv.ParseInt(c.Param("id"), 10, 64)

		if err != nil {
			web.Error(c, 400, "invalid id, must be integer")
			return
		}

		err = s.sellerService.Delete(c, int(id))
		if err != nil {
			web.Error(c, 404, "no seller with the id was found %d", id)
			return
		}

		web.Success(c, 204, nil)
	}
}

func getSellerByParamID(s *Seller, c *gin.Context) (domain.Seller, int, error) {
	// Convierto id en entero y en caso de error lo retorno
	se := domain.Seller{}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return se, 400, fmt.Errorf("invalid id, must be integer")
	}

	// Le pido al service el wh y si no existe retorno un 404
	se, err = s.sellerService.Get(c, id)
	if err != nil {
		return se, 404, fmt.Errorf("no seller with the id was found %d", id)
	}

	return se, 0, nil
}
