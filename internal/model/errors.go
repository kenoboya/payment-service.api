package model

import "errors"

var (
	ErrNotFoundConfigFile              = errors.New("failed to find config file")
	ErrNotFoundEnvFile                 = errors.New("failed to load environment file")
)
