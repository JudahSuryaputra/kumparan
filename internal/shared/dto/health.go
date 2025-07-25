package dto

type HealthCheckResponse struct {
	Status      string `json:"status"`
	RedisStatus string `json:"redis_status"`
	DbStatus    string `json:"db_status"`
}
