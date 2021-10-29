package spam_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/mobocratic/spam"
	"github.com/stretchr/testify/assert"
)

func ExampleSplitWords() {
	words := spam.SplitWords("to be, or not to be.")
	fmt.Println(strings.Join(words, "|"))
	// Output: to|be|or|not|to|be
}

type model map[string]float64

func (m model) Prob(word string) float64 {
	if prob, exists := m[word]; exists {
		return prob
	}
	return 0
}

func TestJudge(t *testing.T) {
	ps := model{"shit": 0.3, "click": 0.7, "fuck": 1.0}
	ph := model{"you": 0.9, "click": 0.2}
	result := spam.Judge("click for more information", ps, ph, 0.6)
	assert.Equal(t, spam.SPAM, result)
}
