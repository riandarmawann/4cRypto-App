package common

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"

	// "fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func TransferETH(client *ethclient.Client, from *ecdsa.PrivateKey, to common.Address, amount *big.Int) error {

	ctx := context.Background()

	publicKey := from.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return errors.New("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	nonce, err := client.PendingNonceAt(ctx, fromAddress)
	if err != nil {
		return err
	}

	gasLimit := uint64(21000)

	// gasPrice := big.NewInt(1000000000)
	// gasPrice := big.NewInt(100000000)

	gasPrice, err := client.SuggestGasPrice(ctx)

	// fmt.Println("Gas Price", gasPrice)

	if err != nil {
		return err
	}

	fmt.Println(amount)

	tx := types.NewTransaction(nonce, to, amount, gasLimit, gasPrice, nil)

	chainID := big.NewInt(1337)

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), from)
	if err != nil {
		return err

	}

	// err = client.SendTransaction(context.Background(), signedTx)
	// if err != nil {
	// 	return err
	// }

	return client.SendTransaction(ctx, signedTx)
}
