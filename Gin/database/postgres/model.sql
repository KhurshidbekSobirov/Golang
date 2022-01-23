CREATE TABLE customers(
    Id serial PRIMARY KEY,
    Name VARCHAR(30),
    Password VARCHAR(20),
    Money INT ,
    Products JSONB
);

CREATE TABLE products(
    Id serial PRIMARY KEY,
    NAME VARCHAR(30),
    Price INT NOT NULL
);

CREATE TABLE customer_products(
    customer_id int REFERENCES customers(id),
    product_id int REFERENCES products(id)
);

INSERT INTO products(name,price)
VALUES ('kartoshka',4000), 
('piyoz',2000),
('sabzi',3000),
('pepsi',10000),
('coco cola',11000),
('non',2500),
('guruch',11000);