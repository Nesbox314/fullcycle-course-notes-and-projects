package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNewClient(t *testing.T) {
	client, err := NewClient("John Doe", "test@test.com")
	assert.Nil(t, err)
	assert.NotNil(t, client)
	assert.Equal(t, "John Doe", client.Name)
	assert.Equal(t, "test@test.com", client.Email)
}

func TestCreateNewClientWhenArgsAreInvalid(t *testing.T) {
	client, err := NewClient("", "")
	assert.NotNil(t, err)
	assert.Nil(t, client)
}

func TestUpdateClient(t *testing.T) {
	client, _ := NewClient("John Doe", "test@test.com")
	err := client.Update("John Doe Jr.", "test@john.com")
	assert.Nil(t, err)
	assert.Equal(t, "John Doe Jr.", client.Name)
	assert.Equal(t, "test@john.com", client.Email)
}

func TestUpdateClientWithInvalidArgs(t *testing.T) {
	client, _ := NewClient("John Doe", "test@test.com")
	err := client.Update("", "test@john.com")
	assert.Error(t, err, "name is requeired")
}

func TestAddAccountToClient(t *testing.T) {
	client, _ := NewClient("Cliente Teste", "teste@teste")
	account := NewAccount(client)
	err := client.AddAccount(account)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(client.Accounts))
}
