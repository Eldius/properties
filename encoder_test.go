package properties

import (
	"bytes"
	"errors"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

type TestStruct struct {
	Value  string `properties:"value"`
	Number int    `properties:"number"`
}

func TestNewEncoder(t *testing.T) {

	t.Run("given a non pointer value encode must return an error", func(t *testing.T) {
		var buf bytes.Buffer

		w := NewEncoder(&buf)

		err := w.Encode(TestStruct{})
		assert.NotNil(t, err)
		assert.Empty(t, buf)
		assert.True(t, errors.Is(err, ErrNotAPointer))
	})

	t.Run("given a non struct type value encode must return an error", func(t *testing.T) {
		var buf bytes.Buffer

		w := NewEncoder(&buf)

		var temp []byte = nil
		err := w.Encode(temp)
		assert.NotNil(t, err)
		assert.Empty(t, buf)
		assert.True(t, errors.Is(err, ErrNotAPointer))
	})

	t.Run("given a pointer value encode must not return an error", func(t *testing.T) {
		var buf bytes.Buffer

		w := NewEncoder(&buf)

		err := w.Encode(&TestStruct{
			Value:  "test-value",
			Number: 123,
		})
		assert.Nil(t, err)
		assert.NotEmpty(t, buf)

		val := buf.String()

		assert.Contains(t, val, `value=test-value`)
		assert.Contains(t, val, `number=123`)
		assert.Len(t, strings.Split(val, "\n"), 2)
	})
}
