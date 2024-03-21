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

func (ch *CommentController) GetById(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("commentId"))
	response, err := ch.commentService.GetById(uint(id))
	if err != nil {
		utils.HandleError(ctx, err)
		return
	}
	utils.HandleSuccess(ctx, "Get comment by id successfully", CommentEntityToCommentResponse(response))
}
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
