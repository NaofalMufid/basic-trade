package helper

import (
	"basic-trade/config"
	"bytes"
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

func UploadFile(fileHeader *multipart.FileHeader, fileName string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cld, err := cloudinary.NewFromParams(config.EnvCloudName(), config.EnvCloudAPIKey(), config.EnvCloudAPISecret())
	if err != nil {
		return "", err
	}

	fileReader, err := convertFile(fileHeader)
	if err != nil {
		return "", err
	}

	uploadParam, err := cld.Upload.Upload(ctx, fileReader, uploader.UploadParams{
		PublicID: fileName,
		Folder:   config.EnvCloudUploadFolder(),
	})
	if err != nil {
		return "", err
	}

	return uploadParam.SecureURL, nil
}

func convertFile(fileHeader *multipart.FileHeader) (*bytes.Reader, error) {
	file, err := fileHeader.Open()
	if err != nil {
		return nil, err
	}
	defer file.Close()

	buffer := new(bytes.Buffer)
	if _, err := io.Copy(buffer, file); err != nil {
		return nil, err
	}

	fileReader := bytes.NewReader(buffer.Bytes())
	return fileReader, nil
}

func RemoveExtension(filename string) string {
	return path.Base(filename[:len(filename)-len(path.Ext(filename))])
}

func DeleteFile(publicID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cld, err := cloudinary.NewFromParams(config.EnvCloudName(), config.EnvCloudAPIKey(), config.EnvCloudAPISecret())
	if err != nil {
		return err
	}

	newPublicID, err := getPublicIDFromURL(publicID)
	if err != nil {
		fmt.Println("Error", err)
	}
	_, err = cld.Upload.Destroy(ctx, uploader.DestroyParams{
		PublicID: newPublicID,
	})
	if err != nil {
		return err
	}
	return nil
}

func getPublicIDFromURL(imageURL string) (string, error) {
	parts := strings.Split(imageURL, "/upload/")
	if len(parts) != 2 {
		return "", fmt.Errorf("invalid image url")
	}

	// Ambil bagian setelah "/upload/" sebagai Public ID
	publicID := parts[1]

	// Hapus ekstensi file jika ada (misalnya ".jpg")
	publicID = strings.TrimSuffix(publicID, filepath.Ext(publicID))

	return publicID, nil
}
