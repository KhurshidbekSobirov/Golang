
CREATE TABLE contacts (
    id uuid DEFAULT uuid_generate_v4(),
    assignee VARCHAR(50) NOT NULL,
    title VARCHAR(50) NOT NULL,
    deadline timestamp,
    status VARCHAR(20),
    created_at timestamp,
    updated_at timestamp,
    deleted_at timestamp,
    PRIMARY KEY (id)
);
