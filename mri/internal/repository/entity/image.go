package entity

import (
	"uzi/internal/domain"

	"github.com/google/uuid"
)

type Image struct {
	Id    uuid.UUID `db:"id"`
	MriID uuid.UUID `db:"mri_id"`
	Page  int       `db:"page"`
}

func (Image) FromDomain(d domain.Image) Image {
	return Image{
		Id:    d.Id,
		MriID: d.MriID,
		Page:  d.Page,
	}
}

func (d Image) ToDomain() domain.Image {
	return domain.Image{
		Id:    d.Id,
		MriID: d.MriID,
		Page:  d.Page,
	}
}

func (Image) SliceToDomain(slice []Image) []domain.Image {
	res := make([]domain.Image, 0, len(slice))
	for _, v := range slice {
		res = append(res, v.ToDomain())
	}
	return res
}
