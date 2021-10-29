package spam

import "regexp"

type Result bool

const (
	SPAM Result = true
	HAM  Result = false

	Strict float64 = 0.6
	Loose  float64 = 0.8

	eps = 1e-6
)

var wordRegex = regexp.MustCompile(`\W+`)

type Prober interface {
	Prob(word string) float64
}

// Judge ...
func Judge(message string, ps, ph Prober, threshold float64) Result {
	coef := 1.0
	for _, w := range SplitWords(message) {
		pspam := ps.Prob(w)
		pham := ph.Prob(w)
		if pham < eps {
			if pspam < eps {
				continue
			}
			pham = 0.001
		}
		coef *= pham / (pspam + pham)
	}
	return Result(coef < threshold)
}

func SplitWords(message string) []string {
	words := []string{}
	for _, w := range wordRegex.Split(message, -1) {
		if w != "" {
			words = append(words, w)
		}
	}
	return words
}
