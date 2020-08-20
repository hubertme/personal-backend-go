# Created on 20 Aug 2020

CREATE TABLE messages (
    id INT NOT NULL AUTO_INCREMENT,
    sender_name VARCHAR(64) NOT NULL,
    sender_email VARCHAR(64) NOT NULL,
    subject varchar(512) NOT NULL,
    content varchar(2048) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    PRIMARY KEY (id)
);