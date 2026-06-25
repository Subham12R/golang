CREATE TABLE IF NOT EXISTS movies (
    id              SERIAL  PRIMARY KEY,
    title           TEXT    NOT NULL,
    director        TEXT    NOT NULL,
    release_year    INTEGER NOT NULL
);