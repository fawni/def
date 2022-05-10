package main

type Word struct {
	Word      string      `json:"word,omitempty"`
	Phonetic  string      `json:"phonetic,omitempty"`
	Phonetics []Phonetics `json:"phonetics,omitempty"`
	Meanings  []Meaning   `json:"meanings,omitempty"`
}

type Phonetics struct {
	Text string `json:"text,omitempty"`
}

type Meaning struct {
	PartOfSpeech string       `json:"partOfSpeech,omitempty"`
	Definitions  []Definition `json:"definitions,omitempty"`
	Synonyms     []string     `json:"synonyms,omitempty"`
	Antonyms     []string     `json:"antonyms,omitempty"`
}

type Definition struct {
	Definition string   `json:"definition,omitempty"`
	Synonyms   []string `json:"synonyms,omitempty"`
	Antonyms   []string `json:"antonyms,omitempty"`
	Example    string   `json:"example,omitempty"`
}
