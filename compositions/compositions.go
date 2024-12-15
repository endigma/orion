package compositions

import (
	"embed"
	"fmt"
	"io"
	"regexp"
	"text/template"
)

//go:embed templates/*
var files embed.FS
var templates = template.Must(template.ParseFS(files, "templates/*"))

type Composition struct {
	Name     string
	Author   string
	Category string
	Init     string
}

func RenderHeader(w io.Writer, data Composition) error {
	err := templates.ExecuteTemplate(w, "header.tmpl", data)
	if err != nil {
		return fmt.Errorf("failed to render header: %w", err)
	}

	return nil
}

func RenderComposition(w io.Writer, data Composition) error {
	data.Init = CleanSQF(data.Init)

	err := templates.ExecuteTemplate(w, "composition.tmpl", data)
	if err != nil {
		return fmt.Errorf("failed to render composition: %w", err)
	}

	return nil
}

// CleanSQF removes characters that are not permitted in the "init" property.
func CleanSQF(s string) string {
	s = removeTabs(s)
	s = removeComments(s)
	s = escapeQuotes(s)
	s = removeNewlines(s)

	return s
}

func escapeQuotes(s string) string {
	quote := regexp.MustCompile(`"((.|\n)*?)"`)
	s = quote.ReplaceAllString(s, `""$1""`)

	return s
}

func removeTabs(s string) string {
	tab := regexp.MustCompile(`\t`)
	s = tab.ReplaceAllString(s, " ")

	return s
}

func removeNewlines(s string) string {
	newline := regexp.MustCompile(`(\r\n|\r|\n)`)
	s = newline.ReplaceAllString(s, ` " \n "`)

	return s
}

func removeComments(s string) string {
	inlineComment := regexp.MustCompile(`(?m)^//.*$`)

	s = inlineComment.ReplaceAllString(s, "")

	blockComment := regexp.MustCompile(`/\*(.|\n)*\*/`)
	s = blockComment.ReplaceAllString(s, "")

	return s
}
