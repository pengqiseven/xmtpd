package blockchain

import (
	"context"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/xmtp/xmtpd/pkg/abi/rateregistry"
	"github.com/xmtp/xmtpd/pkg/config"
	"go.uber.org/zap"
)

/*
*
A RatesAdmin is a struct responsible for calling admin functions on the RatesRegistry contract
*
*/
type RatesAdmin struct {
	client   *ethclient.Client
	signer   TransactionSigner
	contract *rateregistry.RateRegistry
	logger   *zap.Logger
}

func NewRatesAdmin(
	logger *zap.Logger,
	client *ethclient.Client,
	signer TransactionSigner,
	contractsOptions config.ContractsOptions,
) (*RatesAdmin, error) {
	contract, err := rateregistry.NewRateRegistry(
		common.HexToAddress(contractsOptions.SettlementChain.RateRegistryAddress),
		client,
	)
	if err != nil {
		return nil, err
	}

	return &RatesAdmin{
		signer:   signer,
		client:   client,
		logger:   logger.Named("RatesAdmin"),
		contract: contract,
	}, nil
}

/**
*
* AddRates adds a new rate to the rates manager.
* The new rate must have a later start time than the last rate in the contract.
 */
func (r *RatesAdmin) AddRates(
	ctx context.Context,
	rates rateregistry.RateRegistryRates,
) error {
	tx, err := r.contract.AddRates(
		&bind.TransactOpts{
			Signer:  r.signer.SignerFunc(),
			From:    r.signer.FromAddress(),
			Context: ctx,
		},
		rates.MessageFee,
		rates.StorageFee,
		rates.CongestionFee,
		rates.TargetRatePerMinute,
		rates.StartTime,
	)
	if err != nil {
		return err
	}

	_, err = WaitForTransaction(
		ctx,
		r.logger,
		r.client,
		2*time.Second,
		250*time.Millisecond,
		tx.Hash(),
	)

	return err
}

func (r *RatesAdmin) Contract() *rateregistry.RateRegistry {
	return r.contract
}
