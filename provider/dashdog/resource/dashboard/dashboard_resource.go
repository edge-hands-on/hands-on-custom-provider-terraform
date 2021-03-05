package dashboard

import (
	"encoding/json"
	"fmt"
	"github.com/dashdog-provider/dashdog/client"
	types "github.com/dashdog-provider/dashdog/type"
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
			"title": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Dashboard's title.",
			},

			"description": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Dashboard's description.",
			},

			"read_only": {
				Type:        schema.TypeBool,
				Required:    true,
				Description: "If the dashboard is read only or not.",
			},

			"template_variables": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The variable name.",
						},

						"tag": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The associated tag or attribute for this variable.",
						},

						"default": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The variable default value.",
						},
					},
				},
			},

			"widgets": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"title": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The group title.",
						},

						"json": {
							Type:        schema.TypeList,
							Required:    true,
							Elem:        &schema.Schema{Type: schema.TypeString},
							Description: "The widgets' json.",
						},
					},
				},
			},
		},
	}
}

func onCreate(d *schema.ResourceData, m interface{}) error {
	return do("create", d, m)
}

func onRead(d *schema.ResourceData, m interface{}) error {
	return do("read", d, m)
}

func onUpdate(d *schema.ResourceData, m interface{}) error {
	return do("update", d, m)
}

func onDelete(d *schema.ResourceData, m interface{}) error {
	return do("delete", d, m)
}

func do(event string, d *schema.ResourceData, m interface{}) error {
	datadogClient := m.(client.DatadogClient)

	var dashboard types.Dashboard

	title := d.Get("title").(string)
	description := d.Get("description").(string)
	readOnly := d.Get("read_only").(bool)

	templateVariableList := d.Get("template_variables").([]interface{})
	templateVariables := make([]types.TemplateVariable, len(templateVariableList))
	for i, resource := range templateVariableList {
		tResource := resource.(map[string]interface{})

		name := tResource["name"].(string)
		tag := tResource["tag"].(string)
		varDefault := tResource["default"].(string)

		templateVariables[i] = types.TemplateVariable{
			Name:    name,
			Prefix:  tag,
			Default: varDefault,
		}
	}

	widgetResourceList := d.Get("widgets").([]interface{})
	widgetResources := make([]types.GroupWidget, len(widgetResourceList))
	for i, resource := range widgetResourceList {
		tResource := resource.(map[string]interface{})

		groupTitle := tResource["title"].(string)
		jsonResource := tResource["json"].([]interface{})

		widgets := make([]types.Widget, len(jsonResource))
		for i, value := range jsonResource {
			err := json.Unmarshal([]byte(fmt.Sprint(value)), &widgets[i])
			if err != nil {
				log.Fatal(err)
			}
		}

		widgetResources[i] = types.GroupWidget{
			Definition: types.GroupWidgetDefinition{
				LayoutType: "ordered",
				Type:       "group",
				Title:      groupTitle,
				Widgets:    widgets,
			},
		}
	}

	dashboard = types.Dashboard{
		Title:             title,
		Description:       description,
		IsReadOnly:        readOnly,
		LayoutType:        "ordered",
		TemplateVariables: templateVariables,
		GroupWidget:       widgetResources,
	}

	log.Printf("Executing: %s %s", event, dashboard.Title)

	if event == "create" {
		dashboardJson, err := json.Marshal(dashboard)
		if err != nil {
			log.Fatal(err)
		}

		dashboardId, err := datadogClient.CreateDashboard(dashboardJson)
		if err != nil {
			log.Fatal(err)
		}

		d.SetId(dashboardId.(string))
	} else if event == "read" {
		id := d.Id()
		dashboardId, err := datadogClient.GetDashboard(id)
		if err != nil {
			log.Fatal(err)
		}

		d.SetId(dashboardId.(string))
	} else if event == "update" {
		id := d.Id()
		dashboardJson, err := json.Marshal(dashboard)
		if err != nil {
			log.Fatal(err)
		}

		err = datadogClient.UpdateDashboard(id, dashboardJson)
		if err != nil {
			log.Fatal(err)
		}
	} else if event == "delete" {
		id := d.Id()
		err := datadogClient.DeleteDashboard(id)
		if err != nil {
			log.Fatal(err)
		}

		d.SetId("")
	}

	return nil
}
