package node_segment

import (
	"context"

	"github.com/google/uuid"

	adapter "composition-api/internal/adapters/exam"
)

func (s *service) CreateNodeWithSegment(ctx context.Context, arg CreateNodeWithSegmentArg) (uuid.UUID, []uuid.UUID, error) {
	segments := make([]adapter.CreateNodeWithSegmentsIn_Segment, 0, len(arg.Segments))
	for _, segment := range arg.Segments {
		segments = append(segments, adapter.CreateNodeWithSegmentsIn_Segment{
			ImageID:   segment.ImageID,
			Contor:    segment.Contor,
			Tirads_23: segment.Tirads_23,
			Tirads_4:  segment.Tirads_4,
			Tirads_5:  segment.Tirads_5,
		})
	}

	nodeID, segmentIDs, err := s.adapters.Exam.CreateNodeWithSegments(ctx, adapter.CreateNodeWithSegmentsIn{
		MriID:    arg.MriID,
		Node:     adapter.CreateNodeWithSegmentsIn_Node(arg.Node),
		Segments: segments,
	})
	if err != nil {
		return uuid.Nil, nil, err
	}
	return nodeID, segmentIDs, nil
}
