package reader

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
)

type OptCVS func(*CVS)

type JSON struct{}

func NewJSON() JSON {
	return JSON{}
}

type CVS struct {
	Comment rune
	Comma   rune
}

func NewCVS(fns ...OptCVS) *CVS {
	d := DefaultCVS()

	for _, fn := range fns {
		fn(&d)
	}

	return &d
}

func Comma(comma rune) OptCVS {
	return func(c *CVS) {
		c.Comma = comma
	}
}

func Comment(comment rune) OptCVS {
	return func(c *CVS) {
		c.Comment = comment
	}
}

func DefaultCVS() CVS {
	return CVS{
		Comment: '#',
		Comma:   ':',
	}
}

func NewRead(path string, object any) error {
	extension := filepath.Ext(path)

	switch extension[1:] {
	case "json":
		json := NewJSON()
		return json.ReadFile(path, object)
	case "cvs":
		cvs := NewCVS()
		return cvs.ReadFile(path, object)
	default:
		return fmt.Errorf("Can read this file extension")
	}
}

func (JSON) ReadFile(path string, object any) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewDecoder(file).Decode(object)
}

func (c CVS) ReadFile(path string, object any) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	reader.Comment = c.Comment
	reader.Comma = c.Comma

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			if pe, ok := err.(*csv.ParseError); ok {
				fmt.Println("Bad Column: ", pe.Column)
				fmt.Println("Bad Line", pe.Line)
				fmt.Println("Error reported: ", pe.Err)
				if pe.Err == csv.ErrFieldCount {
					continue
				}
			}
		}

		err = parseCSVRecord(record, object)
		if err != nil {
			return err
		}
	}

	return nil
}
func parseCSVRecord(record []string, object interface{}) error {
	objValue := reflect.ValueOf(object)

	if objValue.Kind() != reflect.Ptr || objValue.IsNil() {
		return fmt.Errorf("Objekat mora biti pokazivač na objekat")
	}

	objElem := objValue.Elem()
	if objElem.Kind() != reflect.Slice {
		return fmt.Errorf("Objekat mora biti pokazivač na niz struktura")
	}

	objType := objElem.Type().Elem()
	numFields := objType.NumField()

	if len(record) != numFields {
		return fmt.Errorf("Broj polja u CSV zapisu ne odgovara broju polja u strukturi")
	}

	newElem := reflect.New(objType).Elem()
	for i := 0; i < numFields; i++ {
		field := newElem.Field(i)
		fieldType := objType.Field(i).Type
		fieldValue := record[i]

		switch fieldType.Kind() {
		case reflect.String:
			field.SetString(fieldValue)

		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			numValue, err := strconv.ParseInt(fieldValue, 10, 64)
			if err != nil {
				return err
			}
			field.SetInt(numValue)

		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			numValue, err := strconv.ParseUint(fieldValue, 10, 64)
			if err != nil {
				return err
			}
			field.SetUint(numValue)

		case reflect.Bool:
			boolValue, err := strconv.ParseBool(fieldValue)
			if err != nil {
				return err
			}
			field.SetBool(boolValue)

		case reflect.Float32, reflect.Float64:
			numValue, err := strconv.ParseFloat(fieldValue, 64)
			if err != nil {
				return err
			}
			field.SetFloat(numValue)
		default:
			return fmt.Errorf("Nepodržan tip polja u strukturi")
		}
	}

	// Dodajemo novu strukturu u niz
	objElem.Set(reflect.Append(objElem, newElem))

	return nil
}
