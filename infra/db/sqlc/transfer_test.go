package sqlc

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/ricassiocosta/simplebank/util"
	"github.com/stretchr/testify/require"
)

func createRandomTransfer(t *testing.T, fromAccountID, toAccountID uuid.UUID) Transfer {
	arg := CreateTransferParams{
		FromAccountID: fromAccountID,
		ToAccountID:   toAccountID,
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
	originAcc := createRandomAccount(t)
	destinyAcc := createRandomAccount(t)

	createRandomTransfer(t, originAcc.ID, destinyAcc.ID)
}

func TestGetTransfer(t *testing.T) {
	originAcc := createRandomAccount(t)
	destinyAcc := createRandomAccount(t)

	transfer := createRandomTransfer(t, originAcc.ID, destinyAcc.ID)

	transferFromGet, err := testQueries.GetTransfer(context.Background(), transfer.ID)
	require.NoError(t, err)
	require.NotEmpty(t, transferFromGet)

	require.Equal(t, transfer.FromAccountID, transferFromGet.FromAccountID)
	require.Equal(t, transfer.ToAccountID, transferFromGet.ToAccountID)
	require.Equal(t, transfer.Amount, transferFromGet.Amount)
	require.WithinDuration(t, transfer.CreatedAt, transferFromGet.CreatedAt, time.Second)
}

func TestUpdateTransfer(t *testing.T) {
	originAcc := createRandomAccount(t)
	destinyAcc := createRandomAccount(t)

	transfer := createRandomTransfer(t, originAcc.ID, destinyAcc.ID)

	arg := UpdateTransferParams{
		ID:     transfer.ID,
		Amount: util.RandomMoney(),
	}

	updatedTransfer, err := testQueries.UpdateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, updatedTransfer)

	require.Equal(t, transfer.ID, updatedTransfer.ID)
	require.Equal(t, transfer.FromAccountID, updatedTransfer.FromAccountID)
	require.Equal(t, transfer.ToAccountID, updatedTransfer.ToAccountID)
	require.Equal(t, arg.Amount, updatedTransfer.Amount)
	require.WithinDuration(t, transfer.CreatedAt, updatedTransfer.CreatedAt, time.Second)

}

func TestDeleteTransfer(t *testing.T) {
	originAcc := createRandomAccount(t)
	destinyAcc := createRandomAccount(t)

	transfer := createRandomTransfer(t, originAcc.ID, destinyAcc.ID)

	err := testQueries.DeleteTransfer(context.Background(), transfer.ID)
	require.NoError(t, err)

	transferFromGet, err := testQueries.GetTransfer(context.Background(), transfer.ID)
	require.Error(t, err)
	require.ErrorIs(t, err, sql.ErrNoRows)
	require.Empty(t, transferFromGet)
}

func TestListTransfersBetweenAccounts(t *testing.T) {
	originAcc := createRandomAccount(t)
	destinyAcc := createRandomAccount(t)

	for i := 0; i < 10; i++ {
		createRandomTransfer(t, originAcc.ID, destinyAcc.ID)
	}

	arg := ListTransferBetweenAccountsParams{
		FromAccountID: originAcc.ID,
		ToAccountID:   destinyAcc.ID,
		Limit:         5,
		Offset:        5,
	}

	transfers, err := testQueries.ListTransferBetweenAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfers)

	require.Len(t, transfers, 5)

	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)
	}
}

func TestListTransferFromAccount(t *testing.T) {
	originAcc := createRandomAccount(t)
	destinyAcc := createRandomAccount(t)

	for i := 0; i < 10; i++ {
		createRandomTransfer(t, originAcc.ID, destinyAcc.ID)
	}

	arg := ListTransferFromAccountParams{
		FromAccountID: originAcc.ID,
		Limit:         5,
		Offset:        5,
	}

	transfers, err := testQueries.ListTransferFromAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfers)

	require.Len(t, transfers, 5)

	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)
	}
}

func TestListTransferToAccount(t *testing.T) {
	originAcc := createRandomAccount(t)
	destinyAcc := createRandomAccount(t)

	for i := 0; i < 10; i++ {
		createRandomTransfer(t, originAcc.ID, destinyAcc.ID)
	}

	arg := ListTransferToAccountParams{
		ToAccountID: destinyAcc.ID,
		Limit:       5,
		Offset:      5,
	}

	transfers, err := testQueries.ListTransferToAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfers)

	require.Len(t, transfers, 5)

	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)
	}
}
