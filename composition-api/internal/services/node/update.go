package node

import (
	"context"

	adapter "composition-api/internal/adapters/exam"
	domain "composition-api/internal/domain/exam"
)

func (s *service) UpdateNode(ctx context.Context, arg UpdateNodeArg) (domain.Node, error) {
	node, err := s.adapters.Exam.UpdateNode(ctx, adapter.UpdateNodeIn{
		Id:         arg.Id,
		Validation: arg.Validation,
		Tirads_23:  arg.Tirads_23,
		Tirads_4:   arg.Tirads_4,
		Tirads_5:   arg.Tirads_5,
	})
	if err != nil {
		return domain.Node{}, err
	}
	return node, nil
}
