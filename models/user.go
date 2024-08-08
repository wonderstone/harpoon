package models

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dbinfo Database

type Database struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"pass"`
	Name     string `yaml:"name"`
}

var DB *gorm.DB

type User struct {
	gorm.Model
	Username string `gorm:"uniqueIndex"`
	Password string
}

func ConnectDatabase() {
	

	dsn := dbinfo.User + ":" + dbinfo.Password + "@tcp(" + dbinfo.Host + ":" + dbinfo.Port + ")/" + dbinfo.Name + "?charset=utf8mb4&parseTime=True&loc=Local"

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	database.AutoMigrate(&User{})

	DB = database
}

func init() {
	// read the config file
	yamlFile, err := os.ReadFile("./dbscript/info.yaml")
	if err != nil {
		log.Fatalf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, &dbinfo)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
}