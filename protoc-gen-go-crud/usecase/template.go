package deliveries

import (
	"bytes"
	"strings"
	"text/template"
)

var crudTemplate = `

{{$firstLetter := (GetFirstLetter .ServiceType)}}


type I{{.ServiceType}}Usecase interface {
	Create{{.ServiceType}}(ctx context.Context,in *{{.PackageName}}.Create{{.ServiceType}}Request) error
	Update{{.ServiceType}}(ctx context.Context,in *{{.PackageName}}.Update{{.ServiceType}}Request) error
	Delete{{.ServiceType}}(ctx context.Context,in *{{.PackageName}}.Delete{{.ServiceType}}Request) error
	Get{{.ServiceType}}(ctx context.Context,in *{{.PackageName}}.Get{{.ServiceType}}Request) (*entities.{{.ServiceType}},error)
	List{{.ServiceType}}(ctx context.Context,in *{{.PackageName}}.List{{.ServiceType}}Request) ([]*entities.{{.ServiceType}},int64,error)
	Log(ctx context.Context) *zap.SugaredLogger
}

type {{.ServiceType}}UseCase struct {
}

func New{{.ServiceType}}UseCase() I{{.ServiceType}}Usecase {
	return &{{.ServiceType}}UseCase{}
}

func ({{$firstLetter}} *{{.ServiceType}}UseCase) Create{{.ServiceType}}(ctx context.Context,in *{{.PackageName}}.Create{{.ServiceType}}Request) error {
	panic("todo")
}

func ({{$firstLetter}} *{{.ServiceType}}UseCase) Update{{.ServiceType}}(ctx context.Context,in *{{.PackageName}}.Update{{.ServiceType}}Request) error {
	panic("todo")
}

func ({{$firstLetter}} *{{.ServiceType}}UseCase) Delete{{.ServiceType}}(ctx context.Context,in *{{.PackageName}}.Delete{{.ServiceType}}Request) error{
	panic("todo")
}

func ({{$firstLetter}} *{{.ServiceType}}UseCase) Get{{.ServiceType}}(ctx context.Context,in *{{.PackageName}}.Get{{.ServiceType}}Request) (*entities.{{.ServiceType}},error){
	panic("todo")
}

func ({{$firstLetter}} *{{.ServiceType}}UseCase) List{{.ServiceType}}(ctx context.Context,in *{{.PackageName}}.List{{.ServiceType}}Request) ([]*entities.{{.ServiceType}},int64,error){
	panic("todo")
}

func ({{$firstLetter}} {{.ServiceType}}UseCase) Log(ctx context.Context) *zap.SugaredLogger {
	return global.Logger(ctx).Named("{{.ServiceType}}Repo")
}

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
