-- name: InsertChatThread :one
INSERT INTO chat_threads (thread_id, user_address, thread_name)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetChatThread :one
SELECT * FROM chat_threads
WHERE thread_id = $1
AND user_address = $2
AND is_deleted = FALSE;

-- name: GetChatThreads :many
SELECT thread_id, thread_name FROM chat_threads
WHERE user_address = $1
AND is_deleted = FALSE
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;
