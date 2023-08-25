CREATE TABLE routes (
     id CHAR(36) DEFAULT (UUID()),
     name VARCHAR(255),
     source JSON,
     destination JSON,
     PRIMARY KEY (id)
 );