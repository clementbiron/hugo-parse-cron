package functions

import (
	"github.com/lnquy/cron"
)

func ParseCron(expression string, locale string) string {
	cronLocale, err := cron.ParseLocale(locale)
	if err != nil {
		return expression
	}

	exprDesc, err := cron.NewDescriptor(
		cron.SetLocales(cronLocale),
	)
	if err != nil {
		return expression
	}

	desc, err := exprDesc.ToDescription(expression, cronLocale)
	if err != nil {
		return expression
	}

	return desc
}

func Init() map[string]interface{} {
	m := make(map[string]interface{})
	m["parseCron"] = ParseCron
	return m
}
