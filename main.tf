provider "guacamole" {
  url      = "http://localhost:8088/guacamole"
  user     = "guacadmin"
  password = "guacadmin"
}

resource "guacamole_user" "test" {
  count = 4
  username  = "tf-test-${count.index + 1}"
  password  = "tf-test2-${count.index + 1}"
  full_name = "Daniel Meszaros"
  role      = "lofasz"
}

resource "guacamole_user_connection_permissions" "perms_on_connections" {
  count = 4
  username = "${element(guacamole_user.test.*.username,count.index)}"
  connections = [
    "${guacamole_connection.blah.*.id}"
  ]
  permission = "READ"
}

resource "guacamole_connection" "blah" {
  count = 4
  name                     = "uff-${count.index + 1}"
  parent                   = "ROOT"
  max_connections          = "4"
  max_connections_per_user = "4"
  protocol = "vnc"

  vnc_properties {
    hostname = "1.2.3.${count.index + 1}"
    port     = "5901"
    username = "lofasz"
    password = "hellokitty"
    "color-depth" = "24"
    "clipboard-encoding" = "UTF-8"
  }

  // rdp_properties {
  //   b = "c"
  // }
}

output "hostnames" {
  value = "${guacamole_connection.blah.*.vnc_properties.hostname}"
}