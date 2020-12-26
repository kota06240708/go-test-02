package persistence_test

import (
	"testing"
	"time"

	"gopkg.in/go-playground/assert.v1"
)

const token = "123455678"

func TestAddRefreshToken(t *testing.T) {
	now := time.Now()
	err := refreshTokenPersistence.AddRefreshToken(db, token, &now)

	assert.Equal(t, err, nil)
}

func TestCheckRefreshToken(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		err := refreshTokenPersistence.CheckRefreshToken(db, token)

		assert.Equal(t, err, nil)
	})

	t.Run("error", func(t *testing.T) {
		err := refreshTokenPersistence.CheckRefreshToken(db, "19")

		assert.Equal(t, err != nil, true)
	})
}
