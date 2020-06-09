package app

import (
	"strings"

	"github.com/tsuru/tsuru/app/bind"
)

const suppressedEnv = "****"

var sensitiveEnvs = []string{
	"password",
	"secret",
	"passwd",
	"api_key",
	"apikey",
	"token",
	"credential",
	"pwd",
	"endpoint",
}

func isSensitiveName(name string) bool {
	lowerStr := strings.ToLower(name)

	for _, env := range sensitiveEnvs {
		if strings.Contains(lowerStr, env) {
			return true
		}
	}

	return false
}

func (a *App) SuppressSensitiveEnvs() {
	newEnv := map[string]bind.EnvVar{}
	for key, env := range a.Env {
		if !env.Public || isSensitiveName(env.Name) {
			env.Value = suppressedEnv
		}
		newEnv[key] = env
	}
	a.Env = newEnv

	for i, serviceEnv := range a.ServiceEnvs {
		if !serviceEnv.EnvVar.Public || isSensitiveName(serviceEnv.EnvVar.Name) {
			a.ServiceEnvs[i].Value = suppressedEnv
		}
	}
}
