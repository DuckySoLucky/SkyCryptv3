package utility

import (
	"regexp"
	"skycrypt/src/constants"
	"strconv"
	"strings"
	"time"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var colorCodeRegex = regexp.MustCompile("§[0-9a-fk-or]")

func GetRawLore(text string) string {
	return colorCodeRegex.ReplaceAllString(text, "")
}

var nonAsciiRegex = regexp.MustCompile(`[^\x00-\x7F]`)

func RemoveNonAscii(text string) string {
	return nonAsciiRegex.ReplaceAllString(text, "")
}

func Contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func GetLastValue(m map[int]int) int {
	maxKey := 0
	for key := range m {
		if key > maxKey {
			maxKey = key
		}
	}
	return m[maxKey]
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func TitleCase(s string) string {
	return cases.Title(language.English).String(s)
}

func ParseInt(n string) (int, error) {
	i, err := strconv.Atoi(n)
	if err != nil {
		return 0, err
	}
	return i, nil
}

func RarityNameToInt(rarity string) int {
	for i, r := range constants.RARITIES {
		if strings.EqualFold(r, rarity) {
			return i
		}
	}
	return 0
}

func FormatNumber(n any) string {
	var value float64
	switch v := n.(type) {
	case int:
		value = float64(v)
	case float64:
		value = v
	case float32:
		value = float64(v)
	default:
		return "0"
	}

	abs := value
	if abs < 0 {
		abs = -abs
	}

	var suffix string
	var divisor float64

	switch {
	case abs >= 1e9:
		suffix = "B"
		divisor = 1e9
	case abs >= 1e6:
		suffix = "M"
		divisor = 1e6
	case abs >= 1e3:
		suffix = "K"
		divisor = 1e3
	default:
		if value == float64(int(value)) {
			return strconv.Itoa(int(value))
		}
		return strconv.FormatFloat(value, 'f', -1, 64)
	}

	result := value / divisor
	if result == float64(int(result)) {
		return strconv.Itoa(int(result)) + suffix
	}
	return strconv.FormatFloat(result, 'f', 1, 64) + suffix
}

func ParseTimestamp(timestamp string) int {
	t, err := time.Parse("1/2/06 3:04 PM", timestamp)
	if err != nil {
		return 0
	}

	return int(t.Unix())
}

func Every[T any](slice []T, predicate func(T) bool) bool {
	for _, item := range slice {
		if !predicate(item) {
			return false
		}
	}
	return true
}

func IndexOf(slice []string, item string) int {
	for i, v := range slice {
		if v == item {
			return i
		}
	}

	return -1
}
