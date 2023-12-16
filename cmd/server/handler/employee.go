package handler

import (
	"github.com/extmatperez/meli_bootcamp_go_w5-5/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w5-5/internal/employee"
	"github.com/extmatperez/meli_bootcamp_go_w5-5/pkg/web"
	"github.com/gin-gonic/gin"

	"strconv"
)

type Employee struct {
	employeeService employee.Service
}

type postEmployee struct {
	CardNumberID string `json:"card_number_id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	WarehouseID  int    `json:"warehouse_id"`
}

type patchEmployee struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	WarehouseID int    `json:"warehouse_id"`
}

func NewEmployee(e employee.Service) *Employee {
	return &Employee{
		employeeService: e,
	}
}

//CreateEmployees godoc
//@Summary Create employees
//@Tags Employees
//@Description Create employees
//@Accept json
//@Produce json
//@Param Employees body postEmployee true "Employees to store"
//@Succes 201 {object} web.Response
//@Failure 409 {object} web.errorResponse
//@Failure 422 {object} web.errorResponse
//@Router /employees [post]
func (e *Employee) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req postEmployee
		err := c.ShouldBindJSON(&req)
		if err != nil {
			web.Error(c, 422, err.Error())
			return
		}
		if e.employeeService.Exists(c, req.CardNumberID) {
			web.Error(c, 409, "El employee ya existe")
			return
		}
		if req.CardNumberID == "" {
			web.Error(c, 422, "El CardNumberID es requerido")
			return
		}
		if req.FirstName == "" {
			web.Error(c, 422, "El FirstName es requerido")
			return
		}
		if req.LastName == "" {
			web.Error(c, 422, "El LastName es requerido")
			return
		}
		if req.WarehouseID == 0 {
			web.Error(c, 422, "El WarehouseID es requerido")
			return
		}

		emp := domain.Employee{
			CardNumberID: req.CardNumberID,
			FirstName:    req.FirstName,
			LastName:     req.LastName,
			WarehouseID:  req.WarehouseID,
		}

		id, err := e.employeeService.Save(c, emp)
		if err != nil {
			web.Error(c, 500, "internal server error")
			return
		}

		emp.ID = id
		web.Success(c, 201, emp)
	}
}

//ListEmployees godoc
//@Summary List employees
//@Tags Employees
//@Description get all employees
//@Accept json
//@Produce json
//@Succes 200 {object} web.Response
//@Failure 500 {object} web.errorResponse
//@Router /employees [get]
func (e *Employee) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		employees, err := e.employeeService.GetAll(c)
		if err != nil {
			web.Error(c, 500, err.Error())
			return
		}
		if len(employees) == 0 {
			web.Success(c, 200, []domain.Employee{})
			return
		}
		web.Success(c, 200, employees)
	}
}

//ListEmployee godoc
//@Summary List employee by ID
//@Tags Employees
//@Description get employee by ID
//@Produce json
//@Param id path int true "employees id"
//@Succes 200 {object} web.Response
//@Failure 400 {object} web.errorResponse
//@Failure 404 {object} web.errorResponse
//@Router /employees/{id} [get]
func (e *Employee) Get() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			web.Error(c, 400, "El id es invalido")
			return
		}
		emp, err := e.employeeService.Get(c, id)
		if err != nil {
			web.Error(c, 404, "El id no existe")
			return
		}
		web.Success(c, 200, emp)
	}
}

// GetInboundOrders
func (e *Employee) GetInboundOrders() gin.HandlerFunc {
	return func(c *gin.Context) {

		idQuery, idExists := c.GetQuery("id")

		if idExists {
			id, err := strconv.Atoi(idQuery)
			if err != nil {
				web.Error(c, 400, "error. The id [%d] entered must be of type *integer*", id)
				return
			}

			reportInbOrd, err := e.employeeService.GetInboundOrders(c, id)
			if err != nil {
				web.Error(c, 404, err.Error())
				return
			}

			// Return data Report and Success 200
			web.Success(c, 200, reportInbOrd)

		} else {
			id := 0
			reportInbOrd, err := e.employeeService.GetInboundOrders(c, id)
			if err != nil {
				web.Error(c, 500, err.Error())
				return
			}
			// Return data Report and Success 200
			web.Success(c, 200, reportInbOrd)
		}
	}
}

//ModifyEmployees godoc
//@Summary Modify employees
//@Tags Employees
//@Description Modify employees by ID
//@Accept json
//@Produce json
//@Param id path int true "employees id"
//@Param employees body patchEmployee true "employees to update"
//@Succes 200 {object} web.Response
//@Failure 400 {object} web.errorResponse
//@Failure 404 {object} web.errorResponse
//@Failure 500 {object} web.errorResponse
//@Router /employees/{id} [patch]
func (e *Employee) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req patchEmployee
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			web.Error(c, 400, "El id es invalido")
			return
		}
		emp, err := e.employeeService.Get(c, id)
		if err != nil {
			web.Error(c, 404, "El id no existe")
			return
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			web.Error(c, 404, err.Error())
			return
		}
		if req.FirstName != "" {
			emp.FirstName = req.FirstName
		}
		if req.LastName != "" {
			emp.LastName = req.LastName
		}
		if req.WarehouseID != 0 {
			emp.WarehouseID = req.WarehouseID
		}
		if err := e.employeeService.Update(c, emp); err != nil {
			web.Error(c, 500, err.Error())
			return
		}
		web.Success(c, 200, emp)
	}
}

//DeleteEmployees godoc
//@Summary Delete employee by ID
//@Tags Employees
//@Description Delete employees by ID
//@Param id path int true "employees id"
//@Succes 204 {object} nil
//@Failure 400 {object} web.errorResponse
//@Failure 404 {object} web.errorResponse
//@Router /employees/{id} [delete]
func (e *Employee) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			web.Error(c, 400, "El id es invalido")
			return
		}
		err = e.employeeService.Delete(c, id)
		if err != nil {
			web.Error(c, 404, "El id no existe")
			return
		}
		web.Success(c, 204, nil)
	}
}
