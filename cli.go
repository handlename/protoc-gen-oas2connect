package o2c

import (
	"log/slog"

	"github.com/handlename/protoc-gen-oas2connect/internal/gen"
	_ "github.com/handlename/protoc-gen-oas2connect/internal/logger"
	"google.golang.org/protobuf/compiler/protogen"
)

func Run() int {
	opts := protogen.Options{}

	p := func(plugin *protogen.Plugin) error {
		g, err := gen.NewGenerator(plugin.Files)
		if err != nil {
			slog.Error("failed to create generator", slog.String("err", err.Error()))
			return err
		}

		return g.Generate(func(filename string, content []byte) error {
			if _, err := plugin.NewGeneratedFile(filename, "").Write(content); err != nil {
				return err
			}

			return nil
		})
	}

	opts.Run(p)

	return 0
}
