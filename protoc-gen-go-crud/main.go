package main

import (
	"flag"
	"github.com/fynntang/protobuf-gen-go/protoc-gen-go-crud/deliveries"
	"github.com/fynntang/protobuf-gen-go/protoc-gen-go-crud/entities"
	"github.com/fynntang/protobuf-gen-go/protoc-gen-go-crud/repositories"
	usecase "github.com/fynntang/protobuf-gen-go/protoc-gen-go-crud/usecase"
	_ "github.com/gin-gonic/gin"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/pluginpb"
)

var (
	showVersion = flag.Bool("version", false, "print the version and exit")
	omitempty   = flag.Bool("omitempty", true, "omit if google.api is empty")
)

func main() {
	flag.Parse()
	if *showVersion {
		//fmt.Printf("protoc-gen-g %v\n", release)
		return
	}
	protogen.Options{
		ParamFunc: flag.CommandLine.Set,
	}.Run(func(gen *protogen.Plugin) error {
		gen.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)
		for _, f := range gen.Files {
			if !f.Generate {
				continue
			}
			deliveries.GenerateDeliveriesFile(gen, f, *omitempty)
			usecase.GenerateUseCasesFile(gen, f, *omitempty)
			entities.GenerateEntitiesFile(gen, f, *omitempty)
			repositories.GenerateRespositoriesFile(gen, f, *omitempty)
		}
		return nil
	})
}
