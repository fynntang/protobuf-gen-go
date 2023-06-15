package entities

import (
	"bytes"
	"strings"
	"text/template"
)

var crudTemplate = `

{{$firstLetter := (GetFirstLetter .ServiceType)}}
{{$tableName := (GetTableName .ServiceType)}}

type {{.ServiceType}} struct {
	entity.BaseEntity
}



func ({{$firstLetter}} *{{.ServiceType}}) TableName() string { return "{{$tableName}}" }

`

type serviceDesc struct {
	ServiceType string // Greeter
	ServiceName string // helloworld.Greeter
	Metadata    string // api/helloworld/helloworld.proto
	Methods     []*methodDesc
	MethodSets  map[string]*methodDesc
	PackageName string
}

type methodDesc struct {
	// method
	Name         string
	OriginalName string // The parsed original name
	Num          int
	Request      string
	Reply        string
	Comment      string
	Path         string
	Method       string
	HasVars      bool
	HasBody      bool
	Body         string
	ResponseBody string
	PackageName  string
}

func (s *serviceDesc) execute() string {
	s.MethodSets = make(map[string]*methodDesc)
	for _, m := range s.Methods {
		s.MethodSets[m.Name] = m
	}
	buf := new(bytes.Buffer)
	tmpl, err := template.New("crud").Funcs(template.FuncMap{
		"IsHasPackagePrefix": IsHasPackagePrefix,
		"GetFirstLetter":     GetFirstLetter,
		"GetTableName":       GetTableName,
	}).Parse(strings.TrimSpace(crudTemplate))
	if err != nil {
		panic(err)
	}
	if err := tmpl.Execute(buf, s); err != nil {
		panic(err)
	}
	return strings.Trim(buf.String(), "\r\n")
}

func IsHasPackagePrefix(s string) bool {
	return strings.Contains(s, ".")
}

func GetFirstLetter(s string) string {
	return strings.ToLower(s[0:1])
}

func GetTableName(s string) string {
	return strings.ToLower(s[0:1]) + s[1:] + "s"
}
