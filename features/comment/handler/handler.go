package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/roihan12/h8-mygram/features/comment"
	"github.com/roihan12/h8-mygram/utils"
)

type CommentController struct {
	commentService comment.CommentService
}

func New(srv comment.CommentService) *CommentController {
	return &CommentController{
		commentService: srv,
	}
}

// CreateComment godoc
//
//	@Summary		Create a new comment
//	@Description	create a new comment  with message and photo_id
//	@Tags			Comments
//	@Accept			json
//	@Produce		json
//	@Param			CommentRequest	body		CommentRequest		true	"Create comment request"
//	@Success		200				{object}	CommentResponse		"Comment created"
//	@Failure		400				{object}	utils.ErrorResponse	"Validation error"
//	@Failure		401				{object}	utils.ErrorResponse	"Unauthorized error"
//	@Failure		403				{object}	utils.ErrorResponse	"Forbidden error"
//	@Failure		404				{object}	utils.ErrorResponse	"Data not found error"
//	@Failure		409				{object}	utils.ErrorResponse	"Data conflict error"
//	@Failure		500				{object}	utils.ErrorResponse	"Internal server error"
//	@Router			/comments [post]
//	@Security		BearerAuth
func (ch *CommentController) Create(ctx *gin.Context) {
	var req CommentRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.ValidationError(ctx, err)
		return
	}
	userID, userIDIsExist := ctx.Get("userID")
	if !userIDIsExist {
		err := utils.ErrUnauthorized
		utils.HandleAbort(ctx, err)
		return
	}

	newComment := CommentRequestToCommentEntity(&req)
	newComment.UserID = userID.(uint)

	response, err := ch.commentService.Create(newComment)
	if err != nil {
		utils.HandleError(ctx, err)
		return
	}
	utils.HandleSuccess(ctx, "Create comment successfully", CommentEntityToCommentResponse(response))
}

// UpdateComment godoc
//
//	@Summary		Update a comment
//	@Description	update a comment  with message
//	@Tags			Comments
//	@Accept			json
//	@Produce		json
//	@Param			commentId		path		uint				true	"Comment ID"
//	@Param			CommentRequest	body		CommentUpdate		true	"Update comment request"
//	@Success		200				{object}	CommentResponse		"Comment Update Success"
//	@Failure		400				{object}	utils.ErrorResponse	"Validation error"
//	@Failure		401				{object}	utils.ErrorResponse	"Unauthorized error"
//	@Failure		403				{object}	utils.ErrorResponse	"Forbidden error"
//	@Failure		404				{object}	utils.ErrorResponse	"Data not found error"
//	@Failure		409				{object}	utils.ErrorResponse	"Data conflict error"
//	@Failure		500				{object}	utils.ErrorResponse	"Internal server error"
//	@Router			/comments/{commentId} [put]
//	@Security		BearerAuth
func (ch *CommentController) Update(ctx *gin.Context) {
	var req CommentUpdate
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.ValidationError(ctx, err)
		return
	}

	userID, userIDIsExist := ctx.Get("userID")
	if !userIDIsExist {
		err := utils.ErrUnauthorized
		utils.HandleAbort(ctx, err)
		return
	}

	id, _ := strconv.Atoi(ctx.Param("commentId"))

	response, err := ch.commentService.Update(CommentUpdateToCommentEntity(&req), uint(id), userID.(uint))
	if err != nil {
		utils.HandleError(ctx, err)
		return
	}
	utils.HandleSuccess(ctx, "Update comment successfully", CommentEntityToCommentResponse(response))
}

// DeleteComment godoc
//
//	@Summary		Delete a comment
//	@Description	Delete a comment by id
//	@Tags			Comments
//	@Accept			json
//	@Produce		json
//	@Param			commentId	path		uint				true	"Comment ID"
//	@Success		200			{object}	utils.Response		"Comment deleted"
//	@Failure		400			{object}	utils.ErrorResponse	"Validation error"
//	@Failure		401			{object}	utils.ErrorResponse	"Unauthorized error"
//	@Failure		403			{object}	utils.ErrorResponse	"Forbidden error"
//	@Failure		404			{object}	utils.ErrorResponse	"Data not found error"
//	@Failure		500			{object}	utils.ErrorResponse	"Internal server error"
//	@Router			/comments/{commentId} [delete]
//	@Security		BearerAuth
func (ch *CommentController) Delete(ctx *gin.Context) {
	userID, userIDIsExist := ctx.Get("userID")
	if !userIDIsExist {
		err := utils.ErrUnauthorized
		utils.HandleAbort(ctx, err)
		return
	}

	id, _ := strconv.Atoi(ctx.Param("commentId"))

	err := ch.commentService.Delete(uint(id), userID.(uint))
	if err != nil {
		utils.HandleError(ctx, err)
		return
	}
	utils.HandleSuccess(ctx, "Delete comment successfully", nil)
}

func (ch *CommentController) GetCommentByPhotoID(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("photoId"))
	response, err := ch.commentService.GetByPhotoID(uint(id))
	if err != nil {
		utils.HandleError(ctx, err)
		return
	}
	utils.HandleSuccess(ctx, "Get comment by photo id successfully", ListCommentToCommentResponse(response))
}

// GetComment godoc
//
//	@Summary		Get a comment
//	@Description	get a comment by id
//	@Tags			Comments
//	@Accept			json
//	@Produce		json
//	@Param			commentId	path		uint				true	"Comment ID"
//	@Success		200			{object}	CommentResponse		"Comment retrieved"
//	@Failure		400			{object}	utils.ErrorResponse	"Validation error"
//	@Failure		404			{object}	utils.ErrorResponse	"Data not found error"
//	@Failure		500			{object}	utils.ErrorResponse	"Internal server error"
//	@Router			/comments/{commentId} [get]
//	@Security		BearerAuth
func (ch *CommentController) GetById(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("commentId"))
	response, err := ch.commentService.GetById(uint(id))
	if err != nil {
		utils.HandleError(ctx, err)
		return
	}
	utils.HandleSuccess(ctx, "Get comment by id successfully", CommentEntityToCommentResponse(response))
}

// ListMyComment godoc
//
//	@Summary		List my comments
//	@Description	List my comments
//	@Tags			Comments
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		CommentResponse		"Comment retrieved"
//	@Failure		500	{object}	utils.ErrorResponse	"Internal server error"
//	@Router			/comments [get]
//	@Security		BearerAuth
func (ch *CommentController) MyAllComment(ctx *gin.Context) {
	userID, userIDIsExist := ctx.Get("userID")
	if !userIDIsExist {
		err := utils.ErrUnauthorized
		utils.HandleAbort(ctx, err)
		return
	}
	response, err := ch.commentService.MyAllComment(userID.(uint))
	if err != nil {
		utils.HandleError(ctx, err)
		return
	}
	utils.HandleSuccess(ctx, "Get my all comment successfully", ListCommentToCommentResponse(response))
}
