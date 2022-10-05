-- name: GetQuestions :many

SELECT q.id, q.title, q.answer FROM question q ORDER BY q.id;