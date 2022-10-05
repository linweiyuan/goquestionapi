-- name: GetQuestions :many

SELECT q.id, q.title, q.answer FROM question q ORDER BY q.id;

-- name: GetQuestionScoreByID :one

SELECT (
        q.score ->> sqlc.arg(answer) :: TEXT
    ) :: INT AS score
FROM question q
WHERE q.id = sqlc.arg(id);