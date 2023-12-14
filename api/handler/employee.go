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

// CreateEmployee godoc
// @ID create_employee
// @Router /employee [POST]
// @Summary Create Employee
// @Description Create Employee
// @Tags Employee
// @Accept json
// @Produce json
// @Param object body models.CreateEmployee true "CreateEmployeeRequestBody"
// @Success 200 {object} Response{data=models.Employee} "EmployeeBody"
// @Response 400 {object} Response{data=string} "Invalid Argument"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) CreateEmployee(c *gin.Context) {

	// password := c.GetHeader("Password")
	// if password != "1234" {
	// 	handleResponse(c, http.StatusUnauthorized, "The request requires an user authentication.")
	// 	return
	// }

	var createEmployee models.CreateEmployee
	err := c.ShouldBindJSON(&createEmployee)
	if err != nil {
		handleResponse(c, 400, "ShouldBindJSON err:"+err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), config.CtxTimeout)
	defer cancel()

	resp, err := h.strg.Employee().Create(ctx, createEmployee)
	if err != nil {
		handleResponse(c, http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, http.StatusCreated, resp)
}

// GetByIdEmployee godoc
// @ID get_by_id_employee
// @Router /employee/{id} [GET]
// @Summary Get By Id Employee
// @Description Get By Id Employee
// @Tags Employee
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} Response{data=models.Employee} "EmployeeBody"
// @Response 400 {object} Response{data=string} "Invalid Argument"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) EmployeeGetById(c *gin.Context) {

	var id = c.Param("id")
	if !helpers.IsValidUUID(id) {
		handleResponse(c, http.StatusBadRequest, "id is not uuid")
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), config.CtxTimeout)
	defer cancel()

	resp, err := h.strg.Employee().GetByID(ctx, models.EmployeePrimaryKey{Id: id})
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

// GetListEmployee godoc
// @ID get_list_employee
// @Router /employee [GET]
// @Summary Get List Employee
// @Description Get List Employee
// @Tags Employee
// @Accept json
// @Produce json
// @Param limit query number false "limit"
// @Param offset query number false "offset"
// @Success 200 {object} Response{data=models.EmployeeGetListResponse} "GetListEmployeeResponseBody"
// @Response 400 {object} Response{data=string} "Invalid Argument"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) GetListEmployee(c *gin.Context) {

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

	resp, err := h.strg.Employee().GetList(ctx, models.EmployeeGetListRequest{
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

// UpdateEmployee godoc
// @ID update_employee
// @Router /employee/{id} [PUT]
// @Summary Update Employee
// @Description Update Employee
// @Tags Employee
// @Accept json
// @Produce json
// @Param id path string true "EmployeePrimaryKey_ID"
// @Param object body models.UpdateEmployee true "UpdateBranchBody"
// @Success 200 {object} Response{data=string} "Updated Employee"
// @Response 400 {object} Response{data=string} "Invalid Argument"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) EmployeeUpdate(c *gin.Context) {

	var updateEmployee models.UpdateEmployee

	err := c.ShouldBindJSON(&updateEmployee)
	if err != nil {
		handleResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var id = c.Param("id")
	if !helpers.IsValidUUID(id) {
		handleResponse(c, http.StatusBadRequest, "id is not uuid")
		return
	}
	updateEmployee.Id = id

	ctx, cancel := context.WithTimeout(context.Background(), config.CtxTimeout)
	defer cancel()

	rowsAffected, err := h.strg.Employee().Update(ctx, updateEmployee)
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

	resp, err := h.strg.Employee().GetByID(ctx, models.EmployeePrimaryKey{Id: updateEmployee.Id})
	if err != nil {
		handleResponse(c, http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, http.StatusAccepted, resp)
}

// DeleteEmployee godoc
// @ID delete_employee
// @Router /employee/{id} [DELETE]
// @Summary Delete Employee
// @Description Delete Employee
// @Tags Employee
// @Accept json
// @Produce json
// @Param id path string true "DeleteEmployeePath"
// @Success 200 {object} Response{data=string} "Deleted Employee"
// @Response 400 {object} Response{data=string} "Invalid Argument"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) EmployeeDelete(c *gin.Context) {
	var id = c.Param("id")

	if !helpers.IsValidUUID(id) {
		handleResponse(c, http.StatusBadRequest, "id is not uuid")
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), config.CtxTimeout)
	defer cancel()

	err := h.strg.Employee().Delete(ctx, models.EmployeePrimaryKey{Id: id})
	if err != nil {
		handleResponse(c, http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, http.StatusNoContent, nil)
}
