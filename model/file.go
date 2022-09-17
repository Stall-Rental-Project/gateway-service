package model

import "mime/multipart"

type RentalAttachmentRequest struct {
	Attachment []byte `json:"attachment"`
}

type FileResponse struct {
	Content      string `json:"content"`
	PreSignedUrl string `json:"pre_signed_url"`
	AccessId     string `json:"access_id"`
}

type File struct {
	File multipart.File `json:"file,omitempty" validate:"required"`
}

type Url struct {
	Url string `json:"url,omitempty" validate:"required"`
}
