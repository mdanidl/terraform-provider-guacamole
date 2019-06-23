package guacamole

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform/helper/schema"
	guacapi "github.com/mdanidl/guac-api"
	guactypes "github.com/mdanidl/guac-api/types"
)

func resourceConnection() *schema.Resource {
	return &schema.Resource{
		Create: resourceConnectionCreate,
		Read:   resourceConnectionRead,
		Update: resourceConnectionUpdate,
		Delete: resourceConnectionDelete,

		Schema: map[string]*schema.Schema{
			// "id": &schema.Schema{
			// 	Type:     schema.TypeString,
			// 	Optional: true,
			// 	Computed: true,
			// },
			"parent": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: connectionParentValid,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"protocol": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: connectionProtocolValid,
			},
			"max_connections": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "1",
			},
			"max_connections_per_user": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "1",
			},
			"weight": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"failover_only": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"guacd_hostname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"guacd_port": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"guacd_encryption": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"rdp_properties": &schema.Schema{
				Type:          schema.TypeMap,
				Optional:      true,
				ConflictsWith: []string{"vnc_properties", "telnet_properties", "ssh_properties"},
			},
			"vnc_properties": &schema.Schema{
				Type: schema.TypeMap,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional:      true,
				ConflictsWith: []string{"rdp_properties", "telnet_properties", "ssh_properties"},
			},
			"telnet_properties": &schema.Schema{
				Type:          schema.TypeMap,
				Optional:      true,
				ConflictsWith: []string{"vnc_properties", "rdp_properties", "ssh_properties"},
			},
			"ssh_properties": &schema.Schema{
				Type: schema.TypeMap,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional:      true,
				ConflictsWith: []string{"vnc_properties", "telnet_properties", "rdp_properties"},
			},
		},
	}
}

func resourceConnectionCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*guacapi.Guac)

	chosenProtocolProperties := strings.ToLower(d.Get("protocol").(string)) + "_properties"
	kvs := d.Get(chosenProtocolProperties).(map[string]interface{})
	marshall, err := json.Marshal(kvs)
	if err != nil {
		log.Panicln(err)
	}
	connectionParams := &guactypes.GuacConnectionParameters{}
	err = json.Unmarshal(marshall, &connectionParams)
	if err != nil {
		log.Panicln(err)
	}

	resourceData := &guactypes.GuacConnection{
		Name:             d.Get("name").(string),
		ParentIdentifier: d.Get("parent").(string),
		Protocol:         strings.ToLower(d.Get("protocol").(string)),
		Attributes: guactypes.GuacConnectionAttributes{
			GuacdEncryption:       d.Get("guacd_encryption").(string),
			GuacdHostname:         d.Get("guacd_hostname").(string),
			GuacdPort:             d.Get("guacd_port").(string),
			FailoverOnly:          d.Get("failover_only").(string),
			Weight:                d.Get("weight").(string),
			MaxConnections:        d.Get("max_connections").(string),
			MaxConnectionsPerUser: d.Get("max_connections_per_user").(string),
		},
		Properties: *connectionParams,
	}

	resp, err := client.CreateConnection(resourceData)
	if err != nil {
		return fmt.Errorf("Error creating connection %s", err)
	}

	d.SetId(resp.Identifier)
	return nil
}
func resourceConnectionRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*guacapi.Guac)
	conn := &guactypes.GuacConnection{
		Identifier: d.Id(),
	}
	data, err := client.ReadConnection(conn)
	if err != nil {
		return fmt.Errorf("Error reading connection %s", err)
	}
	chosenProtocolProperties := data.Protocol + "_properties"
	d.Set("name", data.Name)
	d.Set("parent", data.ParentIdentifier)
	d.Set("protocol", data.Protocol)
	d.Set("hostame", data.Properties.Hostname)
	d.Set("weight", data.Attributes.Weight)
	d.Set(chosenProtocolProperties, data.Properties)

	return nil
}
func resourceConnectionUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*guacapi.Guac)

	chosenProtocolProperties := strings.ToLower(d.Get("protocol").(string)) + "_properties"
	kvs := d.Get(chosenProtocolProperties).(map[string]interface{})
	marshall, err := json.Marshal(kvs)
	if err != nil {
		log.Panicln(err)
	}
	connectionParams := &guactypes.GuacConnectionParameters{}
	err = json.Unmarshal(marshall, &connectionParams)
	if err != nil {
		log.Panicln(err)
	}

	resourceData := &guactypes.GuacConnection{
		Name:             d.Get("name").(string),
		Identifier:       d.Id(),
		ParentIdentifier: d.Get("parent").(string),
		Protocol:         strings.ToLower(d.Get("protocol").(string)),
		Attributes: guactypes.GuacConnectionAttributes{
			GuacdEncryption:       d.Get("guacd_encryption").(string),
			GuacdHostname:         d.Get("guacd_hostname").(string),
			GuacdPort:             d.Get("guacd_port").(string),
			FailoverOnly:          d.Get("failover_only").(string),
			Weight:                d.Get("weight").(string),
			MaxConnections:        d.Get("max_connections").(string),
			MaxConnectionsPerUser: d.Get("max_connections_per_user").(string),
		},
		Properties: *connectionParams,
	}

	err = client.UpdateConnection(resourceData)
	if err != nil {
		return fmt.Errorf("Error updating connection %s", err)
	}
	return resourceConnectionRead(d, meta)
}
func resourceConnectionDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*guacapi.Guac)
	connToDelete := &guactypes.GuacConnection{
		Identifier: d.Id(),
	}

	err := client.DeleteConnection(connToDelete)
	if err != nil {
		return fmt.Errorf("Error deleting connection %s", err)
	}
	return nil
}
