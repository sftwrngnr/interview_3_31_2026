package main

import (
	"interview_3_31_2026/config"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Simple test functions to verify config read functionality

func TestConfig(t *testing.T) {
	cfg, err := ReadConfig("config.yaml")
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, *cfg, config.Config{Port: 8080, BindAddress: "127.0.0.1", EnableTelemetry: true})
}

func TestConfigFaile(t *testing.T) {
	_, err := ReadConfig("sludge.yaml")
	assert.NotNil(t, err)
}
