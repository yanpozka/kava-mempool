package mempool

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const (
	txHashIx = iota
	gasIx
	feePerGasIx
	signatureIx
)

var (
	ErrInvalidTransaction = errors.New("invalid transaction")
)

type Transaction struct {
	FeePerGas float64
	Gas       float64
	TxHash    string
	Signature string
}

func (tx *Transaction) String() string {
	return fmt.Sprintf("TxHash=%s Gas=%d FeePerGas=%f Signature=%s", tx.TxHash, int64(tx.Gas), tx.FeePerGas, tx.Signature)
}

func (tx *Transaction) Priority() float64 {
	return tx.Gas * tx.FeePerGas
}

func ParseTransaction(line string) (*Transaction, error) {
	parts := strings.Split(line, " ")
	if len(parts) != 4 {
		return nil, ErrInvalidTransaction
	}

	tx := new(Transaction)

	// TxHash=7F5633784780A119E939A05E3EBF27CD9D8FD7C00EBBAC3AA5AE8BA7DD09B98E
	txHashParts := strings.Split(parts[txHashIx], "=")
	if len(txHashParts) != 2 {
		return nil, ErrInvalidTransaction
	}
	tx.TxHash = txHashParts[1]

	// Gas=808000
	gasParts := strings.Split(parts[gasIx], "=")
	if len(gasParts) != 2 {
		return nil, ErrInvalidTransaction
	}
	gas, err := strconv.ParseFloat(gasParts[1], 64)
	if err != nil {
		return nil, fmt.Errorf("invalid gas value, error: %w", err)
	}
	tx.Gas = gas

	// FeePerGas=0.6498944377795232
	feeParts := strings.Split(parts[feePerGasIx], "=")
	if len(gasParts) != 2 {
		return nil, ErrInvalidTransaction
	}
	fee, err := strconv.ParseFloat(feeParts[1], 64)
	if err != nil {
		return nil, fmt.Errorf("invalid fee per gas value, error: %w", err)
	}
	tx.FeePerGas = fee

	// Signature=13B3186998F58899065989632AF3...
	sigParts := strings.Split(parts[signatureIx], "=")
	if len(txHashParts) != 2 {
		return nil, ErrInvalidTransaction
	}
	tx.Signature = sigParts[1]

	return tx, nil
}
