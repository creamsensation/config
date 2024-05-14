package config

import (
	"github.com/creamsensation/auth"
	"github.com/creamsensation/csrf"
	"github.com/creamsensation/firewall"
)

type Security struct {
	Auth      auth.Config
	Csrf      csrf.Csrf
	Firewalls []firewall.Firewall
}
