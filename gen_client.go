package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"strings"
)

func GenerateTsClient(structs []Struct) error {
	for _, s := range structs {
		file, err := GenerateTsType(s)
		if err != nil {
			return err
		}

		fmt.Println(file)

		// panic("")

		err = ioutil.WriteFile("./client/models/"+ToSnakeCase(s.Name)+".ts", []byte(file), 0777)
		if err != nil {
			return err
		}
	}

	cmd := exec.Command("npm", "install", "--save-dev", "--save-exact", "prettier")
	cmd.Dir = "./client"

	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(out))

	err = ioutil.WriteFile("./client/.prettierrc.json", []byte("{}"), 0777)
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile("./client/.prettierignore", []byte(`# Ignore artifacts
build
coverage`), 0777)
	if err != nil {
		log.Fatal(err)
	}

	cmd = exec.Command("npx", "prettier", "--write", ".")
	cmd.Dir = "./client"

	out, err = cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(out))

	return nil
}

func GenerateTsType(s Struct) (string, error) {
	var t_fields string
	var i_fields string
	var t_constructor_fields string
	var imports string

	// s_snake := ToSnakeCase(s.Name)
	s_upper := ToUpperCase(s.Name)

	for _, c := range s.Columns {

		c_snake := ToSnakeCase(c.Name)

		if c.IsForeignObject {
			fk_type_snake := ToSnakeCase(strings.ReplaceAll(c.Type, "[]", ""))
			fk_type_upper := ToUpperCase(strings.ReplaceAll(c.Type, "[]", ""))
			fk_snake := ToSnakeCase(strings.ReplaceAll(c.Name, "[]", ""))
			// fk_upper := ToUpperCase(strings.ReplaceAll(c.Name, "[]", ""))

			imports += `import {` + "I" + fk_type_upper + ", " + fk_type_upper + "} from './" + fk_type_snake + `';
`

			if strings.Contains(c.Type, "[]") {

				t_fields += fk_snake + "?: " + fk_type_upper + `[];
`
				i_fields += fk_snake + "?: I" + fk_type_upper + `[];
`

				t_constructor_fields += "this." + fk_snake + " = " + "data." + fk_snake + "?.map(i => { return new " + fk_type_upper + `(i) });
`
			} else {
				t_fields += fk_snake + "?: " + fk_type_upper + `;
`
				i_fields += fk_snake + "?: I" + fk_type_upper + `;
`

				t_constructor_fields += "this." + fk_snake + " = data." + fk_snake + " ? " + "new " + fk_type_upper + "(data." + fk_snake + `) : undefined;
`
			}

			continue
		}

		ts_type := ""

		switch c.Type {
		case "string":
			ts_type = "string"
		case "int64":
			ts_type = "number"
		case "Ints":
			ts_type = "number[]"
		case "Strings":
			ts_type = "string[]"
		case "float64":
			ts_type = "number"
		case "bool":
			ts_type = "boolean"
		case "time.Time":
			ts_type = "Date"
		case "map[string]interface{}":
			ts_type = "object"
		case "[]map[string]interface{}":
			ts_type = "object[]"
		default:
			log.Fatal(c)
		}

		if ts_type == "Date" {
			t_fields += c_snake + "?: " + ts_type + `;
`
			i_fields += c_snake + "?: number | string | " + ts_type + `;
`

			t_constructor_fields += "this." + c_snake + " = " + "data." + c_snake + " ? new Date(data." + c_snake + `) : undefined;
`
			continue
		}

		t_fields += c_snake + "?: " + ts_type + `;
`
		i_fields += c_snake + "?: " + ts_type + `;
`

		t_constructor_fields += "this." + c_snake + " = " + "data." + c_snake + `;
`
	}

	constructor := `constructor(data: I` + s_upper + `) {
` + t_constructor_fields + `}`

	intf := `export interface I` + s_upper + ` {
` + i_fields + `}`

	class := `export class ` + s_upper + ` {
` + t_fields + `
` + constructor + `
}`

	file := imports + `
` + intf + `

` + class

	return file, nil
}
