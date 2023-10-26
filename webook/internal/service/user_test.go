package service

import (
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func TestPasswordEncrypt(t *testing.T) {
	password := []byte("123456#123")
	encrypted, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	err = bcrypt.CompareHashAndPassword(encrypted, password)
	require.NoError(t, err)
}
