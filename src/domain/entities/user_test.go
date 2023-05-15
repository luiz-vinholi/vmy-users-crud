package entities

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestUserSetAge(t *testing.T) {
	assert := assert.New(t)

	today := time.Now()
	y, m, d := today.Date()

	date := time.Date(y-18, m, d, 0, 0, 0, 0, time.Local)
	birthDate := date.Format("2006-01-02")
	user := User{
		BirthDate: birthDate,
	}
	err := user.SetAge()

	assert.Equal(user.Age, 18)
	assert.Nil(err)
}

func TestUserSetAgeWithInvalidBirthDate(t *testing.T) {
	assert := assert.New(t)

	user := User{
		BirthDate: "invaliddate",
	}
	err := user.SetAge()

	assert.NotNil(err)
}
