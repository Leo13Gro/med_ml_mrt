package entity

import (
	"time"

	"mri/internal/domain"

	"github.com/google/uuid"
)

type Mri struct {
	Id         uuid.UUID `db:"id"`
	Projection string    `db:"projection"`
	Checked    bool      `db:"checked"`
	PatientID  uuid.UUID `db:"patient_id"`
	DeviceID   int       `db:"device_id"`
	CreateAt   time.Time `db:"create_at"`
}

func (Mri) FromDomain(d domain.Mri) Mri {
	return Mri{
		Id:         d.Id,
		Projection: d.Projection,
		Checked:    d.Checked,
		PatientID:  d.PatientID,
		DeviceID:   d.DeviceID,
		CreateAt:   d.CreateAt,
	}
}

func (d Mri) ToDomain() domain.Mri {
	return domain.Mri{
		Id:         d.Id,
		Projection: d.Projection,
		Checked:    d.Checked,
		PatientID:  d.PatientID,
		DeviceID:   d.DeviceID,
		CreateAt:   d.CreateAt,
	}
}
