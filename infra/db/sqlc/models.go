// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0

package sqlc

import (
	"time"

	"github.com/google/uuid"
)

type Account struct {
	ID        uuid.UUID
	Owner     string
	Balance   int64
	Currency  string
	CreatedAt time.Time
}

type Entry struct {
	ID        uuid.UUID
	AccountID uuid.UUID
	// can be negative or positive
	Amount    int64
	CreatedAt time.Time
}

type Transfer struct {
	ID            uuid.UUID
	FromAccountID uuid.UUID
	ToAccountID   uuid.UUID
	// must be positive
	Amount    int64
	CreatedAt time.Time
}
