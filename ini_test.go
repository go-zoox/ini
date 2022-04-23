package ini

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestUnmarshal(t *testing.T) {
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

func TestMarshal(t *testing.T) {
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

	config := &Config{
		AppName:  "go-zoox web",
		LogLevel: "DEBUG",
		Mysql: struct {
			IP       string `ini:"ip"`
			Port     int    `ini:"port"`
			User     string `ini:"user"`
			Password string `ini:"password"`
			Database string `ini:"database"`
		}{
			IP:       "127.0.0.1",
			Port:     3306,
			User:     "zero",
			Password: "123456",
			Database: "go-zoox",
		},
		Redis: struct {
			IP   string `ini:"ip"`
			Port int    `ini:"port"`
		}{
			IP:   "127.0.0.1",
			Port: 6379,
		},
	}

	if v, err := Marshal(config); err != nil {
		t.Error(err)
	} else {
		// t.Log(string(v))
		fmt.Println(string(v))
	}
}
