package application

import (
	"bytes"
	"context"
	"encoding/base64"
	"net/http"
	"slices"
	"teams_service/internal/core"
	"teams_service/internal/core/cerror"
	"teams_service/internal/core/dto"

	"github.com/minio/minio-go/v7"
)

var allowedMimeTypes = []string{"image/jpeg", "image/png", "image/webp"}

var exts = map[string]string{
	"image/jpeg": ".jpeg",
	"image/png":  ".png",
	"image/webp": ".webp",
}

func (u *useCase) UploadPhoto(c context.Context, dto *dto.UploadPhoto) (string, error) {
	err := u.checkPermission(c, dto.TeamId, dto.MemberId)
	if err != nil {
		return "", cerror.Forbidden()
	}

	file, err := base64.StdEncoding.DecodeString(dto.Photo)
	if err != nil {
		return "", err
	}

	mime := http.DetectContentType(file)
	if !slices.Contains(allowedMimeTypes, mime) {
		return "", cerror.New(cerror.UNSUPPORTED_FORMAT, "this file format is not supported")
	}

	randomStr, err := core.GenerateRandomString(64)
	if err != nil {
		return "", err
	}

	objectName := randomStr + exts[mime]

	opts := minio.PutObjectOptions{ContentType: mime}
	if _, err := u.s3.PutObject(c, "teams", objectName, bytes.NewReader(file), int64(len(file)), opts); err != nil {
		return "", err
	}

	return objectName, u.repo.UpdatePhoto(c, dto.TeamId, objectName)
}
