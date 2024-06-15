resource "vault_policy" "hdwallet_full_access_policy" {
  name = "${local.bucket_name}-full-access-policy"

  policy = <<EOT
path "kv/data/crypto-bundle/${local.bucket_name}/*" {
  capabilities = ["read", "list"]
}
EOT
}

resource "vault_policy" "hdwallet_application_access_policy" {
  name = "${local.bucket_name}-app-policy"

  policy = <<EOT
path "kv/data/crypto-bundle/bc-wallet-common/jwt" {
  capabilities = ["read", "list"]
}

path "kv/data/crypto-bundle/bc-wallet-common/transit" {
  capabilities = ["read", "list"]
}

path "kv/data/crypto-bundle/${local.bucket_name}/*" {
  capabilities = ["read", "list"]
}

path "kv/data/crypto-bundle/${local.bucket_name}/migrator" {
  capabilities = ["deny"]
}

path "kv/data/crypto-bundle/${local.bucket_name}/vault/migrator" {
  capabilities = ["deny"]
}
EOT
}

resource "vault_policy" "hdwallet_encryption_key_policy" {
  name = "${local.bucket_name}-policy"

  policy = <<EOT
path "transit/*/crypto-bundle-*" {
  capabilities = ["create", "read", "update", "patch", "delete", "list"]
}

path "cryptobundle-transit/+/*" {
  capabilities = ["create", "read", "update", "patch", "delete", "list"]
}
EOT
}

resource "vault_transit_secret_backend_key" "encryption_key" {
  backend = "transit"
  name    = local.bucket_name
}

resource "vault_kubernetes_auth_backend_role" "hdwallet_migrator_auth_role" {
  backend                          = "kubernetes"
  role_name                        = local.k8s_migrator_auth_role_name
  bound_service_account_names      = [
    local.k8s_service_account,
  ]
  bound_service_account_namespaces = [
    var.k8s_namespace,
  ]
  token_ttl                        = 240
  token_policies                   = [
    vault_policy.hdwallet_full_access_policy.name,
  ]
  audience                         = ""
}

resource "vault_kubernetes_auth_backend_role" "hdwallet_application_auth_role" {
  backend                          = "kubernetes"
  role_name                        = local.k8s_app_auth_role_name
  bound_service_account_names      = [
    local.k8s_service_account,
  ]
  bound_service_account_namespaces = [
    var.k8s_namespace,
  ]
  token_ttl                        = 240
  token_policies                   = [
    vault_policy.hdwallet_application_access_policy.name,
  ]
  audience                         = ""
}

resource "vault_kv_secret_v2" "common_bucket" {
  mount                      = "kv"
  name                       = "crypto-bundle/${local.bucket_name}/common"
  data_json                  = jsonencode(
      {
        POSTGRESQL_DATABASE_NAME  = postgresql_database.hdwallet-db.name,
        VAULT_APP_ENCRYPTION_KEY  = vault_transit_secret_backend_key.encryption_key.name
      }
  )
}

resource "vault_kv_secret_v2" "controller_bucket" {
  mount                      = "kv"
  name                       = "crypto-bundle/${local.bucket_name}/controller"
  data_json                  = jsonencode(
      {
        POSTGRESQL_PASSWORD  = postgresql_role.hdwallet-controller.password,
        POSTGRESQL_USERNAME  = postgresql_role.hdwallet-controller.name

        NATS_PASSWORD  = data.vault_kv_secret_v2.common_nats.data["NATS_PASSWORD"],
        NATS_USER  = data.vault_kv_secret_v2.common_nats.data["NATS_USER"],

        REDIS_PASSWORD = data.vault_kv_secret_v2.common_redis.data["REDIS_PASSWORD"]
        REDIS_USER = data.vault_kv_secret_v2.common_redis.data["REDIS_USER"]
      }
  )
}

resource "vault_kv_secret_v2" "migrator_bucket" {
  mount                      = "kv"
  name                       = "crypto-bundle/${local.bucket_name}/migrator"
  data_json                  = jsonencode(
      {
        POSTGRESQL_PASSWORD  = postgresql_role.hdwallet-migrator.password,
        POSTGRESQL_USERNAME  = postgresql_role.hdwallet-migrator.name
      }
  )
}