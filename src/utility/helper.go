package utility

import (
	"encoding/base64"
	"fmt"
	"math"
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
	if strings.Contains(s, "_") || strings.Contains(s, "-") {
		parts := strings.FieldsFunc(s, func(r rune) bool {
			return r == '_' || r == '-'
		})
		for i, part := range parts {
			parts[i] = cases.Title(language.English).String(part)
		}
		return strings.Join(parts, " ")
	}

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

func GetSkinHash(base64String string) string {
	if base64String == "" {
		return ""
	}

	data, err := base64.StdEncoding.DecodeString(base64String)
	if err != nil {
		return ""
	}

	parts := strings.Split(string(data), "/")
	if len(parts) == 0 {
		return ""
	}

	return parts[len(parts)-1]
}

func Round(value float64, precision int) float64 {
	if precision < 0 {
		return value
	}
	pow := math.Pow(10, float64(precision))
	return math.Round(value*pow) / pow
}

func ReplaceVariables(template string, variables map[string]float64) string {
	re := regexp.MustCompile(`\{(\w+)\}`)

	return re.ReplaceAllStringFunc(template, func(match string) string {
		name := strings.Trim(match, "{}")

		value, exists := variables[name]
		if !exists {
			return match
		}

		// fmt.Printf("Replacing variable %s with value %.2f\n", name, value)
		if _, err := strconv.ParseFloat(name, 64); err != nil {
			if intValue, err := strconv.Atoi(fmt.Sprintf("%.0f", value)); err == nil && intValue > 0 {
				return "+" + fmt.Sprintf("%.0f", value)
			}
		}

		return fmt.Sprintf("%.0f", value)
	})
}

func CompareInts(a, b int) int {
	if a < b {
		return -1
	} else if a > b {
		return 1
	}
	return 0
}

func CompareStrings(a, b string) int {
	if a < b {
		return -1
	} else if a > b {
		return 1
	}
	return 0
}

func CompareBooleans(a, b bool) int {
	if a == b {
		return 0
	} else if a && !b {
		return 1
	}
	return -1
}

func Filter[T any](slice []T, predicate func(T) bool) []T {
	var result []T
	for _, item := range slice {
		if predicate(item) {
			result = append(result, item)
		}
	}
	return result
}

func SortBy[T any](slice []T, compare func(T, T) int) []T {
	if len(slice) < 2 {
		return slice
	}

	for i := 0; i < len(slice)-1; i++ {
		for j := 0; j < len(slice)-i-1; j++ {
			if compare(slice[j], slice[j+1]) > 0 {
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}

	return slice
}

func Sum(slice []float64) float64 {
	var total float64
	for _, value := range slice {
		total += value
	}
	return total
}

