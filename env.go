// Copyright © 2019 Alexandre Kovac <contact@kovacou.fr>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package env

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"strings"
	"time"
)

// Parse the env variable and trim spaces.
func Parse(key string) string {
	return strings.TrimSpace(os.Getenv(key))
}

// Prefix generate a new variable name based on given prefix and key.
func Prefix(prefix, key string) string {
	return Format("%s_%s", prefix, key)
}

// Format will parse the key based on given format.
func Format(format string, v ...interface{}) string {
	return fmt.Sprintf(format, v...)
}

// Lookup the env variable and trim spaces and if the key is defined.
func Lookup(key string) (string, bool) {
	v, find := os.LookupEnv(key)
	return strings.TrimSpace(v), find
}

// Unmarshal will fill the given struct with the environment variables.
func Unmarshal(v interface{}) error {
	e := reflect.ValueOf(v)
	if e.Kind() == reflect.Ptr && !e.IsNil() {
		e = e.Elem()
		if e.Kind() != reflect.Struct {
			return errors.New("must be a non-nil struct pointer")
		}
	} else {
		return errors.New("must be a non-nil struct pointer")
	}

	t := e.Type()
	for i := 0; i < e.NumField(); i++ {
		vf := e.Field(i)
		if vf.Kind() == reflect.Struct {
			if vf.Addr().CanInterface() {
				v := vf.Addr().Interface()
				if err := Unmarshal(v); err != nil {
					return err
				}
			}
		}

		if vf.CanSet() || vf.Kind() == reflect.Map {
			tf := t.Field(i)
			if tag := tf.Tag.Get("env"); tag != "" {
				if v, ok := os.LookupEnv(tag); ok {
					if err := setValue(tf.Type, vf, v); err != nil {
						return err
					}
				}
			}
		}
	}

	return nil
}

func setValue(t reflect.Type, vf reflect.Value, v string) (err error) {
	switch t.Kind() {
	case reflect.Ptr:
		ptr := reflect.New(t.Elem())
		if err = setValue(t.Elem(), ptr.Elem(), v); err == nil {
			vf.Set(ptr)
		}
	case reflect.Map:
		vm := map[string]string{}

		for _, line := range strings.Split(v, ",") {
			str := strings.Split(line, ":")
			if len(str) == 2 {
				vm[str[0]] = str[1]
			}
		}

		switch t.String() {
		case "map[string]string":
			vf.Set(reflect.ValueOf(vm))
		case "map[string]bool":
			svm := make(map[string]bool, len(vm))
			for k, v := range vm {
				svm[k] = toBool(v)
			}
			vf.Set(reflect.ValueOf(svm))
		case "map[string]int":
			svm := make(map[string]int, len(vm))
			for k, v := range vm {
				svm[k] = toInt(v)
			}
			vf.Set(reflect.ValueOf(svm))
		case "map[string]uint":
			svm := make(map[string]uint, len(vm))
			for k, v := range vm {
				svm[k] = toUint(v)
			}
			vf.Set(reflect.ValueOf(svm))
		case "map[string]int64":
			svm := make(map[string]int64, len(vm))
			for k, v := range vm {
				svm[k] = toInt64(v)
			}
			vf.Set(reflect.ValueOf(svm))
		case "map[string]uint64":
			svm := make(map[string]uint64, len(vm))
			for k, v := range vm {
				svm[k] = toUint64(v)
			}
			vf.Set(reflect.ValueOf(svm))
		case "map[string]float64":
			svm := make(map[string]float64, len(vm))
			for k, v := range vm {
				svm[k] = toFloat64(v)
			}
			vf.Set(reflect.ValueOf(svm))
		case "map[string]interface {}":
			svm := make(map[string]interface{}, len(vm))
			for k, v := range vm {
				svm[k] = v
			}
			vf.Set(reflect.ValueOf(svm))
		default:
			err = errors.New("field type is not supported")
		}
	case reflect.Slice:
		vs := strings.Split(v, ",")
		n := len(vs)

		switch t.String() {
		case "[]string":
			vf.Set(reflect.ValueOf(vs))
		case "[]bool":
			svs := make([]bool, n, n)
			for k, v := range vs {
				svs[k] = toBool(v)
			}
			vf.Set(reflect.ValueOf(svs))
		case "[]int":
			svs := make([]int, n, n)
			for k, v := range vs {
				svs[k] = toInt(v)
			}
			vf.Set(reflect.ValueOf(svs))
		case "[]uint":
			svs := make([]uint, n, n)
			for k, v := range vs {
				svs[k] = toUint(v)
			}
			vf.Set(reflect.ValueOf(svs))
		case "[]int64":
			svs := make([]int64, n, n)
			for k, v := range vs {
				svs[k] = toInt64(v)
			}
			vf.Set(reflect.ValueOf(svs))
		case "[]uint64":
			svs := make([]uint64, n, n)
			for k, v := range vs {
				svs[k] = toUint64(v)
			}
			vf.Set(reflect.ValueOf(svs))
		case "[]float64":
			svs := make([]float64, n, n)
			for k, v := range vs {
				svs[k] = toFloat64(v)
			}
			vf.Set(reflect.ValueOf(svs))
		case "[]interface {}":
			svs := make([]interface{}, n, n)
			for k, v := range vs {
				svs[k] = v
			}
			vf.Set(reflect.ValueOf(svs))
		default:
			err = errors.New("field type is not supported")
		}
	case reflect.String:
		vf.SetString(v)
	case reflect.Bool:
		vf.SetBool(toBool(v))
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		switch t.String() {
		case "time.Duration":
			var d time.Duration
			if d, err = time.ParseDuration(v); err == nil {
				vf.Set(reflect.ValueOf(d))
			}
		default:
			vf.SetInt(toInt64(v))
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		vf.SetUint(toUint64(v))
	case reflect.Float64:
		vf.SetFloat(toFloat64(v))
	case reflect.Float32:
		vf.SetFloat(float64(toFloat32(v)))
	default:

	}

	return
}
