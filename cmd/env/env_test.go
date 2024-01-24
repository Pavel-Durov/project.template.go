package env_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"p3ld3v.dev/template/cmd/env"
)

func TestDependencies(t *testing.T) {
	_ = os.Setenv("HOST", "1.2.3.4")
	_ = os.Setenv("PORT", "1234")
	_ = os.Setenv("LOG_LEVEL", "info")
	_ = os.Setenv("DB_CONNECTION_STRING", "postgres://x:y@z:5432")

	cfg := env.LoadConfig()

	assert.Equal(t, cfg.Port, "1234")
	assert.Equal(t, cfg.Host, "1.2.3.4")
	assert.Equal(t, cfg.LogLevel, "info")
	assert.Equal(t, cfg.DbConnection, "postgres://x:y@z:5432")

	_ = os.Unsetenv("HOST")
	_ = os.Unsetenv("PORT")
}

func TestDefaultDependencies(t *testing.T) {

	cfg := env.LoadConfig()

	assert.Equal(t, cfg.Port, "8080")
	assert.Equal(t, cfg.Host, "0.0.0.0")
}
