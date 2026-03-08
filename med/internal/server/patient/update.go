package patient

import (
	"context"
	"time"

	pb "med/internal/generated/grpc/service"
	"med/internal/server/mappers"
	"med/internal/services/patient"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *handler) UpdatePatient(ctx context.Context, in *pb.UpdatePatientIn) (*pb.UpdatePatientOut, error) {
	patientID, err := uuid.Parse(in.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Неверный формат ID пациента: %s", err.Error())
	}

	var lastExamDate *time.Time
	if in.LastExamDate != nil {
		lastExamDateParsed, err := time.Parse(time.RFC3339, *in.LastExamDate)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "Неверный формат даты последнего обследования: %s", err.Error())
		}
		lastExamDate = &lastExamDateParsed
	}

	patient, err := h.patientSrv.UpdatePatient(
		ctx,
		patientID,
		patient.UpdatePatient{
			Active:       in.Active,
			Malignancy:   in.Malignancy,
			LastExamDate: lastExamDate,
		},
	)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Что то пошло не так: %s", err.Error())
	}

	return &pb.UpdatePatientOut{Patient: mappers.PatientFromDomain(patient)}, nil
}
