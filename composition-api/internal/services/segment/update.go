package segment

import (
	"context"

	adapter "composition-api/internal/adapters/exam"
	domain "composition-api/internal/domain/exam"
)

func (s *service) Update(ctx context.Context, arg UpdateSegmentArg) (domain.Segment, error) {
	segment, err := s.adapters.Exam.UpdateSegment(ctx, adapter.UpdateSegmentIn{
		Id:        arg.Id,
		Contor:    arg.Contor,
		Tirads_23: arg.Tirads_23,
		Tirads_4:  arg.Tirads_4,
		Tirads_5:  arg.Tirads_5,
	})
	if err != nil {
		return domain.Segment{}, err
	}
	return segment, nil
}
