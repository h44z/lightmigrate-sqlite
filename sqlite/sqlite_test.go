package sqlite

import (
	"log"
	"testing"
)

func TestNewDriver_NoDb(t *testing.T) {
	_, err := NewDriver(nil, "")
	if err == nil {
		t.Fatalf("expected error, got: %v", err)
	}
}

func TestNewDriver_NoClient(t *testing.T) {
	_, err := NewDriver(nil, "db")
	if err == nil {
		t.Fatalf("expected error, got: %v", err)
	}
}

func TestWithLogger(t *testing.T) {
	d := &driver{}

	WithLogger(log.Default())(d)
	if d.logger != log.Default() {
		t.Fatalf("failed to set logger")
	}
}

func TestWithMigrationTable(t *testing.T) {
	d := &driver{cfg: &config{}}

	WithMigrationTable("name")(d)
	if d.cfg.MigrationsTable != "name" {
		t.Fatalf("failed to set migration table name")
	}
}

func TestWithVerboseLogging(t *testing.T) {
	d := &driver{}

	WithVerboseLogging(true)(d)
	if d.verbose != true {
		t.Fatalf("failed to set verbose flag")
	}
}

func Test_driver_Close(t *testing.T) {
	d := &driver{}
	err := d.Close()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}
