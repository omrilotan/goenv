package goenv

import (
	"testing"
	"time"
)

func TestSetEnvironment(t *testing.T) {
	goenv := NewGoenv("./test_config.yml", "custom", "nil")
	if goenv.GetPort() != "6711" {
		t.Error("port != 6711")
	}
}

func TestSetEnvironmentNotFound(t *testing.T) {
	goenv := NewGoenv("./test_config.yml", "nonexistent", "nil")
	if goenv.GetPort() != "8080" {
		t.Error("port != 8080")
	}
}

func TestSetConfigFile(t *testing.T) {
	goenv := NewGoenv("./test_config2.yml", "config", "nil")
	if goenv.GetPort() != "4388" {
		t.Error("port != 4388")
	}
}

func TestSetConfigFileNotFound(t *testing.T) {
	defer func() { recover() }()

	NewGoenv("nonexistent", "", "nil")
	t.Error("SetConfigFile didn't panic")
}

func TestGet(t *testing.T) {
	goenv := NewGoenv("./test_config.yml", "config", "nil")
	if goenv.Get("custom", "") != "aoeu" {
		t.Error("custom != aoeu")
	}
}

func TestGetInt(t *testing.T) {
	goenv := NewGoenv("./test_config.yml", "config", "nil")
	if goenv.GetInt("number", 1245) != 1234 {
		t.Error("number != 1234")
	}
}

func TestGetDuration(t *testing.T) {
	goenv := NewGoenv("./test_config.yml", "config", "nil")
	if goenv.GetDuration("duration", "30s") != time.Duration(10*time.Second) {
		t.Error("duration != 10s")
	}
}

func TestRequire(t *testing.T) {
	defer func() { recover() }()
	goenv := NewGoenv("./test_config.yml", "config", "nil")
	goenv.Require("dingdong")
	t.Error("Require didn't panic")
}

func TestGetNotFound(t *testing.T) {
	goenv := NewGoenv("./test_config.yml", "config", "nil")
	if goenv.Get("nonexistent", "snth") != "snth" {
		t.Error("nonexistent != snth")
	}
}

func TestGetEnvName(t *testing.T) {
	goenv := NewGoenv("./test_config.yml", "", "nil")
	if goenv.GetEnvName() != "development" {
		t.Error("wrong environment name returned")
	}
}

func TestCount(t *testing.T) {
	goenv := NewGoenv("./test_config.yml", "list", "nil")
	if goenv.Count("entries") != 3 {
		t.Error("list count wrong")
	}

	if goenv.Require("entries[1]") != "b" {
		t.Error("list entry wrong")
	}

	if goenv.Count("nested") != 2 {
		t.Error("nested list count wrong")
	}

	if goenv.Require("nested[1].value") != "v2" {
		t.Error("nested list entry wrong")
	}
}

// TODO: test new functions
