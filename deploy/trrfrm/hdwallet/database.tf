resource "postgresql_role" "hdwallet-migrator" {
  name     = local.new_migrator_username
  login    = true
  password = local.new_migrator_password
}

resource "postgresql_database" "hdwallet-db" {
  name              = local.database_name
  owner             = postgresql_role.hdwallet-migrator.name
  connection_limit  = -1
  allow_connections = true
}

resource "postgresql_role" "hdwallet-controller" {
  name     = local.new_controller_username
  login    = true
  password = local.new_controller_password
}

resource "postgresql_grant" "controller_user_connect" {
  database    = local.database_name
  role        = postgresql_role.hdwallet-controller.name
  object_type = "database"
  privileges  = ["CONNECT"]
}

resource "postgresql_grant" "controller_user_public_usage" {
  database    = local.database_name
  role        = postgresql_role.hdwallet-controller.name
  schema      = "public"
  object_type = "schema"
  privileges  = ["USAGE"]
}

resource "postgresql_grant" "controller_user_public_sequence_usage" {
  database    = local.database_name
  role        = postgresql_role.hdwallet-controller.name
  schema      = "public"
  object_type = "sequence"
  privileges  = ["USAGE", "SELECT", "UPDATE"]
}

resource "postgresql_default_privileges" "read_only_tables" {
  role     = postgresql_role.hdwallet-controller.name
  database = local.database_name
  schema   = "public"

  owner       = postgresql_role.hdwallet-migrator.name
  object_type = "table"
  privileges  = ["SELECT", "INSERT", "UPDATE", "DELETE"]
}