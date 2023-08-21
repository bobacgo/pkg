package file

import (
	"bufio"
	"io"
	"reflect"
	"strconv"
	"strings"
)

func Parse[T any](src io.Reader, tag string) *T {
	if tag == "" {
		tag = "json"
	}

	rawMap := make(map[string]string)

	scanner := bufio.NewScanner(src)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 && line[0] == '#' { // 跳过注释
			continue
		}
		pivot := strings.IndexAny(line, " ")
		if pivot > 0 && pivot < len(line)-1 {
			key := line[:pivot]
			value := strings.Trim(line[pivot+1:], " ")
			rawMap[strings.ToLower(key)] = value
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	config := new(T)
	t := reflect.TypeOf(config)
	v := reflect.ValueOf(config)
	n := t.Elem().NumField()
	for i := 0; i < n; i++ {
		field := t.Elem().Field(i)
		fValue := v.Elem().Field(i)
		key, ok := field.Tag.Lookup(tag)
		if !ok {
			key = field.Name
		}
		value, ok := rawMap[strings.ToLower(key)]
		if ok {
			switch field.Type.Kind() {
			case reflect.String:
				fValue.SetString(value)
			case reflect.Int:
				if intValue, err := strconv.ParseInt(value, 10, 64); err == nil {
					fValue.SetInt(intValue)
				}
			case reflect.Bool:
				boolValue := "true" == value
				fValue.SetBool(boolValue)
			case reflect.Slice:
				if field.Type.Elem().Kind() == reflect.String {
					split := strings.Split(value, ",")
					fValue.Set(reflect.ValueOf(split))
				}
			}

		}
	}
	return config
}