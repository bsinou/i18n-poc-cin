package main

import (
	"context"

	"github.com/bsinou/i18n-poc-cin/goui"
	"github.com/bsinou/i18n-poc-cin/goui2"
)

func main() {
	goui.LogTranslatedMessage(context.Background())
	goui2.Log2ndTranslatedMessage(context.Background())
}
