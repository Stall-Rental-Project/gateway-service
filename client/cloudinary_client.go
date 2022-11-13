package client

import (
	"errors"
	grpc "gateway-service/client/common"
	"gateway-service/config"
	"gateway-service/model"
	_ "gateway-service/model"
	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
	"github.com/gin-gonic/gin"
	_ "github.com/go-playground/validator/v10"
	"log"
)

type CloudinaryClient struct {
	cloudinary *cloudinary.Cloudinary
}

func NewCloudinaryClient() *CloudinaryClient {
	cld, err := cloudinary.NewFromParams(config.Config.CloudinaryCloudName, config.Config.CloudinaryApiKey, config.Config.CloudinaryApiSecret)
	if err != nil {
		log.Fatalln("Error when loading cloudinary config", err)
	}
	log.Println("Initialized cloudinary instance with config")

	return &CloudinaryClient{
		cloudinary: cld,
	}
}

func (service *CloudinaryClient) UploadFile(ctx *gin.Context, formKey string) (res *model.FileResponse, err error) {
	file, header, err := ctx.Request.FormFile(formKey)

	if err != nil {
		log.Println("Error when extracting file from request", err.Error())
		return nil, err
	}

	if header.Size > config.Config.UploadMaxSize<<(10*2) {
		log.Printf("Exceed max file size. Want %d, Received %d", config.Config.UploadMaxSize<<(10*2), header.Size)
		return nil, errors.New(grpc.ErrorCode_FILE_TOO_LARGE.String())
	}

	defer file.Close()

	fileName := header.Filename

	uploadParam, err := service.cloudinary.Upload.Upload(ctx, file, uploader.UploadParams{Folder: config.Config.CloudinaryUploadFolder})
	if err != nil {
		return nil, err
	}
	return &model.FileResponse{
		Content:      fileName,
		PreSignedUrl: uploadParam.URL,
		AccessId:     uploadParam.AssetID,
	}, nil
}

//func (service *CloudinaryClient) getUploadedFiles(c *gin.Context) {
//	fileName := c.Param("073571da-231c-49a5-90ec-fd81f4c67a3a.jpg")
//
//	// Access the filename using a desired file access id.
//	result, err := service.cloudinary.Admin.Asset(c, admin.AssetParams{
//		PublicID: fileName,
//	})
//	if err != nil {
//		log.Fatalf("Failed to get asset details, %v\n", err)
//	}
//	log.Printf("Public ID: %v, URL: %v\n", result.PublicID, result.SecureURL)
//
//	c.JSON(http.StatusAccepted, gin.H{
//		"message":    "Successfully found the image",
//		"secureURL":  result.SecureURL,
//		"publicURL":  result.URL,
//		"created_at": result.CreatedAt.String(),
//	})
//
//}
//
//func updateFile(c *gin.Context) {
//
//}
//
//func deleteFile(c *gin.Context) {
//
//}
