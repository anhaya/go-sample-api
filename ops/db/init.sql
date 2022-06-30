create database if not exists goapi;
use goapi;

create table if not exists account (
    id varchar(50) NOT NULL, 
    document_number varchar(15) NOT NULL,
    balance DECIMAL(64,8) NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE=INNODB;

CREATE TABLE transaction (
    id INT NOT NULL AUTO_INCREMENT,
    account_id varchar(50) NOT NULL,
    operation_type_id INT NOT NULL,
    amount DECIMAL(64,8) NOT NULL,
    event_date DATETIME NOT NULL,
    PRIMARY KEY (`id`),
    FOREIGN KEY (account_id) REFERENCES account(id)
) ENGINE=INNODB;