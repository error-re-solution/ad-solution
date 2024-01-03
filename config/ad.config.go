package config

import "github.com/spf13/viper"

type LDAP struct {
	Address      string `mapstructure:"ldapAddress"`
	BindDN       string `mapstructure:"bindDN"`
	BindPassword string `mapstructure:"bindPassword"`
	BaseDN       string `mapstructure:"baseDN"`
	Domain       string `mapstructure:"domain"`
}

func LoadLDAPConfig(path string) (config LDAP, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("dev")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		panic(err)
	}
	return
}
