package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"text/template"
)

// Reads environment variables, uses them as values for the tpl_files templates
// and writes the executed templates to the out stream.
func ExecuteTemplates(out io.Writer, tpl_files ...string) error {
	tpl, err := template.ParseFiles(tpl_files...)
	if err != nil {
		return fmt.Errorf("Error parsing template(s): %v", err)
	}

	tpl.Option("missingkey=zero")

	values := make(map[string]string)
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		values[pair[0]] = pair[1]
	}

	err = tpl.Execute(out, values)
	if err != nil {
		return fmt.Errorf("Failed to parse standard input: %v", err)
	}
	return nil
}

func main() {
	err := ExecuteTemplates(os.Stdout, os.Args[1:]...)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
