package o2c

import (
	"flag"
	"log/slog"
	"path"
	"path/filepath"
	"strings"

	"github.com/handlename/protoc-gen-oas2connect/internal/gen"
	_ "github.com/handlename/protoc-gen-oas2connect/internal/logger"
	"google.golang.org/protobuf/compiler/protogen"
)

const extension = ".oas2connect.go"
const packageSuffix = "oas2connect"

func Run() int {
	var generatedDir string
	var packageName string
	var connectPackagePath string

	var flags flag.FlagSet
	flags.StringVar(&connectPackagePath, "connect_package_path", "", "path to package path generated by connect-go")

	protogen.Options{
		ParamFunc: flags.Set,
	}.Run(func(plugin *protogen.Plugin) error {
		for _, file := range plugin.Files {
			if !file.Generate {
				continue
			}

			modifyGeneratedGoPackageName(file)
			modifyGeneratedFilenamePrefix(file)

			generatedFile := plugin.NewGeneratedFile(file.GeneratedFilenamePrefix+extension, "")

			if err := gen.Generate(
				file,
				strings.Trim(file.GoImportPath.String(), `"`),
				connectPackagePath,
				generatedFile,
			); err != nil {
				slog.Error("failed to generate", slog.String("err", err.Error()), slog.String("file", file.Desc.Path()))
				return err
			}

			if generatedDir == "" {
				generatedDir = path.Dir(string(file.GeneratedFilenamePrefix))
			}

			if packageName == "" {
				packageName = string(file.GoPackageName)
			}
		}

		if err := generateConvertFile(plugin.NewGeneratedFile(generatedDir+"/convert.go", ""), packageName); err != nil {
			return err
		}

		return nil
	})

	return 0
}

func modifyGeneratedGoPackageName(file *protogen.File) {
	file.GoPackageName += packageSuffix
}

func modifyGeneratedFilenamePrefix(file *protogen.File) {
	slashed := filepath.ToSlash(file.GeneratedFilenamePrefix)
	file.GeneratedFilenamePrefix = path.Join(
		path.Dir(slashed),
		string(file.GoPackageName),
		path.Base(slashed),
	)
}

func generateConvertFile(g *protogen.GeneratedFile, packageName string) error {
	if err := gen.GenerateOther("Convert", packageName, g); err != nil {
		return err
	}

	return nil
}
