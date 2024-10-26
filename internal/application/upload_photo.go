package application

import (
	"bytes"
	"context"
	"encoding/base64"
	"image"
	"image/jpeg"
	"image/png"
	"math"
	"net/http"
	"slices"
	"teams_service/internal/core/cerror"
	"teams_service/internal/core/dto"
	"teams_service/internal/core/models"

	"golang.org/x/image/draw"
	"golang.org/x/image/webp"
)

var allowedMimeTypes = []string{"image/jpeg", "image/png", "image/webp"}

func (u *useCase) UploadPhoto(c context.Context, user *models.User, dto *dto.UploadPhoto) (string, error) {
	err := u.checkPermission(c, dto.TeamId, user.ID, UPDATE_TEAM_PERMISSION)
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

	newFile, err := thumbnail(file, 512)
	if err != nil {
		return "", err
	}

	filename, err := u.s3.Put(c, newFile)
	if err != nil {
		return "", err
	}

	return filename, u.repo.UpdatePhoto(c, dto.TeamId, filename)
}

func thumbnail(file []byte, width int) ([]byte, error) {
	var src image.Image
	var output bytes.Buffer
	var err error

	mimetype := http.DetectContentType(file)

	r := bytes.NewReader(file)

	switch mimetype {
	case "image/jpeg":
		src, err = jpeg.Decode(r)
	case "image/png":
		src, err = png.Decode(r)
	case "image/webp":
		src, err = webp.Decode(r)
	}

	if err != nil {
		return []byte{}, err
	}

	if src.Bounds().Dx() <= 512 {
		return file, nil
	}

	ratio := (float64)(src.Bounds().Max.Y) / (float64)(src.Bounds().Max.X)
	height := int(math.Round(float64(width) * ratio))

	dst := image.NewRGBA(image.Rect(0, 0, width, height))

	draw.NearestNeighbor.Scale(dst, dst.Rect, src, src.Bounds(), draw.Over, nil)

	err = jpeg.Encode(&output, dst, nil)
	if err != nil {
		return []byte{}, err
	}

	return output.Bytes(), err
}
