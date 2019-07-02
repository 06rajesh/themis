package token

import (
	"crypto/rand"
	"key"
	"random"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewFactory(t *testing.T) {
	var (
		assert   = assert.New(t)
		require  = require.New(t)
		registry = key.NewRegistry(rand.Reader)
		noncer   = random.NewBase64Noncer(rand.Reader, 128, nil)
	)

	factory, err := NewFactory(noncer, registry, Descriptor{
		Alg: "RS256",
		Key: key.Descriptor{
			Kid:  "test",
			Bits: 512,
		},
		Nonce: true,
	})

	require.NoError(err)
	require.NotNil(factory)

	token, err := factory.NewToken(Request{})
	require.NoError(err)
	assert.True(len(token) > 0)
}