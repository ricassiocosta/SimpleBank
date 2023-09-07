-- name: CreateTransfer :one
INSERT INTO transfers (
    from_account_id,
    to_account_id,
    amount
) VALUES (
    $1, $2, $3
)
RETURNING *;

-- name: GetTransfer :one
SELECT * FROM transfers
WHERE id = $1 LIMIT 1;

-- name: ListTransferFromAccount :many
SELECT * FROM transfers
WHERE from_account_id = $1
ORDER BY id
LIMIT $2
OFFSET $3;

-- name: ListTransferToAccount :many
SELECT * FROM transfers
WHERE to_account_id = $1
ORDER BY id
LIMIT $2
OFFSET $3;

-- name: ListTransferBetweenAccounts :many
SELECT * FROM transfers
WHERE from_account_id = $1
AND to_account_id = $2
ORDER BY id
LIMIT $3
OFFSET $4;

-- name: UpdateTransfer :one
UPDATE transfers
SET amount = $2
WHERE id = $1
RETURNING *;

-- name: DeleteTransfer :exec
DELETE FROM transfers WHERE id = $1;