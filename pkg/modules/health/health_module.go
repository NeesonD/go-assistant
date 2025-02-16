package health

import (
    "context"
    "fmt"
)

type HealthDevice struct {
    ID       string
    Status   string
    // Add more fields as necessary
}

type HealthAlert struct {
    Type    string
    Device  HealthDevice
    Message string
}

type HealthModule struct {
    devices map[string]HealthDevice
    alerts  chan HealthAlert
}

func NewHealthModule() *HealthModule {
    return &HealthModule{
        devices: make(map[string]HealthDevice),
        alerts:  make(chan HealthAlert, 10),
    }
}

func (m *HealthModule) Start(ctx context.Context) {
    go m.monitorDevices(ctx)
}

func (m *HealthModule) monitorDevices(ctx context.Context) {
    for {
        select {
        case alert := <-m.alerts:
            m.processAlert(alert)
        case <-ctx.Done():
            return
        }
    }
}

func (m *HealthModule) processAlert(alert HealthAlert) {
    if alert.Type == "abnormal_heartrate" {
        sendEmergencyNotification(alert)
        eventBus <- Event{
            Type: "system.emergency",
            Payload: alert,
        }
    }
}

func sendEmergencyNotification(alert HealthAlert) {
    fmt.Printf("Emergency alert: %s from device %s\n", alert.Message, alert.Device.ID)
}