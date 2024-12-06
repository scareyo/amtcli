package amt

import (
  "fmt"
  oamt "github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman"
  oamtclient "github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/client"
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
    fmt.Println(err)
  }
  settings := response.Body.GetResponse

  return DeviceInfo {
    AmtFqdn: settings.HostName + "." + settings.DomainName,
    HostFqdn: settings.HostOSFQDN,
  }
}

func (c *Client) GetPowerState() PowerState {
  return Off
}

func (c *Client) SetPowerState(state PowerState) {

}
