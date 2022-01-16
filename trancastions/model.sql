CREATE TABLE bookShop(
    id serial PRIMARY KEY,
    name text,
    capasity INTEGER,
);

CREATE TABLE section(
    name JSONB,
    name text
);


CREATE TABLE outlets(
    id serial PRIMARY KEY,
    name text
);

CREATE TABLE aboutBookshop(
    shop_id int REFERENCES bookShop(id),
    section_id int REFERENCES section(id),
    outlets_id int REFERENCES outlets(id)
);

