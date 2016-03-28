package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

// {{.Type}} Int16
// {{.NativeType}} int16
// {{.Indefinite}} an
// {{.NullType}} NullInt64
// {{.NullTypeField}} Int64
// {{.NullTypeVal}} int64
// {{.ZeroVal}} 0
// {{.Var}} n

func main() {
	gopath := os.Getenv("GOPATH")
	relfile := filepath.FromSlash("src/github.com/spkg/nullable/template")
	for _, godir := range filepath.SplitList(gopath) {
		templateDir := filepath.Join(godir, relfile)
		for _, fileName := range []string{"_type_test.go", "_type.go"} {
			filePath := filepath.Join(templateDir, fileName)
			_, err := os.Stat(filePath)
			if err == nil {
				for _, p := range params {
					tmpl, err := template.ParseFiles(filePath)
					if err != nil {
						fmt.Println(err)
						os.Exit(1)
					}
					outFileName := strings.Replace(fileName, "_type", strings.ToLower(p.Type), 1)
					outFile, err := os.Create(outFileName)
					if err != nil {
						fmt.Println(err)
						os.Exit(1)
					}
					if err := tmpl.Execute(outFile, p); err != nil {
						fmt.Println(fileName+":", err)
						os.Exit(1)
					}
					outFile.Close()
					println(outFile.Name())
				}
			}
		}
	}
}

// params contains parameters for the template.
var params = []struct {
	Type          string
	NativeType    string
	Indefinite    string
	NullType      string
	NullTypeField string
	NullTypeVal   string
	ZeroVal       string
	Var           string
	Imports       []string
	NoScan        bool
	NoValue       bool
}{
	{
		Type:          "String",
		NativeType:    "string",
		Indefinite:    "a",
		NullType:      "NullString",
		NullTypeField: "String",
		NullTypeVal:   "string",
		ZeroVal:       "\"\"",
		Var:           "s",
	},
	{
		Type:          "Bool",
		NativeType:    "bool",
		Indefinite:    "a",
		NullType:      "NullBool",
		NullTypeField: "Bool",
		NullTypeVal:   "bool",
		ZeroVal:       "false",
		Var:           "b",
	},
	{
		Type:          "Time",
		NativeType:    "time.Time",
		Indefinite:    "a",
		NullType:      "-",
		NullTypeField: "-",
		NullTypeVal:   "",
		ZeroVal:       "time.Time{}",
		Var:           "tm",
		NoScan:        true,
		Imports:       []string{"time"},
	},
	{
		Type:          "Float64",
		NativeType:    "float64",
		Indefinite:    "a",
		NullType:      "NullFloat64",
		NullTypeField: "Float64",
		NullTypeVal:   "float64",
		ZeroVal:       "0",
		Var:           "n",
	},
	{
		Type:          "Float32",
		NativeType:    "float32",
		Indefinite:    "a",
		NullType:      "NullFloat64",
		NullTypeField: "Float64",
		NullTypeVal:   "float64",
		ZeroVal:       "0",
		Var:           "n",
	},
	{
		Type:          "Int",
		NativeType:    "int",
		Indefinite:    "an",
		NullType:      "NullInt64",
		NullTypeField: "Int64",
		NullTypeVal:   "int64",
		ZeroVal:       "0",
		Var:           "n",
	},
	{
		Type:          "Uint",
		NativeType:    "uint",
		Indefinite:    "a",
		NullType:      "NullInt64",
		NullTypeField: "Int64",
		NullTypeVal:   "int64",
		ZeroVal:       "0",
		Var:           "n",
	},
	{
		Type:          "Int64",
		NativeType:    "int64",
		Indefinite:    "an",
		NullType:      "NullInt64",
		NullTypeField: "Int64",
		NullTypeVal:   "int64",
		ZeroVal:       "0",
		Var:           "n",
	},
	{
		Type:          "Uint64",
		NativeType:    "uint64",
		Indefinite:    "a",
		NullType:      "NullInt64",
		NullTypeField: "Int64",
		NullTypeVal:   "int64",
		ZeroVal:       "0",
		Var:           "n",
	},
	{
		Type:          "Int32",
		NativeType:    "int32",
		Indefinite:    "an",
		NullType:      "NullInt64",
		NullTypeField: "Int64",
		NullTypeVal:   "int64",
		ZeroVal:       "0",
		Var:           "n",
	},
	{
		Type:          "Uint32",
		NativeType:    "uint32",
		Indefinite:    "a",
		NullType:      "NullInt64",
		NullTypeField: "Int64",
		NullTypeVal:   "int64",
		ZeroVal:       "0",
		Var:           "n",
	},
	{
		Type:          "Int16",
		NativeType:    "int16",
		Indefinite:    "an",
		NullType:      "NullInt64",
		NullTypeField: "Int64",
		NullTypeVal:   "int64",
		ZeroVal:       "0",
		Var:           "n",
	},
	{
		Type:          "Uint16",
		NativeType:    "uint16",
		Indefinite:    "a",
		NullType:      "NullInt64",
		NullTypeField: "Int64",
		NullTypeVal:   "int64",
		ZeroVal:       "0",
		Var:           "n",
	},
	{
		Type:          "Int8",
		NativeType:    "int8",
		Indefinite:    "an",
		NullType:      "NullInt64",
		NullTypeField: "Int64",
		NullTypeVal:   "int64",
		ZeroVal:       "0",
		Var:           "n",
	},
	{
		Type:          "Byte",
		NativeType:    "byte",
		Indefinite:    "a",
		NullType:      "NullInt64",
		NullTypeField: "Int64",
		NullTypeVal:   "int64",
		ZeroVal:       "0",
		Var:           "n",
	},
}
