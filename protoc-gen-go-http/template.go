package main

import (
	"bytes"
	"strings"
	"text/template"
)

var httpTemplate = `
{{$svrType := .ServiceType}}
{{$svrName := .ServiceName}}

{{- range .MethodSets}}
const Route{{$svrType}}{{.OriginalName}} = "{{.Path}}"
{{- end}}

type {{.ServiceType}}HTTPServer interface {
{{- range .MethodSets}}
	{{- if ne .Comment ""}}
	{{.Comment}}
	{{- end}}
	{{.Name}}(*gin.Context, *{{.Request}}) (*{{.Reply}}, error)
{{- end}}
}

func Register{{.ServiceType}}HTTPServer(r gin.IRouter, srv {{.ServiceType}}HTTPServer) {
	{{- range .Methods}}
	r.{{.Method}}(Route{{$svrType}}{{.OriginalName}}, _{{$svrType}}_{{.Name}}{{.Num}}_HTTP_Handler(srv))
	{{- end}}
}

{{range .Methods}}
func _{{$svrType}}_{{.Name}}{{.Num}}_HTTP_Handler(srv {{$svrType}}HTTPServer) gin.HandlerFunc {
	return func(c *gin.Context) {
		var in {{.Request}}
		{{- if .HasBody}}

		if err := c.Bind(&in{{.Body}}); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		
		{{- if not (eq .Body "")}}
		if err := c.BindQuery(&in); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		{{- end}}
		{{- else}}
		if err := c.BindQuery(&in{{.Body}}); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		{{- end}}
		{{- if .HasVars}}
		if err := c.BindUri(&in); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		{{- end}}

		out, err := srv.{{.Name}}(c, &in)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		if c.IsAborted() {
			return
		}

		var buffer bytes.Buffer
		var data interface{}
		var newData interface{}

		data = out
		if message, ok := data.(proto.Message); ok {
			t := jsonpb.Marshaler{EmitDefaults: true, OrigName: true, EnumsAsInts:true}
			t.Marshal(&buffer, message)
			json.Unmarshal(buffer.Bytes(), &newData)
		} else {
			newData = data
		}

		c.JSON(http.StatusOK, newData{{.ResponseBody}})
	}
}
{{end}}
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
	// http_rule
	Path         string
	Method       string
	HasVars      bool
	HasBody      bool
	Body         string
	ResponseBody string
}

func (s *serviceDesc) execute() string {
	s.MethodSets = make(map[string]*methodDesc)
	for _, m := range s.Methods {
		s.MethodSets[m.Name] = m
	}
	buf := new(bytes.Buffer)
	tmpl, err := template.New("http").Parse(strings.TrimSpace(httpTemplate))
	if err != nil {
		panic(err)
	}
	if err := tmpl.Execute(buf, s); err != nil {
		panic(err)
	}

	return strings.Trim(buf.String(), "\r\n")
}
