DROP TABLE Board;
DROP TABLE Post;

CREATE TABLE Board (
Id SERIAL,
longname TEXT,
shortname TEXT
);

CREATE TABLE Post (
Id SERIAL,
content TEXT
);
