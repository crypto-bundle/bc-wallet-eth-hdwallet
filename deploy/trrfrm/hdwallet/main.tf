provider "vault" {
  skip_tls_verify = true
}

provider "postgresql" {
    sslmode  = "disable"
}

provider "jetstream" {
  servers     = var.nats_server_address
}

data "vault_kv_secret_v2" "trrfmr" {
    mount = "kv"
    name = "crypto-bundle/bc-wallet-common/trrfrm"
#     path =
}

data "vault_kv_secret_v2" "common_redis" {
  mount = "kv"
  name = "crypto-bundle/bc-wallet-common/redis"
  #     path =
}

data "vault_kv_secret_v2" "common_nats" {
  mount = "kv"
  name = "crypto-bundle/bc-wallet-common/nats"
  #     path =
}

resource "random_password" "password" {
  length           = 16
  special          = true
  override_special = "!#$%&*()-_=+[]{}<>:?"
}

resource "random_password" "migrator_password" {
  length           = 16
  special          = true
  override_special = "!#$%&*()-_=+[]{}<>:?"
}

resource "random_password" "controller_password" {
  length           = 16
  special          = true
  override_special = "!#$%&*()-_=+[]{}<>:?"
}