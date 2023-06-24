package properties

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReadToMap(t *testing.T) {
	t.Run("given a simple and valid properties file should return right values", func(t *testing.T) {
		r := bytes.NewReader([]byte(`key0=value0
key1=value1`))

		values, err := readToMap(r)
		assert.Nil(t, err)
		assert.Len(t, values, 2)
		if v, ok := values["key0"]; assert.True(t, ok) {
			assert.Equal(t, "value0", v)
		}
		if v, ok := values["key1"]; assert.True(t, ok) {
			assert.Equal(t, "value1", v)
		}
	})

	t.Run("given a valid properties file with a comment line should ignore it return right values", func(t *testing.T) {
		r := bytes.NewReader([]byte(`# it's a line comment
key0=value0
key1=value1`))

		values, err := readToMap(r)
		assert.Nil(t, err)
		assert.Len(t, values, 2)
		if v, ok := values["key0"]; assert.True(t, ok) {
			assert.Equal(t, "value0", v)
		}
		if v, ok := values["key1"]; assert.True(t, ok) {
			assert.Equal(t, "value1", v)
		}
		if v, ok := values["key1"]; assert.True(t, ok) {
			assert.Equal(t, "value1", v)
		}
	})

	t.Run("given a valid properties file with two comment lines should ignore it return right values", func(t *testing.T) {
		r := bytes.NewReader([]byte(`# it's a line comment
key0=value0
# another comment
key1=value1`))

		values, err := readToMap(r)
		assert.Nil(t, err)
		assert.Len(t, values, 2)
		if v, ok := values["key0"]; assert.True(t, ok) {
			assert.Equal(t, "value0", v)
		}
		if v, ok := values["key1"]; assert.True(t, ok) {
			assert.Equal(t, "value1", v)
		}
	})

	t.Run("given a valid properties file with an empty line should ignore it return right values", func(t *testing.T) {
		r := bytes.NewReader([]byte(`
key0=value0

key1=value1`))

		values, err := readToMap(r)
		assert.Nil(t, err)
		assert.Len(t, values, 2)
		if v, ok := values["key0"]; assert.True(t, ok) {
			assert.Equal(t, "value0", v)
		}
		if v, ok := values["key1"]; assert.True(t, ok) {
			assert.Equal(t, "value1", v)
		}
	})
}
