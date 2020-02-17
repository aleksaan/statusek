--CREATE ROLE statusek_admin NOSUPERUSER NOCREATEDB NOCREATEROLE NOINHERIT LOGIN PASSWORD 'st@tusek_@dmin';

--set role statusek_admin;


-- DROP SCHEMA statusek;

--CREATE SCHEMA statuses AUTHORIZATION statusek_admin;

--set role postgres;
--create extension if not exists "uuid-ossp";
--set role statusek_admin;

-------------------------------------------------------------------
--Object 
-------------------------------------------------------------------
drop table if exists statuses.objects cascade;

create table statuses.objects 
(
object_id serial primary key,
object_name text not null
);

insert into statuses.objects (object_name)
values 	('ROMB_TASK_1'),
		('ROMB_TASK_2'),
		('ROMB_TASK_3');

select * from statuses.objects;

-------------------------------------------------------------------
--Statuses 
-------------------------------------------------------------------
drop table if exists statuses.statuses cascade;

create table statuses.statuses
(
object_id int references statuses.objects (object_id),
status_id serial primary key,
status_name text not null,
status_is_mandatory boolean not null,
status_desc text not null,
unique (object_id, status_name)
);

insert into statuses.statuses (object_id, status_name, status_desc, status_is_mandatory)
values 
(1, 'RUN',  '1 - mandatory', true),
(1, 'NODE_2',  '2 - mandatory', true),
(1, 'NODE_3',  '3 - optional', false),
(1, 'NODE_4',  '4 - mandatory', true),
(1, 'NODE_5',  '5 - mandatory', true),
(1, 'NODE_6',  '6 - mandatory', true)
;

--insert into statuses.statuses (object_id, status_name, status_desc, status_is_optional)
--values 
--(2, 'NODE_1',   '1 - mandatory', false),
--(2, 'NODE_2',   '2 - mandatory', false),
--(2, 'NODE_3',   '3 - mandatory', false),
--(2, 'NODE_4',   '4 - mandatory', false)
--;
--
--
--insert into statuses.statuses (object_id, status_name, status_desc, status_is_optional)
--values 
--(3, 'NODE_1',  '1 - mandatory', false),
--(3, 'NODE_2',  '2 - mandatory', false),
--(3, 'NODE_3',  '3 - optional', true),
--(3, 'NODE_4',  '4 - mandatory', false)
--;

select * from statuses.statuses;


-------------------------------------------------------------------
--Workflow 
-------------------------------------------------------------------

drop table if exists statuses.workflows cascade;

create table statuses.workflows 
(
workflow_id serial primary key,
status_id_prev int references statuses.statuses (status_id) not null,
status_id_next int references statuses.statuses (status_id) not null,
unique (status_id_prev, status_id_next)
);

insert into statuses.workflows (status_id_prev, status_id_next)
values
(1, 2),
(1, 3),
(2, 4),
(3, 5),
(4, 5),
(4, 6);


select * from statuses.workflows;


drop view if exists statuses.v_workflows;

create view statuses.v_workflows as
	with recursive cte as (
	   select distinct 
	   		s0.object_id,
	   		null::int as status_id_prev, 
	   		null::text as status_name_prev, 
	   		null::bool as status_is_mandatory_prev,
	   		s0.status_id as status_id_next, 
	   		s0.status_name as status_name_next, 
	   		s0.status_is_mandatory as status_is_mandatory_next,
	   		1 as level
	   from   statuses.statuses s0
	   where  s0.status_id not in (select status_id_next from statuses.workflows)
	
	   union  all
	   select distinct 
	   		s2.object_id,
	   		s1.status_id as status_id_prev, 
	   		s1.status_name as status_name_prev, 
	   		s1.status_is_mandatory as status_is_mandatory_prev,
	   		w1.status_id_next, 
	   		s2.status_name as status_name_next,
	   		s2.status_is_mandatory as status_is_mandatory_next,
	   		c.level + 1
	   from  cte c 
	   		join  statuses.statuses s1 
	   			on s1.status_id = c.status_id_next
	   		left join statuses.workflows w1
	   			on s1.status_id = w1.status_id_prev
	   		left join statuses.statuses s2
	   			on w1.status_id_next = s2.status_id 
	   )
	select 
		object_id,
		status_id_prev,
		status_name_prev,
		status_is_mandatory_prev,
		status_id_next as status_id,
		status_name_next as status_name,
		status_is_mandatory_next as status_is_mandatory,
		level as status_level
	from cte
	where 
		status_id_next is not null
	order by status_level, status_id_next;


select * from statuses.v_workflows where object_id=1;

--create view statuses.v_workflows as
--select w.workflow_id,
--	o.object_id,
--	o.object_name,
--	w.status_id_prev,
--	s1.status_name as status_name_prev,
--	s1.status_is_optional as status_is_optional_prev,
--	w.status_id_next,
--	s2.status_name as status_name_next,
--	s2.status_is_optional as status_is_optional_next
--from statuses.workflows w, statuses.statuses s1, statuses.statuses s2, statuses.objects o
--where w.status_id_prev = s1.status_id and w.status_id_next =s2.status_id and s1.object_id = o.object_id;



--drop view if exists statuses.v_statuses;
--create view statuses.v_statuses as 
--select distinct 
--	s.object_id , 
--	--предыдущий статус
--	w.status_id_prev,
--	w.status_name_prev,
--	w.status_is_optional_prev,
--	--нужный статус
--	s.status_id,
--	s.status_name,
--	s.status_desc,
--	--следующий статус
--	w2.status_id_next,
--	w2.status_name_next,
--	w2.status_is_optional_next,
--	case when w.status_id_next is null
--		then true
--		else false
--	end as status_is_first,
--	case when w2.status_id_prev is null
--		then true
--		else false
--	end as status_is_last
--from
--	statuses.statuses s 
--		left join statuses.v_workflows w
--			on s.status_id=w.status_id_next
--		left join statuses.v_workflows w2
--			on s.status_id=w2.status_id_prev
--order by s.status_id ;


select * from statuses.v_statuses




-------------------------------------------------------------------
--Instances 
-------------------------------------------------------------------

drop table if exists statuses.instances cascade;

create table statuses.instances 
(
instance_id bigserial primary key,
instance_token uuid DEFAULT uuid_generate_v4(),
object_id int references statuses.objects (object_id) not null,
instance_creation_dt timestamp default pg_catalog.clock_timestamp() 
);

select * from statuses.instances;
------------------

-------------------------------------------------------------------
--Events 
-------------------------------------------------------------------

drop table if exists statuses.events cascade;

create table statuses.events 
(
event_id bigserial primary key,
instance_id bigint references statuses.instances (instance_id) not null,
status_id int references statuses.statuses (status_id) not null,
status_asterisk_text text,
event_creation_dt timestamp default pg_catalog.clock_timestamp()
);

select * from statuses.events;

--insert into statuses.events (instance_id, status_id)
--values(1,1);

drop view if exists statuses.v_events;

create view statuses.v_events as
select  
		o.object_id,
		o.object_name,
		s.status_name,
		i.instance_token,
		i.instance_creation_dt,
		i.instance_id,
		e.event_id,
		s.status_id,
		
from 	statuses.events e,
		statuses.statuses s,
	    statuses.instances i,
	    statuses.objects o--,
	    --statuses.v_workflows w 
where e.instance_id = i.instance_id and 
	  e.status_id = s.status_id and 
	  o.object_id = i.object_id and 
	  --o.object_id = w.object_id and 
	  --s.status_id  = w.status_id
	  ;


select * from statuses.v_events;

