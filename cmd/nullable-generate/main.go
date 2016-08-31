package main

import (
	"bytes"
	"fmt"
	"go/format"
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
					if p.ZeroVal == "" {
						p.ZeroVal = "0"
					}
					if p.NonZeroVal == "" {
						p.NonZeroVal = "1"
					}
					if p.Var == "" {
						p.Var = "n"
					}
					if p.Indefinite == "" {
						p.Indefinite = "a"
					}
					tmpl, err := template.ParseFiles(filePath)
					if err != nil {
						fmt.Println(err)
						os.Exit(1)
					}
					var outBuf bytes.Buffer
					if err = tmpl.Execute(&outBuf, p); err != nil {
						fmt.Println(fileName+":", err)
						os.Exit(1)
					}
					outFileName := strings.Replace(fileName, "_type", strings.ToLower(p.Type), 1)
					outFile, err := os.Create(outFileName)
					if err != nil {
						fmt.Println(err)
						os.Exit(1)
					}
					outBytes, err := format.Source(outBuf.Bytes())
					if err != nil {
						fmt.Println("format failed:", err)
						outBytes = outBuf.Bytes()
					}
					if _, err = outFile.Write(outBytes); err != nil {
						fmt.Println(err)
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
	NonZeroVal    string
	Var           string
	Imports       []string
	NoScan        bool
	NoValue       bool
	NoCast        bool
}{
	{
		Type:          "String",
		NativeType:    "string",
		Indefinite:    "a",
		NullType:      "NullString",
		NullTypeField: "String",
		NullTypeVal:   "string",
		ZeroVal:       `""`,
		NonZeroVal:    `"non-zero"`,
		Var:           "s",
		NoCast:        true,
	},
	{
		Type:          "Bool",
		NativeType:    "bool",
		Indefinite:    "a",
		NullType:      "NullBool",
		NullTypeField: "Bool",
		NullTypeVal:   "bool",
		ZeroVal:       "false",
		NonZeroVal:    "true",
		Var:           "b",
		NoCast:        true,
	},
	{
		Type:          "Time",
		NativeType:    "time.Time",
		Indefinite:    "a",
		NullType:      "-",
		NullTypeField: "-",
		NullTypeVal:   "",
		ZeroVal:       "time.Time{}",
		NonZeroVal:    "time.Now()",
		Var:           "tm",
		NoScan:        true,
		NoCast:        true,
		Imports:       []string{"time"},
	},
	{
		Type:          "Float64",
		NativeType:    "float64",
		Indefinite:    "a",
		NullType:      "NullFloat64",
		NullTypeField: "Float64",
		NullTypeVal:   "float64",
		NoCast:        true,
	},
	{
		Type:          "Float32",
		NativeType:    "float32",
		Indefinite:    "a",
		NullType:      "NullFloat64",
		NullTypeField: "Float64",
		NullTypeVal:   "float64",
	},
	{
		Type:          "Int",
		NativeType:    "int",
		Indefinite:    "an",
		NullType:      "NullInt64",
		NullTypeField: "Int64",
		NullTypeVal:   "int64",
	},
	{
		Type:          "Uint",
		NativeType:    "uint",
		Indefinite:    "a",
		NullType:      "NullInt64",
		NullTypeField: "Int64",
		NullTypeVal:   "int64",
	},
	{
		Type:          "Int64",
		NativeType:    "int64",
		Indefinite:    "an",
		NullType:      "NullInt64",
		NullTypeField: "Int64",
		NullTypeVal:   "int64",
		NoCast:        true,
	},
	{
		Type:          "Uint64",
		NativeType:    "uint64",
		Indefinite:    "a",
		NullType:      "NullInt64",
		NullTypeField: "Int64",
		NullTypeVal:   "int64",
	},
	{
		Type:          "Int32",
		NativeType:    "int32",
		Indefinite:    "an",
		NullType:      "NullInt64",
		NullTypeField: "Int64",
		NullTypeVal:   "int64",
	},
	{
		Type:          "Uint32",
		NativeType:    "uint32",
		Indefinite:    "a",
		NullType:      "NullInt64",
		NullTypeField: "Int64",
		NullTypeVal:   "int64",
	},
	{
		Type:          "Int16",
		NativeType:    "int16",
		Indefinite:    "an",
		NullType:      "NullInt64",
		NullTypeField: "Int64",
		NullTypeVal:   "int64",
	},
	{
		Type:          "Uint16",
		NativeType:    "uint16",
		Indefinite:    "a",
		NullType:      "NullInt64",
		NullTypeField: "Int64",
		NullTypeVal:   "int64",
	},
	{
		Type:          "Int8",
		NativeType:    "int8",
		Indefinite:    "an",
		NullType:      "NullInt64",
		NullTypeField: "Int64",
		NullTypeVal:   "int64",
	},
	{
		Type:          "Byte",
		NativeType:    "byte",
		Indefinite:    "a",
		NullType:      "NullInt64",
		NullTypeField: "Int64",
		NullTypeVal:   "int64",
	},
}
