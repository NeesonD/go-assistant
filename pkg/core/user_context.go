type UserContext struct {
    ID          string
    Preferences map[string]interface{}
    HealthData  HealthStatus
}