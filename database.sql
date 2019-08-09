create table mine.hotData (
    id int auto_increment primary key,
    str text collate utf8mb4_unicode_ci null,
    dataType varchar(45) collate utf8mb4_bin not null,
    name varchar(45) collate utf8mb4_bin not null,
    isShow int default 0 null
) charset=utf8mb4;
create index hotData__index_key on mine.hotData (dataType);
create index hotData__index_name on mine.hotData (name);