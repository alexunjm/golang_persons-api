package config

import (
	"fmt"
)

// LoadAppEnv creates ApplicationConfig instance
func LoadAppEnv() {

	ApplicationConfig = app{
		appEnv: GetApplicationConfig(),
		sqlEnv: GetSQLDataConnection(),
	}
}

type app struct {
	appEnv APPConfigContract
	sqlEnv SQLDataConnectionContract
}

// ApplicationConfig handles all infrastructure default config
var ApplicationConfig app

func (config *app) GetAPIPort() string {
	return fmt.Sprintf(":%v", config.appEnv.Port())
}

func (config *app) GetSQLConfig() SQLDataConnectionContract {
	return config.sqlEnv
}
