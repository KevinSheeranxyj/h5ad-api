package config

type AppConfig struct {
	App
	Log
	Redis
	Mysql
}

type App struct {
	Debug bool
	Env   string
	Host  string
	Port  string
	Url   string
}
type Log struct {
	Dir    string
	ToFile bool
}
type Redis struct {
	Hostname string
	Password string
	HostPort string
	Database int
}
type Mysql struct {
	Hostname string
	Username string
	Password string
	HostPort string
	Database string
}
