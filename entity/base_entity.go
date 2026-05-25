package entity

import (
	"time"

	"github.com/google/uuid"
)

// BaseEntity collects common persistence fields for sqlx/xsql-based entities.
type BaseEntity struct {
	ID         uuid.UUID  `db:"id" json:"id"`
	CreateBy   *uuid.UUID `db:"create_by" json:"createBy"`
	CreateTime NullTime   `db:"create_time" json:"createTime"`
	UpdateBy   *uuid.UUID `db:"update_by" json:"updateBy"`
	UpdateTime NullTime   `db:"update_time" json:"updateTime"`
	Version    int        `db:"version" json:"version"`
	OwnerID    *uuid.UUID `db:"owner_id" json:"ownerId"`
	EntryType  *string    `db:"entry_type" json:"entryType"`
}

func (e *BaseEntity) PrepareForCreate(actorID *uuid.UUID) {
	now := time.Now().UTC()
	e.ID = uuid.Must(uuid.NewV7())
	e.CreateBy = actorID
	e.CreateTime = NewTime(now)
	e.UpdateBy = actorID
	e.UpdateTime = NewTime(now)
	e.Version = 1
}

func (e *BaseEntity) PrepareForUpdate(actorID *uuid.UUID) {
	now := time.Now().UTC()
	e.UpdateBy = actorID
	e.UpdateTime = NewTime(now)
}
