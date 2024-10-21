package enums

import (
	"fmt"
	"os"
	"text/template"

	adddoublequotestxt "github.com/rolniuq/mypackage/add-double-quotes-txt"
	"github.com/rolniuq/mypackage/lor"
)

const EnumTemplate = `package {{.PackageName}}

type {{.TypeName}} string

const (
{{- range .Enums }}
	ENUM_{{ . }} {{$.TypeName}} = "{{ . }}"
{{- end }}
)
`

type EnumGenerator struct {
	PackageName string
	TypeName    string
	Enums       []string
}

func GenerateGoCodeFromJSON(enumName, fileName, outputPath string) error {
	out, err := lor.ReadJsonFile[adddoublequotestxt.Output](fileName)
	if err != nil {
		return err
	}

	g := EnumGenerator{
		PackageName: "enums",
		TypeName:    enumName,
		Enums:       out.Data,
	}

	tmpl, err := template.New("enum").Parse(EnumTemplate)
	if err != nil {
		return fmt.Errorf("failed to parse template: %w", err)
	}

	outputFile, err := os.Create(fmt.Sprintf("%s/%s.go", ".out", outputPath))
	if err != nil {
		return fmt.Errorf("failed to create output file: %w", err)
	}
	defer outputFile.Close()

	err = tmpl.Execute(outputFile, g)
	if err != nil {
		return fmt.Errorf("failed to execute template: %w", err)
	}

	return nil
}
