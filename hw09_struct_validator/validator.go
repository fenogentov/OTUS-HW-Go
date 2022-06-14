package hw09structvalidator

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

var ErrNoStruct = errors.New("error input parameter is not a structure")

type ValidationError struct {
	Field string
	Err   error
}

type ValidationErrors []ValidationError

func (v ValidationErrors) Error() string {
	mess := ""
	for i, err := range v {
		mess += fmt.Sprintf("%d. field %v : %v\n", i, err.Field, err.Err)
	}
	return mess
}

func intValidate(v int, tag string) error {
	for _, c := range strings.Split(tag, "|") {
		switch {
		case strings.HasPrefix(c, "max"):
			max := 0
			fmt.Sscanf(c, "max:%d", &max)
			if v > max {
				return fmt.Errorf("error maximum value (%v)", tag)
			}
		case strings.HasPrefix(c, "min"):
			min := 0
			fmt.Sscanf(c, "min:%d", &min)
			if v < min {
				return fmt.Errorf("error minimum value (%v)", tag)
			}
		case strings.HasPrefix(c, "in"):
			tag = strings.TrimLeft(tag, "in:")
			flagOK := false
			for _, c := range strings.Split(tag, ",") {
				i, _ := strconv.Atoi(c)
				if i == v {
					flagOK = true
				}
			}
			if !flagOK {
				return fmt.Errorf("error not in the list")
			}
		}
	}
	return nil
}

func stringValidate(str string, tag string) error {
	for _, c := range strings.Split(tag, "|") {
		switch {
		case strings.HasPrefix(c, "len"):
			length := 0
			fmt.Sscanf(c, "len:%d", &length)
			if len(str) != length {
				return fmt.Errorf("error string length (%v)", tag)
			}
		case strings.HasPrefix(c, "max"):
			maxLen := 0
			fmt.Sscanf(c, "max:%d", &maxLen)
			if len(str) > maxLen {
				return fmt.Errorf("error maximum string length (%v)", tag)
			}
		case strings.HasPrefix(c, "min"):
			minLen := 0
			fmt.Sscanf(c, "min:%d", &minLen)
			if len(str) < minLen {
				return fmt.Errorf("error maximum string length (%v)", tag)
			}
		case strings.HasPrefix(c, "regexp"):
			tag = strings.TrimPrefix(tag, "regexp:")
			if rez, _ := regexp.MatchString(tag, str); !rez {
				return fmt.Errorf("error regular expression (%v)", tag)
			}
		case strings.HasPrefix(c, "in"):
			tag = strings.TrimPrefix(tag, "in:")
			arrIn := strings.Split(tag, ",")
			flagOk := false
			for _, a := range arrIn {
				if a == str {
					flagOk = true
				}
			}
			if !flagOk {
				return fmt.Errorf("error not in the list (%v)", tag)
			}
		}
	}
	return nil
}

func Validate(v interface{}) error {
	var validationErrors ValidationErrors

	t := reflect.ValueOf(v)

	if t.Type().Kind() != reflect.Struct {
		return ErrNoStruct
	}

	for i := 0; i < t.NumField(); i++ {
		var errField error

		tag, ok := t.Type().Field(i).Tag.Lookup("validate")
		if !ok {
			continue
		}

		field := t.Field(i)
		fieldName := t.Type().Field(i).Name

		if !field.CanInterface() {
			continue
		}

		switch field.Kind() {
		case reflect.Int:
			intgr := field.Interface().(int)
			errField = intValidate(intgr, tag)

			if errField != nil {
				validationErrors = append(validationErrors, ValidationError{
					Field: fieldName,
					Err:   errField,
				})
			}
		case reflect.String:
			str := field.String()
			errField = stringValidate(str, tag)
			if errField != nil {
				validationErrors = append(validationErrors, ValidationError{
					Field: fieldName,
					Err:   errField,
				})
			}
		case reflect.Slice:
			switch field.Interface().(type) {
			case []string:
				vl := field.Interface().([]string)
				for _, str := range vl {
					errField = stringValidate(str, tag)
					if errField != nil {
						validationErrors = append(validationErrors, ValidationError{
							Field: fieldName,
							Err:   errField,
						})
					}
				}
			case []int:
				vl := field.Interface().([]int)
				for _, intgr := range vl {
					errField = intValidate(intgr, tag)
					if errField != nil {
						validationErrors = append(validationErrors, ValidationError{
							Field: fieldName,
							Err:   errField,
						})
					}
				}
			}
		}
	}

	if validationErrors != nil {
		return validationErrors
	}
	return nil
}
