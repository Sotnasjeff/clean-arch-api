CREATE TABLE orders (
    id VARCHAR(255) NOT NULL,
    price FLOAT NOT NULL,
    tax FLOAT NOT NULL,
    final_price FLOAT NOT NULL
)

INSERT INTO orders (id, price, tax, final_price) VALUES ("aaa", 25.0, 0.5, 25.5)
INSERT INTO orders (id, price, tax, final_price) VALUES ("bbb", 30.0, 0.5, 30.5)
INSERT INTO orders (id, price, tax, final_price) VALUES ("ccc", 35.0, 0.5, 35.5)
INSERT INTO orders (id, price, tax, final_price) VALUES ("ddd", 40.0, 0.5, 40.5)
INSERT INTO orders (id, price, tax, final_price) VALUES ("eee", 15.0, 0.5, 15.5)
INSERT INTO orders (id, price, tax, final_price) VALUES ("fff", 10.0, 0.5, 10.5)