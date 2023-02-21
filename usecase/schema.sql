-- SQLite
-- SQLite
CREATE TABLE address (
	id_address INTEGER PRIMARY KEY AUTOINCREMENT,
    id_customer INTEGER NOT NULL,
    state TEXT NOT NULL, 
    city TEXT NOT NULL,
	street TEXT NOT NULL,
	number INTEGER NOT NULL,
	zip_number TEXT NOT NULL
);


CREATE TABLE customer (
	id_customer INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL, 
	phone_number TEXT NOT NULL,
    state TEXT NOT NULL, 
    city TEXT NOT NULL,
	street TEXT NOT NULL,
	number INTEGER NOT NULL,
	zip_number TEXT NOT NULL
);


CREATE TABLE payment (
	id_payment INTEGER PRIMARY KEY AUTOINCREMENT,
    payment_form TEXT NOT NULL

);

CREATE TABLE product_type (
    id_type INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    description TEXT

);

CREATE TABLE product (
	id_product INTEGER PRIMARY KEY AUTOINCREMENT,
    id_type INTEGER NOT NULL,
    value_per_meter DOUBLE NOT NULL, 
    total_value DOUBLE NOT NULL,
    thickness DOUBLE NOT NULL,
    color TEXT NOT NULL

);


CREATE TABLE item (
	id_item INTEGER PRIMARY KEY AUTOINCREMENT,
    id_invoice INTEGER NOT NULL,
    id_product INTEGER NOT NULL, 
    amount DOUBLE NOT NULL,
	quantity INTEGER NOT NULL,
	discount DOUBLE NOT NULL,
	metreage DOUBLE NOT NULL
);


insert into item (id_invoice, id_product, amount, quantity, discount, metreage) VALUES (1, 1, 10.25, 2, 0, 2)

CREATE TABLE invoice (
	id_invoice INTEGER PRIMARY KEY AUTOINCREMENT,
    id_payment INTEGER DEFAULT 0 NOT NULL,
    id_customer INTEGER DEFAULT 0 NOT NULL, 
    id_delivery_address INTEGER DEFAULT 0 NOT NULL,
	invoice_type  TEXT DEFAULT "n/a" NOT NULL,
	date text DEFAULT "n/a" NOT NULL,
	total_amount DOUBLE DEFAULT 0 NOT NULL,
    total_discount DOUBLE DEFAULT 0 NOT NULL
);





insert into customer(name, cpf, phone_number) VALUES ("Amanda", "019283983", "978866458")



INSERT INTO address (state, city, street, number, zip_number) VALUES ("AP", "Cutias", "Rua Manoel Raimundo, s/n", 992, "68973-970")
INSERT INTO address (state, city, street, number, zip_number) 
VALUES ("SP", "São Paulo", "Travessa Flor do Agreste", 421, "08235-085")
INSERT INTO address (state, city, street, number, zip_number) VALUES ("SP", "São Paulo", "Rua Manoel Raimundo, s/n", 992, "68973-970")
INSERT INTO address (state, city, street, number, zip_number) VALUES ("SP", "Cutias", "Rua Manoel Raimundo, s/n", 992, "68973-970")
INSERT INTO address (state, city, street, number, zip_number) VALUES ("SP", "Cutias", "Rua Manoel Raimundo, s/n", 992, "68973-970")
INSERT INTO address (state, city, street, number, zip_number) VALUES ("SP", "Cutias", "Rua Manoel Raimundo, s/n", 992, "68973-970")