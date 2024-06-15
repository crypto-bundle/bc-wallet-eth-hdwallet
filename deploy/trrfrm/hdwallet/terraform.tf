terraform {
  required_providers {
    postgresql = {
        source = "cyrilgdn/postgresql"
    }
    vault = {
      source  = "hashicorp/vault"
    }
    random = {
        source  = "hashicorp/random"
    }
    jetstream = {
      source = "nats-io/jetstream"
      version = "0.1.1"
    }
  }

  backend "pg" {
  }
}