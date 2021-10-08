package mempool

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMempool(t *testing.T) {
	mp := NewMempool(3)

	require.Equal(t, 0, len(mp.transactions))
	mp.AddTransaction(&Transaction{Gas: 10, FeePerGas: 0.1})
	mp.AddTransaction(&Transaction{Gas: 30, FeePerGas: 0.1})
	mp.AddTransaction(&Transaction{Gas: 20, FeePerGas: 0.1})

	// fill the mempool
	require.Equal(t, 3, len(mp.transactions))
	require.Equal(t, Transaction{Gas: 30, FeePerGas: 0.1}, *mp.transactions[0])
	require.Equal(t, Transaction{Gas: 20, FeePerGas: 0.1}, *mp.transactions[1])
	require.Equal(t, Transaction{Gas: 10, FeePerGas: 0.1}, *mp.transactions[2])

	// drop transaction with lowest fee and insert as last one
	mp.AddTransaction(&Transaction{Gas: 15, FeePerGas: 0.1})
	require.Equal(t, 3, len(mp.transactions))
	require.Equal(t, Transaction{Gas: 30, FeePerGas: 0.1}, *mp.transactions[0])
	require.Equal(t, Transaction{Gas: 20, FeePerGas: 0.1}, *mp.transactions[1])
	require.Equal(t, Transaction{Gas: 15, FeePerGas: 0.1}, *mp.transactions[2])

	// drop transaction with lowest fee and insert as second one
	mp.AddTransaction(&Transaction{Gas: 22, FeePerGas: 0.1})
	require.Equal(t, 3, len(mp.transactions))
	require.Equal(t, Transaction{Gas: 30, FeePerGas: 0.1}, *mp.transactions[0])
	require.Equal(t, Transaction{Gas: 22, FeePerGas: 0.1}, *mp.transactions[1])
	require.Equal(t, Transaction{Gas: 20, FeePerGas: 0.1}, *mp.transactions[2])

	// drop transaction with lowest fee and insert as first one
	mp.AddTransaction(&Transaction{Gas: 40, FeePerGas: 0.1, TxHash: "a"})
	require.Equal(t, 3, len(mp.transactions))
	require.Equal(t, Transaction{Gas: 40, FeePerGas: 0.1, TxHash: "a"}, *mp.transactions[0])
	require.Equal(t, Transaction{Gas: 30, FeePerGas: 0.1}, *mp.transactions[1])
	require.Equal(t, Transaction{Gas: 22, FeePerGas: 0.1}, *mp.transactions[2])

	// insert transacation with duplicated fee
	mp.AddTransaction(&Transaction{Gas: 40, FeePerGas: 0.1, TxHash: "b"})
	require.Equal(t, 3, len(mp.transactions))
	require.Equal(t, Transaction{Gas: 40, FeePerGas: 0.1, TxHash: "a"}, *mp.transactions[0])
	require.Equal(t, Transaction{Gas: 40, FeePerGas: 0.1, TxHash: "b"}, *mp.transactions[1])
	require.Equal(t, Transaction{Gas: 30, FeePerGas: 0.1}, *mp.transactions[2])

	// all transactions with same fee
	mp.AddTransaction(&Transaction{Gas: 40, FeePerGas: 0.1, TxHash: "c"})
	require.Equal(t, 3, len(mp.transactions))
	require.Equal(t, Transaction{Gas: 40, FeePerGas: 0.1, TxHash: "a"}, *mp.transactions[0])
	require.Equal(t, Transaction{Gas: 40, FeePerGas: 0.1, TxHash: "c"}, *mp.transactions[1])
	require.Equal(t, Transaction{Gas: 40, FeePerGas: 0.1, TxHash: "b"}, *mp.transactions[2])

	// all transactions with same fee and it should drop the last one "a"
	mp.AddTransaction(&Transaction{Gas: 40, FeePerGas: 0.1, TxHash: "d"})
	require.Equal(t, 3, len(mp.transactions))
	require.Equal(t, Transaction{Gas: 40, FeePerGas: 0.1, TxHash: "a"}, *mp.transactions[0])
	require.Equal(t, Transaction{Gas: 40, FeePerGas: 0.1, TxHash: "d"}, *mp.transactions[1])
	require.Equal(t, Transaction{Gas: 40, FeePerGas: 0.1, TxHash: "c"}, *mp.transactions[2])
}
