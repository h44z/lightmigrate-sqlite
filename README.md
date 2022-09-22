# LightMigrate - sqlite3 migration driver

[![codecov](https://codecov.io/gh/h44z/lightmigrate-sqlite/branch/master/graph/badge.svg?token=S7E18P04CY)](https://codecov.io/gh/h44z/lightmigrate-sqlite)
[![License: MIT](https://img.shields.io/badge/license-MIT-green.svg)](https://opensource.org/licenses/MIT)
[![GoDoc](https://pkg.go.dev/badge/github.com/h44z/lightmigrate-sqlite/sqlite)](https://pkg.go.dev/github.com/h44z/lightmigrate-sqlite/sqlite)
![GitHub last commit](https://img.shields.io/github/last-commit/h44z/lightmigrate-sqlite)
[![Go Report Card](https://goreportcard.com/badge/github.com/h44z/lightmigrate-sqlite)](https://goreportcard.com/report/github.com/h44z/lightmigrate-sqlite)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/h44z/lightmigrate-sqlite)
![GitHub code size in bytes](https://img.shields.io/github/languages/code-size/h44z/lightmigrate-sqlite)
[![GitHub Release](https://img.shields.io/github/release/h44z/lightmigrate-sqlite.svg)](https://github.com/h44z/lightmigrate-sqlite/releases)

This module is part of the [LightMigrate](https://github.com/h44z/lightmigrate) library.
It provides a migration driver for sqlite3.

## Features
 * Driver work with sqlite3 (uses cgo, [https://github.com/mattn/go-sqlite3](https://github.com/mattn/go-sqlite3)).
 * [Examples](./examples)

## Configuration Options

Configuration options can be passed to the constructor using the `With<Config-Option>` functions.

| Config Value      | Defaults          | Description                                        |
|-------------------|-------------------|----------------------------------------------------|
| `MigrationsTable` | schema_migrations | Name of the migrations table.                      |
| `Logger`          | log.Default()     | The logger instance that should be used.           |
| `VerboseLogging`  | false             | If set to true, more log messages will be printed. |