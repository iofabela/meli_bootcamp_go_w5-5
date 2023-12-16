package handler

import (
	"fmt"
	"strconv"

	"github.com/extmatperez/meli_bootcamp_go_w5-5/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w5-5/internal/warehouse"
	"github.com/extmatperez/meli_bootcamp_go_w5-5/pkg/web"
	"github.com/gin-gonic/gin"
)

type Warehouse struct {
	warehouseService warehouse.Service
}

type postRequestWH struct {
	Address            string `json:"address"`
	Telephone          string `json:"telephone"`
	WarehouseCode      string `json:"warehouse_code"`
	MinimumCapacity    int    `json:"minimum_capacity"`
	MinimumTemperature int    `json:"minimum_temperature"`
}

type patchRequestWH struct {
	Address            string `json:"address"`
	Telephone          string `json:"telephone"`
	WarehouseCode      string `json:"warehouse_code"`
	MinimumCapacity    *int   `json:"minimum_capacity"`
	MinimumTemperature *int   `json:"minimum_temperature"`
}

func NewWarehouse(w warehouse.Service) *Warehouse {
	return &Warehouse{
		warehouseService: w,
	}
}

// GetWarehouses godoc
// @Summary Get Warehouse by id
// @Tags Warehouses
// @Description get warehouse by id
// @Produce json
// @Success 200 {object} web.response
// @Failure 400 {object} web.errorResponse
// @Failure 404 {object} web.errorResponse
// @Param id path string true "Warehouse id"
// @Router /warehouses/{id} [get]
func (w *Warehouse) Get() gin.HandlerFunc {
	return func(c *gin.Context) {
		//Obtengo el wh con la funcion auxiliar o retorno un error
		wh, errCode, err := getWHByParamID(w, c)
		if err != nil {
			web.Error(c, errCode, err.Error())
			return
		}
		web.Success(c, 200, wh)
	}
}

// GetWarehouse godoc
// @Summary Get all Warehouses
// @Tags Warehouses
// @Description get all Warehouses
// @Produce json
// @Success 200 {object} web.response
// @Failure 500 {object} web.errorResponse
// @Router /warehouses [get]
func (w *Warehouse) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		//Pido al service todos los Warehouses, si hay error devuelvo un 500
		whs, err := w.warehouseService.GetAll(c)
		if err != nil {
			web.Error(c, 500, "internal server error")
			// fmt.Printf("ERROR GetAll(): %v\n", err.Error())
			return
		}
		// Retorno una lista de WHs o una lista vacia si no hay ninguno en la BBDD.
		if len(whs) == 0 {
			web.Success(c, 200, []domain.Warehouse{})
			return
		}
		web.Success(c, 200, whs)
	}
}

// CreateWarehouses godoc
// @Summary Create Warehouses
// @Tags Warehouses
// @Description create Warehouses
// @Accept json
// @Produce json
// @Param warehouse body postRequestWH true "Warehouse to store"
// @Success 201 {object} web.response
// @Failure 409 {object} web.errorResponse
// @Failure 422 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /warehouses [post]
func (w *Warehouse) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req postRequestWH

		// Obtengo el Request del body y si hay error lo retorno
		if err := c.ShouldBindJSON(&req); err != nil {
			web.Error(c, 422, "%s", err.Error())
			return
		}

		// Comprobación de existencia del campo WarehouseCode
		if req.WarehouseCode == "" {
			web.Error(c, 422, "%s", "field warehouse_code is required")
			return
		}

		// Comprobación de existencia del codigo de Warehouse en la BBDD
		if w.warehouseService.Exists(c, req.WarehouseCode) {
			web.Error(c, 409, "warehouse with code %v already exists", req.WarehouseCode)
			return
		}

		// Seteo el WH con el request
		wh := domain.Warehouse{
			Address:            req.Address,
			Telephone:          req.Telephone,
			WarehouseCode:      req.WarehouseCode,
			MinimumCapacity:    req.MinimumCapacity,
			MinimumTemperature: req.MinimumTemperature,
		}

		// Guardo el WH en la BBDD
		id, err := w.warehouseService.Save(c, wh)
		if err != nil {
			web.Error(c, 500, "internal server error")
			// fmt.Printf("ERROR Create(): %v\n", err.Error())
			return
		}

		// Agrego el id generado y lo retorno
		wh.ID = id
		web.Success(c, 201, wh)
	}
}

// UpadteWarehouses godoc
// @Summary Update Warehouses
// @Tags Warehouses
// @Description update warehouses by id
// @Accept json
// @Produce json
// @Param id path string true "Warehouse id"
// @Param warehouse body patchRequestWH true "Warehouse to update"
// @Success 200 {object} web.response
// @Failure 400 {object} web.errorResponse
// @Failure 404 {object} web.errorResponse
// @Failure 409 {object} web.errorResponse
// @Failure 422 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /warehouses/{id} [patch]
func (w *Warehouse) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req patchRequestWH

		//Obtengo el wh con la funcion auxiliar o retorno un error
		wh, errCode, err := getWHByParamID(w, c)
		if err != nil {
			web.Error(c, errCode, err.Error())
			return
		}

		// Obtengo el Request del body y si hay error lo retorno
		if err := c.ShouldBindJSON(&req); err != nil {
			web.Error(c, 422, err.Error())
			return
		}

		// Si se envio un Warehouse y es diferente al actual, verifico que no exista en la BBDD
		if req.WarehouseCode != "" && req.WarehouseCode != wh.WarehouseCode {
			if w.warehouseService.Exists(c, req.WarehouseCode) {
				web.Error(c, 409, "warehouse with code %v already exists", req.WarehouseCode)
				return
			}
			wh.WarehouseCode = req.WarehouseCode
		}

		// Actualizo los campos que se hayan enviado
		updateWHFields(req, &wh)

		// Envio el update al service y retorno un error 500 si existe
		if err := w.warehouseService.Update(c, wh); err != nil {
			web.Error(c, 500, "internal server error")
			// fmt.Printf("ERROR Update(): %v\n", err.Error())
			return
		}

		// Retorno el wh actualizado
		web.Success(c, 200, wh)
	}
}

// DeleteWarehouse godoc
// @Summary Delete Warehouse by id
// @Tags Warehouses
// @Description delete warehouse by id
// @Param id path string true "Warehouse id"
// @Success 204 {object} nil
// @Failure 404 {object} web.errorResponse
// @Router /warehouses/{id} [delete]
func (w *Warehouse) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Convierto id en entero y en caso de error lo retorno
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			web.Error(c, 400, "%s", "id must be integer")
			return
		}
		// Llamo al Delete del service y en caso de error retorno un 404
		if err := w.warehouseService.Delete(c, id); err != nil {
			web.Error(c, 404, "warehouse not found")
			return
		}
		// Si todo salió bien retorno una 204 y una respuesta vacia
		web.Success(c, 204, nil)
	}
}

func getWHByParamID(w *Warehouse, c *gin.Context) (domain.Warehouse, int, error) {
	// Convierto id en entero y en caso de error lo retorno
	wh := domain.Warehouse{}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return wh, 400, fmt.Errorf("id must be integer")
	}

	// Le pido al service el wh y si no existe retorno un 404
	wh, err = w.warehouseService.Get(c, id)
	if err != nil {
		return wh, 404, fmt.Errorf("warehouse not found")
	}

	return wh, 0, nil
}

func updateWHFields(req patchRequestWH, wh *domain.Warehouse) {
	if req.Address != "" {
		wh.Address = req.Address
	}

	if req.Telephone != "" {
		wh.Telephone = req.Telephone
	}

	if req.MinimumCapacity != nil {
		wh.MinimumCapacity = *req.MinimumCapacity
	}

	if req.MinimumTemperature != nil {
		wh.MinimumTemperature = *req.MinimumTemperature
	}
}
