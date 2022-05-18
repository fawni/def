package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/spf13/cobra"
)

var (
	full     bool
	related  bool
	examples bool

	cmd = &cobra.Command{
		Use:   "def []",
		Short: "def is a command-line dictionary client",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return execute(args)
		},
	}
)

func init() {
	cmd.PersistentFlags().BoolVarP(&full, "full", "f", false, "displays synonyms, antonyms and examples if found")
	cmd.PersistentFlags().BoolVarP(&related, "related", "r", false, "displays synonyms and antonyms if found")
	cmd.PersistentFlags().BoolVarP(&examples, "examples", "e", false, "displays examples if found")
}

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Println(errorStyle.Render(err.Error()))
	}
}

func execute(args []string) error {
	resp, err := http.Get("https://api.dictionaryapi.dev/api/v2/entries/en/" + args[0])
	if err != nil {
		return err
	}

	switch resp.StatusCode {
	case 200:
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		var words []Word
		if err := json.Unmarshal(body, &words); err != nil {
			return err
		}
		words[0].render()
	case 404:
		fmt.Println(errorStyle.Render("defeinitions not found"))
	default:
		fmt.Println(errorStyle.Render("failed with status: " + resp.Status))
	}

	return nil
}

func (w Word) render() {
	res := titleMargin.Render(titleStyle.Render(w.Word) + w.getPhonetic())
	res += w.getMeanings()
	fmt.Println(res + "\n")
}

// the api is somewhat inconsistent for phonems so we have to
//  search for the phonetic of a word, that is if it exists
func (w Word) getPhonetic() string {
	if w.Phonetic != "" {
		return " - " + phoneticStyle.Render(w.Phonetic)
	} else {
		for _, phonetic := range w.Phonetics {
			if phonetic.Text != "" {
				return " - " + phoneticStyle.Render(phonetic.Text)
			}
		}
	}
	return ""
}

func (w Word) getMeanings() (s string) {
	for _, meaning := range w.Meanings {
		s += "\n" + posStyle.Render(meaning.PartOfSpeech)
		s += meaning.getMeaningSynonymAndAnyonym()
		for i, definition := range meaning.Definitions {
			n := fmt.Sprintf("%d. ", i+1)
			s += "\n" + textStyle.Render(n+definition.Definition)
			s += definition.getExample()
		}

	}
	return
}

func (m Meaning) getMeaningSynonymAndAnyonym() (s string) {
	if !full && !related {
		return
	}

	const max = 4
	switch {
	case len(m.Synonyms) > max:
		s += " " + synonymStyle.Render(strings.Join(m.Synonyms[:max], ", "))
	default:
		s += " " + synonymStyle.Render(strings.Join(m.Synonyms, ", "))
	}

	switch {
	case len(m.Antonyms) > max:
		s += " " + antonymStyle.Render(strings.Join(m.Antonyms[:max], ", "))
	default:
		s += " " + antonymStyle.Render(strings.Join(m.Antonyms, ", "))
	}

	return
}

func (d Definition) getExample() (s string) {
	if !full && !examples {
		return
	}

	if d.Example != "" {
		s = exampleStyle.Render("\n\"" + d.Example + "\"")
	}

	return
}
