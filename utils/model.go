package utils

import (
	"simplify-go/model/entity"
	"simplify-go/model/web"
)

func ToPhotosResponse(example entity.Photos) web.PhotoResponse {
	return web.PhotoResponse{
		Id:        example.Id,
		Name:      example.Name,
		Email:     example.Email,
		CreatedAt: example.CreatedAt,
		UpdatedAt: example.UpdatedAt,
	}
}
