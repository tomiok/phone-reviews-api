package models

// Smartphone model structure for smartphone
type Smartphone struct {
	Id            int64 // similar to a long in Java - ID autoincremental in SQL
	Name          string
	Price         int
	CountryOrigin string
	OS            string
}

// CreateSmartphoneCMD
type CreateSmartphoneCMD struct {
	Name          string `json:"name"`
	Price         int    `json:"price"`
	CountryOrigin string `json:"country_origin"`
	OS            string `json:"os"`
}
