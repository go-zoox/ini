# ini - Simple INI Format Config Parse with Marshal and Unmarshal

[![PkgGoDev](https://pkg.go.dev/badge/github.com/go-zoox/ini)](https://pkg.go.dev/github.com/go-zoox/ini)
[![Build Status](https://github.com/go-zoox/ini/actions/workflows/ci.yml/badge.svg?branch=master)](https://github.com/go-zoox/ini/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-zoox/ini)](https://goreportcard.com/report/github.com/go-zoox/ini)
[![Coverage Status](https://coveralls.io/repos/github/go-zoox/ini/badge.svg?branch=master)](https://coveralls.io/github/go-zoox/ini?branch=master)
[![GitHub issues](https://img.shields.io/github/issues/go-zoox/ini.svg)](https://github.com/go-zoox/ini/issues)
[![Release](https://img.shields.io/github/tag/go-zoox/ini.svg?label=Release)](https://github.com/go-zoox/ini/tags)

## Installation
To install the package, run:
```bash
go get github.com/go-zoox/ini
```

## Getting Started

```go
func main() {
  	iniText := `app_name = go-zoox web

# possible values: DEBUG, INFO, WARNING, ERROR, FATAL
log_level = DEBUG

[mysql]
ip = 127.0.0.1
port = 3306
user = zero
password = 123456
database = go-zoox

[redis]
ip = 127.0.0.1
port = 6379`

	type Config struct {
		AppName  string `ini:"app_name"`
		LogLevel string `ini:"log_level"`
		Mysql    struct {
			IP       string `ini:"ip"`
			Port     int    `ini:"port"`
			User     string `ini:"user"`
			Password string `ini:"password"`
			Database string `ini:"database"`
		} `ini:"mysql"`
		Redis struct {
			IP   string `ini:"ip"`
			Port int    `ini:"port"`
		} `ini:"redis"`
	}

	var config Config
	if err := Unmarshal([]byte(iniText), &config); err != nil {
		t.Error(err)
	} else {
		j, _ := json.MarshalIndent(config, "", "  ")
		t.Log(string(j))
	}
}
```

## License
GoZoox is released under the [MIT License](./LICENSE).
