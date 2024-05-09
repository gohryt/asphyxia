package env

import (
	"errors"
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

type (
	Path struct {
		Value string
	}

	Prefix struct {
		Value string
	}
)

var ErrNoReqiredVariable = errors.New("no required variable")

// TODO: error while parse slice
//
// tags from .env goes to lowercase. ENV_NAME == env_name
func Parse[T_configuration any](optionList ...any) (*T_configuration, error) {
	pathList, prefix := []string(nil), ""

	for _, option := range optionList {
		switch typed := option.(type) {
		case Prefix:
			prefix = typed.Value
		case Path:
			pathList = append(pathList, typed.Value)
		}
	}

	err := godotenv.Load(pathList...)
	if err != nil {
		return nil, err
	}

	configuration := new(T_configuration)

	err = parse(reflect.ValueOf(configuration), prefix, "", "")
	if err != nil {
		return configuration, err
	}

	return configuration, err
}

func parse(value reflect.Value, name, _default, required string) error {
	kind := value.Kind()

	switch kind {
	case reflect.Pointer:
		return parse(value.Elem(), name, _default, required)
	case reflect.Struct:
		_type := value.Type()

		if name != "" && !strings.HasSuffix(name, "_") {
			name += "_"
		}

		for i := 0; i < _type.NumField(); i++ {
			field := _type.Field(i)

			err := parse(value.Field(i), (name + field.Tag.Get("env")), field.Tag.Get("default"), field.Tag.Get("required"))
			if err != nil {
				return err
			}
		}

		return nil
	}

	if name == "" {
		return nil
	}

	read := os.Getenv(name)
	if read == "" {
		read = _default
	}

	if read == "" {
		if required == "true" {
			return fmt.Errorf("%w: %s", ErrNoReqiredVariable, name)
		}

		return nil
	}

	typeDuration := reflect.TypeOf(time.Duration(0))
	typeSliceString := reflect.TypeOf(reflect.TypeOf([]string{}))

	switch kind {
	case reflect.String:
		value.SetString(read)
	case reflect.Int64:
		if value.Type() == typeDuration {
			d, err := time.ParseDuration(read)
			if err != nil {
				return err
			}

			value.SetInt(int64(d))
		} else {
			i, err := strconv.ParseInt(read, 0, 64)
			if err != nil {
				return err
			}

			value.SetInt(i)
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32:
		i, err := strconv.ParseInt(read, 0, 64)
		if err != nil {
			return err
		}

		value.SetInt(i)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		i, err := strconv.ParseUint(read, 0, 64)
		if err != nil {
			return err
		}

		value.SetUint(i)
	case reflect.Float32, reflect.Float64:
		i, err := strconv.ParseFloat(read, 64)
		if err != nil {
			return err
		}

		value.SetFloat(i)
	case reflect.Complex64, reflect.Complex128:
		i, err := strconv.ParseComplex(read, 64)
		if err != nil {
			return err
		}

		value.SetComplex(i)
	case reflect.Bool:
		i, err := strconv.ParseBool(read)
		if err != nil {
			return err
		}

		value.SetBool(i)
	case reflect.Slice:
		log.Println(value.Type(), typeSliceString)
		if value.Type() == typeSliceString {
			data := reflect.MakeSlice(typeSliceString, 0, 0)

			for _, i := range strings.Split(read, "") {
				data = reflect.Append(data, reflect.ValueOf(i))
			}

			value.Set(data)
		}
	}

	return nil
}
