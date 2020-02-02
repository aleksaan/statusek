CREATE ROLE statusek_admin NOSUPERUSER NOCREATEDB NOCREATEROLE NOINHERIT LOGIN PASSWORD 'st@tusek_@dmin';

set role statusek_admin;


-- DROP SCHEMA statusek;

CREATE SCHEMA statuses AUTHORIZATION statusek_admin;

-------------------------------------------------------------------
--Object 
-------------------------------------------------------------------
drop  table statuses.objects;

create table statuses.objects 
(
object_id serial,
object_name text
);


-------------------------------------------------------------------
--Statuses 
-------------------------------------------------------------------
create table statuses.statuses
(
object_id references on statuses.objects (object_id)
status_id serial,

)