package deliveries

import (
	"bytes"
	"strings"
	"text/template"
)

var crudTemplate = `
type {{.ServiceType}}Service struct {
}

{{$firstLetter := (GetFirstLetter .ServiceType)}}
{{$svrType := .ServiceType}}
{{$svrName := .ServiceName}}
{{- range .MethodSets}}
{{$request := .Request}}
{{$reply := .Reply}}


{{- if not (IsHasPackagePrefix .Request)}}
{{$request = printf "%s.%s" .PackageName .Request}}
{{- end}}

{{- if not (IsHasPackagePrefix .Reply)}}
{{$reply = printf "%s.%s" .PackageName .Reply}}
{{- end}}

func ({{$firstLetter}} {{$svrType}}Service){{.OriginalName}}(c *gin.Context, in *{{$request}})(*{{$reply}},error) {
	panic("todo")
}
{{- end}}

func ({{$firstLetter}} {{.ServiceType}}Service) Log(c *gin.Context) *zap.SugaredLogger {
	return global.Logger(c).Named("{{.ServiceType}}Repo")
}

`

type serviceDesc struct {
	ServiceType string // Greeter
	ServiceName string // helloworld.Greeter
	Metadata    string // api/helloworld/helloworld.proto
	Methods     []*methodDesc
	MethodSets  map[string]*methodDesc
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
