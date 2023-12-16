package handler

import (
	"github.com/extmatperez/meli_bootcamp_go_w5-5/internal/domain"
	purchaseOrder "github.com/extmatperez/meli_bootcamp_go_w5-5/internal/purchase_orders"
	"github.com/extmatperez/meli_bootcamp_go_w5-5/pkg/web"
	"github.com/gin-gonic/gin"
)

type requestPurchaseOrders struct {
	OrderNumber     string `json:"order_number" binding:"required"`
	OrderDate       string `json:"order_date" binding:"required"`
	TrackingCode    string `json:"tracking_code" binding:"required"`
	BuyerId         int    `json:"buyer_id" binding:"required"`
	ProductRecordId int    `json:"product_record_id" binding:"required"`
	OrderStatusId   int    `json:"order_status_id" binding:"required"`
}

type PurchaseOrder struct {
	purchaseOrderService purchaseOrder.Service
}

func NewPurchaseOrder(po purchaseOrder.Service) *PurchaseOrder {
	return &PurchaseOrder{
		purchaseOrderService: po,
	}
}

//Create a Purchase Order
//@Summary Create a purchase order in the list of them
//@Tags Purchase Order
//@description Create a purchase order indicating its parameters.
//@Accept json
//@Produce json
//@Param buyer body requestPurchaseOrders true "Create a Purchase Order"
//@Success 201 {object} web.response
//@Failed 400 {object} web.errorResponse
//@Router /purchaseOrders [post]
func (po *PurchaseOrder) Create() gin.HandlerFunc {
	return func(c *gin.Context) {

		var req requestPurchaseOrders
		if err := c.ShouldBindJSON(&req); err != nil {
			web.Error(c, 422, "error: JSON keys required are not included.")
			return
		}

		exist, err := po.purchaseOrderService.Exists(c, req.OrderNumber)
		if err != nil {
			web.Error(c, 500, err.Error())
			return
		}
		if exist {
			web.Error(c, 409, "error: purchase order with order_number:%s already exist", req.OrderNumber)
			return
		}

		purchaseOrder := domain.PurchaseOrders{
			OrderNumber:     req.OrderNumber,
			OrderDate:       req.OrderDate,
			TrackingCode:    req.TrackingCode,
			BuyerId:         req.BuyerId,
			ProductRecordId: req.ProductRecordId,
			OrderStatusId:   req.OrderStatusId,
		}

		id, err := po.purchaseOrderService.Save(c, purchaseOrder)
		if err != nil {
			web.Error(c, 400, err.Error())
			return
		}
		purchaseOrder.ID = id
		web.Success(c, 201, purchaseOrder)
	}
}
