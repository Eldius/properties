package properties

import (
	"fmt"
	"github.com/pkg/errors"
	"io"
	"reflect"
)

// http://www.inanzzz.com/index.php/post/5p8q/creating-custom-struct-tags-with-golang

var (
	ErrNotAPointer = errors.New("value must be a pointer")
	ErrNotAStruct  = errors.New("value must be a struct")
)

// Encoder encodes an struct pointer to a properties file
type Encoder struct {
	w io.Writer
}

// Encode encodes a struct pointer to a properties file
func (e *Encoder) Encode(v any) error {
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

		_, err := fmt.Fprintf(e.w, "%s=%v", fieldTag, fieldValue)
		if err != nil {
			err = fmt.Errorf("writing property line (%s=%v): %w", fieldTag, fieldValue, err)
			return err
		}

		if i < (valueType.NumField() - 1) {
			_, err := fmt.Fprintln(e.w)
			if err != nil {
				err = fmt.Errorf("writing line break: %w", err)
				return err
			}
		}
	}

	return nil
}

func NewEncoder(w io.Writer) *Encoder {
	return &Encoder{
		w: w,
	}
}
