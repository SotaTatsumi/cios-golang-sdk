package ciossdmock_test

import (
	"testing"

	ciossdk "github.com/optim-corp/cios-golang-sdk/sdk"
	ciossdmock "github.com/optim-corp/cios-golang-sdk/sdk/mock"

	"github.com/stretchr/testify/assert"
)

func TestMockClient_Account_Mock(t *testing.T) {
	assert.Implements(t, (*ciossdk.Account)(nil), ciossdmock.NoImplementAccount{})
}
