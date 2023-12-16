package handler

import (
	"github.com/extmatperez/meli_bootcamp_go_w5-5/internal/domain"
	inboundorder "github.com/extmatperez/meli_bootcamp_go_w5-5/internal/inbound_order"
	"github.com/extmatperez/meli_bootcamp_go_w5-5/pkg/web"
	"github.com/gin-gonic/gin"
)

// InboundOrder structure
type InboundOrder struct {
	inboundOrderService inboundorder.Service
}

type postInboundOrder struct {
	OrderDate      string `json:"order_date"`
	OrderNumber    string `json:"order_number"`
	EmployeeID     int    `json:"employee_id"`
	ProductBatchID int    `json:"product_batch_id"`
	WarehouseID    int    `json:"warehouse_id"`
}

// NewInboundOrder
func NewInboundOrder(i inboundorder.Service) *InboundOrder {
	return &InboundOrder{
		inboundOrderService: i,
	}
}

//Create godoc
//@Summary Create inboundOrder
//@Tags InboundOrders
//@Description Create inboundOrdes
//@Accept json
//@Produce json
//@Param InboundOrders body postInboundOrder true "InboundOrders to store"
//@Succes 201 {object} web.Response
//@Failure 409 {object} web.errorResponse
//@Failure 422 {object} web.errorResponse
//@Router /inboundOrders [post]
func (i *InboundOrder) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req postInboundOrder

		if err := c.ShouldBindJSON(&req); err != nil {
			web.Error(c, 400, err.Error())
			return
		}
		if req.OrderDate == "" {
			web.Error(c, 422, "El OrderDate es requerido")
			return
		}
		if req.OrderNumber == "" {
			web.Error(c, 422, "El OrderNumber es requerido")
			return
		}
		if req.EmployeeID == 0 {
			web.Error(c, 422, "El EmployeeID es requerido")
			return
		}
		if req.ProductBatchID == 0 {
			web.Error(c, 422, "El ProductBatchID es requerido")
			return
		}
		if req.WarehouseID == 0 {
			web.Error(c, 422, "El WarehouseID es requerido")
			return
		}
		exist, err := i.inboundOrderService.Exists(c, req.EmployeeID)
		if err != nil {
			web.Error(c, 500, err.Error())
			return
		}
		if !exist {
			web.Error(c, 409, "El employee no existe")
			return
		}

		inbOrd := domain.InboundOrder{
			OrderDate:      req.OrderDate,
			OrderNumber:    req.OrderNumber,
			EmployeeID:     req.EmployeeID,
			ProductBatchID: req.ProductBatchID,
			WarehouseID:    req.WarehouseID,
		}

		id, err := i.inboundOrderService.Save(c, inbOrd)
		if err != nil {
			web.Error(c, 500, err.Error())
			return
		}

		inbOrd.ID = id
		web.Success(c, 201, inbOrd)
	}
}
