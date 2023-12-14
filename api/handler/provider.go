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

// CreateProvider godoc
// @ID create_provider
// @Router /provider [POST]
// @Summary Create Provider
// @Description Create Provider
// @Tags Provider
// @Accept json
// @Produce json
// @Param object body models.CreateProvider true "CreateProviderRequestBody"
// @Success 200 {object} Response{data=models.Provider} "ProviderBody"
// @Response 400 {object} Response{data=string} "Invalid Argument"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) CreateProvider(c *gin.Context) {

	// password := c.GetHeader("Password")
	// if password != "1234" {
	// 	handleResponse(c, http.StatusUnauthorized, "The request requires an user authentication.")
	// 	return
	// }

	var createProvider models.CreateProvider
	err := c.ShouldBindJSON(&createProvider)
	if err != nil {
		handleResponse(c, 400, "ShouldBindJSON err:"+err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), config.CtxTimeout)
	defer cancel()

	resp, err := h.strg.Provider().Create(ctx, createProvider)
	if err != nil {
		handleResponse(c, http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, http.StatusCreated, resp)
}

// GetByIdProvider godoc
// @ID get_by_id_provider
// @Router /provider/{id} [GET]
// @Summary Get By Id Provider
// @Description Get By Id Provider
// @Tags Provider
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} Response{data=models.Provider} "ProviderBody"
// @Response 400 {object} Response{data=string} "Invalid Argument"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) ProviderGetById(c *gin.Context) {

	var id = c.Param("id")
	if !helpers.IsValidUUID(id) {
		handleResponse(c, http.StatusBadRequest, "id is not uuid")
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), config.CtxTimeout)
	defer cancel()

	resp, err := h.strg.Provider().GetByID(ctx, models.ProviderPrimaryKey{Id: id})
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

// GetListProvider godoc
// @ID get_list_provider
// @Router /provider [GET]
// @Summary Get List Provider
// @Description Get List Provider
// @Tags Provider
// @Accept json
// @Produce json
// @Param limit query number false "limit"
// @Param offset query number false "offset"
// @Success 200 {object} Response{data=models.ProviderGetListResponse} "GetListProviderResponseBody"
// @Response 400 {object} Response{data=string} "Invalid Argument"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) GetListProvider(c *gin.Context) {

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

	resp, err := h.strg.Provider().GetList(ctx, models.ProviderGetListRequest{
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

// UpdateProvider godoc
// @ID update_provider
// @Router /provider/{id} [PUT]
// @Summary Update Provider
// @Description Update Provider
// @Tags Provider
// @Accept json
// @Produce json
// @Param id path string true "ProviderPrimaryKey_ID"
// @Param object body models.UpdateProvider true "UpdateBranchBody"
// @Success 200 {object} Response{data=string} "Updated Provider"
// @Response 400 {object} Response{data=string} "Invalid Argument"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) ProviderUpdate(c *gin.Context) {

	var updateProvider models.UpdateProvider

	err := c.ShouldBindJSON(&updateProvider)
	if err != nil {
		handleResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var id = c.Param("id")
	if !helpers.IsValidUUID(id) {
		handleResponse(c, http.StatusBadRequest, "id is not uuid")
		return
	}
	updateProvider.Id = id

	ctx, cancel := context.WithTimeout(context.Background(), config.CtxTimeout)
	defer cancel()

	rowsAffected, err := h.strg.Provider().Update(ctx, updateProvider)
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

	resp, err := h.strg.Provider().GetByID(ctx, models.ProviderPrimaryKey{Id: updateProvider.Id})
	if err != nil {
		handleResponse(c, http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, http.StatusAccepted, resp)
}

// DeleteProvider godoc
// @ID delete_provider
// @Router /provider/{id} [DELETE]
// @Summary Delete Provider
// @Description Delete Provider
// @Tags Provider
// @Accept json
// @Produce json
// @Param id path string true "DeleteProviderPath"
// @Success 200 {object} Response{data=string} "Deleted Provider"
// @Response 400 {object} Response{data=string} "Invalid Argument"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) ProviderDelete(c *gin.Context) {
	var id = c.Param("id")

	if !helpers.IsValidUUID(id) {
		handleResponse(c, http.StatusBadRequest, "id is not uuid")
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), config.CtxTimeout)
	defer cancel()

	err := h.strg.Provider().Delete(ctx, models.ProviderPrimaryKey{Id: id})
	if err != nil {
		handleResponse(c, http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, http.StatusNoContent, nil)
}
