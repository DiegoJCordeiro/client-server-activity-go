-- name: QueryAllQuotation :many
SELECT * FROM Quotations ORDER BY created_at;

-- name: QueryQuotationById :one
SELECT * FROM Quotations WHERE id = ?;

-- name: InsertQuotation :exec
INSERT INTO Quotations(id, bid, ask, timestamp, created_at) Values (?, ?, ?, ?, ?);

-- name: UpdateQuotation :exec
UPDATE Quotations SET bid = ? and ask = ? and updated_at = ? WHERE id = ?;

-- name: DeleteQuotation :exec
UPDATE Quotations SET deleted_at = ? WHERE id = ?;