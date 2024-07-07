package main

import (
	"fmt"
	"strings"

	"github.com/muesli/reflow/wordwrap"
)

type word struct {
	Word      string      `json:"word,omitempty"`
	Phonetic  string      `json:"phonetic,omitempty"`
	Phonetics []phonetics `json:"phonetics,omitempty"`
	Meanings  []meaning   `json:"meanings,omitempty"`
}

type phonetics struct {
	Text string `json:"text,omitempty"`
}

type meaning struct {
	PartOfSpeech string       `json:"partOfSpeech,omitempty"`
	Definitions  []definition `json:"definitions,omitempty"`
	Synonyms     []string     `json:"synonyms,omitempty"`
	Antonyms     []string     `json:"antonyms,omitempty"`
}

type definition struct {
	Definition string   `json:"definition,omitempty"`
	Synonyms   []string `json:"synonyms,omitempty"`
	Antonyms   []string `json:"antonyms,omitempty"`
	Example    string   `json:"example,omitempty"`
}

func (w word) render() {
	res := titleMargin.Render(titleStyle.Render(w.Word) + w.renderPhonetic())
	res += w.renderMeanings()
	fmt.Println(res + "\n")
}

// the api is somewhat inconsistent for phonems so we have to search for the phonetic of a word, that is if it exists
func (w word) renderPhonetic() string {
	if w.Phonetic != "" {
		return " - " + phoneticStyle.Render(w.Phonetic)
	}
	for _, phonetic := range w.Phonetics {
		if phonetic.Text != "" {
			return " - " + phoneticStyle.Render(phonetic.Text)
		}
	}
	return ""
}

func (w word) renderMeanings() (s string) {
	for _, meaning := range w.Meanings {
		s += "\n" + posStyle.Render(meaning.PartOfSpeech)
		s += meaning.renderMeaningSynonymAndAntonym()
		for i, definition := range meaning.Definitions {
			n := fmt.Sprintf("%d. ", i+1)
			s += "\n" + textStyle.Render(n+wordwrap.String(definition.Definition, 70))
			s += definition.renderExample()
		}
	}
	return
}

func (m meaning) renderMeaningSynonymAndAntonym() (s string) {
	if !full && !related {
		return
	}

	const MAX = 4
	if len(m.Synonyms) > MAX {
		s += " " + synonymStyle.Render(strings.Join(m.Synonyms[:MAX], ", "))
	} else {
		s += " " + synonymStyle.Render(strings.Join(m.Synonyms, ", "))
	}

	if len(m.Antonyms) > MAX {
		s += " " + antonymStyle.Render(strings.Join(m.Antonyms[:MAX], ", "))
	} else {
		s += " " + antonymStyle.Render(strings.Join(m.Antonyms, ", "))
	}

	return
}

func (d definition) renderExample() (s string) {
	if !full && !examples {
		return
	}

	if d.Example != "" {
		s = exampleStyle.Render("\n\"" + d.Example + "\"")
	}

	return
}
