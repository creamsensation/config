package config

import (
	"github.com/creamsensation/translator"
	
	"github.com/creamsensation/form"
)

type Localization struct {
	Enabled    bool
	Path       bool
	Languages  []Language
	Translator translator.Translator
	Form       form.Messages
}

type Language struct {
	Main bool
	Code string
}
