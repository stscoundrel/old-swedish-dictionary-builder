package dictionary

import "encoding/xml"

type feat struct {
	Name  string `xml:"att,attr"`
	Value string `xml:"val,attr"`
}

type sense struct {
	Prefix     feat `xml:"feat"`
	Definition feat `xml:"Definition>feat"`
}

type rawDictionaryEntry struct {
	LemmaFormPresentation []feat  `xml:"Lemma>FormRepresentation>feat"`
	WordForm              []feat  `xml:"WordForm>feat"`
	Sense                 []sense `xml:"Sense"`
}

type rawDictionaryEntries struct {
	XMLName xml.Name             `xml:"LexicalResource"`
	Entries []rawDictionaryEntry `xml:"Lexicon>LexicalEntry"`
}
