package utils

import (
	"context"
	"fmt"
	"log"
	"mime/multipart"
	"path/filepath"
	"strings"
	"time"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/roihan12/h8-mygram/app/config"
)

type Uploader interface {
	Upload(file *multipart.FileHeader) (string, error)
	Destroy(publicID string) error
}

type claudinaryUploader struct {
	cld *cloudinary.Cloudinary
}

func NewCloudinary(cfg *config.AppConfig) Uploader {
	cld, err := cloudinary.NewFromParams(cfg.CLOUDINARY_CLOUD_NAME, cfg.CLOUDINARY_API_KEY, cfg.CLOUDINARY_API_SECRET)
	if err != nil {
		log.Println("init cloudinary gagal", err)
		return nil
	}

	return &claudinaryUploader{cld: cld}
}

func GetPublicID(secureURL string) string {
	// Mengambil bagian terakhir dari URL sebagai nama file
	fileName := filepath.Base(secureURL)
	// Menghapus ekstensi file jika ada
	fileNameWithoutExt := fileName[:len(fileName)-len(filepath.Ext(fileName))]
	// Mengembalikan nama file tanpa ekstensi
	return fileNameWithoutExt
}

func (cu *claudinaryUploader) Upload(file *multipart.FileHeader) (string, error) {

	src, _ := file.Open()
	defer src.Close()

	publicID := fmt.Sprintf("%d-%s", int(file.Size), time.Now().Format("20060102-150405")) // Format  "file_size-(YY-MM-DD)-(hh-mm-ss)""

	uploadResult, err := cu.cld.Upload.Upload(
		context.Background(),
		src,
		uploader.UploadParams{
			PublicID: publicID,
		})
	if err != nil {
		return "", err
	}

	return uploadResult.SecureURL, nil
}

func (cu *claudinaryUploader) Destroy(publicID string) error {
	res, err := cu.cld.Upload.Destroy(
		context.Background(),
		uploader.DestroyParams{
			PublicID: publicID,
		},
	)

	if err != nil {
		return err
	}

	if strings.Contains(res.Result, "not found") {
		return ErrDataNotFound
	}

	fmt.Printf("%+v\n", res)

	return nil
}
