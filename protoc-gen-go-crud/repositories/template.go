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
}


type {{.ServiceType}}Repo struct {
}

func (m *{{.ServiceType}}Repo) DB(ctx context.Context) *gorm.DB {
	return global.DBFromContext(ctx).Model(&entities.{{.ServiceType}}{})
}

func (m {{.ServiceType}}Repo) Create{{.ServiceType}}(ctx context.Context, {{.ServiceType}} *entities.{{.ServiceType}}) error {
	panic("todo")
}

func (m {{.ServiceType}}Repo) Update{{.ServiceType}}(ctx context.Context, updateFields []string, {{.ServiceType}} *entities.{{.ServiceType}}) error {
	panic("todo")
}

func (m {{.ServiceType}}Repo) Get{{.ServiceType}}(ctx context.Context, id entity.ID) (*entities.{{.ServiceType}}, error) {
	panic("todo")
}

func (m {{.ServiceType}}Repo) Delete{{.ServiceType}}(ctx context.Context, id entity.ID) error {
	panic("todo")
}

func (m {{.ServiceType}}Repo) Get{{.ServiceType}}s(ctx context.Context, filter *database.Filter) (res []*entities.{{.ServiceType}}, count int64, err error) {
	panic("todo")
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
