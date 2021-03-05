package dashdog

import (
	"errors"
	"github.com/dashdog-provider/dashdog/client"
	"github.com/dashdog-provider/dashdog/resource/dashboard"
	"github.com/dashdog-provider/dashdog/resource/widget"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"host": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("DATADOG_HOST", ""),
			},
			"api_key": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("DATADOG_API_KEY", ""),
			},
			"app_key": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("DATADOG_APP_KEY", ""),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"dashdog_widget": widget.Resource(),
			"dashdog_dashboard": dashboard.Resource(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	host := d.Get("host").(string)
	apiKey := d.Get("api_key").(string)
	appKey := d.Get("app_key").(string)

	datadogClient := client.New(host, apiKey, appKey)
	if !datadogClient.ValidateApiKey() {
		return nil, errors.New("informed API KEY is invalid")
	}

	return datadogClient, nil
}
