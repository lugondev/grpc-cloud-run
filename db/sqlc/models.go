// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package db

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type StatusUserGroup string

const (
	StatusUserGroupAccepted StatusUserGroup = "accepted"
	StatusUserGroupRejected StatusUserGroup = "rejected"
	StatusUserGroupPending  StatusUserGroup = "pending"
)

func (e *StatusUserGroup) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = StatusUserGroup(s)
	case string:
		*e = StatusUserGroup(s)
	default:
		return fmt.Errorf("unsupported scan type for StatusUserGroup: %T", src)
	}
	return nil
}

type NullStatusUserGroup struct {
	StatusUserGroup StatusUserGroup
	Valid           bool // Valid is true if String is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullStatusUserGroup) Scan(value interface{}) error {
	if value == nil {
		ns.StatusUserGroup, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.StatusUserGroup.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullStatusUserGroup) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return ns.StatusUserGroup, nil
}

type Account struct {
	ID        int64     `json:"id"`
	Owner     string    `json:"owner"`
	Balance   int64     `json:"balance"`
	Currency  string    `json:"currency"`
	CreatedAt time.Time `json:"created_at"`
}

type Approval struct {
	ID         uuid.UUID `json:"id"`
	GroupID    int64     `json:"group_id"`
	Blockchain string    `json:"blockchain"`
	Network    string    `json:"network"`
	Amount     int64     `json:"amount"`
	Status     bool      `json:"status"`
	ExpiresAt  time.Time `json:"expires_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	CreatedAt  time.Time `json:"created_at"`
}

type Entry struct {
	ID        int64 `json:"id"`
	AccountID int64 `json:"account_id"`
	// can be negative or positive
	Amount    int64     `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

type Group struct {
	ID          int64     `json:"id"`
	GroupName   string    `json:"group_name"`
	Owner       string    `json:"owner"`
	Status      bool      `json:"status"`
	Deactivated bool      `json:"deactivated"`
	UpdatedAt   time.Time `json:"updated_at"`
	CreatedAt   time.Time `json:"created_at"`
}

type Transfer struct {
	ID            int64 `json:"id"`
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	// must be positive
	Amount    int64     `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

type UsersGroup struct {
	UserID    string          `json:"user_id"`
	GroupID   int64           `json:"group_id"`
	Weight    sql.NullString  `json:"weight"`
	Threshold string          `json:"threshold"`
	Status    StatusUserGroup `json:"status"`
}
