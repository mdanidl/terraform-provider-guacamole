package guacamole

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	guacapi "github.com/mdanidl/guac-api"
	guactypes "github.com/mdanidl/guac-api/types"
)

func resourceUserConnectionPermissions() *schema.Resource {
	return &schema.Resource{
		Create: resourceUserConnectionPermissionsCreate,
		Read:   resourceUserConnectionPermissionsRead,
		Update: resourceUserConnectionPermissionsUpdate,
		Delete: resourceUserConnectionPermissionsDelete,

		Schema: map[string]*schema.Schema{
			"username": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"connections": &schema.Schema{
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Required: true,
			},
			"permission": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceUserConnectionPermissionsCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*guacapi.Guac)
	reqDataSlice := []guactypes.GuacPermissionItem{}

	for _, connid := range d.Get("connections").([]interface{}) {
		reqData := guactypes.GuacPermissionItem{
			Op:    "add",
			Path:  "/connectionPermissions/" + connid.(string),
			Value: d.Get("permission").(string),
		}
		reqDataSlice = append(reqDataSlice, reqData)
	}
	err := client.SendUserConnectionPermissionChanges(d.Get("username").(string), reqDataSlice)
	if err != nil {
		return fmt.Errorf("Error creating permissions to connections for user %s. Error: %s", d.Get("username"), err)
	}
	d.SetId(d.Get("username").(string))
	return nil
}
func resourceUserConnectionPermissionsRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}
func resourceUserConnectionPermissionsUpdate(d *schema.ResourceData, meta interface{}) error {
	return resourceUserConnectionPermissionsCreate(d, meta)
}
func resourceUserConnectionPermissionsDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*guacapi.Guac)
	reqDataSlice := []guactypes.GuacPermissionItem{}

	for _, connid := range d.Get("connections").([]interface{}) {
		reqData := guactypes.GuacPermissionItem{
			Op:    "remove",
			Path:  "/connectionPermissions/" + connid.(string),
			Value: d.Get("permission").(string),
		}
		reqDataSlice = append(reqDataSlice, reqData)
	}
	err := client.SendUserConnectionPermissionChanges(d.Get("username").(string), reqDataSlice)
	if err != nil {
		return fmt.Errorf("Error creating permissions to connections for user %s. Error: %s", d.Get("username"), err)
	}
	return nil
}
