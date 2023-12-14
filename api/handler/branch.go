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

// CreateBranch godoc
// @ID create_branch
// @Router /branch [POST]
// @Summary Create Branch
// @Description Create Branch
// @Tags Branch
// @Accept json
// @Produce json
// @Param object body models.CreateBranch true "CreateBranchRequestBody"
// @Success 200 {object} Response{data=models.Branch} "BranchBody"
// @Response 400 {object} Response{data=string} "Invalid Argument"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) CreateBranch(c *gin.Context) {

	// password := c.GetHeader("Password")
	// if password != "1234" {
	// 	handleResponse(c, http.StatusUnauthorized, "The request requires an user authentication.")
	// 	return
	// }

	var createBranch models.CreateBranch
	err := c.ShouldBindJSON(&createBranch)
	if err != nil {
		handleResponse(c, 400, "ShouldBindJSON err:"+err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), config.CtxTimeout)
	defer cancel()

	resp, err := h.strg.Branch().Create(ctx, createBranch)
	if err != nil {
		handleResponse(c, http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, http.StatusCreated, resp)
}

// GetByIdBranch godoc
// @ID get_by_id_branch
// @Router /branch/{id} [GET]
// @Summary Get By Id Branch
// @Description Get By Id Branch
// @Tags Branch
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} Response{data=models.Branch} "BranchBody"
// @Response 400 {object} Response{data=string} "Invalid Argument"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) BranchGetById(c *gin.Context) {

	var id = c.Param("id")
	if !helpers.IsValidUUID(id) {
		handleResponse(c, http.StatusBadRequest, "id is not uuid")
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), config.CtxTimeout)
	defer cancel()

	resp, err := h.strg.Branch().GetByID(ctx, models.BranchPrimaryKey{Id: id})
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

// GetListBranch godoc
// @ID get_list_branch
// @Router /branch [GET]
// @Summary Get List Branch
// @Description Get List Branch
// @Tags Branch
// @Accept json
// @Produce json
// @Param limit query number false "limit"
// @Param offset query number false "offset"
// @Success 200 {object} Response{data=models.GetListBranchResponse} "GetListBranchResponseBody"
// @Response 400 {object} Response{data=string} "Invalid Argument"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) GetListBranch(c *gin.Context) {

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

	resp, err := h.strg.Branch().GetList(ctx, models.GetListBranchRequest{
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

// UpdateBranch godoc
// @ID update_branch
// @Router /branch/{id} [PUT]
// @Summary Update Branch
// @Description Update Branch
// @Tags Branch
// @Accept json
// @Produce json
// @Param id path string true "BranchPrimaryKey_ID"
// @Param object body models.UpdateBranch true "UpdateBranchBody"
// @Success 200 {object} Response{data=string} "Updated Branch"
// @Response 400 {object} Response{data=string} "Invalid Argument"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) BranchUpdate(c *gin.Context) {

	var updateBranch models.UpdateBranch

	err := c.ShouldBindJSON(&updateBranch)
	if err != nil {
		handleResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var id = c.Param("id")
	if !helpers.IsValidUUID(id) {
		handleResponse(c, http.StatusBadRequest, "id is not uuid")
		return
	}
	updateBranch.Id = id

	ctx, cancel := context.WithTimeout(context.Background(), config.CtxTimeout)
	defer cancel()

	rowsAffected, err := h.strg.Branch().Update(ctx, updateBranch)
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

	resp, err := h.strg.Branch().GetByID(ctx, models.BranchPrimaryKey{Id: updateBranch.Id})
	if err != nil {
		handleResponse(c, http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, http.StatusAccepted, resp)
}

// DeleteBranch godoc
// @ID delete_branch
// @Router /branch/{id} [DELETE]
// @Summary Delete Branch
// @Description Delete Branch
// @Tags Branch
// @Accept json
// @Produce json
// @Param id path string true "DeleteBranchPath"
// @Success 200 {object} Response{data=string} "Deleted Branch"
// @Response 400 {object} Response{data=string} "Invalid Argument"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) BranchDelete(c *gin.Context) {
	var id = c.Param("id")

	if !helpers.IsValidUUID(id) {
		handleResponse(c, http.StatusBadRequest, "id is not uuid")
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), config.CtxTimeout)
	defer cancel()

	err := h.strg.Branch().Delete(ctx, models.BranchPrimaryKey{Id: id})
	if err != nil {
		handleResponse(c, http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, http.StatusNoContent, nil)
}
