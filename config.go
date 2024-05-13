package config

import (
	"github.com/creamsensation/filesystem"
	
	"github.com/creamsensation/mailer"
	"github.com/creamsensation/quirk"
)

type Config struct {
	App          App
	Cache        Cache
	Database     map[string]*quirk.DB
	Export       Export
	Filesystem   filesystem.Config
	Localization Localization
	Parser       Parser
	Router       Router
	Security     Security
	Smtp         mailer.Config
}
