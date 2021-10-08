package mempool

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseValidTransaction(t *testing.T) {
	tx, err := ParseTransaction("TxHash=hashy Gas=729000 FeePerGas=0.11134106816568039 Signature=signy")

	require.NoError(t, err)
	require.NotEmpty(t, tx)
	require.Equal(t, "hashy", tx.TxHash)
	require.Equal(t, 729000.0, tx.Gas)
	require.Equal(t, 0.11134106816568039, tx.FeePerGas)
	require.Equal(t, "signy", tx.Signature)
}

func TestParseInvalidValidTransactions(t *testing.T) {
	for _, rawTx := range []string{
		"",
		"TxHash=hashy Gas= FeePerGas=0.11134106816568039 ",
		"TxHash=hashy Gas= Signature=signy",
		"TxHash=hashy FeePerGas=0.11134106816568039 Signature=signy",
		"Gas=123 FeePerGas=0.11134106816568039 Signature=signy",
		"TxHash=hashy Gas= FeePerGas=0.11134106816568039 Signature=signy",
		"TxHash=hashy Gas=1 FeePerGas=0.a Signature=signy",
	} {
		tx, err := ParseTransaction(rawTx)
		require.Error(t, err)
		require.Empty(t, tx)
	}

}
