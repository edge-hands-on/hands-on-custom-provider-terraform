package widget

import (
	"encoding/json"
	types "github.com/dashdog-provider/dashdog/type"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func Resource() *schema.Resource {
	return &schema.Resource{
		Create: onCreate,
		Read:   onRead,
		Update: onUpdate,
		Delete: onDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		SchemaVersion: 1,

		Schema: map[string]*schema.Schema{
			"type": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The widget type, one of ('timeseries', 'toplist', 'query_value').",
			},

			"title": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Title for the widget.",
			},

			"autoscale": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Description: "Autoscale value for query value widget.",
			},

			"precision": {
				Type:        schema.TypeInt,
				Optional:    true,
				ForceNew:    true,
				Description: "The precision for query value widget.",
			},

			"json": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The widget json.",
			},

			"requests": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"query": {
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
							Description: "The query to be executed.",
						},

						"display_type": {
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Description: "Used for the time series widget.",
						},

						"aggregator": {
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Description: "Used for the query value widget.",
						},
					},
				},
			},
		},
	}
}

func onCreate(d *schema.ResourceData, _ interface{}) error {
	return do("create", d)
}

func onRead(d *schema.ResourceData, _ interface{}) error {
	return do("read", d)
}

func onUpdate(d *schema.ResourceData, _ interface{}) error {
	return do("update", d)
}

func onDelete(d *schema.ResourceData, _ interface{}) error {
	return do("delete", d)
}

func do(event string, d *schema.ResourceData) error {
	var widget types.Widget
	widgetType := d.Get("type").(string)
	if widgetType == "timeseries" {
		timeserieWidget := types.TimeseriesWidget{
			Type:  "timeseries",
			Title: d.Get("title").(string),
		}

		requestsList := d.Get("requests").([]interface {})
		requests := make([]types.TimeseriesRequest, len(requestsList))
		for i, resource := range requestsList {
			tResource := resource.(map[string]interface{})

			query := tResource["query"].(string)
			displayType := tResource["display_type"].(string)

			requests[i] = types.TimeseriesRequest{
				DisplayType: displayType,
				Q:           query,
			}
		}

		timeserieWidget.Requests = requests

		widget = types.Widget{
			Definition: timeserieWidget,
		}
	}

	widgetKey := uuid.New()

	log.Printf("Executing: %s %s", event, widgetType)

	if event == "create" {
		widgetJson, err := json.Marshal(widget)
		if err != nil {
			panic(err)
		}

		d.SetId(widgetKey.String())
		err = d.Set("json", string(widgetJson))
		if err != nil {
			panic(err)
		}
	} else if event == "read" {
	} else if event == "update" {
		widgetJson, err := json.Marshal(widget)
		if err != nil {
			panic(err)
		}

		err = d.Set("json", string(widgetJson))
		if err != nil {
			panic(err)
		}
	} else if event == "delete" {
		d.SetId("")
	}

	return nil
}
