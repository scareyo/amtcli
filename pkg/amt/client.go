package amt

import (
  "log"

  wsman "github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman"
  wsmanclient "github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/client"
  amtboot "github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/amt/boot"
  cimboot "github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/cim/boot"
  cimpower "github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/cim/power"
  cimservice "github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/cim/service"
)

type Client struct {
  msg wsman.Messages
}

type ClientParameters struct {
  Host string
  Username string
  Password string
  UseTls bool
}

func Create(params ClientParameters) Client {
  return Client {
    msg: wsman.NewMessages(
      wsmanclient.Parameters {
        Target: params.Host,
        Username: params.Username,
        Password: params.Password,
        UseTLS: params.UseTls,
        UseDigest: true,
        SelfSignedAllowed: true,
        LogAMTMessages: true,
      }),
  }
}

func (c *Client) GetInfo() DeviceInfo {
  settings := c.getGeneralSettings()
  
  return DeviceInfo {
    State: c.GetPowerState(),
    AmtFqdn: settings.HostName + "." + settings.DomainName,
    HostFqdn: settings.HostOSFQDN,
  }
}

func (c *Client) GetPowerState() PowerState {
  state := c.getPowerState()

  switch state {
  case cimservice.PowerStateOn:
    return PowerStateOn
  case cimservice.PowerStateOffSoft:
    return PowerStateOff
  case cimservice.PowerStateOffHard:
    return PowerStateOff
  }

  return PowerStateUnknown
}

func (c *Client) SetPowerState(state PowerState) bool {
  var powerCode cimpower.PowerState

  switch state {
    case PowerStateOn:
      powerCode = cimpower.PowerOn
    case PowerStateOff:
      powerCode = cimpower.PowerOffSoft
    case PowerStateReset:
      powerCode = cimpower.MasterBusReset
    case PowerStateRestart:
      powerCode = cimpower.PowerCycleOffSoft
  }

  response := c.setPowerState(powerCode)
  if response != cimpower.ReturnValueCompletedWithNoError {
    log.Println(response)
    return false
  }

  return true
}

func (c *Client) SetBootTarget(target BootTarget) bool {
  var source cimboot.Source

  switch target {
    case BootTargetCd:
      source = cimboot.CD
    case BootTargetHdd:
      source = cimboot.HardDrive
    case BootTargetPxe:
      source = cimboot.PXE
    case BootTargetBios:
      log.Fatal("BIOS boot target not yet supported")
  }

  settings := c.getBootSettings()
  c.setBootConfigRole("Intel(r) AMT: Boot Configuration 0", 1)
  c.setBootOrder(source)

  newSettings := amtboot.BootSettingDataRequest{
    BIOSLastStatus:         settings.BIOSLastStatus,
    BIOSPause:              false,
    BIOSSetup:              false,
    BootMediaIndex:         0,
    BootguardStatus:        settings.BootguardStatus,
    ConfigurationDataReset: false,
    ElementName:            settings.ElementName,
    EnforceSecureBoot:      settings.EnforceSecureBoot,
    FirmwareVerbosity:      0,
    ForcedProgressEvents:   false,
    InstanceID:             settings.InstanceID,
    LockKeyboard:           false,
    LockPowerButton:        false,
    LockResetButton:        false,
    LockSleepButton:        false,
    OptionsCleared:         true,
    OwningEntity:           settings.OwningEntity,
    ReflashBIOS:            false,
    UseIDER:                false,
    UseSOL:                 false,
    UseSafeMode:            false,
    UserPasswordBypass:     false,
    SecureErase:            false,
    IDERBootDevice:         1,
  }

  c.setBootSettings(newSettings)

  return false
}
