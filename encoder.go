package properties

import (
	"errors"
	"fmt"
	"io"
	"reflect"
)

var (
	ErrNotAPointer       = errors.New("parameter must be a pointer")
	ErrNotAStruct        = errors.New("value type must be a struct")
	ErrInvalidFieldValue = errors.New("field value is not valid")
)

// Encoder encodes a struct pointer to a properties file
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
		if err := e.parseField(valueType.Field(i), valueSource); err != nil {
			return err
		}
	}

	return nil
}

func (e *Encoder) parseField(field reflect.StructField, src reflect.Value) error {
	fieldTag, ok := field.Tag.Lookup(propertiesTag)
	if !ok {
		fieldTag = field.Name
	}
	if fieldTag == "-" {
		return nil
	}

	fieldValue := src.FieldByName(field.Name)
	if !fieldValue.IsValid() {
		return fmt.Errorf("%w: field '%s' is not valid", ErrInvalidFieldValue, fieldTag)
	}

	_, err := fmt.Fprintf(e.w, "%s=%v\n", fieldTag, fieldValue)
	if err != nil {
		err = fmt.Errorf("writing property line (%s=%v): %w", fieldTag, fieldValue, err)
		return err
	}

	return nil
}

func NewEncoder(w io.Writer) *Encoder {
	return &Encoder{
		w: w,
	}
}
