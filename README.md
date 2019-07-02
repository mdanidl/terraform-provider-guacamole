# terraform-provider-guacamole

This custom provider is for creating resources on an [Apache Guacamole](http://guacamole.apache.org) server.

The capabililties so far are somewhat limited, but enough to get started.

## Configuration

Configuration of the provider is very simple:

```ruby
provider "guacamole" {
  url      = "http://localhost:8080/guacamole"
  user     = "guacadmin"
  password = "guacadmin"
}
```

## Resources

Currently the provider supports 3 resources:

### guacamole_user

```ruby
resource "guacamole_user" "training_users" {
  username  = "user"
  password  = "password"
  full_name = "Training User"
}
```
Note: The user database's unique ID is the username, so if you want to change the username, it will be a destroy/create.

### guacamole_connection

This is the main part "where the magic happens". 

```ruby
resource "guacamole_connection" "vnc" {
  count                    = "${var.count}"
  name                     = "VNC-${count.index + 1}"
  parent                   = "ROOT"
  max_connections          = "1"
  max_connections_per_user = "1"
  protocol                 = "vnc"

  vnc_properties {
    hostname           = "${element(module.linux_instances.ip_addresses,count.index)}"
    port               = "5901"
    password           = "12345678"
    color-depth        = "24"
    clipboard-encoding = "UTF-8"
  }
}
```

It's important that after choosing protocol, you create a block with the corresponding properties block. The list of available property names I was able to identify so far: https://github.com/mdanidl/guac-api/blob/1567fead4f1658cbaeac4b4d69a213e703ff1fb1/types/connections.go#L22

There are overlaps, I might create a matrix table for this later.

### guacamole_user_connection_permission

```ruby
resource "guacamole_user_connection_permissions" "perms" {
  count    = "${var.count}"
  username = "${element(guacamole_user.dpg_users.*.username, count.index)}"

  connections = [
    "${element(guacamole_connection.vnc.*.id, count.index)}",
    "${element(guacamole_connection.ssh_user.*.id, count.index)}",
  ]

  permission = "READ"
}
```

Note: when adding new permissions, this works just fine, but I haven't tested the "destroy" path too much, and since it's a json-patch at the backend, it might not actually have the desired effect. 

