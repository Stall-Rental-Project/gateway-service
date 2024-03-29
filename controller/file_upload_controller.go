package controller

import (
	"gateway-service/client"
	grpc "gateway-service/client/common"
	"gateway-service/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type FileUploadController struct {
	fileClient *client.CloudinaryClient
}

func NewFileController(
	fileClient *client.CloudinaryClient) FileUploadController {
	return FileUploadController{
		fileClient: fileClient,
	}
}

// UploadAttachment
// @Router   /api/v2/files/upload [POST]
// @Summary  Upload attachment
// @Param    _  body model.RentalAttachmentRequest  true  "Body"
// @Tags     File
// @Accept   mpfd
// @Produce  json
// @Success  200          {object}  model.FileResponse
// @Failure  400,401,500  {object}  model.ErrorResponse
func (controller *FileUploadController) UploadAttachment(ctx *gin.Context) {
	resp, err := controller.fileClient.UploadFile(ctx, "attachment")
	if err != nil {
		if err.Error() == grpc.ErrorCode_FILE_TOO_LARGE.String() {
			ctx.JSON(http.StatusBadRequest, &model.ErrorResponse{
				ErrorCode:        grpc.ErrorCode_FILE_TOO_LARGE.String(),
				ErrorDescription: "Exceeds maximum file size",
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, &model.ErrorResponse{
				ErrorCode:        grpc.ErrorCode_INTERNAL_SERVER_ERROR.String(),
				ErrorDescription: err.Error(),
			})
		}
	} else {
		ctx.JSON(http.StatusOK, resp)
	}
}

// GetStaticFile
// @Summary Get static file
// @Tags    File
// @Accept  mpfd
// @Param   name    query    string true "filename"
// @Success 200     {object} object
// @Failure 400,500 {object} model.ErrorResponse
// @Router  /api/v2/files/static [GET]
func (controller *FileUploadController) GetStaticFile(ctx *gin.Context) {
	ctx.FileAttachment("assets/static", ctx.Query("name"))
}

func (controller *FileUploadController) GetFile(ctx *gin.Context) {
	path := ctx.Param("path")
	ctx.File("assets/images/" + path)
}
