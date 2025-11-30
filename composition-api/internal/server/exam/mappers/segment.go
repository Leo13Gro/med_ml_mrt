package mappers

import (
	domain "composition-api/internal/domain/exam"
	api "composition-api/internal/generated/http/api"
)

type Segment struct{}

func (Segment) Domain(segment domain.Segment) (api.Segment, error) {
	contor, err := Contor(segment.Contor)
	if err != nil {
		return api.Segment{}, err
	}

	return api.Segment{
		ID:       segment.Id,
		ImageID:  segment.ImageID,
		NodeID:   segment.NodeID,
		Contor:   contor,
		Ai:       segment.Ai,
		Knosp012: segment.Knosp012,
		Knosp3:   segment.Knosp3,
		Knosp4:   segment.Knosp4,
	}, nil
}

func (Segment) SliceDomain(segments []domain.Segment) ([]api.Segment, error) {
	result := make([]api.Segment, 0, len(segments))
	for _, segment := range segments {
		segment, err := Segment{}.Domain(segment)
		if err != nil {
			return nil, err
		}
		result = append(result, segment)
	}
	return result, nil
}
