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
values 	
	('1-POINT TASK'),
	('2-POINT LINE TASK'),
	('3-POINT LINE TASK'),
	('3-POINT TRIANGLE TASK'),
	('3-POINT BACK TRIANGLE TASK'),
	('4-POINT ROMB TASK')
;

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
(1, 'FINISHED', 'Task is finished', true),
(2, 'RUN',  'Task is running', true),
(2, 'FINISHED',  'Task is finished', true),
(3, 'RUN',  'Task is running', true),
(3, 'IN THE MIDDLE',  'In the middle of the task', true),
(3, 'FINISHED',  'Task is finished', true),
(4, 'RUN',  'Task is running', true),
(4, 'FINISHED_SUCCESS',  'Task is finished succesfully', false),
(4, 'FINISHED_ERROR',  'Task is finished with errors', false),
(5, 'NODE 1 RUN',  'Node 1 is running', false),
(5, 'NODE 2 RUN',  'Node 2 is running', false),
(5, 'FINISHED',  'Task is finished with errors', true),
(6, 'START',  'Task is running', true),
(6, 'NODE 1 RUN',  'Node 1 is running', false),
(6, 'NODE 2 RUN',  'Node 2 is running', false),
(6, 'FINISHED',  'Finished', true)
;




select * from statuses.statuses where object_id=2;
select * from statuses.instances where instance_token ='4354ba9c-c112-4e45-8791-dbeff78480a1' where instance_id=5;
select * from statuses.workflows w where status_id_next = 3
select * from statuses.v_last_statuses w
select * from statuses.events e where instance_id=6
select * from  statuses.v_last_statuses WHERE object_id = 2

SELECT * FROM "statuses"."statuses"  WHERE (status_name::text = 'RUN' and object_id = 2) LIMIT 1 
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
(2, 3);


select * from statuses.workflows;

--drop view if exists statuses.v_workflows;
--create view statuses.v_workflows as
--	select w.*, s.object_id 
--		from statuses.workflows w 
--			left join statuses.statuses s
--				on w.status_id_next = s.status_id;
--
--select * from statuses.v_workflows where object_id=2;			
			

drop view if exists statuses.v_last_statuses;

create view statuses.v_last_statuses as
	select s.* 
	from statuses.statuses s 
	where s.status_id  not in (select w.status_id_prev  from statuses.workflows w);
			
--select * from statuses.v_last_statuses where object_id=2;					
			

--drop view if exists statuses.v_workflows;

--create view statuses.v_workflows as
--	with recursive cte as (
--	   select distinct 
--	   		s0.object_id,
--	   		null::int as status_id_prev, 
--	   		null::text as status_name_prev, 
--	   		null::bool as status_is_mandatory_prev,
--	   		s0.status_id as status_id_next, 
--	   		s0.status_name as status_name_next, 
--	   		s0.status_is_mandatory as status_is_mandatory_next,
--	   		1 as level
--	   from   statuses.statuses s0
--	   where  s0.status_id not in (select status_id_next from statuses.workflows)
--	
--	   union  all
--	   select distinct 
--	   		s2.object_id,
--	   		s1.status_id as status_id_prev, 
--	   		s1.status_name as status_name_prev, 
--	   		s1.status_is_mandatory as status_is_mandatory_prev,
--	   		w1.status_id_next, 
--	   		s2.status_name as status_name_next,
--	   		s2.status_is_mandatory as status_is_mandatory_next,
--	   		c.level + 1
--	   from  cte c 
--	   		join  statuses.statuses s1 
--	   			on s1.status_id = c.status_id_next
--	   		left join statuses.workflows w1
--	   			on s1.status_id = w1.status_id_prev
--	   		left join statuses.statuses s2
--	   			on w1.status_id_next = s2.status_id 
--	   )
--	select 
--		object_id,
--		status_id_prev,
--		status_name_prev,
--		status_is_mandatory_prev,
--		status_id_next as status_id,
--		status_name_next as status_name,
--		status_is_mandatory_next as status_is_mandatory,
--		level as status_level
--	from cte c 
--	--where 
--	--	status_id_next is not null
--	order by status_level, status_id_next;
--
--select * 
--from statuses.v_workflows w1 
--where w1.status_id not in (select status_id_prev as status_id from statuses.v_workflows w2 where w1.object_id = w2.object_id )

select * from statuses.v_workflows where object_id=2;

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


--select * from statuses.v_statuses




-------------------------------------------------------------------
--Instances 
-------------------------------------------------------------------

drop table if exists statuses.instances cascade;

create table statuses.instances 
(
instance_id bigserial primary key,
instance_token text not null,
instance_timeout int not null,
object_id int references statuses.objects (object_id) not null,
instance_creation_dt timestamp with time zone default pg_catalog.clock_timestamp() ,
check(instance_timeout > 0)
);

delete from statuses.instances; 
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
event_creation_dt timestamp with time zone default pg_catalog.clock_timestamp()
);

delete from statuses.events;
select * from statuses.events;

--insert into statuses.events (instance_id, status_id)
--values(1,1);
--
--drop view if exists statuses.v_events;
--
--create view statuses.v_events as
--select  
--		o.object_id,
--		o.object_name,
--		s.status_name,
--		i.instance_token,
--		i.instance_creation_dt,
--		i.instance_id,
--		e.event_id,
--		s.status_id,
--		
--from 	statuses.events e,
--		statuses.statuses s,
--	    statuses.instances i,
--	    statuses.objects o--,
--	    --statuses.v_workflows w 
--where e.instance_id = i.instance_id and 
--	  e.status_id = s.status_id and 
--	  o.object_id = i.object_id and 
--	  --o.object_id = w.object_id and 
--	  --s.status_id  = w.status_id
--	  ;
--
--
--select * from statuses.v_events;

--------------------------------------------------------------------------------------------------------
--------------------------------------------------------------------------------------------------------
--version 1.2
--25/03/2020
--------------------------------------------------------------------------------------------------------
--------------------------------------------------------------------------------------------------------

select * from statuses.statuses s 
--alter table statuses.statuses drop column status_is_stopstatus;
alter table statuses.statuses add column status_is_stopstatus bool default false;

--alter table statuses.statuses add constraint con1 check ((status_is_stopstatus = true and status_is_mandatory=true) or status_is_stopstatus is false);


--test for stop-statuses
insert into statuses.objects (object_name)
values 	
	('EXAMPLE_3')
	returning *;
	
--7

insert into statuses.statuses (object_id, status_name, status_desc, status_is_mandatory, status_is_stopstatus)
values 
(7, 'CALLED', 'Service called', true, false),
(7, 'WRONG PARAMS',  'Wrong parameters', true, true),
(7, 'RUNNED',  'Task runned', true, false),
(7, 'FAILED',  'Task failed', true, true),
(7, 'START SAVING',  'Saving started', true, false),
(7, 'RESULTS SAVED',  'Results saved', true, false),
insert into statuses.statuses (object_id, status_name, status_desc, status_is_mandatory, status_is_stopstatus)
values 
(7, 'OPTIONAL_1',  'Optional 1 set', false, false),
(7, 'OPTIONAL_2',  'Optional 2 set', false, false)
returning *;

select * from statuses.statuses s where s.object_id =7

insert into statuses.workflows (status_id_prev, status_id_next)
values
(17, 18),
(17, 19),
(19, 20),
(19, 21),
(19, 22);
;


insert into statuses.workflows (status_id_prev, status_id_next)
values
(22, 23),
(22, 24);


--select * from statuses.instances i
alter table statuses.instances add column instance_is_finished bool default false;
alter table statuses.instances add column instance_is_finished_description text;
--
--
--select * from statuses.instances where instance_token='66f51f5c-6c40-4c09-bc9e-45ff89f0ef9b'
--select * from statuses.events where instance_id=5
--select * from statuses.statuses s where s.object_id =2

--alter table statuses.statuses drop column status_type
alter table statuses.statuses add column status_type text default 'MANDATORY' not null;
alter table statuses.statuses add constraint cst1 check(status_type in ('MANDATORY', 'OPTIONAL','STOP-STATUS'));

update  statuses.statuses 
set status_type='OPTIONAL' 
where status_is_mandatory =false and status_is_stopstatus=false;

update  statuses.statuses 
set status_type='STOP-STATUS' 
where status_is_stopstatus=true;

drop view if exists statuses.v_last_statuses;

alter table statuses.statuses drop column status_is_mandatory;
alter table statuses.statuses drop column status_is_stopstatus;

-- SELECT relation::regclass, * FROM pg_locks WHERE NOT GRANTED;
--select * from pg_stat_activity
--select pg_terminate_backend(10780)
