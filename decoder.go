package properties

import (
	"fmt"
	"io"
	"reflect"
	"strconv"
	"strings"
)

type Decoder struct {
	r io.Reader
}

func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{r: r}
}

// Decode decodes properties file to a struct pointer
func (d *Decoder) Decode(v any) error {
	values, err := readToMap(d.r)
	if err != nil {
		err = fmt.Errorf("reading input content: %w", err)
		return err
	}

	valueSource := reflect.ValueOf(v)
	if valueSource.Kind() != reflect.Ptr {
		return ErrNotAPointer
	}
	valueSource = valueSource.Elem()
	if valueSource.Kind() != reflect.Struct {
		return ErrNotAStruct
	}

	valueType := valueSource.Type()

	for i := 0; i < valueType.NumField(); i++ {
		fieldTag, ok := valueType.Field(i).Tag.Lookup(propertiesTag)
		if !ok {
			continue
		}

		fieldName := valueType.Field(i).Name
		fieldValue := valueSource.FieldByName(fieldName)
		if !fieldValue.IsValid() {
			continue
		}

		if !fieldValue.CanSet() {
			continue
		}

		v, ok := values[fieldTag]
		if !ok {
			continue
		}
		switch valueSource.Field(i).Kind() {
		case reflect.String:
			fieldValue.SetString(v)
		case reflect.Int:
			iv, err := strconv.ParseInt(v, 10, 64)
			if err != nil {
				err = fmt.Errorf("failed to parse value for field '%s':%w", fieldTag, err)
				return err
			}
			fieldValue.SetInt(iv)
		}
	}

	return nil
}

func readToMap(r io.Reader) (map[string]string, error) {
	b, err := io.ReadAll(r)
	if err != nil {
		err = fmt.Errorf("reading content: %w", err)
		return nil, err
	}

	values := make(map[string]string)
	for _, l := range strings.Split(string(b), "\n") {
		if strings.HasPrefix(l, "#") {
			continue
		}
		if len(l) == 0 {
			continue
		}

		tmp := strings.Split(l, "=")
		values[tmp[0]] = tmp[1]
	}

	return values, nil
}
