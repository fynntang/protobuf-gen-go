package repositories

import (
	"fmt"
	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
)

var methodSets = make(map[string]int)

func GenerateRespositoriesFile(gen *protogen.Plugin, file *protogen.File, omitempty bool) *protogen.GeneratedFile {
	if len(file.Services) == 0 || (omitempty && !hasHTTPRule(file.Services)) {
		return nil
	}

	filePrefix := strings.Split(file.GeneratedFilenamePrefix, "/")
	filename := "../internal/" + filePrefix[1] + "/repositories/" + filePrefix[2] + ".go"

	// 判断文件是否存在
	if _, err := os.Stat(strings.TrimPrefix(filename, "../")); err == nil {
		log.Println("文件已存在，跳过生成：", filename)
		return nil
	}

	g := gen.NewGeneratedFile(filename, file.GoImportPath)

	g.P()
	g.P("package ", "repositories")
	g.P()

	generateFileContent(gen, file, g, omitempty)
	return g
}

func generateFileContent(gen *protogen.Plugin, file *protogen.File, g *protogen.GeneratedFile, omitempty bool) {
	if len(file.Services) == 0 {
		return
	}
	g.P()
	for _, service := range file.Services {
		genService(gen, file, g, service, omitempty)
	}
}

func genService(_ *protogen.Plugin, file *protogen.File, g *protogen.GeneratedFile, service *protogen.Service, omitempty bool) {
	sd := &serviceDesc{
		ServiceType: service.GoName,
		ServiceName: string(service.Desc.FullName()),
		Metadata:    file.Desc.Path(),
		PackageName: string(file.GoPackageName),
	}
	//
	//for _, method := range service.Methods {
	//	if method.Desc.IsStreamingClient() || method.Desc.IsStreamingServer() {
	//		continue
	//	}
	//	rule, ok := proto.GetExtension(method.Desc.Options(), annotations.E_Http).(*annotations.HttpRule)
	//	if rule != nil && ok {
	//		sd.Methods = append(sd.Methods, buildCrudRule(g, method, rule, file.GoPackageName))
	//	}
	//}
	//if len(sd.Methods) != 0 {
	g.P(sd.execute())
	//}
}

func hasHTTPRule(services []*protogen.Service) bool {
	for _, service := range services {
		for _, method := range service.Methods {
			if method.Desc.IsStreamingClient() || method.Desc.IsStreamingServer() {
				continue
			}
			rule, ok := proto.GetExtension(method.Desc.Options(), annotations.E_Http).(*annotations.HttpRule)
			if rule != nil && ok {
				return true
			}
		}
	}
	return false
}

func buildCrudRule(g *protogen.GeneratedFile, m *protogen.Method, rule *annotations.HttpRule, packageName protogen.GoPackageName) *methodDesc {
	var (
		path         string
		method       string
		body         string
		responseBody string
	)

	switch pattern := rule.Pattern.(type) {
	case *annotations.HttpRule_Get:
		path = pattern.Get
		method = http.MethodGet
	case *annotations.HttpRule_Put:
		path = pattern.Put
		method = http.MethodPut
	case *annotations.HttpRule_Post:
		path = pattern.Post
		method = http.MethodPost
	case *annotations.HttpRule_Delete:
		path = pattern.Delete
		method = http.MethodDelete
	case *annotations.HttpRule_Patch:
		path = pattern.Patch
		method = http.MethodPatch
	case *annotations.HttpRule_Custom:
		path = pattern.Custom.Path
		method = pattern.Custom.Kind
	}
	body = rule.Body
	responseBody = rule.ResponseBody
	md := buildMethodDesc(g, m, method, path, packageName)

	if method == http.MethodGet || method == http.MethodDelete {
		if body != "" {
			_, _ = fmt.Fprintf(os.Stderr, "\u001B[31mWARN\u001B[m: %s %s body should not be declared.\n", method, path)
		}
	} else {
		if body == "" {
			_, _ = fmt.Fprintf(os.Stderr, "\u001B[31mWARN\u001B[m: %s %s does not declare a body.\n", method, path)
		}
	}
	if body == "*" {
		md.HasBody = true
		md.Body = ""
	} else if body != "" {
		md.HasBody = true
		md.Body = "." + camelCaseVars(body)
	} else {
		md.HasBody = false
	}
	if responseBody == "*" {
		md.ResponseBody = ""
	} else if responseBody != "" {
		md.ResponseBody = "." + camelCaseVars(responseBody)
	}
	return md
}

func buildMethodDesc(g *protogen.GeneratedFile, m *protogen.Method, method, path string, packageName protogen.GoPackageName) *methodDesc {

	defer func() { methodSets[m.GoName]++ }()

	vars := buildPathVars(path)

	for v, s := range vars {
		fields := m.Input.Desc.Fields()

		if s != nil {
			path = replacePath(v, *s, path)
		}
		for _, field := range strings.Split(v, ".") {
			if strings.TrimSpace(field) == "" {
				continue
			}
			if strings.Contains(field, ":") {
				field = strings.Split(field, ":")[0]
			}
			fd := fields.ByName(protoreflect.Name(field))
			if fd == nil {
				fmt.Fprintf(os.Stderr, "\u001B[31mERROR\u001B[m: The corresponding field '%s' declaration in message could not be found in '%s'\n", v, path)
				os.Exit(2)
			}
			if fd.IsMap() {
				fmt.Fprintf(os.Stderr, "\u001B[31mWARN\u001B[m: The field in path:'%s' shouldn't be a map.\n", v)
			} else if fd.IsList() {
				fmt.Fprintf(os.Stderr, "\u001B[31mWARN\u001B[m: The field in path:'%s' shouldn't be a list.\n", v)
			} else if fd.Kind() == protoreflect.MessageKind || fd.Kind() == protoreflect.GroupKind {
				fields = fd.Message().Fields()
			}
		}
	}
	comment := m.Comments.Leading.String() + m.Comments.Trailing.String()
	if comment != "" {
		comment = "// " + m.GoName + strings.TrimPrefix(strings.TrimSuffix(comment, "\n"), "//")
	}

	return &methodDesc{
		Name:         m.GoName,
		OriginalName: string(m.Desc.Name()),
		Num:          methodSets[m.GoName],
		Request:      g.QualifiedGoIdent(m.Input.GoIdent),
		Reply:        g.QualifiedGoIdent(m.Output.GoIdent),
		Comment:      comment,
		Path:         replacePathToGinPath(path),
		Method:       method,
		HasVars:      len(vars) > 0,
		PackageName:  string(packageName),
	}
}

func buildPathVars(path string) (res map[string]*string) {
	if strings.HasSuffix(path, "/") {
		fmt.Fprintf(os.Stderr, "\u001B[31mWARN\u001B[m: Path %s should not end with \"/\" \n", path)
	}
	pattern := regexp.MustCompile(`(?i){([a-z.0-9_\s]*)=?([^{}]*)}`)
	matches := pattern.FindAllStringSubmatch(path, -1)
	res = make(map[string]*string, len(matches))
	for _, m := range matches {
		name := strings.TrimSpace(m[1])
		if len(name) > 1 && len(m[2]) > 0 {
			res[name] = &m[2]
		} else {
			res[name] = nil
		}
	}
	return
}

func replacePath(name string, value string, path string) string {
	pattern := regexp.MustCompile(fmt.Sprintf(`(?i){([\s]*%s[\s]*)=?([^{}]*)}`, name))
	idx := pattern.FindStringIndex(path)
	if len(idx) > 0 {
		path = fmt.Sprintf("%s{%s:%s}%s",
			path[:idx[0]], // The start of the match
			name,
			strings.ReplaceAll(value, "*", ".*"),
			path[idx[1]:],
		)
	}
	return path
}

func camelCaseVars(s string) string {
	subs := strings.Split(s, ".")
	vars := make([]string, 0, len(subs))
	for _, sub := range subs {
		vars = append(vars, camelCase(sub))
	}
	return strings.Join(vars, ".")
}

// camelCase returns the CamelCased name.
// If there is an interior underscore followed by a lower case letter,
// drop the underscore and convert the letter to upper case.
// There is a remote possibility of this rewrite causing a name collision,
// but it's so remote we're prepared to pretend it's nonexistent - since the
// C++ generator lowercase names, it's extremely unlikely to have two fields
// with different capitalization.
// In short, _my_field_name_2 becomes XMyFieldName_2.
func camelCase(s string) string {
	if s == "" {
		return ""
	}
	t := make([]byte, 0, 32)
	i := 0
	if s[0] == '_' {
		// Need a capital letter; drop the '_'.
		t = append(t, 'X')
		i++
	}
	// Invariant: if the next letter is lower case, it must be converted
	// to upper case.
	// That is, we process a word at a time, where words are marked by _ or
	// upper case letter. Digits are treated as words.
	for ; i < len(s); i++ {
		c := s[i]
		if c == '_' && i+1 < len(s) && isASCIILower(s[i+1]) {
			continue // Skip the underscore in s.
		}
		if isASCIIDigit(c) {
			t = append(t, c)
			continue
		}
		// Assume we have a letter now - if not, it's a bogus identifier.
		// The next word is a sequence of characters that must start upper case.
		if isASCIILower(c) {
			c ^= ' ' // Make it a capital letter.
		}
		t = append(t, c) // Guaranteed not lower case.
		// Accept lower case sequence that follows.
		for i+1 < len(s) && isASCIILower(s[i+1]) {
			i++
			t = append(t, s[i])
		}
	}
	return string(t)
}

// Is c an ASCII lower-case letter?
func isASCIILower(c byte) bool {
	return 'a' <= c && c <= 'z'
}

// Is c an ASCII digit?
func isASCIIDigit(c byte) bool {
	return '0' <= c && c <= '9'
}

func protocVersion(gen *protogen.Plugin) string {
	v := gen.Request.GetCompilerVersion()
	if v == nil {
		return "(unknown)"
	}
	var suffix string
	if s := v.GetSuffix(); s != "" {
		suffix = "-" + s
	}
	return fmt.Sprintf("v%d.%d.%d%s", v.GetMajor(), v.GetMinor(), v.GetPatch(), suffix)
}

// replacePathToGinPath Conversion parameter routing {xx} --> :xx
func replacePathToGinPath(path string) string {
	paths := strings.Split(path, "/")
	for i, p := range paths {
		if len(p) > 0 && (p[0] == '{' && p[len(p)-1] == '}' || p[0] == ':') {
			paths[i] = ":" + p[1:len(p)-1]
		}
	}
	return strings.Join(paths, "/")
}
