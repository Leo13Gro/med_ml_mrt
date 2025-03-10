package domain

import (
	"time"

	"github.com/google/uuid"
)

type Uzi struct {
	Id         uuid.UUID
	Projection string
	Checked    bool
	PatientID  uuid.UUID
	DeviceID   int
	CreateAt   time.Time
}

type Image struct {
	Id    uuid.UUID
	MriID uuid.UUID
	Page  int
}

type Node struct {
	Id       uuid.UUID
	Ai       bool
	MriID    uuid.UUID
	Knosp012 float64
	Knosp3   float64
	Knosp4   float64
}

type Segment struct {
	Id       uuid.UUID
	ImageID  uuid.UUID
	NodeID   uuid.UUID
	Contor   string
	Knosp012 float64
	Knosp3   float64
	Knosp4   float64
}

type Device struct {
	Id   int
	Name string
}

type Echographic struct {
	Id              uuid.UUID
	Contors         *string
	LeftLobeLength  *float64
	LeftLobeWidth   *float64
	LeftLobeThick   *float64
	LeftLobeVolum   *float64
	RightLobeLength *float64
	RightLobeWidth  *float64
	RightLobeThick  *float64
	RightLobeVolum  *float64
	GlandVolum      *float64
	Isthmus         *float64
	Struct          *string
	Echogenicity    *string
	RegionalLymph   *string
	Vascularization *string
	Location        *string
	Additional      *string
	Conclusion      *string
}
