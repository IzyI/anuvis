goose --dir migrations  create user sql
goose --dir migrations  up
goose --dir migrations  down-to 0
