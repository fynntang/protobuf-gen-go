package repositories

import (
	"bytes"
	"strings"
	"text/template"
)

var crudTemplate = `

type I{{.ServiceType}}Repo interface {
	Create{{.ServiceType}}(ctx context.Context, {{.ServiceType}} *entities.{{.ServiceType}}) error
	Update{{.ServiceType}}(ctx context.Context, updateFields []string,{{.ServiceType}} *entities.{{.ServiceType}}) error
	Get{{.ServiceType}}(ctx context.Context, id entity.ID) (*entities.{{.ServiceType}}, error)
	Get{{.ServiceType}}s(ctx context.Context, filter *database.Filter) (res []*entities.User, count int64, err error)
	Delete{{.ServiceType}}(ctx context.Context, id entity.ID) error
	Log(ctx context.Context) *zap.SugaredLogger
}


{{$firstLetter := (GetFirstLetter .ServiceType)}}
{{$entityName := (GetEntityName .ServiceType)}}


type {{.ServiceType}}Repo struct {
}

func ({{$firstLetter}} *{{.ServiceType}}Repo) DB(ctx context.Context) *gorm.DB {
	return global.DBFromContext(ctx).Model(&entities.{{.ServiceType}}{})
}

func ({{$firstLetter}} {{.ServiceType}}Repo) Create{{.ServiceType}}(ctx context.Context, {{$entityName}} *entities.{{.ServiceType}}) error {
	panic("todo")
}

func ({{$firstLetter}} {{.ServiceType}}Repo) Update{{.ServiceType}}(ctx context.Context, updateFields []string, {{$entityName}} *entities.{{.ServiceType}}) error {
	panic("todo")
}

func ({{$firstLetter}} {{.ServiceType}}Repo) Get{{.ServiceType}}(ctx context.Context, id entity.ID) (*entities.{{.ServiceType}}, error) {
	panic("todo")
}

func ({{$firstLetter}} {{.ServiceType}}Repo) Delete{{.ServiceType}}(ctx context.Context, id entity.ID) error {
	panic("todo")
}

func ({{$firstLetter}} {{.ServiceType}}Repo) Get{{.ServiceType}}s(ctx context.Context, filter *database.Filter) (res []*entities.{{.ServiceType}}, count int64, err error) {
	panic("todo")
}

func ({{$firstLetter}} {{.ServiceType}}Repo) Log(ctx context.Context) *zap.SugaredLogger {
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
		"GetEntityName":      GetEntityName,
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

func GetEntityName(s string) string {
	return strings.ToLower(s[0:1]) + s[1:]
}
