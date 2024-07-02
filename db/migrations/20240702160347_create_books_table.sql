-- +goose Up
-- +goose StatementBegin
CREATE TABLE books
(
    id         SERIAL PRIMARY KEY,
    title      VARCHAR(255) NOT NULL,
    author_id  INT          NOT NULL,
    genre_id   INT          NOT NULL,
    year       INT          NOT NULL,
    CONSTRAINT fk_author FOREIGN KEY (author_id)
        REFERENCES authors (id)
        ON DELETE CASCADE,
    CONSTRAINT fk_genre
        FOREIGN KEY (genre_id)
            REFERENCES genres (id)
            ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE books;
-- +goose StatementEnd
