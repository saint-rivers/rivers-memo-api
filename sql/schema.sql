create table tag ( tag_name varchar(24) not null primary key );

create table note (
    id bigint not null primary key,
    link varchar not null,
    title varchar,
    description varchar,
    image_url varchar
);

create table note_tag (
    note bigint,
    tag varchar(24),
    Foreign Key (note) REFERENCES note (id),
    Foreign Key (tag) REFERENCES tag (tag_name),
);