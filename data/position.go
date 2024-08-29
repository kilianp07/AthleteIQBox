package data

import (
	"fmt"
	"reflect"
	"strings"
)

type Position struct {
	Timestamp  int64   `json:"timestamp_unix"`
	Latitude   float64 `json:"latitude"`
	Longitude  float64 `json:"longitude"`
	Altitude_M float64 `json:"altitude_m"`
	Speed_kMh  float64 `json:"speed_kmh"`
	Course     float64 `json:"course_degrees"`
}

// Copy creates a deep copy of the Position object.
// It returns a new Position object with the same field values as the original.
func (p *Position) Copy() Position {
	originalValue := reflect.ValueOf(p).Elem()
	copyValue := reflect.New(originalValue.Type()).Elem()

	for i := range originalValue.NumField() {
		copyValue.Field(i).Set(originalValue.Field(i))
	}

	return copyValue.Interface().(Position)
}

// SQLCreateQuery generates the SQL CREATE TABLE query for the Position struct.
// It inspects the fields of the Position struct using reflection and generates
// the corresponding column definitions based on the field types.
// The generated query includes the column names, types, and any additional
// constraints such as primary key.
// The returned query can be used to create the "positions" table in a database.
func (p *Position) SQLCreateQuery() string {
	t := reflect.TypeOf(*p)
	var columns []string

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		columnName := field.Tag.Get("json")
		if columnName == "" {
			columnName = strings.ToLower(field.Name)
		}

		var columnType string
		switch field.Type.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			columnType = "INTEGER"
		case reflect.Float32, reflect.Float64:
			columnType = "REAL"
		default:
			columnType = "TEXT" // Default type if needed for other types
		}

		columnDef := fmt.Sprintf("%s %s", columnName, columnType)
		if columnName == "timestamp_unix" {
			columnDef += " PRIMARY KEY"
		}

		columns = append(columns, columnDef)
	}

	return fmt.Sprintf("CREATE TABLE IF NOT EXISTS positions (%s)", strings.Join(columns, ", "))
}

// SQLInsertQuery generates an SQL INSERT query for the Position struct.
//
// The generated query includes the column names and corresponding values
// based on the struct tags and field values of the Position instance.
//
// Returns the generated SQL INSERT query as a string.
func (p *Position) SQLInsertQuery() string {
	t := reflect.TypeOf(*p)
	v := reflect.ValueOf(*p)
	var columns []string
	var values []string

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		columnName := field.Tag.Get("json")
		if columnName == "" {
			columnName = strings.ToLower(field.Name)
		}

		columns = append(columns, columnName)

		fieldValue := v.Field(i)
		switch fieldValue.Kind() {
		case reflect.String:
			values = append(values, fmt.Sprintf("'%v'", fieldValue.Interface()))
		default:
			values = append(values, fmt.Sprintf("%v", fieldValue.Interface()))
		}
	}

	return fmt.Sprintf("INSERT INTO positions (%s) VALUES (%s)", strings.Join(columns, ", "), strings.Join(values, ", "))
}
