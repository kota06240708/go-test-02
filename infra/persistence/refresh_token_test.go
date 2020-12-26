package persistence_test

import (
	"testing"
	"time"

	"gopkg.in/go-playground/assert.v1"
)

const token = "123455678"
const tokenError = "12345"

func TestAddRefreshToken(t *testing.T) {
	maxRefresh := time.Hour * 24
	now := time.Now()
	expire := now.Add(maxRefresh)
	nowError := time.Date(2014, time.December, 31, 12, 13, 24, 0, time.UTC)
	err := refreshTokenPersistence.AddRefreshToken(db, token, &expire)

	errE := refreshTokenPersistence.AddRefreshToken(db, tokenError, &nowError)

	assert.Equal(t, err, nil)
	assert.Equal(t, errE, nil)
}

func TestCheckRefreshToken(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		err := refreshTokenPersistence.CheckRefreshToken(db, token)

		assert.Equal(t, err, nil)
	})

	t.Run("error token", func(t *testing.T) {
		err := refreshTokenPersistence.CheckRefreshToken(db, "19")

		assert.Equal(t, err != nil, true)
	})

	t.Run("error expire", func(t *testing.T) {
		err := refreshTokenPersistence.CheckRefreshToken(db, tokenError)

		assert.Equal(t, err != nil, true)
	})
}
