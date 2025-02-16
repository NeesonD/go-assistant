type HealthDevice struct {
    ID          string
    Type        string
    Status      string
    LastChecked time.Time
}

func (d *HealthDevice) CheckStatus() error {
    // Implement the logic to check the device status
    // This is a placeholder for actual device status checking logic
    return nil
}

func (d *HealthDevice) UpdateStatus(newStatus string) {
    d.Status = newStatus
    d.LastChecked = time.Now()
}