package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type ChromeVerifyConfig struct {
	ChromeDriverPath string `yaml:"chromeDriverPath"`
	Port             int    `yaml:"port"`
	URL              string `yaml:"url"`
	ChromeBinaryPath string `yaml:"chromeBinaryPath"`
}

type Config struct {
	Wechat struct {
		AppID           string `yaml:"appid"`
		AppSecret       string `yaml:"appsecret"`
		Code2SessionURL string `yaml:"code2sessionURL"`
	} `yaml:"wechat"`
	ChromeVerify struct {
		ChromeDriverPath string `yaml:"chromeDriverPath"`
		Port             int    `yaml:"port"`
		DriverURL        string `yaml:"driverURL"`
		URL              string `yaml:"url"`
		ChromeBinaryPath string `yaml:"chromeBinaryPath"`
		LogFilePath      string `yaml:"logFilePath"`
		VerifySign       string `yaml:"verifySign"`
	} `yaml:"chromeVerify"`
	RSA struct {
		PrivateKey string `yaml:"privateKey"`
		PublicKey  string `yaml:"publicKey"`
	} `yaml:"RSA"`
	Database struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		DBName   string `yaml:"dbname"`
	} `yaml:"database"`
	Server struct {
		Addr string `yaml:"addr"`
	} `yaml:"server"`
	Task struct {
		ReminderTime string `yaml:"reminderTime"`
	} `yaml:"task"`
	Redis struct {
		Host         string `yaml:"host"`
		Port         int    `yaml:"port"`
		Password     string `yaml:"password"`
		DB           int    `yaml:"db"`
		PoolSize     int    `yaml:"poolSize"`
		MinIdleConns int    `yaml:"MinIdleConns"`
		MaxRetries   int    `yaml:"MaxRetries"`
		DialTimeout  int    `yaml:"DialTimeout"`
		ReadTimeout  int    `yaml:"ReadTimeout"`
		WriteTimeout int    `yaml:"WriteTimeout"`
		Prefix       string `yaml:"prefix"`
	} `yaml:"redis"`
	MachineID int `yaml:"machineID"`
	Token     struct {
		SecretKey  string `yaml:"secretKey"`
		ExpireTime int    `yaml:"expireTime"`
	} `yaml:"token"`
}

func NewConfig() *Config {
	yamlData, err := os.ReadFile("config.yaml")
	if err != nil {
		return nil
	}
	var cfg Config
	if err := yaml.Unmarshal(yamlData, &cfg); err != nil {
		return nil
	}
	return &cfg
}
