package eliab

// Get doc ...
var Get Configuration

// Configuration doc ...
type Configuration struct {
	Mode                string
	General             map[string]General `toml:"general"`
	DataBase            DataBase
	OnePushNotification OnePushNotification `toml:"onePushNotification"`
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

// OnePushNotification doc ...
type OnePushNotification struct {
	URL string
}
