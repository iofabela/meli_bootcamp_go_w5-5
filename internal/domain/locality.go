package domain

type Locality struct {
	ID           int    `json:"id"`
	LocalityName string `json:"locality_name"`
	ProvinceName string `json:"province_name"`
	CountryName  string `json:"country_name"`
}

type ReportSeller struct {
	LocalityID   int    `json:"locality_id"`
	LocalityName string `json:"locality_name"`
	SellersCount int    `json:"sellers_count"`
}

type LocalityNew struct {
	ID           int    `json:"id"`
	LocalityName string `json:"locality_name"`
	ProvinceID   int    `json:"province_id"`
}
type LocalityCarries struct {
	LocalityID   int    `json:"locality_id"`
	LocalityName string `json:"locality_name"`
	CarriesCount int    `json:"carries_count"`
}
