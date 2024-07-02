-- +goose Up
-- +goose StatementBegin
INSERT INTO genres (name)
VALUES ('Fiction'),
       ('Non-Fiction'),
       ('Mystery'),
       ('Fantasy'),
       ('Science Fiction'),
       ('Romance'),
       ('Thriller'),
       ('Horror'),
       ('Historical'),
       ('Young Adult'),
       ('Biography'),
       ('Self-Help'),
       ('Graphic Novel'),
       ('Adventure'),
       ('Crime'),
       ('Drama'),
       ('Poetry'),
       ('Satire'),
       ('Children'),
       ('Western');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
TRUNCATE genres CASCADE;
-- +goose StatementEnd
