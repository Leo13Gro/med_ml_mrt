package mri

import (
	"context"

	"mri/internal/domain"
	pb "mri/internal/generated/grpc/service"
	"mri/internal/services/mri"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type MriHandler interface {
	CreateMri(ctx context.Context, req *pb.CreateMriIn) (*pb.CreateMriOut, error)
	GetMri(ctx context.Context, in *pb.GetMriIn) (*pb.GetMriOut, error)
	GetPatientMris(ctx context.Context, in *pb.GetPatientMrisIn) (*pb.GetPatientMrisOut, error)
	GetEchographic(ctx context.Context, in *pb.GetEchographicIn) (*pb.GetEchographicOut, error)
	UpdateMri(ctx context.Context, req *pb.UpdateMriIn) (*pb.UpdateMriOut, error)
	UpdateEchographic(ctx context.Context, in *pb.UpdateEchographicIn) (*pb.UpdateEchographicOut, error)
}

type handler struct {
	mriSrv mri.Service
}

func New(
	mriSrv mri.Service,
) MriHandler {
	return &handler{
		mriSrv: mriSrv,
	}
}

func (h *handler) CreateMri(ctx context.Context, in *pb.CreateMriIn) (*pb.CreateMriOut, error) {
	uuid, err := h.mriSrv.CreateMri(ctx, domain.Mri{
		Projection: in.Projection,
		PatientID:  uuid.MustParse(in.PatientId),
		DeviceID:   int(in.DeviceId),
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Что то пошло не так: %s", err.Error())
	}

	return &pb.CreateMriOut{Id: uuid.String()}, nil
}

func (h *handler) GetEchographic(ctx context.Context, in *pb.GetEchographicIn) (*pb.GetEchographicOut, error) {
	echographic, err := h.mriSrv.GetMriEchographicsByID(ctx, uuid.MustParse(in.Id))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Что то пошло не так: %s", err.Error())
	}

	pbEchographic := domainEchographicToPb(&echographic)

	return &pb.GetEchographicOut{
		Echographic: pbEchographic,
	}, nil
}

func (h *handler) GetMri(ctx context.Context, in *pb.GetMriIn) (*pb.GetMriOut, error) {
	mri, err := h.mriSrv.GetMriByID(ctx, uuid.MustParse(in.Id))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Что то пошло не так: %s", err.Error())
	}

	pbMri := domainMriToPbMri(&mri)

	return &pb.GetMriOut{
		Mri: pbMri,
	}, nil
}

func (h *handler) GetPatientMris(ctx context.Context, in *pb.GetPatientMrisIn) (*pb.GetPatientMrisOut, error) {
	mris, err := h.mriSrv.GetMrisByPatientID(ctx, uuid.MustParse(in.PatientId))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Что то пошло не так: %s", err.Error())
	}

	resp := make([]*pb.Mri, 0, len(mris))
	for _, v := range mris {
		resp = append(resp, domainMriToPbMri(&v))
	}

	return &pb.GetPatientMrisOut{Mris: resp}, nil
}

func (h *handler) UpdateMri(ctx context.Context, in *pb.UpdateMriIn) (*pb.UpdateMriOut, error) {
	mri, err := h.mriSrv.UpdateMri(ctx,
		uuid.MustParse(in.Id),
		mri.UpdateMri{
			Projection: in.Projection,
			Checked:    in.Checked,
		},
	)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Что то пошло не так: %s", err.Error())
	}

	return &pb.UpdateMriOut{
		Mri: domainMriToPbMri(&mri),
	}, nil
}

func (h *handler) UpdateEchographic(ctx context.Context, in *pb.UpdateEchographicIn) (*pb.UpdateEchographicOut, error) {
	echographic, err := h.mriSrv.UpdateEchographic(
		ctx,
		uuid.MustParse(in.Echographic.Id),
		mri.UpdateEchographic{
			Contors:         in.Echographic.Contors,
			LeftLobeLength:  in.Echographic.LeftLobeLength,
			LeftLobeWidth:   in.Echographic.LeftLobeWidth,
			LeftLobeThick:   in.Echographic.LeftLobeThick,
			LeftLobeVolum:   in.Echographic.LeftLobeVolum,
			RightLobeLength: in.Echographic.RightLobeLength,
			RightLobeWidth:  in.Echographic.RightLobeWidth,
			RightLobeThick:  in.Echographic.RightLobeThick,
			RightLobeVolum:  in.Echographic.RightLobeVolum,
			GlandVolum:      in.Echographic.GlandVolum,
			Isthmus:         in.Echographic.Isthmus,
			Struct:          in.Echographic.Struct,
			Echogenicity:    in.Echographic.Echogenicity,
			RegionalLymph:   in.Echographic.RegionalLymph,
			Vascularization: in.Echographic.Vascularization,
			Location:        in.Echographic.Location,
			Additional:      in.Echographic.Additional,
			Conclusion:      in.Echographic.Conclusion,
		})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Что то пошло не так: %s", err.Error())
	}

	return &pb.UpdateEchographicOut{
		Echographic: domainEchographicToPb(&echographic),
	}, nil
}
