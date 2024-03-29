package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/fatih/camelcase"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
)

var (
	tags           string
	folder         string
	ignore         string
	relativeFolder string
	pbFiles        []string
	fset           *token.FileSet
	tagsValue      []string
	ignoreFields   []string
	ignoreFieldVal = make(map[string]string)
)

func init() {
	flag.StringVar(&tags, "tags", "", "tags for struct eg: json,xml")
	flag.StringVar(&folder, "folder", "", "paths for struct eg: ./pb")
	flag.StringVar(&ignore, "ignore", "", "ignore fields for struct eg: form:file,json:file")
	fset = token.NewFileSet()
}
func main() {
	flag.Parse()

	checkArgs()
	walkFolder(relativeFolder)

	for _, f := range pbFiles {
		parsePbFile(f)
	}
}

func checkArgs() {
	if tags == "" {
		log.Fatal("tags is empty")
	}

	if folder == "" {
		log.Fatal("folder is empty")
	}

	pwd, _ := os.Getwd()
	relativeFolder = path.Join(pwd, folder)
	if _, err := os.Stat(relativeFolder); os.IsNotExist(err) {
		log.Fatalf("folder %s is not exist", relativeFolder)
	}

	tagsValue = strings.Split(tags, ",")

	if ignore != "" {
		for _, s := range strings.Split(ignore, ",") {
			i := strings.Split(s, ":")
			if len(i) != 2 {
				log.Fatalf("ignore field %s format error", s)
			}
			key := fmt.Sprintf(`%s:"%s,omitempty"`, i[0], i[1])
			ignoreFields = append(ignoreFields, key)
			ignoreFieldVal[key] = fmt.Sprintf(`%s:"-"`, i[0])
		}
	}
}

func walkFolder(folderPath string) {
	if err := filepath.Walk(folderPath, func(fp string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if strings.HasSuffix(info.Name(), ".pb.go") {
			pbFiles = append(pbFiles, fp)
		}
		return nil
	}); err != nil {
		log.Fatalf("walk folder %s error: %v", folderPath, err)
	}
}

func parsePbFile(pbFile string) {
	pbContent, err := os.ReadFile(pbFile)
	if err != nil {
		log.Fatalf("read file %s error: %v", pbFile, err)
	}
	f, err := parser.ParseFile(fset, pbFile, pbContent, parser.ParseComments)
	if err != nil {
		log.Fatalf("parse file %s error: %v", pbFile, err)
	}

	ast.Inspect(f, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.TypeSpec:
			switch t := x.Type.(type) {
			case *ast.StructType:
				handleTags(t.Fields)
			}
		}
		return true
	})

	var buf bytes.Buffer
	if err := format.Node(&buf, fset, f); err != nil {
		log.Fatalf("format file %s error: %v", pbFile, err)
	}

	if err := os.WriteFile(pbFile, buf.Bytes(), 0644); err != nil {
		log.Fatalf("write file %s error: %v", pbFile, err)
	}
}

type fieldTag struct {
	tagName  string
	tagValue string
}

func handleTags(fields *ast.FieldList) {
	for _, field := range fields.List {
		if len(field.Names) == 0 || field.Tag == nil {
			continue
		}

		var newName string
		for _, name := range field.Names {
			newName = strings.ToLower(strings.Join(camelcase.Split(name.Name), "_"))
		}

		tv := strings.ReplaceAll(field.Tag.Value, "`", "")
		tagList := strings.Split(tv, " ")
		fieldTags := make(map[string]*fieldTag, 0)
		var sortTags []string
		for _, t := range tagList {
			tf := strings.Split(t, ":")
			if len(tf) != 2 {
				continue
			}
			fieldTags[tf[0]] = &fieldTag{
				tagName:  strings.TrimSpace(tf[0]),
				tagValue: strings.Replace(tf[1], "\"", "", -1),
			}
			sortTags = append(sortTags, tf[0])
		}

		for _, tag := range tagsValue {
			if _, ok := fieldTags[tag]; ok {
				continue
			}

			fieldTags[tag] = &fieldTag{tagName: tag, tagValue: fmt.Sprintf("%s,omitempty", newName)}
			sortTags = append(sortTags, tag)
		}

		var newTags string
		for _, tag := range sortTags {
			newTags += fmt.Sprintf("%s:\"%s\" ", fieldTags[tag].tagName, fieldTags[tag].tagValue)
		}

		for _, v := range ignoreFields {
			if vv, ok := ignoreFieldVal[v]; ok {
				newTags = strings.ReplaceAll(newTags, v, vv)
			}
		}

		field.Tag.Value = fmt.Sprintf("`%s`", newTags)
	}
}
