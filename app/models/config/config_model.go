package config

import (
	"liu/app/models"
	"liu/pkg/database"
)

type Config struct {
	models.BaseModel

	Title     string `json:"title,omitempty"`
	Name      string `json:"name,omitempty"`
	Icp       string `json:"icp,omitempty"`
	Copyright string `json:"copyright,omitempty"`

	models.TimestampsField
}

func (config *Config) Create() {
	database.DB.Create(&config)
}

func (config *Config) Save() (rowsAffected int64) {
	result := database.DB.Save(&config)
	return result.RowsAffected
}

func (config *Config) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&config)
	return result.RowsAffected
}
