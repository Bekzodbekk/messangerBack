package handlers

import (
	"api-gateway/genproto/userpb"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *HandlerST) Register(ctx *gin.Context) {
	var req userpb.CreateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.service.Register(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, resp)
}

func (h *HandlerST) Verify(ctx *gin.Context) {
	var user userpb.VerifyRequest
	err := ctx.BindJSON(&user)
	if err != nil {
		ctx.JSON(400, err)
		return
	}

	resp, err := h.service.UserClient.Verify(ctx, &user)
	if err != nil {
		ctx.JSON(400, err)
		return
	}

	ctx.JSON(201, resp)
}

func (h *HandlerST) SignIn(ctx *gin.Context) {
	var req userpb.SignInRequest
	err := ctx.BindJSON(&req)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	resp, err := h.service.Login(ctx, &req)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	ctx.JSON(200, resp)
}

func (h *HandlerST) UpdateUser(ctx *gin.Context) {
	id := ctx.Param("id")
	var req userpb.UpdateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req.Id = id

	resp, err := h.service.UpdateUser(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *HandlerST) DeleteUser(ctx *gin.Context) {
	id := ctx.Param("id")
	req := &userpb.DeleteUserRequest{Id: id}

	resp, err := h.service.DeleteUser(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *HandlerST) GetUserById(ctx *gin.Context) {
	id := ctx.Param("id")
	req := &userpb.GetUserByIdRequest{Id: id}

	resp, err := h.service.GetUserById(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *HandlerST) GetUserByFilter(ctx *gin.Context) {
	// Create a new GetUserByFilterRequest
	req := &userpb.GetUserByFilterRequest{
		FirstName: ctx.Query("first_name"),
		LastName:  ctx.Query("last_name"),
		Email:     ctx.Query("email"),
		Username:  ctx.Query("username"),
	}

	// Handle DeletedAt separately as it's an int64
	if deletedAt, err := strconv.ParseInt(ctx.Query("deleted_at"), 10, 64); err == nil {
		req.DeletedAt = deletedAt
	}

	// Call the service
	resp, err := h.service.GetUserByFilter(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}
func (h *HandlerST) GetUsers(ctx *gin.Context) {
	req := &userpb.Void{}

	resp, err := h.service.GetUsers(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}
