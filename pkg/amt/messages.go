package amt

import (
  "log"

  amtboot "github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/amt/boot"
  amtgeneral "github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/amt/general"
  cimboot "github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/cim/boot"
  cimpower "github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/cim/power"
  cimservice "github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/cim/service"
)

func (c *Client) getGeneralSettings() amtgeneral.GeneralSettingsResponse {
  response, err := c.msg.AMT.GeneralSettings.Get()
  if err != nil {
    log.Fatal(err)
  }
  return response.Body.GetResponse
}

func (c *Client) getPowerState() cimservice.PowerState {
  response, err := c.msg.CIM.ServiceAvailableToElement.Enumerate()
  if err != nil {
    log.Fatal(err)
  }

  response, err = c.msg.CIM.ServiceAvailableToElement.Pull(response.Body.EnumerateResponse.EnumerationContext)
  if err != nil {
    log.Fatal(err)
  }

  return response.Body.PullResponse.AssociatedPowerManagementService[0].PowerState
}

func (c *Client) setPowerState(state cimpower.PowerState) cimpower.ReturnValue {
  response, err := c.msg.CIM.PowerManagementService.RequestPowerStateChange(state)
  if err != nil {
    log.Fatal(err)
  }
  return response.Body.RequestPowerStateChangeResponse.ReturnValue
}

func (c *Client) getBootCapabilities() amtboot.BootCapabilitiesResponse {
  capabilities, err := c.msg.AMT.BootCapabilities.Get()
  if err != nil {
    log.Fatal(err)
  }
  return capabilities.Body.BootCapabilitiesGetResponse
}

func (c *Client) getBootSettings() amtboot.BootSettingDataResponse {
  capabilities, err := c.msg.AMT.BootSettingData.Get()
  if err != nil {
    log.Fatal(err)
  }
  return capabilities.Body.BootSettingDataGetResponse
}

func (c *Client) setBootSettings(settings amtboot.BootSettingDataRequest) amtboot.Response {
  putBootSettingData, err := c.msg.AMT.BootSettingData.Put(settings)
  if err != nil {
    log.Println(err)
  }
  return putBootSettingData
}

func (c *Client) setBootConfigRole(instanceId string, role int) cimboot.ReturnValue {
  response, err := c.msg.CIM.BootService.SetBootConfigRole(instanceId, role)
  if err != nil {
    log.Println(err)
  }
  return response.Body.SetBootConfigRole_OUTPUT.ReturnValue
}

func (c *Client) setBootOrder(source cimboot.Source) cimboot.ReturnValue {
  response, err := c.msg.CIM.BootConfigSetting.ChangeBootOrder(source)
  if err != nil {
    log.Println(err)
  }
  return response.Body.ChangeBootOrder_OUTPUT.ReturnValue
}
