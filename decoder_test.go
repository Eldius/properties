package properties

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

type testStruct0 struct {
	Key0 string `properties:"key0"`
	Key1 string `properties:"key1"`
}

type complexTestStruct struct {
	BoolValue0  bool   `properties:"bool-value-0"`
	BoolValue1  bool   `properties:"bool-value-1"`
	BoolValue2  bool   `properties:"bool-value-2"`
	StringValue string `properties:"string.value"`
	IntValue    int    `properties:"int_value"`
	IntValue32  int32  `properties:"int32_value"`
	IntValue64  int64  `properties:"int64_value"`
}

func TestReadToMap(t *testing.T) {
	t.Run("tests reading from a string", func(t *testing.T) {
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
	})

	t.Run("tests reading from a file", func(t *testing.T) {
		t.Run("given a simple and valid properties file should return right values", func(t *testing.T) {
			f, err := os.Open("test_samples/props0.properties")
			assert.Nil(t, err)

			var out testStruct0
			if err := NewDecoder(f).Decode(&out); !assert.Nil(t, err) {
				assert.FailNow(t, err.Error())
			}

			assert.Equal(t, "value0", out.Key0)
			assert.Equal(t, "value1", out.Key1)
		})

		t.Run("given a valid properties file with a comment line should ignore it return right values", func(t *testing.T) {
			f, err := os.Open("test_samples/props1.properties")
			assert.Nil(t, err)

			var out testStruct0
			if err := NewDecoder(f).Decode(&out); !assert.Nil(t, err) {
				assert.FailNow(t, err.Error())
			}

			assert.Equal(t, "value0", out.Key0)
			assert.Equal(t, "value1", out.Key1)
		})

		t.Run("given a valid properties file with two comment lines should ignore it return right values", func(t *testing.T) {
			f, err := os.Open("test_samples/props2.properties")
			assert.Nil(t, err)

			var out testStruct0
			if err := NewDecoder(f).Decode(&out); !assert.Nil(t, err) {
				assert.FailNow(t, err.Error())
			}

			assert.Equal(t, "value0", out.Key0)
			assert.Equal(t, "value1", out.Key1)
		})

		t.Run("given a valid properties file with an empty line should ignore it return right values", func(t *testing.T) {
			f, err := os.Open("test_samples/props3.properties")
			assert.Nil(t, err)

			var out testStruct0
			if err := NewDecoder(f).Decode(&out); !assert.Nil(t, err) {
				assert.FailNow(t, err.Error())
			}

			assert.Equal(t, "value0", out.Key0)
			assert.Equal(t, "value1", out.Key1)
		})

		t.Run("given a complex struct with multiple attribute types should return the right values", func(t *testing.T) {
			f, err := os.Open("test_samples/props4.properties")
			assert.Nil(t, err)

			var out complexTestStruct
			if err := NewDecoder(f).Decode(&out); !assert.Nil(t, err) {
				assert.FailNow(t, err.Error())
			}

			assert.True(t, out.BoolValue0)
			assert.False(t, out.BoolValue1)
			assert.False(t, out.BoolValue2)
			assert.Equal(t, "Can you read me?", out.StringValue)
			assert.Equal(t, int(42), out.IntValue)
			assert.Equal(t, int32(32), out.IntValue32)
			assert.Equal(t, int64(64), out.IntValue64)
		})
	})
}
