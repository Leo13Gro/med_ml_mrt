package node

import (
	"context"

	"github.com/AlekSi/pointer"

	domain "composition-api/internal/domain/exam"
	api "composition-api/internal/generated/http/api"
	mappers "composition-api/internal/server/exam/mappers"
	apimappers "composition-api/internal/server/mappers"
	nodeSrv "composition-api/internal/services/node"
)

func (h *handler) MriNodesIDPatch(ctx context.Context, req *api.MriNodesIDPatchReq, params api.MriNodesIDPatchParams) (api.MriNodesIDPatchRes, error) {
	var validation *domain.NodeValidation
	switch {
	case req.Validation.Null:
		validation = pointer.To(domain.NodeValidationNull)
	case req.Validation.IsSet():
		validation = pointer.To(domain.NodeValidation(req.Validation.Value))
	}

	node, err := h.services.NodeService.UpdateNode(ctx, nodeSrv.UpdateNodeArg{
		Id:         params.ID,
		Validation: validation,
		Tirads_23:  apimappers.FromOptFloat64(req.Tirads23),
		Tirads_4:   apimappers.FromOptFloat64(req.Tirads4),
		Tirads_5:   apimappers.FromOptFloat64(req.Tirads5),
	})
	if err != nil {
		return nil, err
	}

	return pointer.To(mappers.Node{}.Domain(node)), nil
}
