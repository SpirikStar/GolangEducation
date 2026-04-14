package main

import (
	"context"
	"fmt"
	"strings"

	"gopkg.gilang.dev/translator/v2/googletranslate"
	"gopkg.gilang.dev/translator/v2/params"
)

func main() {

	text := strings.TrimSpace("hello world!")

	translated, err := translateEnToRu(text)
	if err != nil {
		fmt.Printf("Ошибка перевода: %v\n\n", err)

	}

	fmt.Printf("Русский: %s\n\n", translated)
}

func translateEnToRu(text string) (string, error) {
	ctx := context.Background()

	client := googletranslate.New()

	result, err := client.Translate(ctx, text, params.ENGLISH, params.RUSSIAN)
	if err != nil {
		return "", err
	}

	return result.Text, nil
}
