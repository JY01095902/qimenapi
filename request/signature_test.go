package request

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSign(t *testing.T) {
	tests := []struct {
		given []string
		want  string
	}{
		{given: []string{"hello1234", "key123"}, want: "ufYU7rvXhHY3IDyZgyt6SA=="},
	}

	for _, tc := range tests {
		content := tc.given[0]
		secret := tc.given[1]
		signature := sign(content, secret)
		assert.Equal(t, tc.want, signature)
	}
}
