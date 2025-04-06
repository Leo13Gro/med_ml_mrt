package repository

import (
	"github.com/WantBeASleep/goooool/daolib"

	"mri/internal/repository/entity"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

const mriTable = "mri"

type MriQuery interface {
	InsertMri(mri entity.Mri) error
	CheckExist(id uuid.UUID) (bool, error)
	GetMriByPK(id uuid.UUID) (entity.Mri, error)
	GetMrisByPatientID(patientID uuid.UUID) ([]entity.Mri, error)
	UpdateMri(mri entity.Mri) (int64, error)
}

type mriQuery struct {
	*daolib.BaseQuery
}

func (q *mriQuery) SetBaseQuery(baseQuery *daolib.BaseQuery) {
	q.BaseQuery = baseQuery
}

func (q *mriQuery) InsertMri(mri entity.Mri) error {
	query := q.QueryBuilder().
		Insert(mriTable).
		Columns(
			"id",
			"projection",
			"checked",
			"patient_id",
			"device_id",
			"create_at",
		).
		Values(
			mri.Id,
			mri.Projection,
			mri.Checked,
			mri.PatientID,
			mri.DeviceID,
			mri.CreateAt,
		)

	_, err := q.Runner().Execx(q.Context(), query)
	if err != nil {
		return err
	}

	return nil
}

func (q *mriQuery) GetMriByPK(id uuid.UUID) (entity.Mri, error) {
	query := q.QueryBuilder().
		Select(
			"id",
			"projection",
			"checked",
			"patient_id",
			"device_id",
			"create_at",
		).
		From(mriTable).
		Where(sq.Eq{
			"id": id,
		})

	var mri entity.Mri
	if err := q.Runner().Getx(q.Context(), &mri, query); err != nil {
		return entity.Mri{}, err
	}

	return mri, nil
}

func (q *mriQuery) GetMrisByPatientID(patientID uuid.UUID) ([]entity.Mri, error) {
	query := q.QueryBuilder().
		Select(
			"id",
			"projection",
			"checked",
			"patient_id",
			"device_id",
			"create_at",
		).
		From(mriTable).
		Where(sq.Eq{
			"patient_id": patientID,
		})

	var mri []entity.Mri
	if err := q.Runner().Selectx(q.Context(), &mri, query); err != nil {
		return nil, err
	}

	return mri, nil
}

func (q *mriQuery) CheckExist(id uuid.UUID) (bool, error) {
	query := q.QueryBuilder().
		Select(
			"id",
			"projection",
			"checked",
			"patient_id",
			"device_id",
			"create_at",
		).
		Prefix("SELECT EXISTS (").
		From(mriTable).
		Where(sq.Eq{
			"id": id,
		}).
		Suffix(")")

	var exists bool
	if err := q.Runner().Getx(q.Context(), &exists, query); err != nil {
		return false, err
	}

	return exists, nil
}

func (q *mriQuery) UpdateMri(mri entity.Mri) (int64, error) {
	query := q.QueryBuilder().
		Update(mriTable).
		SetMap(sq.Eq{
			"projection": mri.Projection,
			"checked":    mri.Checked,
		}).
		Where(sq.Eq{
			"id": mri.Id,
		})

	rows, err := q.Runner().Execx(q.Context(), query)
	if err != nil {
		return 0, err
	}

	return rows.RowsAffected()
}
