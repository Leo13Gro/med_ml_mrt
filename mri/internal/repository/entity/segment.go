package entity

import (
	"uzi/internal/domain"

	"github.com/google/uuid"
)

type Segment struct {
	Id       uuid.UUID `db:"id"`
	ImageID  uuid.UUID `db:"image_id"`
	NodeID   uuid.UUID `db:"node_id"`
	Contor   string    `db:"contor"`
	Knosp012 float64   `db:"knosp_012"`
	Knosp3   float64   `db:"knosp_3"`
	Knosp4   float64   `db:"knosp_4"`
}

func (Segment) FromDomain(d domain.Segment) Segment {
	return Segment{
		Id:       d.Id,
		ImageID:  d.ImageID,
		NodeID:   d.NodeID,
		Contor:   d.Contor,
		Knosp012: d.Knosp012,
		Knosp3:   d.Knosp3,
		Knosp4:   d.Knosp4,
	}
}

func (d Segment) ToDomain() domain.Segment {
	return domain.Segment{
		Id:       d.Id,
		ImageID:  d.ImageID,
		NodeID:   d.NodeID,
		Contor:   d.Contor,
		Knosp012: d.Knosp012,
		Knosp3:   d.Knosp3,
		Knosp4:   d.Knosp4,
	}
}

func (Segment) SliceToDomain(slice []Segment) []domain.Segment {
	res := make([]domain.Segment, 0, len(slice))
	for _, v := range slice {
		res = append(res, v.ToDomain())
	}
	return res
}
