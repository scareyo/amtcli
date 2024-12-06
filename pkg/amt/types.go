package amt

type PowerState int

const (
  On  PowerState = iota
  Off
  Reset
  Restart
  Custom
)

type DeviceInfo struct {
  AmtFqdn string
  HostFqdn string
}
