package env

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
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

var typeForDuration = reflect.TypeFor[time.Duration]()

func Load(pathList []string) error {
	if pathList == nil {
		pathList = []string{".env"}
	}

	var (
		path    string
		errList []error
		err     error
	)

	for _, path = range pathList {
		err = LoadFile(path)
		if err != nil {
			errList = append(errList, err)
		}
	}

	return errors.Join(errList...)
}

func LoadFile(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var (
		text           string
		equalSignIndex int
		key            string
		value          string
	)

	for scanner.Scan() {
		text = scanner.Text()

		if text == "" || strings.HasPrefix(text, "#") {
			continue
		}

		equalSignIndex = strings.IndexByte(text, '=')

		if equalSignIndex < 0 {
			continue
		}

		key, value = text[:equalSignIndex], text[equalSignIndex+1:]

		if len(value) > 2 && strings.HasPrefix(value, `"`) && strings.HasSuffix(value, `"`) {
			value = value[1 : len(value)-1]
		}

		if key == "" || value == "" {
			continue
		}

		err = os.Setenv(key, value)
		if err != nil {
			return err
		}
	}

	return nil
}

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

	err := Load(pathList)
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

func parse(value reflect.Value, n, d, r string) error {
	k := value.Kind()
	t := value.Type()

	switch k {
	case reflect.Pointer:
		return parse(value.Elem(), n, d, r)
	case reflect.Struct:
		if n != "" && !strings.HasSuffix(n, "_") {
			n += "_"
		}

		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)

			err := parse(value.Field(i), (n + field.Tag.Get("name")), field.Tag.Get("default"), field.Tag.Get("required"))
			if err != nil {
				return err
			}
		}

		return nil
	}

	if n == "" {
		return nil
	}

	read := os.Getenv(n)
	if read == "" {
		read = d
	}

	if read == "" {
		if r == "true" {
			return fmt.Errorf("%w: %s", ErrNoReqiredVariable, n)
		}

		return nil
	}

	switch k {
	case reflect.Slice:
		e := t.Elem()

		switch e.Kind() {
		case reflect.String:
			data := reflect.MakeSlice(t, 0, strings.Count(read, ","))

			for _, i := range strings.Split(read, ",") {
				data = reflect.Append(data, reflect.ValueOf(i))
			}

			value.Set(data)
		}
	case reflect.String:
		value.SetString(read)
	case reflect.Int64:
		if value.Type() == typeForDuration {
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
	}

	return nil
}
