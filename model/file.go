package model

type RentalAttachmentRequest struct {
	Attachment []byte `json:"attachment"`
	Type       int64  `json:"type"`
}

type FileResponse struct {
	Content          string `json:"content"`
	OriginalFilename string `json:"original_filename"`
	PreSignedUrl     string `json:"pre_signed_url"`
}
