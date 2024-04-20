package collections

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_ZeroValue(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		assert.Equal(t, 0, ZeroValue[int]())
		assert.Equal(t, "", ZeroValue[string]())
		assert.Equal(t, float32(0), ZeroValue[float32]())
	})
}
