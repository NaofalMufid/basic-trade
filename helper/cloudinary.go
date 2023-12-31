package helper

import (
	"basic-trade/config"
	"bytes"
	"context"
	"errors"
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
	maxSize := int64(3 << 20)
	if fileHeader.Size > maxSize {
		return "", errors.New("file size exceeds the maximum allowed size 3 mb")
	}

	allowedExtensions := map[string]bool{"jpg": true, "jpeg": true, "png": true}
	fileExt := strings.ToLower(filepath.Ext(fileHeader.Filename)[1:])
	if !allowedExtensions[fileExt] {
		return "", errors.New("invalid file extension, allowed extensions are jpg, jpeg, png")
	}

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
	parts := strings.Split(imageURL, "/")
	publicID := parts[len(parts)-1]
	publicID = strings.TrimSuffix(publicID, filepath.Ext(publicID))
	publicID = config.EnvCloudUploadFolder() + "/" + publicID

	return publicID, nil
}
