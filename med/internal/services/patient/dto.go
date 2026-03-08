package patient

import (
	"time"

	"med/internal/domain"
)

type UpdatePatient struct {
	Active       *bool
	Malignancy   *bool
	LastExamDate *time.Time
}

func (u UpdatePatient) Update(d *domain.Patient) {
	if u.Active != nil {
		d.Active = *u.Active
	}
	if u.Malignancy != nil {
		d.Malignancy = *u.Malignancy
	}
	if u.LastExamDate != nil {
		d.LastExamDate = u.LastExamDate
	}
}
