package entity

import (
	"uzi/internal/domain"

	"github.com/google/uuid"
)

type Node struct {
	Id       uuid.UUID `db:"id"`
	Ai       bool      `db:"ai"`
	MriID    uuid.UUID `db:"mri_id"`
	Knosp012 float64   `db:"knosp_012"`
	Knosp3   float64   `db:"knosp_3"`
	Knosp4   float64   `db:"knosp_4"`
}

func (Node) FromDomain(d domain.Node) Node {
	return Node{
		Id:       d.Id,
		Ai:       d.Ai,
		MriID:    d.MriID,
		Knosp012: d.Knosp012,
		Knosp3:   d.Knosp3,
		Knosp4:   d.Knosp4,
	}
}

func (d Node) ToDomain() domain.Node {
	return domain.Node{
		Id:       d.Id,
		Ai:       d.Ai,
		MriID:    d.MriID,
		Knosp012: d.Knosp012,
		Knosp3:   d.Knosp3,
		Knosp4:   d.Knosp4,
	}
}

func (Node) SliceToDomain(slice []Node) []domain.Node {
	res := make([]domain.Node, 0, len(slice))
	for _, v := range slice {
		res = append(res, v.ToDomain())
	}
	return res
}
