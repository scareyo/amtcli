package amt

import (
  "log"

  oamt "github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman"
  oamtclient "github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/client"
  oamtpower "github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/cim/power"
  oamtservice "github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/cim/service"
)

type Client struct {
  msg oamt.Messages
}

type ClientParameters struct {
  Host string
  Username string
  Password string
  UseTls bool
}

func Create(params ClientParameters) Client {
  return Client {
    msg: oamt.NewMessages(
      oamtclient.Parameters {
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
  response, err := c.msg.AMT.GeneralSettings.Get()
  if err != nil {
    log.Fatal(err)
  }
  settings := response.Body.GetResponse

  return DeviceInfo {
    State: c.GetPowerState(),
    AmtFqdn: settings.HostName + "." + settings.DomainName,
    HostFqdn: settings.HostOSFQDN,
  }
}

func (c *Client) GetPowerState() PowerState {
  response, err := c.msg.CIM.ServiceAvailableToElement.Enumerate()
  if err != nil {
    log.Fatal(err)
  }

  response, err = c.msg.CIM.ServiceAvailableToElement.Pull(response.Body.EnumerateResponse.EnumerationContext)
  if err != nil {
    log.Fatal(err)
  }

  state := response.Body.PullResponse.AssociatedPowerManagementService

  switch state[0].PowerState {
  case oamtservice.PowerStateOn:
    return PowerStateOn
  case oamtservice.PowerStateOffSoft:
    return PowerStateOff
  case oamtservice.PowerStateOffHard:
    return PowerStateOff
  }

  log.Println("Unhandled power state: " + string(state[0].PowerState))

  return PowerStateUnknown
}

func (c *Client) SetPowerState(state PowerState) bool {
  var powerCode oamtpower.PowerState

  switch state {
    case PowerStateOn:
      powerCode = oamtpower.PowerOn
    case PowerStateOff:
      powerCode = oamtpower.PowerOffSoft
    case PowerStateReset:
      powerCode = oamtpower.MasterBusReset
    case PowerStateRestart:
      powerCode = oamtpower.PowerCycleOffSoft
  }

  response, err := c.msg.CIM.PowerManagementService.RequestPowerStateChange(powerCode)
  if err != nil {
    log.Fatal(err)
  }

  success := (response.Body.RequestPowerStateChangeResponse.ReturnValue == oamtpower.ReturnValueCompletedWithNoError)
  if !success {
    log.Fatal(response.Body.RequestPowerStateChangeResponse.ReturnValue)
  }

  return success
}

func (c *Client) GetBootTarget() BootTarget {
  return BootTargetHdd
}

func (c *Client) SetBootTarget() bool {
  return false
}
