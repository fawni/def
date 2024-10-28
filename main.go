package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/hbollon/go-edlib"
	"github.com/spf13/cobra"
)

var (
	full     bool
	related  bool
	examples bool

	cmd = &cobra.Command{
		Use:   "def <word>",
		Short: "def is a command-line dictionary client",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return define(args)
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
		os.Exit(1)
	}
}

func define(args []string) error {
	res, err := http.Get("https://api.dictionaryapi.dev/api/v2/entries/en/" + args[0])
	if err != nil {
		return err
	}

	switch res.StatusCode {
	case 200:
		defer res.Body.Close()
		body, err := io.ReadAll(res.Body)
		if err != nil {
			return err
		}
		var words []word
		if err := json.Unmarshal(body, &words); err != nil {
			return err
		}
		for _, word := range words {
			word.render()
		}
	case 404:
		fmt.Println(errorStyle.Render("no definitions found."))
		suggest(args[0])
		os.Exit(1)
	default:
		fmt.Println(errorStyle.Render("failed with status " + res.Status))
		os.Exit(1)
	}

	return nil
}

func suggest(word string) {
	// TODO: cache this?
	res, err := http.Get("https://raw.githubusercontent.com/meetDeveloper/freeDictionaryAPI/refs/heads/master/meta/wordList/english.txt")
	if err != nil {
		return
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return
	}
	var words []string
	for _, word := range strings.Split(string(body), "\n") {
		words = append(words, strings.TrimSpace(word))
	}

	matches, err := edlib.FuzzySearchSet(word, words, 3, edlib.JaroWinkler)
	if err != nil {
		return
	}

	var mstr string
	for i, match := range matches {
		m := "'" + match + "'"
		if i == 0 {
			mstr = m
			continue
		}
		if i == len(matches)-1 {
			mstr += " or " + m
			break
		}
		mstr += ", " + m
	}

	fmt.Println(errorStyle.Render("did you mean: " + mstr + "?"))
}
