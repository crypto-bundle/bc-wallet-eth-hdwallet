package wallet_manager

import (
	"context"
	"github.com/crypto-bundle/bc-wallet-tron-hdwallet/internal/app"
	"github.com/crypto-bundle/bc-wallet-tron-hdwallet/internal/types"
	tronCore "github.com/fbsobreira/gotron-sdk/pkg/proto/core"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type singleMnemonicWalletUnit struct {
	logger *zap.Logger

	walletUUID    uuid.UUID
	walletTitle   string
	walletPurpose string

	cfgSrv         configService
	cryptoSrv      encryptService
	walletsDataSrv walletsDataService

	mnemonicWalletsDataSrv mnemonicWalletsDataService

	mnemonicUnit walletPoolMnemonicUnitService
}

func (u *singleMnemonicWalletUnit) Init(ctx context.Context) error {
	return u.mnemonicUnit.Init(ctx)
}

func (u *singleMnemonicWalletUnit) Run(ctx context.Context) error {
	return u.mnemonicUnit.Run(ctx)
}

func (u *singleMnemonicWalletUnit) Shutdown(ctx context.Context) error {
	return u.mnemonicUnit.Shutdown(ctx)
}

func (u *singleMnemonicWalletUnit) GetWalletUUID() uuid.UUID {
	return u.walletUUID
}

func (u *singleMnemonicWalletUnit) GetWalletTitle() string {
	return u.walletTitle
}

func (u *singleMnemonicWalletUnit) GetWalletPurpose() string {
	return u.walletPurpose
}

func (u *singleMnemonicWalletUnit) GetWalletPublicData() *types.PublicWalletData {
	return &types.PublicWalletData{
		UUID:     u.walletUUID,
		Title:    u.walletTitle,
		Purpose:  u.walletPurpose,
		Strategy: types.WalletMakerSingleMnemonicStrategy,
		MnemonicWallets: []*types.PublicMnemonicWalletData{
			u.mnemonicUnit.GetPublicData(),
		},
	}
}

func (u *singleMnemonicWalletUnit) SignTransaction(ctx context.Context,
	mnemonicUUID uuid.UUID,
	account, change, index uint32,
	transaction *tronCore.Transaction,
) ([]byte, error) {
	if mnemonicUUID != u.mnemonicUnit.GetMnemonicUUID() {
		return nil, ErrPassedMnemonicWalletNotFound
	}

	return u.mnemonicUnit.SignTransaction(ctx, account, change, index, transaction)
}

func (u *singleMnemonicWalletUnit) AddMnemonicUnit(unit walletPoolMnemonicUnitService) error {
	if u.mnemonicUnit != nil {
		return ErrMnemonicAlreadySet
	}

	u.mnemonicUnit = unit

	return nil
}

func (u *singleMnemonicWalletUnit) GetAddressByPath(ctx context.Context,
	mnemonicUUID uuid.UUID,
	account, change, index uint32,
) (string, error) {
	if mnemonicUUID != u.mnemonicUnit.GetMnemonicUUID() {
		return "", ErrPassedMnemonicWalletNotFound
	}

	return u.mnemonicUnit.GetAddressByPath(ctx, account, change, index)
}

func (u *singleMnemonicWalletUnit) GetAddressesByPathByRange(ctx context.Context,
	mnemonicWalletUUID uuid.UUID,
	accountIndex uint32,
	internalIndex uint32,
	addressIndexFrom uint32,
	addressIndexTo uint32,
) ([]*types.PublicDerivationAddressData, error) {
	if mnemonicWalletUUID != u.mnemonicUnit.GetMnemonicUUID() {
		return nil, ErrPassedMnemonicWalletNotFound
	}

	return u.mnemonicUnit.GetAddressesByPathByRange(ctx, accountIndex, internalIndex, addressIndexFrom, addressIndexTo)
}

func newSingleMnemonicWalletPoolUnit(logger *zap.Logger,
	walletUUID uuid.UUID,
	walletTitle string,
	walletPurpose string,
) *singleMnemonicWalletUnit {
	return &singleMnemonicWalletUnit{
		logger:        logger.With(zap.String(app.WalletUUIDTag, walletUUID.String())),
		walletUUID:    walletUUID,
		walletTitle:   walletTitle,
		walletPurpose: walletPurpose,
	}
}