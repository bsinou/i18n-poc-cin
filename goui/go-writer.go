package goui

import (
	"context"
	"fmt"

	"github.com/nicksnyder/go-i18n/i18n"

	"github.com/bsinou/i18n-poc-cin/goui/lang"
	"github.com/bsinou/i18n-poc-cin/utils"
)

type holder struct{ CreationCount int }

// LogTranslatedMessage simply prints some messages in the various supported languages to standrad out.
func LogTranslatedMessage(ctx context.Context) error {

	locales := []string{utils.GetDefaultLanguage(), "fr", "de", "es"}

	for _, k := range locales {
		T := lang.Bundle().GetTranslationFunc(k)
		doPrint(k, T)
	}

	return nil
}

type deletionHolder struct {
	Count int
	Name  string
}

func newDeletionHolder(a int, b string) deletionHolder {
	return deletionHolder{
		Count: a,
		Name:  b,
	}
}

func doPrint(l string, t i18n.TranslateFunc) {

	fmt.Println("### ", l, ": ", t("Goui.Page.Title"), "-", t("Goui.Page.OtherTitle"))
	fmt.Println("Msg #1:", t("Goui.Page.Added", 1))
	fmt.Println("Msg #2:", t("Goui.Page.Added", 8))
	fmt.Println("Msg #3:", t("Goui.Page.Deleted", newDeletionHolder(1, "Jack")))
	fmt.Println("Msg #4:", t("Goui.Page.Deleted", newDeletionHolder(6, "Fanny")))

	fmt.Println()
}
