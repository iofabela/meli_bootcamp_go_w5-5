package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/iofabela/meli_bootcamp_go_w5-5/cmd/server/handler"
	"github.com/iofabela/meli_bootcamp_go_w5-5/docs"
	"github.com/iofabela/meli_bootcamp_go_w5-5/internal/buyer"
	"github.com/iofabela/meli_bootcamp_go_w5-5/internal/carry"
	"github.com/iofabela/meli_bootcamp_go_w5-5/internal/employee"
	inboundorder "github.com/iofabela/meli_bootcamp_go_w5-5/internal/inbound_order"
	"github.com/iofabela/meli_bootcamp_go_w5-5/internal/locality"
	"github.com/iofabela/meli_bootcamp_go_w5-5/internal/product"
	product_batch "github.com/iofabela/meli_bootcamp_go_w5-5/internal/product_batches"
	purchaseOrder "github.com/iofabela/meli_bootcamp_go_w5-5/internal/purchase_orders"
	"github.com/iofabela/meli_bootcamp_go_w5-5/internal/section"
	"github.com/iofabela/meli_bootcamp_go_w5-5/internal/seller"
	"github.com/iofabela/meli_bootcamp_go_w5-5/internal/warehouse"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Router interface {
	MapRoutes()
}

type router struct {
	r  *gin.Engine
	rg *gin.RouterGroup
	db *sql.DB
}

func NewRouter(r *gin.Engine, db *sql.DB) Router {
	return &router{r: r, db: db}
}

func (r *router) MapRoutes() {
	r.setGroup()

	r.buildSellerRoutes()
	r.buildProductRoutes()
	r.buildSectionRoutes()
	r.buildWarehouseRoutes()
	r.buildEmployeeRoutes()
	r.buildInboundOrderRoutes()
	r.buildBuyerRoutes()
	r.buildProductBatchRoutes()
	r.buildPurchaseOrdersRoutes()
	r.buildSwaggerRoutes()
	r.buildLocalitiesRoutes()
	r.buildCarryRoutes()
}

func (r *router) setGroup() {
	r.rg = r.r.Group("/api/v1")
}

func (r *router) buildSellerRoutes() {
	// Example
	repo := seller.NewRepository(r.db)
	service := seller.NewService(repo)
	handler := handler.NewSeller(service)
	sellerRoutes := r.rg.Group("/sellers")
	{
		sellerRoutes.GET("/", handler.GetAll())
		sellerRoutes.GET("/:id", handler.Get())
		sellerRoutes.POST("/", handler.Create())
		sellerRoutes.PATCH("/:id", handler.Update())
		sellerRoutes.DELETE("/:id", handler.Delete())
	}
}

func (r *router) buildProductRoutes() {
	repo := product.NewRepository(r.db)
	service := product.NewService(repo)
	handler := handler.NewProduct(service)
	prdRoutes := r.rg.Group("/products")
	{
		prdRoutes.GET("/", handler.GetAll())
		prdRoutes.GET("/:id", handler.Get())
		prdRoutes.POST("/", handler.Create())
		prdRoutes.PATCH("/:id", handler.Update())
		prdRoutes.DELETE("/:id", handler.Delete())
	}
}

func (r *router) buildSectionRoutes() {
	repo := section.NewRepository(r.db)
	service := section.NewService(repo)
	handler := handler.NewSection(service)
	section := r.rg.Group("/sections")
	{
		section.GET("/", handler.GetAll())
		section.GET("/:id", handler.Get())
		section.POST("/", handler.Create())
		section.PATCH("/:id", handler.Update())
		section.DELETE("/:id", handler.Delete())
		section.GET("/reportProducts", handler.GetProductReport())
	}

}

func (r *router) buildWarehouseRoutes() {
	repo := warehouse.NewRepository(r.db)
	service := warehouse.NewService(repo)
	handler := handler.NewWarehouse(service)
	whRoutes := r.rg.Group("/warehouses")
	{
		whRoutes.GET("/", handler.GetAll())
		whRoutes.POST("/", handler.Create())
		whRoutes.GET("/:id", handler.Get())
		whRoutes.PATCH("/:id", handler.Update())
		whRoutes.DELETE("/:id", handler.Delete())
	}
}

func (r *router) buildEmployeeRoutes() {
	repo := employee.NewRepository(r.db)
	service := employee.NewService(repo)
	handler := handler.NewEmployee(service)
	employeeRoutes := r.rg.Group("/employees")
	{
		employeeRoutes.POST("/", handler.Create())
		employeeRoutes.GET("/", handler.GetAll())
		employeeRoutes.GET("/:id", handler.Get())
		employeeRoutes.GET("/reportInboundOrders", handler.GetInboundOrders())
		employeeRoutes.PATCH("/:id", handler.Update())
		employeeRoutes.DELETE("/:id", handler.Delete())
	}
}

func (r *router) buildInboundOrderRoutes() {

	repo := inboundorder.NewRepository(r.db)
	service := inboundorder.NewService(repo)
	handler := handler.NewInboundOrder(service)
	inboundOrdersRoutes := r.rg.Group("/inboundOrders")

	inboundOrdersRoutes.POST("/", handler.Create())
}

func (r *router) buildBuyerRoutes() {

	repo := buyer.NewRepository(r.db)
	service := buyer.NewService(repo)
	handler := handler.NewBuyer(service)
	buyersRoutes := r.rg.Group("/buyers")

	docs.SwaggerInfo.Host = "localhost:8080"
	r.r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	buyersRoutes.GET("/", handler.GetAll())
	buyersRoutes.GET("/:id", handler.Get())
	buyersRoutes.POST("/", handler.Create())
	buyersRoutes.PATCH("/:id", handler.Update())
	buyersRoutes.DELETE("/:id", handler.Delete())
	buyersRoutes.GET("/reportPurchaseOrders", handler.PurchaseOrders())
}

func (r *router) buildPurchaseOrdersRoutes() {
	repo := purchaseOrder.NewRepository(r.db)
	service := purchaseOrder.NewService(repo)
	handler := handler.NewPurchaseOrder(service)
	purchaseOrderRoutes := r.rg.Group("/purchaseOrders")
	{
		purchaseOrderRoutes.POST("/", handler.Create())
	}
}

func (r *router) buildLocalitiesRoutes() {

	repo := locality.NewRepository(r.db)
	service := locality.NewService(repo)
	handler := handler.NewLocality(service)

	localitiesRoutes := r.rg.Group("/localities")
	{
		localitiesRoutes.POST("/", handler.CreateLocality())
		localitiesRoutes.GET("/reportSellers", handler.ReportSellers())
		localitiesRoutes.GET("/reportCarries", handler.ReportCarries())
	}
}

func (r *router) buildCarryRoutes() {

	repo := carry.NewRepository(r.db)
	service := carry.NewService(repo)
	handler := handler.NewCarry(service)
	carrieRoutes := r.rg.Group("/carries")
	{
		carrieRoutes.POST("/", handler.Create())
	}
}

func (r *router) buildProductBatchRoutes() {
	repo := product_batch.NewRepository(r.db)
	service := product_batch.NewService(repo)
	handler := handler.NewProductBatch(service)
	section := r.rg.Group("/productBatches")
	{
		section.POST("/", handler.Create())
	}

}

func (r *router) buildSwaggerRoutes() {
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/api/v1"
	r.rg.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
