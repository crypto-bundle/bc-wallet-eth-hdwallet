resource "jetstream_stream" "hd_wallet_single_wallet_create_stream_v1" {
  name     = local.nats_wallet_management_stream_name
  subjects = [
    "${local.nats_wallet_create_subject_name_v1}.ACCEPT.${local.nats_v1_postfix}",
    "${local.nats_wallet_create_subject_name_v1}.VALIDATE.${local.nats_v1_postfix}",
    "${local.nats_wallet_create_subject_name_v1}.PROCESS.${local.nats_v1_postfix}",
    "${local.nats_wallet_create_subject_name_v1}.CHOREOGRAPHY.${local.nats_v1_postfix}",
    "${local.nats_wallet_create_subject_name_v1}.REPLY.${local.nats_v1_postfix}",

    "${local.nats_wallet_import_subject_name_v1}.ACCEPT.${local.nats_v1_postfix}",
    "${local.nats_wallet_import_subject_name_v1}.VALIDATE.${local.nats_v1_postfix}",
    "${local.nats_wallet_import_subject_name_v1}.PROCESS.${local.nats_v1_postfix}",
    "${local.nats_wallet_import_subject_name_v1}.CHOREOGRAPHY.${local.nats_v1_postfix}",
    "${local.nats_wallet_import_subject_name_v1}.REPLY.${local.nats_v1_postfix}",

    "${local.nats_wallet_enable_subject_name_v1}.ACCEPT.${local.nats_v1_postfix}",
    "${local.nats_wallet_enable_subject_name_v1}.VALIDATE.${local.nats_v1_postfix}",
    "${local.nats_wallet_enable_subject_name_v1}.PROCESS.${local.nats_v1_postfix}",
    "${local.nats_wallet_enable_subject_name_v1}.CHOREOGRAPHY.${local.nats_v1_postfix}",
    "${local.nats_wallet_enable_subject_name_v1}.REPLY.${local.nats_v1_postfix}",

    "${local.nats_wallet_disable_subject_name_v1}.ACCEPT.${local.nats_v1_postfix}",
    "${local.nats_wallet_disable_subject_name_v1}.VALIDATE.${local.nats_v1_postfix}",
    "${local.nats_wallet_disable_subject_name_v1}.PROCESS.${local.nats_v1_postfix}",
    "${local.nats_wallet_disable_subject_name_v1}.CHOREOGRAPHY.${local.nats_v1_postfix}",
    "${local.nats_wallet_disable_subject_name_v1}.REPLY.${local.nats_v1_postfix}",

    "${local.nats_wallets_enabled_subject_name_v1}.ACCEPT.${local.nats_v1_postfix}",
    "${local.nats_wallets_enabled_subject_name_v1}.VALIDATE.${local.nats_v1_postfix}",
    "${local.nats_wallets_enabled_subject_name_v1}.PROCESS.${local.nats_v1_postfix}",
    "${local.nats_wallets_enabled_subject_name_v1}.CHOREOGRAPHY.${local.nats_v1_postfix}",
    "${local.nats_wallets_enabled_subject_name_v1}.REPLY.${local.nats_v1_postfix}",

    "${local.nats_wallets_disable_subject_name_v1}.ACCEPT.${local.nats_v1_postfix}",
    "${local.nats_wallets_disable_subject_name_v1}.VALIDATE.${local.nats_v1_postfix}",
    "${local.nats_wallets_disable_subject_name_v1}.PROCESS.${local.nats_v1_postfix}",
    "${local.nats_wallets_disable_subject_name_v1}.CHOREOGRAPHY.${local.nats_v1_postfix}",
    "${local.nats_wallets_disable_subject_name_v1}.REPLY.${local.nats_v1_postfix}",
  ]
  retention = "interests"
  storage  = "memory"
  replicas = 1
}

#resource "jetstream_stream" "hd_wallet_single_wallet_disable_stream_v1" {
#name     = local.nats_wallet_disable_stream_name_v1
#subjects = [
#"${local.nats_wallet_disable_stream_name_v1}.ACCEPT",
#"${local.nats_wallet_disable_stream_name_v1}.VALIDATE",
#"${local.nats_wallet_disable_stream_name_v1}.PROCESS",
#"${local.nats_wallet_disable_stream_name_v1}.CHOREOGRAPHY",
#"${local.nats_wallet_disable_stream_name_v1}.REPLY",
#]
#retention = "interests"
#storage  = "memory"
#replicas = 1
#}
#
#resource "jetstream_stream" "hd_wallet_multiple_wallets_enable_stream_v1" {
#name     = local.nats_wallets_enable_stream_name_v1
#subjects = [
#"${local.nats_wallets_enable_stream_name_v1}.ACCEPT",
#"${local.nats_wallets_enable_stream_name_v1}.VALIDATE",
#"${local.nats_wallets_enable_stream_name_v1}.PROCESS",
#"${local.nats_wallets_enable_stream_name_v1}.CHOREOGRAPHY",
#"${local.nats_wallets_enable_stream_name_v1}.REPLY",
#]
#retention = "interests"
#storage  = "memory"
#replicas = 1
#}
#
#resource "jetstream_stream" "hd_wallet_multiple_wallets_disable_stream_v1" {
#name     = local.nats_wallets_disable_stream_name_v1
#subjects = [
#"${local.nats_wallets_disable_stream_name_v1}.ACCEPT",
#"${local.nats_wallets_disable_stream_name_v1}.VALIDATE",
#"${local.nats_wallets_disable_stream_name_v1}.PROCESS",
#"${local.nats_wallets_disable_stream_name_v1}.CHOREOGRAPHY",
#"${local.nats_wallets_disable_stream_name_v1}.REPLY",
#]
#retention = "interests"
#storage  = "memory"
#replicas = 1
#}
#
#resource "jetstream_stream" "hd_wallet_single_wallet_import_stream_v1" {
#name     = local.nats_wallet_import_stream_name_v1
#subjects = [
#"${local.nats_wallet_import_stream_name_v1}.ACCEPT",
#"${local.nats_wallet_import_stream_name_v1}.VALIDATE",
#"${local.nats_wallet_import_stream_name_v1}.PROCESS",
#"${local.nats_wallet_import_stream_name_v1}.CHOREOGRAPHY",
#"${local.nats_wallet_import_stream_name_v1}.REPLY",
#]
#retention = "interests"
#storage  = "memory"
#replicas = 1
#}
#
#resource "jetstream_stream" "hd_wallet_single_wallet_enable_stream_v1" {
#name     = local.nats_wallet_enable_stream_name_v1
#subjects = [
#"${local.nats_wallet_enable_stream_name_v1}.ACCEPT",
#"${local.nats_wallet_enable_stream_name_v1}.VALIDATE",
#"${local.nats_wallet_enable_stream_name_v1}.PROCESS",
#"${local.nats_wallet_enable_stream_name_v1}.CHOREOGRAPHY",
#"${local.nats_wallet_enable_stream_name_v1}.REPLY",
#]
#retention = "interests"
#storage  = "memory"
