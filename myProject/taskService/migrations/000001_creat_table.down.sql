CREATE TABLE IF NOT EXISTS tasks(
    Id text,
    Name text NOT NULL,
    Deadline timestamp NOT NULL,
    Summary text,
    AssigneId text NOT NULL,
    Status text NOT NULL,
    Create_at timestamp NOT NULL,
    Update_at timestamp,
    Delete_at timestamp

);
