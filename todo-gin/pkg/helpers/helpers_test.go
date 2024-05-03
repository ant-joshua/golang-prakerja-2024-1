package helpers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSumSuccess(t *testing.T) {
	result := Sum(1, 2)

	assert.Equal(t, result, 3, "Sum(1, 2) = 3; want 3")
}

func TestSumFailed(t *testing.T) {
	result := Sum(1, 2)

	assert.NotEqual(t, result, 4, "Sum(1, 2) = 3; want not 4")

	t.Log("TestSumFailed done")
}

func TestHashPassword(t *testing.T) {
	password := "password"
	hashedPassword, err := HashPassword(password)

	assert.Nil(t, err, "HashPassword(%s) = %v; want nil", password, err)

	assert.NotEmpty(t, hashedPassword, "HashPassword(%s) = %v; want not empty", password, hashedPassword)
}

func TestComparePasswordSuccess(t *testing.T) {
	password := "password"
	hashedPassword, _ := HashPassword(password)

	result, err := ComparePassword(hashedPassword, password)

	assert.Nil(t, err, "ComparePassword(%s, %s) = %v; want nil", hashedPassword, password, err)

	assert.True(t, result, "ComparePassword(%s, %s) = %v; want true, err", hashedPassword, password, result)
}

func TestComparePasswordFailed(t *testing.T) {
	password := "password"
	hashedPassword, _ := HashPassword(password)

	result, err := ComparePassword(hashedPassword, "wrong_password")

	assert.NotNil(t, err, "ComparePassword(%s, %s) = %v; want not nil", hashedPassword, "wrong_password", err)

	assert.False(t, result, "ComparePassword(%s, %s) = %v; want false", hashedPassword, "wrong_password", result)
}
