create table posts (
    id  serial primary key,
    content text,
    author varchar(255)
);

/* 
To execute the script, run this command on the console
psql -U gwp -f setup.sql -d gwp
*/