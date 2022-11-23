package config

import (
	"github.com/go-kit/log"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)


func TestLoadFile(t *testing.T) {
	logger := log.NewLogfmtLogger(os.Stderr)
	cfg, err := LoadFile("../testdata/config_test.yaml", logger)
	if err != nil {
		assert.Error(t, err)
	}
	assert.Equal(t, "localhost", cfg.ServerConfig.Host)
}
