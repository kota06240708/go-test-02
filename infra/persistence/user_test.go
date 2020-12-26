package persistence_test

import (
	"testing"

	"github.com/api/domain/model"

	"gopkg.in/go-playground/assert.v1"
)

func TestGetAllUser(t *testing.T) {
	users, err := userPersistence.GetAll(db)

	assert.Equal(t, err, nil)
	assert.Equal(t, len(users), 2)
}

func TestGetCurrentUser(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		user, err := userPersistence.GetCurrentUser(db, "example01@gmail.com")

		assert.Equal(t, err, nil)
		assert.Equal(t, user.Name, "test01")
	})

	t.Run("error", func(t *testing.T) {
		_, err := userPersistence.GetCurrentUser(db, "exam1@gmail.com")

		assert.Equal(t, err != nil, true)
	})
}

func TestGetCurrentUserID(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		user, err := userPersistence.GetCurrentUserID(db, 1)

		assert.Equal(t, err, nil)
		assert.Equal(t, user.Name, "test01")
	})

	t.Run("error", func(t *testing.T) {
		_, err := userPersistence.GetCurrentUserID(db, 10000000)

		assert.Equal(t, err != nil, true)
	})
}

func TestAddUser(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		err := userPersistence.AddUser(db, "test00001", 1000, "", "121121", "example0001.gmail.com")

		var user model.User

		db.Where("email = ?", "example0001.gmail.com").First(&user)

		assert.Equal(t, err, nil)
		assert.Equal(t, user.Name, "test00001")
	})

	t.Run("error email double", func(t *testing.T) {
		err := userPersistence.AddUser(db, "test00001", 1000, "", "121121", "example01@gmail.com")

		assert.Equal(t, err != nil, true)
	})
}

func TestUpdateUser(t *testing.T) {
	var user model.User
	const text = "更新"

	db.Where("email = ?", "example01@gmail.com").First(&user)

	user.Name = text

	err := userPersistence.UpdateUser(db, &user)

	var userUpdate model.User

	db.Where("email = ?", "example01@gmail.com").First(&userUpdate)

	assert.Equal(t, err, nil)
	assert.Equal(t, userUpdate.Name, text)
}

func TestDeleteUser(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		var user model.User

		err := userPersistence.DeleteUser(db, 1)

		errDB := db.Where("id = ?", 1).First(&user).Error

		assert.Equal(t, err, nil)
		assert.Equal(t, errDB != nil, true)
	})

	t.Run("error", func(t *testing.T) {
		err := userPersistence.DeleteUser(db, 10000000)

		assert.Equal(t, err != nil, true)
	})
}
