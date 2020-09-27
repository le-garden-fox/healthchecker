package routes

// HealthCheckResponse base response
type HealthCheckResponse struct {
	Alive        bool
	Host         string
	Port         string
	ErrorMessage string
}
