package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewUser(t *testing.T) {
	usr, err := NewUser("","", "", "")
	assert.NotNil(t, err)

	usr, err = NewUser("Satoshi","satoshi@btc.com", "btc@pswd#", "1975/12/31")
	assert.Nil(t, err)
	assert.Equal(t, usr.Name, "Satoshi")
	assert.Greater(t, usr.ID, int64(0))
	assert.NotEqual(t, usr.Password, "btc@pswd#")
}

func TestValidatePassword(t *testing.T) {
	usr, _ := NewUser("Satoshi","satoshi@btc.com", "btc@pswd#", "1975/12/31")
	valid := usr.ValidatePassword("btc@pswd#")
	assert.True(t, valid)

	valid = usr.ValidatePassword("wrong_password")
	assert.False(t, valid)
}