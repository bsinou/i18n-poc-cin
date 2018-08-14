package main

import (
	"context"

	"github.com/bsinou/i18n-poc-cin/goui"
)

func main() {
	goui.LogTranslatedMessage(context.Background())
}
