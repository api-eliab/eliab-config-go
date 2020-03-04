package eliab

// Get doc ...
var Get Configuration

// Configuration doc ...
type Configuration struct {
	Mode                string
	AppVersion          string             `toml:"appVersion"`
	General             map[string]General `toml:"general"`
	DataBase            DataBase
	OneSignal           OneSignal           `toml:"oneSignal"`
	Services            map[string]Services `toml:"services"`
	AuthenticationToken string              `toml:"authenticationToken"`
}

// General doc ...
type General struct {
	PortServer string
}

// DataBase doc ...
type DataBase struct {
	User     string
	Password string
	Server   string
	DataBase string
	Port     int64
}

// OneSignal doc ...
type OneSignal struct {
	AppID string
	Key   string
}

// Services doc ..
type Services struct {
	URL         string
	Description string
}
