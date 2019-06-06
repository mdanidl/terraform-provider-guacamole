provider "guacamole" {
  url = "http://localhost:8088/guacamole"
  user = "guacadmin"
  password = "guacadmin"
}

resource "guacamole_user" "test" {
  username = "tf-test"
  password = "tf-test2"
}

