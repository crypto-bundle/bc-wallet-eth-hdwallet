locals {
    k8s_service_account = "bc-wallet-${var.network}-hdwallet"

    stage_name = upper(var.stage_name)

    application_name = "hdwallet"
    bucket_name = "bc-wallet-${var.network}-hdwallet"

    database_name = "bc-wallet-${var.network}-hdwallet"
    new_password = random_password.password

    new_migrator_username = "bc-wallet-${var.network}-hdwallet-migrator"
    new_migrator_password = random_password.migrator_password.result

    new_controller_username = "bc-wallet-${var.network}-hdwallet-controller"
    new_controller_password = random_password.controller_password.result

    k8s_app_auth_role_name = "bc-wallet-${var.network}-hdwallet-app-auth-role"
    k8s_migrator_auth_role_name = "bc-wallet-${var.network}-hdwallet-migrator-auth-role"


    nats_wallet_management_name_part = "MANAGE_WALLET"
    nats_wallet_create_name_part = "CREATE_ONE"
    nats_wallet_import_name_part = "IMPORT_ONE"
    nats_wallet_enable_name_part = "ENABLE_ONE"
    nats_wallet_disable_name_part = "DISABLE_ONE"

    nats_wallets_enable_name_part = "ENABLE_MULTIPLE"
    nats_wallets_disable_name_part = "DISABLE_MULTIPLE"

    nats_v1_postfix = "V1"

    nats_wallet_management_stream_name = upper(format("%s__%s__%s__%s__%s",
        local.stage_name,
        local.application_name,
        local.nats_wallet_management_name_part,
        "cryptobundle",
        var.network,
    ))

    nats_wallet_create_subject_name_v1 = upper(format("%s.%s",
        local.nats_wallet_management_stream_name,
        local.nats_wallet_create_name_part,
    ))
    nats_wallet_import_subject_name_v1 = upper(format("%s.%s",
        local.nats_wallet_management_stream_name,
        local.nats_wallet_import_name_part,
    ))
    nats_wallet_enable_subject_name_v1 = upper(format("%s.%s",
        local.nats_wallet_management_stream_name,
        local.nats_wallet_enable_name_part,
    ))
    nats_wallet_disable_subject_name_v1 = upper(format("%s.%s",
        local.nats_wallet_management_stream_name,
        local.nats_wallet_disable_name_part,
    ))

    nats_wallets_enabled_subject_name_v1 = upper(format("%s.%s",
        local.nats_wallet_management_stream_name,
        local.nats_wallets_enable_name_part,
    ))

    nats_wallets_disable_subject_name_v1 = upper(format("%s.%s",
        local.nats_wallet_management_stream_name,
        local.nats_wallets_disable_name_part,
    ))
}