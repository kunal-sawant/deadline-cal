-- name: ListTasks :many
SELECT * FROM Tasks;

-- name: CreateTask :one
INSERT INTO Tasks (TASK_NAME,START_DATE,END_DATE) VALUES
(?, ?, ?) RETURNING *;

-- name: DeleteTask :exec
DELETE FROM Tasks
WHERE ID = ?;