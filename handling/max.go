package handling

import (
	"reflect"
	"strconv"

	"github.com/theTardigrade/golang-validation/data"
)

func init() {
	addHandler("max", maxDatum{})
}

type maxDatum struct{}

func (d maxDatum) Test(m *data.Main, t *data.Tag) (success bool, err error) {
	switch m.FieldKind {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		success, err = d.testInts(m, t)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		success, err = d.testUints(m, t)
	case reflect.Float32, reflect.Float64:
		success, err = d.testFloats(m, t)
	default:
		err = ErrUnexpectedType
	}

	return
}

func (d maxDatum) testInts(m *data.Main, t *data.Tag) (success bool, err error) {
	tagValueInt, err := strconv.ParseInt(t.Value, 10, 64)
	if err == nil && m.FieldValue.Int() <= tagValueInt {
		success = true
	}

	return
}

func (d maxDatum) testUints(m *data.Main, t *data.Tag) (success bool, err error) {
	tagValueUint, err := strconv.ParseUint(t.Value, 10, 64)
	if err == nil && m.FieldValue.Uint() <= tagValueUint {
		success = true
	}

	return
}

func (d maxDatum) testFloats(m *data.Main, t *data.Tag) (success bool, err error) {
	tagValueFloat, err := strconv.ParseFloat(t.Value, 64)
	if err == nil && m.FieldValue.Float() <= tagValueFloat {
		success = true
	}

	return
}

func (d maxDatum) FailureMessage(m *data.Main, t *data.Tag) string {
	return m.FormattedFieldName + " cannot be greater than " + t.Value + "."
}
