package handler

import (
	"strconv"

	"github.com/extmatperez/meli_bootcamp_go_w5-5/internal/buyer"
	"github.com/extmatperez/meli_bootcamp_go_w5-5/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w5-5/pkg/web"
	"github.com/gin-gonic/gin"
)

type requestBuyer struct {
	CardNumberID string `json:"card_number_id" binding:"required"`
	FirstName    string `json:"first_name" binding:"required"`
	LastName     string `json:"last_name" binding:"required"`
}

type requestToUpdate struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type Buyer struct {
	buyerService buyer.Service
}

func NewBuyer(b buyer.Service) *Buyer {
	return &Buyer{
		buyerService: b,
	}
}

//Get Buyer by id
//@Summary Get buyer by id
//@Tags Buyer
//@Description Get buyer indicating its id.
//@Produce json
//@Param id path string true "id"
//@Success 200 {object} web.response
//@Failed 404 {object} web.errorResponse
//@Router /buyers/{id} [get]
func (b *Buyer) Get() gin.HandlerFunc {
	return func(c *gin.Context) {

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			web.Error(c, 400, "%s", err.Error())
			return
		}

		buyer, err := b.buyerService.Get(c, id)
		if err != nil {
			web.Error(c, 404, "error: buyer with id:%v not found", id)
			return
		}
		web.Success(c, 200, buyer)
	}
}

//List of Buyers
//@Summary Obtain list of buyers.
//@Tags Buyer
//@description Get all buyers.
//@Produce json
//@Success 200 {object} web.response
//@Failed 500 {object} web.errorResponse
//@Router /buyers [get]
func (b *Buyer) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {

		b, err := b.buyerService.GetAll(c)
		if err != nil {
			web.Error(c, 500, err.Error())
			return
		}
		if len(b) == 0 {
			web.Success(c, 200, []domain.Buyer{})
			return
		}
		web.Success(c, 200, b)
	}
}

//Create a Buyer
//@Summary Create a buyer in the list of them
//@Tags Buyer
//@description Create a buyer indicating its parameters.
//@Accept json
//@Produce json
//@Param buyer body requestBuyer true "Create Buyer"
//@Success 201 {object} web.response
//@Failed 400 {object} web.errorResponse
//@Router /buyers [post]
func (b *Buyer) Create() gin.HandlerFunc {
	return func(c *gin.Context) {

		var req requestBuyer
		if err := c.ShouldBindJSON(&req); err != nil {
			web.Error(c, 422, "error: JSON keys required are not included.")
			return
		}

		exist := b.buyerService.Exists(c, req.CardNumberID)
		if exist {
			web.Error(c, 409, "error: buyer with card_number_id:%s already exist", req.CardNumberID)
			return
		}

		buyer := domain.Buyer{
			CardNumberID: req.CardNumberID,
			FirstName:    req.FirstName,
			LastName:     req.LastName,
		}

		id, err := b.buyerService.Save(c, buyer)
		if err != nil {
			web.Error(c, 400, err.Error())
			return
		}
		buyer.ID = id
		web.Success(c, 201, buyer)
	}
}

//Update Buyer Patch
//@Summary Update by id
//@Tags Buyer
//@Description Update buyer modifying only name and lastname parameters.
//@Accept json
//@Produce json
//@Param buyer body requestToUpdate true "Buyer to uptdate"
//@Param id path string true "id"
//@Success 200 {object} web.response
//@Failed 404 {object} web.errorResponse
//@Router /buyers/{id} [patch]
func (b *Buyer) Update() gin.HandlerFunc {
	return func(c *gin.Context) {

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			web.Error(c, 400, "%s", err.Error())
			return
		}

		buyer, err := b.buyerService.Get(c, id)
		if err != nil {
			web.Error(c, 404, "error: buyer with id:%v not found", id)
			return
		}

		var req requestToUpdate
		if err := c.ShouldBindJSON(&req); err != nil {
			web.Error(c, 422, "%s", err.Error())
			return
		}

		if req.FirstName == "" && req.LastName == "" {
			web.Error(c, 400, "error: both keys {first_name, last_name} are empty. At least one of them need be included.")
			return
		}

		if req.FirstName != "" {
			buyer.FirstName = req.FirstName
		}
		if req.LastName != "" {
			buyer.LastName = req.LastName
		}

		error := b.buyerService.Update(c, buyer)
		if error != nil {
			web.Error(c, 404, "error: buyer with id:%v not found", id)
			return
		}

		web.Success(c, 200, buyer)

	}
}

//Delete Buyer
//@Summary Delete buyer by id
//@Tags Buyer
//@Description Delete buyer modifying indicating its id.
//@Produce json
//@Param id path string true "id"
//@Success 204 {object} web.response
//@Failed 404 {object} web.errorResponse
//@Router /buyers/{id} [delete]
func (b *Buyer) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			web.Error(c, 400, "%s", err.Error())
			return
		}

		error := b.buyerService.Delete(c, id)
		if error != nil {
			web.Error(c, 404, "error: buyer with id:%v not found", id)
			return
		}

		web.Success(c, 204, nil)

	}
}

//Get Purchase Orders
//@Summary Get Purchase Orders by id or Get All
//@Tags Buyer
//@Description Get purchase orders by id or get all non indicating an id.
//@Produce json
//@Param id query int false "Purchase Order Id"
//@Success 200 {object} web.response
//@Failed 404 {object} web.errorResponse
//@Router /buyers/reportPurchaseOrders [get]
func (b *Buyer) PurchaseOrders() gin.HandlerFunc {
	return func(c *gin.Context) {

		idQuery, idExist := c.GetQuery("id")

		if idExist {
			id, err := strconv.Atoi(idQuery)
			if err != nil {
				web.Error(c, 400, "%s", err.Error())
				return
			}
			buyersOrders, err := b.buyerService.GetPurchaseOrders(c, id)
			if err != nil {
				web.Error(c, 404, "error: buyer with id:%v not found", id)
				return
			}
			web.Success(c, 200, buyersOrders)
		} else {
			id := 0
			buyersOrders, err := b.buyerService.GetPurchaseOrders(c, id)
			if err != nil {
				web.Error(c, 404, "error: buyer with id:%v not found", id)
				return
			}
			web.Success(c, 200, buyersOrders)

		}

	}
}
