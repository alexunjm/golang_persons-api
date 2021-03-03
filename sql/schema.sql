CREATE TABLE person
(
    id        VARCHAR(64) NOT NULL,
    firstname VARCHAR(64) NOT NULL,
    lastname  VARCHAR(64) NOT NULL,
    age       INT NOT NULL,

    PRIMARY KEY (id)

) CHARACTER SET utf8mb4
  COLLATE utf8mb4_bin;
