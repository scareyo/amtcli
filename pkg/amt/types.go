package amt

type PowerState string; const (
  PowerStateOn        PowerState = "On"
  PowerStateOff       PowerState = "Off" 
  PowerStateReset     PowerState = "Reset" 
  PowerStateRestart   PowerState = "Restart"
  PowerStateUnknown   PowerState = "Unknown"
)

type BootTarget string; const (
  BootTargetHdd   BootTarget = "HDD"
  BootTargetCd    BootTarget = "CD/DVD"
  BootTargetPxe   BootTarget = "PXE"
  BootTargetBios  BootTarget = "BIOS"
)

type DeviceInfo struct {
  State       PowerState  `json:"state"`
  AmtFqdn     string      `json:"amtFqdn"`
  HostFqdn    string      `json:"hostFqdn"` 
}
