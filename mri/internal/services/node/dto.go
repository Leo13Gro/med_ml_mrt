package node

import "uzi/internal/domain"

// TODO: починить баг при запросе со всеми полями nil
type UpdateNode struct {
	Knosp012 *float64
	Knosp3   *float64
	Knosp4   *float64
}

func (u UpdateNode) Update(d *domain.Node) {
	if u.Knosp012 != nil {
		d.Knosp012 = *u.Knosp012
	}
	if u.Knosp3 != nil {
		d.Knosp3 = *u.Knosp3
	}
	if u.Knosp4 != nil {
		d.Knosp4 = *u.Knosp4
	}
}
