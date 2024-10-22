package common

import (
	"fmt"
	"reflect"
)

func MapFields(dp interface{}, sp interface{}) error {
	dv := reflect.ValueOf(dp)
	sv := reflect.ValueOf(sp)

	if dv.Kind() != reflect.Ptr || sv.Kind() != reflect.Ptr {
		return fmt.Errorf("both dp and sp must be pointers")
	}

	dv = dv.Elem()
	sv = sv.Elem()

	if dv.Kind() != reflect.Struct || sv.Kind() != reflect.Struct {
		return fmt.Errorf("both dp and sp must point to structs")
	}

	dt := dv.Type()
	for i := 0; i < dt.NumField(); i++ {
		sf := dt.Field(i)
		v := sv.FieldByName(sf.Name)
		if !v.IsValid() || !v.Type().AssignableTo(sf.Type) {
			continue
		}
		if dv.Field(i).CanSet() {
			dv.Field(i).Set(v)
		}
	}
	return nil
}
