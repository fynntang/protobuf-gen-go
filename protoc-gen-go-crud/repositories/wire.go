package repositories

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type TemplateParams struct {
	FunctionNames []string
}

func GenerateWireFile(path string, sets []string) error {
	// 定义模板
	tmpl := `package repositories

import "github.com/google/wire"

var ProviderSet = wire.NewSet({{ range $index, $element := .FunctionNames }}{{ if $index }}, {{ end }}{{ $element }}{{ end }})`

	// 解析模板
	t, err := template.New("generated-file").Parse(tmpl)
	if err != nil {
		return err
	}

	f := strings.Split(path, "V1")
	filename := fmt.Sprintf("./internal/%s/repositories/wire.go", strings.ToLower(f[0]))
	oldWireName := fmt.Sprintf("./internal/%s/repositories/repository.go", strings.ToLower(f[0]))

	// 判断文件是否存在
	if _, err := os.Stat(strings.TrimPrefix(filename, "../")); err == nil {
		log.Println("wire已存在，跳过生成：", filename)
		return nil
	}
	// 兼容历史数据
	if _, err := os.Stat(strings.TrimPrefix(oldWireName, "../")); err == nil {
		log.Println("wire文件已存在，跳过生成：", filename)
		return nil
	}

	functions := []string{}
	for _, set := range sets {
		functions = append(functions, fmt.Sprintf("New%sRepo", set))
	}

	params := TemplateParams{
		FunctionNames: functions,
	}

	// 目录检查
	dir := filepath.Dir(filename)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			log.Fatalf("failed to create directory: %v", err)
		}
	}

	// 检查文件是否存在，如果不存在则创建它
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		file, err := os.Create(filename)
		if err != nil {
			return err
		}
		defer file.Close()
	}

	// 打开文件
	file, err := os.OpenFile(filename, os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// 生成新内容并写入文件
	err = t.Execute(file, params)
	if err != nil {
		return err
	}

	return nil
}
