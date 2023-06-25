package writer

import (
	"bufio"
	"log"
	"os"
	"strings"

	"github.com/stscoundrel/old-swedish-dictionary-builder/dictionary"
)

var DSL_INDENT string = "    "

func formatEntry(entry dictionary.DictionaryEntry) []string {
	entryLines := []string{}

	entryLines = append(entryLines, entry.Headword)

	for _, definition := range entry.Definitions {
		entryLines = append(entryLines, DSL_INDENT+""+strings.Trim(definition, " "))
	}

	if len(entry.GrammaticalAspect) > 0 {
		entryLines = append(entryLines, DSL_INDENT+"Grammatical aspect: "+entry.GrammaticalAspect)
	}
	if len(entry.Information) > 0 {
		entryLines = append(entryLines, DSL_INDENT+"Additional info: "+entry.Information)
	}

	if len(entry.AlternativeForms) > 0 {
		synonymLine := DSL_INDENT + "Alternative forms: " + strings.Join(entry.AlternativeForms, ", ")
		entryLines = append(entryLines, synonymLine)
	}

	return entryLines
}

func buildHeaders(name string) []string {
	return []string{
		"#NAME " + name,
		"#INDEX_LANGUAGE \"Old Swedish\"",
		"#CONTENTS_LANGUAGE \"Swedish\"",
	}
}

func formatDSL(entries []dictionary.DictionaryEntry, name string) []string {
	dsl := buildHeaders(name)

	for _, entry := range entries {
		entryLines := formatEntry(entry)
		dsl = append(dsl, entryLines...)
	}

	return dsl
}

func WriteDsl(filePath string, entries []dictionary.DictionaryEntry, name string) {
	dslLines := formatDSL(entries, name)

	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatalf("Error creating DSL file: %s", err)
	}

	datawriter := bufio.NewWriter(file)

	for _, line := range dslLines {
		_, _ = datawriter.WriteString(line + "\n")
	}

	datawriter.Flush()
	file.Close()
}
