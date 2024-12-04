-- name: GetLastSendedPoll :one
SELECT * FROM SendedPolls ORDER BY id DESC LIMIT 1;

-- name: AddSendedVote :exec
INSERT INTO SendedPolls(poll_tg_id) VALUES(?);

-- name: GetLastIdMorningPoll :one
SELECT id FROM MorningPolls ORDER BY id DESC LIMIT 1;

-- name: GetMorningPollById :one
SELECT * FROM MorningPolls WHERE id == ?;

-- name: GetOpionsForPoll :many
SELECT * FROM MorningPollsOption WHERE morning_poll_id = ?;