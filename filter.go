package toxicfilter
import (
	"bufio"
	"os"
	"regexp"
	"strings"
)

var (
	nonAlphanumericRegex = regexp.MustCompile(`[^a-zA-Z0-9 ]+`)
	badWords             = make(map[string]struct{})
)

// LoadBadWords loads bad words list from a text file into memory.
func LoadBadWords(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := strings.TrimSpace(scanner.Text())
		if word != "" {
			badWords[strings.ToLower(word)] = struct{}{}
		}
	}
	return scanner.Err()
}

// cleanString removes punctuation and special characters.
func cleanString(str string) string {
	return nonAlphanumericRegex.ReplaceAllString(str, "")
}

// Check returns true if the text contains a banned word.
func Check(text string) bool {
	clearText := cleanString(text)
	words := strings.Fields(strings.ToLower(clearText))

	for _, w := range words {
		if _, exists := badWords[w]; exists {
			return true
		}
	}
	return false
}

// Censor replaces banned words with asterisks (****)
func Censor(text string) string {
	originalWords := strings.Fields(text)
	lowered := strings.Fields(strings.ToLower(cleanString(text)))

	for i, w := range lowered {
		if _, exists := badWords[w]; exists {
			originalWords[i] = strings.Repeat("*", len(originalWords[i]))
		}
	}
	return strings.Join(originalWords, " ")
}
