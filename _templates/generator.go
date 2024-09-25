package main

import (
	"errors"
	"fmt"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"os"
	"regexp"
	"strings"
	"text/template"
	"unicode"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("Object name is required.")
		fmt.Println("Usage: go run YourModelName (without brackets)")
		os.Exit(2)
	}
	modelName := os.Args[1]
	rm := false
	if len(os.Args) >= 3 {
		rm = true
	}
	if modelName != "" {
		tv := &TemplateVariables{
			ModelName:      cases.Title(language.English, cases.NoLower).String(modelName),
			LowerModelName: toLowerCaseFirst(modelName),
		}
		err := createFromTemplates(tv)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(2)
		}
		if rm {
			err := registerModel(tv)
			if err != nil {
				fmt.Println(err.Error())
				os.Exit(2)
			}
		}
		return
	} else {
		fmt.Println("Object name is required.")
		fmt.Println("Usage: go run YourModelName (without brackets)")
	}
}

type TemplateVariables struct {
	ModelName      string
	LowerModelName string
}

func createFromTemplates(tv *TemplateVariables) error {
	snakeCaseName := toSnakeCase(tv.ModelName)
	// files to create
	var create map[string]string = map[string]string{
		"controller.tpl":     fmt.Sprintf("src/controller/internal/%s.go", snakeCaseName),
		"handler.tpl":        fmt.Sprintf("src/handler/api_%s.go", snakeCaseName),
		"model.tpl":          fmt.Sprintf("src/storage/model/%s.go", snakeCaseName),
		"request_id.tpl":     fmt.Sprintf("src/http/request/%s_id.go", snakeCaseName),
		"request_list.tpl":   fmt.Sprintf("src/http/request/%s_list.go", snakeCaseName),
		"request_create.tpl": fmt.Sprintf("src/http/request/%s_create.go", snakeCaseName),
		"request_update.tpl": fmt.Sprintf("src/http/request/%s_update.go", snakeCaseName),
		"response.tpl":       fmt.Sprintf("src/http/response/%s.go", snakeCaseName),
		"service_impl.tpl":   fmt.Sprintf("src/service/internal/%s.go", snakeCaseName),
		"storage_impl.tpl":   fmt.Sprintf("src/storage/internal/%s.go", snakeCaseName),
	}

	// ensure we can create all pack of files
	for _, dest := range create {
		if _, err := os.Stat(dest); !errors.Is(err, os.ErrNotExist) {
			fmt.Println(fmt.Sprintf("File '%s' already exist. Cannot proceed", dest))
			os.Exit(2)
		}
	}

	// creating files
	for src, dest := range create {
		fmt.Println(fmt.Sprintf("Processing template '%s'", src))
		tmp, err := template.ParseFiles(src)
		if err != nil {
			return err
		}
		fd, err := os.OpenFile("../"+dest, os.O_CREATE|os.O_APPEND|os.O_RDWR|os.O_EXCL, 0666)
		fmt.Println(fmt.Sprintf("Creating file '%s'", dest))
		if err != nil {
			return err
		}
		fmt.Println(fmt.Sprintf("Writing to file '%s'", dest))
		err = tmp.Execute(fd, tv)
		if err != nil {
			return err
		}
		fmt.Println(fmt.Sprintf("File '%s' complite.", dest))
	}

	return nil
}

func registerModel(tv *TemplateVariables) error {
	//TODO:
	return nil
	tmp, err := template.ParseFiles("storage_interface.tpl")
	if err != nil {
		return err
	}
	fd, err := os.OpenFile("../src/storage/storages.go", os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		return err
	}
	err = tmp.Execute(fd, tv)
	if err != nil {
		return err
	}
	return nil
}

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func toSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

func toLowerCaseFirst(str string) string {
	a := []rune(str)
	a[0] = unicode.ToLower(a[0])
	return string(a)
}
