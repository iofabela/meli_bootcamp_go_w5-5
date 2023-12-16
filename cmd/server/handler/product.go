package handler

import (
	"fmt"
	"strconv"

	"github.com/extmatperez/meli_bootcamp_go_w5-5/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w5-5/internal/product"
	"github.com/extmatperez/meli_bootcamp_go_w5-5/pkg/web"
	"github.com/gin-gonic/gin"
)

// Product ... | Product struct for Service
type Product struct {
	productService product.Service
}

type postReq struct {
	Description    string  `json:"description"`
	ExpirationRate int     `json:"expiration_rate"`
	FreezingRate   int     `json:"freezing_rate" `
	Height         float32 `json:"height"`
	Length         float32 `json:"length"`
	Netweight      float32 `json:"netweight"`
	ProductCode    string  `json:"product_code"`
	RecomFreezTemp float32 `json:"recommended_freezing_temperature"`
	Width          float32 `json:"width"`
	ProductTypeID  int     `json:"product_type_id"`
	SellerID       int     `json:"seller_id"`
}

type patchReq struct {
	Description    string   `json:"description"`
	ExpirationRate *int     `json:"expiration_rate"`
	FreezingRate   *int     `json:"freezing_rate"`
	Height         *float32 `json:"height"`
	Length         *float32 `json:"length"`
	Netweight      *float32 `json:"netweight"`
	ProductCode    string   `json:"product_code"`
	Width          *float32 `json:"width"`
}

// NewProduct ...
func NewProduct(p product.Service) *Product {
	return &Product{
		productService: p,
	}
}

// Get | List-Products godoc
// @Summary List of all products from database
// @Tags Products
// @Description Get all products from database
// @Produce  json
// @Success 200 {object} web.response
// @Failure 500 {object} web.errorResponse
// @Router /products [GET]
func (p *Product) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		prd, err := p.productService.GetAll(c)
		if err != nil {
			web.Error(c, 500, "%s", err.Error())
			return
		}
		if len(prd) == 0 {
			web.Success(c, 200, []domain.Warehouse{})
			return
		}
		web.Success(c, 200, prd)
	}
}

// Get by Id godoc
// @Summary Get a product with ID
// @Tags Products
// @Description Get product information using ID
// @Produce  json
// @Param id path int true "id"
// @Success 200 {object} web.response
// @Failure 400 {object} web.errorResponse
// @Failure 404 {object} web.errorResponse
// @Router /products/{id} [GET]
func (p *Product) Get() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			web.Error(c, 400, "error. The id [%d] entered must be of type *integer*", id)
			return
		}

		prd, err := p.productService.Get(c, id)
		if err != nil {
			web.Error(c, 404, "error. No product found with the entered id: %d", id)
			return
		}

		web.Success(c, 200, prd)
	}
}

// Post | Create a product godoc
// @Summary Create a new Product with Service
// @Tags Products
// @Description Create a product for POST
// @Accept json
// @Produce  json
// @Param product body postReq true "Create a product"
// @Success 200 {object} web.response
// @Failure 400 {object} web.errorResponse
// @Failure 422 {object} web.errorResponse
// @Failure 409 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /products [POST]
func (p *Product) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req postReq

		if err := c.ShouldBindJSON(&req); err != nil {
			web.Error(c, 422, "%s", err.Error())
			return
		}

		if p.productService.Exists(c, req.ProductCode) {
			web.Error(c, 409, "error. the product with the code: %v, already exists", req.ProductCode)
			return
		}

		prd := domain.Product{}

		errs := validatePost(&req, &prd)
		if len(errs) > 0 {
			errors := ""
			for _, error := range errs {
				errors += error.Error()
			}

			web.Error(c, 400, "error: %s", errors)

			return
		}

		id, err := p.productService.Save(c, prd)
		if err != nil {
			web.Error(c, 500, "error. %s", err.Error())
			return
		}

		prd.ID = id
		web.Success(c, 201, prd)
	}
}

func validatePost(req *postReq, pr *domain.Product) []error {
	errors := []error{}

	if req.Description != "" {
		pr.Description = req.Description
	} else {
		errors = append(errors, fmt.Errorf("\n*Description* field must not be null"))
	}

	if &req.ExpirationRate != nil {
		if req.ExpirationRate > 0 {
			pr.ExpirationRate = req.ExpirationRate
		} else {
			errors = append(errors, fmt.Errorf("\n*ExpirationRate* field must not be null & greater than [zero]"))
		}
	}

	if &req.FreezingRate != nil {
		if req.FreezingRate > 0 {
			pr.FreezingRate = req.FreezingRate
		} else {
			errors = append(errors, fmt.Errorf("\n*FreezingRate* field must not be null"))
		}
	}

	if &req.Height != nil {
		if req.Length > 0 {
			pr.Height = req.Height
		} else {
			errors = append(errors, fmt.Errorf("\n*Height* field must not be null"))
		}
	}

	if &req.Length != nil {
		if req.Length > 0 {
			pr.Length = req.Length
		} else {
			errors = append(errors, fmt.Errorf("\n*Lenght* field must not be null & greater than [zero]"))
		}
	}

	if &req.Netweight != nil {
		if req.Netweight > 0 {
			pr.Netweight = req.Netweight
		} else {
			errors = append(errors, fmt.Errorf("\n*Netweight* field must not be null & greater than [zero]"))
		}
	}

	if req.ProductCode != "" {
		pr.ProductCode = req.ProductCode
	} else {
		errors = append(errors, fmt.Errorf("\n*ProductCode* field must not be null"))
	}

	if &req.RecomFreezTemp != nil {
		if req.RecomFreezTemp > 0 {
			pr.RecomFreezTemp = req.RecomFreezTemp
		} else {
			errors = append(errors, fmt.Errorf("\n*RecomFreezTemp* field must not be null"))
		}
	}

	if &req.Width != nil {
		if req.Width > 0 {
			pr.Width = req.Width
		} else {
			errors = append(errors, fmt.Errorf("\n*Width* field must not be null & greater than [zero]"))
		}
	}

	if &req.ProductTypeID != nil {
		if req.ProductTypeID > 0 {
			pr.ProductTypeID = req.ProductTypeID
		} else {
			errors = append(errors, fmt.Errorf("\n*ProductTypeID* field must not be null & greater than [zero]"))
		}
	}

	if &req.SellerID != nil {
		if req.SellerID > 0 {
			pr.SellerID = req.SellerID
		} else {
			errors = append(errors, fmt.Errorf("\n*SellerID* field must not be null & greater than [zero]"))
		}
	}

	return errors

}

// Patch | Update a product godoc
// @Summary Update a Product with Service
// @Tags Products
// @Description Update a product with the requirements required.
// @Accept json
// @Produce  json
// @Param id path int true "id"
// @Param product body patchReq true "Create a product"
// @Success 200 {object} web.response
// @Failure 400 {object} web.errorResponse
// @Failure 404 {object} web.errorResponse
// @Failure 422 {object} web.errorResponse
// @Failure 409 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /products/{id} [PATCH]
func (p *Product) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req patchReq

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			web.Error(c, 400, "error. The entered id must be of type *integer*")
			return
		}

		prd, err := p.productService.Get(c, id)
		if err != nil {
			web.Error(c, 404, "error. No product found with the entered id: %d", id)
			return
		}

		if err := c.ShouldBindJSON(&req); err != nil {
			web.Error(c, 422, "%s", err.Error())
			return
		}

		if req.ProductCode != "" {
			if p.productService.Exists(c, req.ProductCode) {
				web.Error(c, 409, "error. the product with code: %v, already exists", req.ProductCode)
				return
			}
			prd.ProductCode = req.ProductCode
		}

		errs := validatePatch(&req, &prd)
		if len(errs) > 0 {
			errors := ""
			for _, error := range errs {
				errors += error.Error()
			}

			web.Error(c, 400, "error: %s", errors)

			return
		}

		if err := p.productService.Update(c, prd); err != nil {
			web.Error(c, 500, "error: %s", err.Error())
			return
		}

		web.Success(c, 200, prd)
	}
}

func validatePatch(req *patchReq, p *domain.Product) []error {
	errors := []error{}

	if req.Description != "" {
		p.Description = req.Description
	} else {
		errors = append(errors, fmt.Errorf("\n*Description* field must not be null"))
	}

	if req.ExpirationRate != nil {
		if *req.ExpirationRate > 0 {
			p.ExpirationRate = *req.ExpirationRate
		} else {
			errors = append(errors, fmt.Errorf("\n*ExpirationRate* field must not be null & greater than [zero]"))
		}
	}

	if req.FreezingRate != nil {
		if *req.FreezingRate > 0 {
			p.FreezingRate = *req.FreezingRate
		} else {
			errors = append(errors, fmt.Errorf("\n*FreezingRate* field must not be null"))
		}
	}

	if req.Height != nil {
		if *req.Length > 0 {
			p.Height = *req.Height
		} else {
			errors = append(errors, fmt.Errorf("\n*Height* field must not be null"))
		}
	}

	if req.Length != nil {
		if *req.Length > 0 {
			p.Length = *req.Length
		} else {
			errors = append(errors, fmt.Errorf("\n*Lenght* field must not be null & greater than [zero]"))
		}
	}

	if req.Netweight != nil {
		if *req.Netweight > 0 {
			p.Netweight = *req.Netweight
		} else {
			errors = append(errors, fmt.Errorf("\n*Netweight* field must not be null & greater than [zero]"))
		}
	}

	if req.Width != nil {
		if *req.Width > 0 {
			p.Width = *req.Width
		} else {
			errors = append(errors, fmt.Errorf("\n*Width* field must not be null & greater than [zero]"))
		}
	}

	return errors
}

// Delete | Delete a product godoc
// @Summary Delete a Product with Service
// @Tags Products
// @Description Delete a product from database
// @Param id path int true "id"
// @Success 204 {object} web.response
// @Failure 400 {object} web.errorResponse
// @Failure 404 {object} web.errorResponse
// @Router /products/{id} [DELETE]
func (p *Product) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			web.Error(c, 400, "error. The entered id must be of type *integer*")
			return
		}
		if err := p.productService.Delete(c, id); err != nil {
			web.Error(c, 404, "error. No product found with the entered id: %d", id)
			return
		}
		web.Success(c, 204, nil)
	}
}
