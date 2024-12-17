package config

import (
	"os"
	"path/filepath"

	"github.com/dagu-org/dagu/internal/build"
	"github.com/dagu-org/dagu/internal/util"
)

type resolver struct {
	configDir       string
	dagsDir         string
	suspendFlagsDir string
	dataDir         string
	logsDir         string
	adminLogsDir    string
	baseConfigFile  string
}

type XDGConfig struct {
	DataHome   string
	ConfigHome string
}

func newResolver(appHomeEnv, legacyPath string, xdg XDGConfig) resolver {
	var (
		r           resolver
		useXDGRules = true
	)

	// For backward compatibility.
	// If the environment variable is set, use it.
	// Use the legacy ~/.<app name> directory if it exists.
	if v := os.Getenv(appHomeEnv); v != "" {
		r.configDir = v
		useXDGRules = false
	} else if util.FileExists(legacyPath) {
		r.configDir = legacyPath
		useXDGRules = false
	} else {
		r.configDir = filepath.Join(xdg.ConfigHome, build.Slug)
	}

	if useXDGRules {
		setXDGPaths(&r, xdg)
	} else {
		setLegacyPaths(&r)
	}

	return r
}

func setXDGPaths(r *resolver, xdg XDGConfig) {
	r.dataDir = filepath.Join(xdg.DataHome, build.Slug, "history")
	r.logsDir = filepath.Join(xdg.DataHome, build.Slug, "logs")
	r.baseConfigFile = filepath.Join(xdg.ConfigHome, build.Slug, "base.yaml")
	r.adminLogsDir = filepath.Join(xdg.DataHome, build.Slug, "logs", "admin")
	r.suspendFlagsDir = filepath.Join(xdg.DataHome, build.Slug, "suspend")
	r.dagsDir = filepath.Join(xdg.ConfigHome, build.Slug, "dags")
}

func setLegacyPaths(r *resolver) {
	r.dataDir = filepath.Join(r.configDir, "data")
	r.logsDir = filepath.Join(r.configDir, "logs")
	r.baseConfigFile = filepath.Join(r.configDir, "base.yaml")
	r.adminLogsDir = filepath.Join(r.configDir, "logs", "admin")
	r.suspendFlagsDir = filepath.Join(r.configDir, "suspend")
	r.dagsDir = filepath.Join(r.configDir, "dags")
}
