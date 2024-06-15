variable "k8s_namespace" {
  type = string
  default = "default"
}

variable "stage_name" {
  type = string
  default = "DEV"
}

variable "nats_server_address" {
  type = string
  default = "nats.default.svc.cluster.local:4222"
}

variable "network" {
  type = string
  default = "ethereum"
}