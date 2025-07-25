// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0

package swodb

import (
	"database/sql/driver"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
)

type EnumSwitchoverState string

const (
	EnumSwitchoverStateIdle       EnumSwitchoverState = "idle"
	EnumSwitchoverStateInProgress EnumSwitchoverState = "in_progress"
	EnumSwitchoverStateUseNextDb  EnumSwitchoverState = "use_next_db"
)

func (e *EnumSwitchoverState) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = EnumSwitchoverState(s)
	case string:
		*e = EnumSwitchoverState(s)
	default:
		return fmt.Errorf("unsupported scan type for EnumSwitchoverState: %T", src)
	}
	return nil
}

type NullEnumSwitchoverState struct {
	EnumSwitchoverState EnumSwitchoverState
	Valid               bool // Valid is true if EnumSwitchoverState is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullEnumSwitchoverState) Scan(value interface{}) error {
	if value == nil {
		ns.EnumSwitchoverState, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.EnumSwitchoverState.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullEnumSwitchoverState) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.EnumSwitchoverState), nil
}

type ChangeLog struct {
	ID        int64
	TableName string
	RowID     string
}

type PgStatActivity struct {
	State           pgtype.Text
	XactStart       pgtype.Timestamptz
	ApplicationName pgtype.Text
}

type SwitchoverLog struct {
	ID        int64
	Timestamp pgtype.Timestamptz
	Data      []byte
}

type SwitchoverState struct {
	Ok           bool
	CurrentState EnumSwitchoverState
	DbID         pgtype.UUID
}
