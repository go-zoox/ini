package ini

import (
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
			Port     int64  `ini:"port"`
			User     string `ini:"user"`
			Password string `ini:"password"`
			Database string `ini:"database"`
		} `ini:"mysql"`
		Redis struct {
			IP   string `ini:"ip"`
			Port int64  `ini:"port"`
		} `ini:"redis"`
	}

	var config Config
	if err := Unmarshal([]byte(iniText), &config); err != nil {
		t.Error(err)
	} else {
		// j, _ := json.MarshalIndent(config, "", "  ")
		// t.Log(string(j))

		if config.AppName != "go-zoox web" {
			t.Error("app_name != go-zoox web")
		}

		if config.LogLevel != "DEBUG" {
			t.Error("log_level != DEBUG")
		}

		if config.Mysql.IP != "127.0.0.1" {
			t.Error("mysql.ip != 127.0.0.1")
		}

		if config.Mysql.Port != 3306 {
			t.Error("mysql.port != 3306")
		}

		if config.Mysql.User != "zero" {
			t.Error("mysql.user != zero")
		}

		if config.Mysql.Password != "123456" {
			t.Error("mysql.password != 123456")
		}

		if config.Mysql.Database != "go-zoox" {
			t.Error("mysql.database != go-zoox")
		}

		if config.Redis.IP != "127.0.0.1" {
			t.Error("redis.ip != 127.0.0.1")
		}

		if config.Redis.Port != 6379 {
			t.Error("redis.port != 6379")
		}
	}
}

func TestMarshal(t *testing.T) {
	type Config struct {
		AppName  string `ini:"app_name"`
		LogLevel string `ini:"log_level"`
		Mysql    struct {
			IP       string `ini:"ip"`
			Port     int64  `ini:"port"`
			User     string `ini:"user"`
			Password string `ini:"password"`
			Database string `ini:"database"`
		} `ini:"mysql"`
		Redis struct {
			IP   string `ini:"ip"`
			Port int64  `ini:"port"`
		} `ini:"redis"`
	}

	config := &Config{
		AppName:  "go-zoox web",
		LogLevel: "DEBUG",
		Mysql: struct {
			IP       string `ini:"ip"`
			Port     int64  `ini:"port"`
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
			Port int64  `ini:"port"`
		}{
			IP:   "127.0.0.1",
			Port: 6379,
		},
	}

	if v, err := Marshal(config); err != nil {
		t.Error(err)
	} else {
		// t.Log(string(v))
		// fmt.Println(string(v))
		if string(v) != `appname = go-zoox web
loglevel = DEBUG

[mysql]
database = go-zoox
ip = 127.0.0.1
password = 123456
port = 3306
user = zero

[redis]
ip = 127.0.0.1
port = 6379` {
			t.Error("unexpected ini text")
		}
	}
}
