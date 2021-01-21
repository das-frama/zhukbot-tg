package zhuk

import (
	"fmt"
	"strings"
)

func AnalyzeText(text string) (string, error) {
	var response string
	str := strings.TrimSpace(strings.Trim(strings.ToLower(text), "."))

	if strings.HasPrefix(str, "найти себе нового жука") {
		response = findZhuk(strings.TrimLeft(str, "найти себе нового жука"))
	} else if strings.HasPrefix(str, "найти себе жука") {
		response = findZhuk(strings.TrimLeft(str, "найти себе жука"))
	} else if strings.HasPrefix(str, "найти жука") {
		response = findZhuk(strings.TrimLeft(str, "найти жука"))
	}

	return response, nil
}

func findZhuk(location string) string {
	t, ok := Search(location)
	if !ok {
		return "Вы не смогли найти жука."
	}

	return fmt.Sprintf("Вы нашли отличного, прекрасного жука %d.", t)
}
