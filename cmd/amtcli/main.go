package main

import (
  "github.com/scareyo/amtcli/cmd/amtcli/cli"
)

func main() {
  cli.Execute()
}

/*func old() {
  host := flag.String("h", "", "hostname")
  username := flag.String("u", "", "username")
  password := flag.String("p", "", "password")

  flag.Parse()

  fmt.Println(*host)

  clientParams := client.Parameters {
    Target: *host,
    Username: *username,
    Password: *password,
    UseDigest: true,
    UseTLS: true,
    SelfSignedAllowed: true,
    LogAMTMessages: true,
  }

  wsmanMessages := wsman.NewMessages(clientParams)

  fmt.Println("**************************************************")
  fmt.Println("General settings")
  fmt.Println("**************************************************")
  response, err := wsmanMessages.AMT.GeneralSettings.Get()
  if err != nil {
    fmt.Println(err)
  }
  fmt.Println(response.Body.GetResponse)
  
  fmt.Println("**************************************************")
  fmt.Println("Get boot capabilities")
  fmt.Println("**************************************************")
  capabilities, err := wsmanMessages.AMT.BootCapabilities.Get()
  if err != nil {
    fmt.Println(err)
  }
  fmt.Println(capabilities)

  fmt.Println("**************************************************")
  fmt.Println("Get boot settings")
  fmt.Println("**************************************************")
  bootConfig, err := wsmanMessages.AMT.BootSettingData.Get()
  if err != nil {
  }

  bootData := bootConfig.Body.BootSettingDataGetResponse

  newData := boot.BootSettingDataRequest{
    BIOSLastStatus:         bootData.BIOSLastStatus,
    BIOSPause:              false,
    BIOSSetup:              false,
    BootMediaIndex:         0,
    BootguardStatus:        bootData.BootguardStatus,
    ConfigurationDataReset: false,
    ElementName:            bootData.ElementName,
    EnforceSecureBoot:      bootData.EnforceSecureBoot,
    FirmwareVerbosity:      0,
    ForcedProgressEvents:   false,
    InstanceID:             bootData.InstanceID,
    LockKeyboard:           false,
    LockPowerButton:        false,
    LockResetButton:        false,
    LockSleepButton:        false,
    OptionsCleared:         true,
    OwningEntity:           bootData.OwningEntity,
    ReflashBIOS:            false,
    UseIDER:                false,
    UseSOL:                 false,
    UseSafeMode:            false,
    UserPasswordBypass:     false,
    SecureErase:            false,
    IDERBootDevice:         1,
  }

  fmt.Println("**************************************************")
  fmt.Println("Set boot config role")
  fmt.Println("**************************************************")
  configRole, err := wsmanMessages.CIM.BootService.SetBootConfigRole("Intel(r) AMT: Boot Configuration 0", 1)
  if err != nil {
    fmt.Println(err)
  }
  fmt.Println(configRole)

  fmt.Println("**************************************************")
  fmt.Println("Change boot order")
  fmt.Println("**************************************************")
  bootOrder, err := wsmanMessages.CIM.BootConfigSetting.ChangeBootOrder("Intel(r) AMT: Force PXE Boot")
  if err != nil {
    fmt.Println(err)
  }
  fmt.Println(bootOrder)
  
  fmt.Println("**************************************************")
  fmt.Println("Set boot settings")
  fmt.Println("**************************************************")
  putBootSettingData, err := wsmanMessages.AMT.BootSettingData.Put(newData)
  if err != nil {
    fmt.Println(err)
  }
  fmt.Println(putBootSettingData)

  /***************************************************
  // Power On.
  PowerOn PowerState = 2 // Verified Hardware Power On

  // Sleep - Light.
  SleepLight PowerState = 3 // ?

  // Sleep - Deep.
  SleepDeep PowerState = 4 // ?

  // Power Cycle (Off Soft).
  PowerCycleOffSoft PowerState = 6 // ?

  // Power Off - Hard.
  PowerOffHard PowerState = 8 // Verfied Hardware Power Off

  // Hibernate.
  Hibernate PowerState = 7 // ?

  // Power Off - Soft.
  PowerOffSoft PowerState = 9 // ?

  // Power Cycle (Off Hard).
  PowerCycleOffHard PowerState = 5 // Verified Hardware Power Cycle (off then on)

  // Master Bus Reset.
  MasterBusReset PowerState = 10 // Verified Hardware Reboot

  // Diagnostic Interrupt (NMI).
  DiagnosticInterruptNMI PowerState = 11 // ?

  // Power Off - Soft Graceful.
  PowerOffSoftGraceful PowerState = 12 // ?

  // Power Off - Hard Graceful.
  PowerOffHardGraceful PowerState = 13 // ?

  // Master Bus Reset Graceful.
  MasterBusResetGraceful PowerState = 14 // ?

  // Power Cycle (Off - Soft Graceful).
  PowerCycleOffSoftGraceful PowerState = 15 // ?

  // Power Cycle (Off - Hard Graceful).
  PowerCycleOffHardGraceful PowerState = 16 // ?
  ***************************************************
  fmt.Println("**************************************************")
  fmt.Println("Power State Change")
  fmt.Println("**************************************************")
  powerResponse, err := wsmanMessages.CIM.PowerManagementService.RequestPowerStateChange(power.PowerState(5))
  if err != nil {
    fmt.Println(err)
  }
  fmt.Println(powerResponse.Body.RequestPowerStateChangeResponse.ReturnValue)
  
  
}*/
