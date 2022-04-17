SELECT * FROM CUSTOMERS


CREATE TABLE endereco (
        id_endereco INTEGER PRIMARY KEY AUTOINCREMENT,
        rua TEXT NOT NULL,
        numero INTEGER NOT NULL,
        cep TEXT,
        cidade TEXT DEFAULT "Campo Mourao" NOT NULL,
        estado TEXT DEFAULT "Parana" NOT NULL     
        )
             

drop table endereco

INSERT INTO endereco (rua, numero, cep, cidade, estado)
 VALUES ("Praça Ovídio Ribeiro de Abreu", 980, "28220-970", "Barcelos", "RJ")


SELECT * FROM endereco

CREATE TABLE empresa (
        id_empresa INTEGER PRIMARY KEY AUTOINCREMENT,
        id_endereco INTEGER NOT NULL,
        nome TEXT NOT NULL,
        cnpj TEXT NOT NULL,
        telefone TEXT NOT NULL,
        FOREIGN KEY (id_endereco) REFERENCES endereco (id_endereco) 
        )
INSERT INTO empresa (id_endereco, nome, cnpj, telefone)
 VALUES (1, "Empresa teste", "64.994.384/0001-98", "(22) 98604-6848" )

select * from empresa

select e.nome, e.cnpj, en.cidade
from empresa as e
join endereco as en on en.id_endereco = e.id_endereco
where id_empresa = 1


CREATE TABLE cliente (
        id_cliente INTEGER PRIMARY KEY AUTOINCREMENT,
        id_endereco INTEGER NOT NULL,
        nome TEXT NOT NULL,
        cpf TEXT NOT NULL,
        telefone TEXT NOT NULL,
        FOREIGN KEY (id_endereco) REFERENCES endereco (id_endereco) 
        )
             


CREATE TABLE pagamento (
        id_pagamento INTEGER PRIMARY KEY AUTOINCREMENT,
        forma_pagamento TEXT NOT NULL
        )


CREATE TABLE nota (
	     id_nota INTEGER PRIMARY KEY AUTOINCREMENT,
             id_pagamento INTEGER NOT NULL,
             id_cliente INTEGER NOT NULL,
             id_endereco_entrega INTEGER NOT NULL,
             tipo_nota TEXT NOT NULL,
             data NUMERIC NOT NULL,
             valor_total REAL NOT NULL,
             desconto_total REAL NOT NULL,
             FOREIGN KEY (id_pagamento) REFERENCES pagamento (id_pagamento) 
             FOREIGN KEY (id_cliente) REFERENCES cliente (id_cliente) 
             FOREIGN KEY (id_endereco_entrega) REFERENCES endereco (id_endereco) 
             )

CREATE TABLE tipo_produto (
	        id_tipo_produto INTEGER PRIMARY KEY AUTOINCREMENT,
            nome TEXT NOT NULL,
            descricao TEXT
        )



CREATE TABLE produto (
	        id_produto INTEGER PRIMARY KEY AUTOINCREMENT,
            id_tipo INTEGER NOT NULL,
            valor_metragem REAL NOT NULL,
            espessura REAL,
            cor TEXT,
            FOREIGN KEY (id_tipo) REFERENCES tipo_produto (id_tipo_produto) 
        )

CREATE TABLE item (
	         id_item INTEGER PRIMARY KEY AUTOINCREMENT,
             id_produto INTEGER NOT NULL,
             id_nota INTEGER NOT NULL,
             quantidade INTEGER NOT NULL,
             valor INTEGER NOT NULL,
             desconto INTEGER,
             metragem_produto REAL NOT NULL,
             FOREIGN KEY (id_produto) REFERENCES produto (id_produto) 
             FOREIGN KEY (id_nota) REFERENCES nota (id_nota) 
             )
             
