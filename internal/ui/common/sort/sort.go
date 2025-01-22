package sort

import (
	"cmp"
	"fmt"
	"reflect"
	"slices"
	"strings"

	"github.com/vadimklimov/cpi-navigator/internal/config"
)

type Options struct {
	Field string
	Order config.SortOrder
}

type Direction int

const (
	DirectionAscending Direction = iota
	DirectionDescending
)

func Sort[T any](items []T, options Options) {
	var direction Direction

	switch options.Order {
	case config.SortOrderAscending:
		direction = DirectionAscending
	case config.SortOrderDescending:
		direction = DirectionDescending
	}

	sort(items, options.Field, direction)
}

func sort[T any](items []T, name string, direction Direction) {
	if len(items) == 0 || name == "" {
		return
	}

	field, err := fieldByName[T](name)
	if err != nil {
		return
	}

	slices.SortFunc(items, comparator[T](field, direction))
}

func fieldByName[T any](name string) (reflect.StructField, error) {
	t := reflect.TypeOf((*T)(nil)).Elem()

	if t.Kind() != reflect.Struct {
		return reflect.StructField{}, fmt.Errorf("argument must be struct, got %s", t.Kind())
	}

	field, ok := t.FieldByNameFunc(func(field string) bool {
		return strings.EqualFold(name, field)
	})

	if !ok {
		return reflect.StructField{}, fmt.Errorf("field %s not found in struct type %s", name, t.Name())
	}

	return field, nil
}

func comparator[T any](field reflect.StructField, direction Direction) func(a, b T) int {
	return func(a, b T) int {
		var result int

		valA := reflect.ValueOf(a).FieldByName(field.Name)
		valB := reflect.ValueOf(b).FieldByName(field.Name)

		switch field.Type.Kind() {
		case reflect.String:
			result = cmp.Compare(strings.ToLower(valA.String()), strings.ToLower(valB.String()))
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
			reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			result = cmp.Compare(valA.Int(), valB.Int())
		case reflect.Float32, reflect.Float64:
			result = cmp.Compare(valA.Float(), valB.Float())
		default:
			result = cmp.Compare(
				strings.ToLower(fmt.Sprintf("%v", valA.Interface())),
				strings.ToLower(fmt.Sprintf("%v", valB.Interface())),
			)
		}

		if direction == DirectionDescending {
			result = -result
		}

		return result
	}
}
