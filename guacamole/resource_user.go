package guacamole

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	guacapi "github.com/mdanidl/guac-api"
	guactypes "github.com/mdanidl/guac-api/types"
)

func resourceUser() *schema.Resource {
	return &schema.Resource{
		Create: resourceUserCreate,
		Read:   resourceUserRead,
		Update: resourceUserUpdate,
		Delete: resourceUserDelete,

		Schema: map[string]*schema.Schema{
			"username": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"password": &schema.Schema{
				Type:      schema.TypeString,
				Required:  true,
				Sensitive: true,
			},
			"full_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"role": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceUserCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*guacapi.Guac)

	user := &guactypes.GuacUser{
		Username: d.Get("username").(string),
		Password: d.Get("password").(string),
		Attributes: guactypes.GuacUserAttributes{
			GuacFullName:           d.Get("full_name").(string),
			GuacOrganizationalRole: d.Get("role").(string),
		},
	}
	resp, err := client.CreateUser(user)
	if err != nil {
		return fmt.Errorf("Error creating user %s", err)
	}
	d.SetId(resp.Username)
	return resourceUserRead(d, meta)
}
func resourceUserRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*guacapi.Guac)
	user := &guactypes.GuacUser{
		Username: d.Id(),
	}
	resp, err := client.ReadUser(user)
	if err != nil {
		return fmt.Errorf("Error reading user %s", err)
	}
	d.Set("username", resp.Username)
	d.Set("full_name", resp.Attributes.GuacFullName)
	d.Set("role", resp.Attributes.GuacOrganizationalRole)

	return nil
}
func resourceUserUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*guacapi.Guac)
	user := &guactypes.GuacUser{
		Username: d.Get("username").(string),
		Password: d.Get("password").(string),
		Attributes: guactypes.GuacUserAttributes{
			GuacFullName:           d.Get("full_name").(string),
			GuacOrganizationalRole: d.Get("role").(string),
		},
	}
	err := client.UpdateUser(user)
	if err != nil {
		return fmt.Errorf("Error updating user %s", err)
	}

	return resourceUserRead(d, meta)
}
func resourceUserDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*guacapi.Guac)
	user := &guactypes.GuacUser{
		Username: d.Id(),
	}
	err := client.DeleteUser(user)
	if err != nil {
		return fmt.Errorf("Error deleting user %s", err)
	}
	return nil
}
