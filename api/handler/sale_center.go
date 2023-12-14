package handler

import (
	"context"
	"database/sql"
	"market/config"

	models "market/models/organization"
	"market/pkg/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateCenter godoc
// @ID create_center
// @Router /center [POST]
// @Summary Create Center
// @Description Create Center
// @Tags Center
// @Accept json
// @Produce json
// @Param object body models.CreateSaleCenter true "CreateCenterRequestBody"
// @Success 200 {object} Response{data=models.SaleCenter} "CenterBody"
// @Response 400 {object} Response{data=string} "Invalid Argument"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) CreateCenter(c *gin.Context) {

	// password := c.GetHeader("Password")
	// if password != "1234" {
	// 	handleResponse(c, http.StatusUnauthorized, "The request requires an user authentication.")
	// 	return
	// }

	var createCenter models.CreateSaleCenter
	err := c.ShouldBindJSON(&createCenter)
	if err != nil {
		handleResponse(c, 400, "ShouldBindJSON err:"+err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), config.CtxTimeout)
	defer cancel()

	resp, err := h.strg.Center().Create(ctx, createCenter)
	if err != nil {
		handleResponse(c, http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, http.StatusCreated, resp)
}

// GetByIdCenter godoc
// @ID get_by_id_center
// @Router /center/{id} [GET]
// @Summary Get By Id Center
// @Description Get By Id Center
// @Tags Center
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} Response{data=models.SaleCenter} "CenterBody"
// @Response 400 {object} Response{data=string} "Invalid Argument"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) CenterGetById(c *gin.Context) {

	var id = c.Param("id")
	if !helpers.IsValidUUID(id) {
		handleResponse(c, http.StatusBadRequest, "id is not uuid")
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), config.CtxTimeout)
	defer cancel()

	resp, err := h.strg.Center().GetByID(ctx, models.SaleCenterPrimaryKey{Id: id})
	if err == sql.ErrNoRows {
		handleResponse(c, http.StatusBadRequest, "no rows in result set")
		return
	}

	if err != nil {
		handleResponse(c, http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, http.StatusOK, resp)
}

// GetListCenter godoc
// @ID get_list_center
// @Router /center [GET]
// @Summary Get List Center
// @Description Get List Center
// @Tags Center
// @Accept json
// @Produce json
// @Param limit query number false "limit"
// @Param offset query number false "offset"
// @Success 200 {object} Response{data=models.SaleCenterGetListResponse} "GetListCenterResponseBody"
// @Response 400 {object} Response{data=string} "Invalid Argument"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) GetListCenter(c *gin.Context) {

	limit, err := getIntegerOrDefaultValue(c.Query("limit"), 10)
	if err != nil {
		handleResponse(c, http.StatusBadRequest, "invalid query limit")
		return
	}

	offset, err := getIntegerOrDefaultValue(c.Query("offset"), 0)
	if err != nil {
		handleResponse(c, http.StatusBadRequest, "invalid query offset")
		return
	}

	search := c.Query("search")
	if err != nil {
		handleResponse(c, http.StatusBadRequest, "invalid query search")
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), config.CtxTimeout)
	defer cancel()

	resp, err := h.strg.Center().GetList(ctx, models.SaleCenterGetListRequest{
		Limit:  limit,
		Offset: offset,
		Search: search,
	})

	if err != nil {
		handleResponse(c, http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, http.StatusOK, resp)
}

// UpdateCenter godoc
// @ID update_center
// @Router /center/{id} [PUT]
// @Summary Update Center
// @Description Update Center
// @Tags Center
// @Accept json
// @Produce json
// @Param id path string true "CenterPrimaryKey_ID"
// @Param object body models.UpdateSaleCenter true "UpdateCenterBody"
// @Success 200 {object} Response{data=string} "Updated Center"
// @Response 400 {object} Response{data=string} "Invalid Argument"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) CenterUpdate(c *gin.Context) {

	var updateCenter models.UpdateSaleCenter

	err := c.ShouldBindJSON(&updateCenter)
	if err != nil {
		handleResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var id = c.Param("id")
	if !helpers.IsValidUUID(id) {
		handleResponse(c, http.StatusBadRequest, "id is not uuid")
		return
	}
	updateCenter.Id = id

	ctx, cancel := context.WithTimeout(context.Background(), config.CtxTimeout)
	defer cancel()

	rowsAffected, err := h.strg.Center().Update(ctx, updateCenter)
	if err != nil {
		handleResponse(c, http.StatusInternalServerError, err)
		return
	}

	if rowsAffected == 0 {
		handleResponse(c, http.StatusBadRequest, "no rows affected")
		return
	}

	ctx, cancel = context.WithTimeout(context.Background(), config.CtxTimeout)
	defer cancel()

	resp, err := h.strg.Center().GetByID(ctx, models.SaleCenterPrimaryKey{Id: updateCenter.Id})
	if err != nil {
		handleResponse(c, http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, http.StatusAccepted, resp)
}

// DeleteCenter godoc
// @ID delete_center
// @Router /center/{id} [DELETE]
// @Summary Delete Center
// @Description Delete Center
// @Tags Center
// @Accept json
// @Produce json
// @Param id path string true "DeleteCenterPath"
// @Success 200 {object} Response{data=string} "Deleted Center"
// @Response 400 {object} Response{data=string} "Invalid Argument"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) CenterDelete(c *gin.Context) {
	var id = c.Param("id")

	if !helpers.IsValidUUID(id) {
		handleResponse(c, http.StatusBadRequest, "id is not uuid")
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), config.CtxTimeout)
	defer cancel()

	err := h.strg.Center().Delete(ctx, models.SaleCenterPrimaryKey{Id: id})
	if err != nil {
		handleResponse(c, http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, http.StatusNoContent, nil)
}
