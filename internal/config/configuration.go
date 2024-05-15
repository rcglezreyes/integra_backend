package config

import (
	"path/filepath"
	"runtime"
	"strconv"

	"github.com/spf13/viper"
)

func ConfigEnv() (error, bool) {
	viper.SetConfigName(CONFIG_FILE_NAME)
	viper.AddConfigPath(RootDir())
	viper.AutomaticEnv()
	viper.SetConfigType(CONFIG_FILE_TYPE)
	if err := viper.ReadInConfig(); err != nil {
		return err, false
	}
	return nil, true
}

func RootDir() string {
	_, b, _, _ := runtime.Caller(0)
	return filepath.Join(filepath.Dir(b), "../..")
}

func DBCredentials() (string, int, string, string, string, error) {
	port, err := strconv.Atoi(viper.GetString(Port))
	if err != nil {
		return "", 0, "", "", "", err
	}
	return viper.GetString(Host), port, viper.GetString(User), viper.GetString(Password), viper.GetString(Dbname), nil
}
