package graphql

import (
	"audienceviral-api/connections"
	"fmt"
	"reflect"
	"strings"
	"unicode"

	"github.com/dosco/graphjin/core"
)

var Graph *core.GraphJin

func ToUpperCamelCase(s string) string {
	rs := []rune(s)
	new := []rune{}
	next_upper := true

	for _, r := range rs {
		if next_upper {
			new = append(new, unicode.ToUpper(r))
			next_upper = false
			continue
		}
		if unicode.IsPunct(r) {
			next_upper = true
			continue
		}

		new = append(new, r)
	}

	return string(new)

}

func ToSnakeCase(s string) string {
	var res = make([]rune, 0, len(s))
	var p = '_'
	for i, r := range s {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
			res = append(res, '_')
		} else if unicode.IsUpper(r) && i > 0 {
			if unicode.IsLetter(p) && !unicode.IsUpper(p) || unicode.IsDigit(p) {
				res = append(res, '_', unicode.ToLower(r))
			} else {
				res = append(res, unicode.ToLower(r))
			}
		} else {
			res = append(res, unicode.ToLower(r))
		}

		p = r
	}
	return string(res)
}

func ReflectToFragment(data interface{}) string {
	var fragment string
	var fieldsList string
	var structName string
	var structSnake string

	v := reflect.ValueOf(data)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	structName = v.Type().Name()
	if strings.Contains(structName, "Data") {
		structName = strings.Replace(structName, "Data", "", -1)
	}
	structSnake = ToSnakeCase(structName)

	for i := 0; i < v.NumField(); i++ {
		var fieldName string
		var fieldSnake string

		fieldName = v.Type().Field(i).Name
		fieldSnake = ToSnakeCase(fieldName)

		if i == 0 {
			fieldsList = fieldSnake
			continue
		}
		fieldsList = fieldsList + `
		` + fieldSnake

	}

	fragment = fmt.Sprintf(`
	
	fragment %s on %s {
		%s
	}
	
	`, structName, structSnake, fieldsList)

	fmt.Println(fragment)

	return fragment
}

func Init() error {
	var err error
	Graph, err = core.NewGraphJin(nil, connections.Postgres)

	if err != nil {
		return err
	}

	return nil
}
