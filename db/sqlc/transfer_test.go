package db

import (
	"context"
	"testing"

	"github.com/BrunoBiz/simplebank/util"
	"github.com/stretchr/testify/require"
)

func CreateRandomTransfer(t *testing.T, accountFrom Account, accountTo Account) Transfer {
	if accountFrom.ID == 0 {
		accountFrom = createRandomAccount(t)
	}
	if accountTo.ID == 0 {
		accountTo = createRandomAccount(t)
	}

	arg := CreateTransferParams{
		FromAccountID: accountFrom.ID,
		ToAccountID:   accountTo.ID,
		Amount:        util.RandomMoney(),
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, arg.FromAccountID, transfer.FromAccountID)
	require.Equal(t, arg.ToAccountID, transfer.ToAccountID)
	require.Equal(t, arg.Amount, transfer.Amount)

	return transfer
}

func TestCreateTransfer(t *testing.T) {
	CreateRandomTransfer(t, Account{}, Account{})
}

func TestGetTransfer(t *testing.T) {
	transfer1 := CreateRandomTransfer(t, Account{}, Account{})

	transfer2, err := testQueries.GetTransfer(context.Background(), transfer1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, transfer2)

	require.Equal(t, transfer1.FromAccountID, transfer2.FromAccountID)
	require.Equal(t, transfer1.ToAccountID, transfer2.ToAccountID)
	require.Equal(t, transfer1.Amount, transfer2.Amount)
}

func TestListTransfers(t *testing.T) {
	accountFrom := createRandomAccount(t)
	accountTo := createRandomAccount(t)

	for i := 0; i < 10; i++ {
		CreateRandomTransfer(t, accountFrom, accountTo)
	}

	arg := ListTransfersParams{
		FromAccountID: accountFrom.ID,
		ToAccountID:   accountTo.ID,
		Limit:         5,
		Offset:        5,
	}

	transfers, err := testQueries.ListTransfers(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, transfers)

	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)
	}

}
