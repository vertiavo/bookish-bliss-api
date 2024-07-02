-- +goose Up
-- +goose StatementBegin
INSERT INTO authors (first_name, last_name)
VALUES ('Jane', 'Austen'),
       ('Mark', 'Twain'),
       ('Charles', 'Dickens'),
       ('Leo', 'Tolstoy'),
       ('George', 'Orwell'),
       ('Agatha', 'Christie'),
       ('J.K.', 'Rowling'),
       ('F. Scott', 'Fitzgerald'),
       ('Ernest', 'Hemingway'),
       ('J.R.R.', 'Tolkien');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
TRUNCATE authors CASCADE;
-- +goose StatementEnd
