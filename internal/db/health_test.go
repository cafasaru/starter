// +build integration

package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPingDB(t *testing.T) {
	t.Run("test ping database", func(t *testing.T) {
		db, err := NewDatabase()
		assert.NoError(t, err)

		err = db.PingDB(context.Background())
		assert.NoError(t, err)

		fmt.Println("tessting the health of the database connection.")
	})
}
