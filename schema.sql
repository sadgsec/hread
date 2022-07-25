DROP TABLE Board;
DROP TABLE Post;

CREATE TABLE Board (
Id serial primary key,
longname text,
shortname text
);

CREATE TABLE Post (
Id serial,
boardid serial references Board(Id),
content text
);
