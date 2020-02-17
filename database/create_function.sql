------------------------------------------------------------------------------------------------
------------------------------------------fnCreateInstance---------------------------
------------------------------------------------------------------------------------------------

create or replace function statuses.fnCreateInstance(in_objectName text) returns text as $$
declare 
	v_object_id int;
	v_rows int;
	v_instance_token text;
begin
	-----------------------------------------------------------------
	--check if object_name is exists
	select object_id into v_object_id
	from statuses.objects
	where object_name=in_objectName;

	get diagnostics v_rows = ROW_COUNT;

	if v_rows=0 then
		raise exception 'Object name (%) doesn''t exists', in_objectName using hint='Try use existing object name';
	end if;
	

	-----------------------------------------------------------------
	--creating instance of object
	insert into statuses.instances (object_id)
	values(v_object_id)
	returning instance_token into v_instance_token;

	get diagnostics v_rows = ROW_COUNT;

	if v_rows=0 then
		raise exception 'Instance creating was failed (%)', in_objectName;
	end if;


	return v_instance_token;

end;
$$ language plpgsql;

----------------test---------------------




--select * from statuses.objects;
--select * from statuses.instances;
--select * from statuses.v_events;
--select * from statuses.v_workflows;
--select * from statuses.statuses s ;

--select statuses.fnSetStatus('9c18e4d3-ea63-4c14-8c6f-68a3795bb013', 'TEST');
--select statuses.fnSetStatus('15452e55-2f05-4087-827c-199f0089873e', 'FINISHED');

--Представление для текущего статуса





------------------------------------------------------------------------------------------------
------------------------------------------fnGetLastStatusIdOfInstance---------------------------
------------------------------------------------------------------------------------------------


create or replace function statuses.fnGetLastStatusIdOfInstance(in_instance_id bigint) returns int as $$
declare 
	v_status_id int;

begin 
	
	select e.status_id into v_status_id
	from 
		statuses.v_events e 
	where 
		instance_id=in_instance_id
	order by e.event_id desc 
	limit 1;

	return v_status_id;
	
end;
$$ language plpgsql; 

------------------------------------------------------------------------------------------------
------------------------------------------fnCheckPairOfStatusesIsValid---------------------------------------
------------------------------------------------------------------------------------------------

create or replace function statuses.fnCheckPairOfStatusesIsValid(in_object_id int, in_last_status_id int, in_new_status_id int) returns bool as $$
declare 
	v_resuts bool;

begin 
	
	/*
	 * Function returns: 
	 * 		True - when pair of statuses is valid
	 * 		False - when pair of statuses is not valid
	 * */
	
	select true into v_resuts
	from 
		statuses.v_statuses s
	where
		s.object_id = in_object_id and
		coalesce(s.status_id_prev,-1) = coalesce(in_last_status_id,-1) and
		s.status_id = in_new_status_id
		;
	
	if v_resuts is null then
		v_resuts = false;
	end if;

	return v_resuts;
	
end;
$$ language plpgsql; 

--select statuses.fnCheckPairOfStatusesIsValid(2, null, null)
--
--select statuses.fnGetWorkflowId(1, 4)
--select * from statuses.v_workflows w where w.workflow_id=statuses.fnGetWorkflowId(1, 4)
--select * from statuses.instances i
--select * from statuses.v_events
--select * from statuses.statuses s  where s.status_id=statuses.fnGetLastStatusIdOfInstance(1)
--select * from statuses.v_statuses
	
------------------------------------------------------------------------------------------------
------------------------------------------fnGetStatusId---------------------------------------
------------------------------------------------------------------------------------------------

create or replace function statuses.fnGetStatusId(in_object_id int, in_status_name text) returns int as $$
declare 
	v_status_id int;
begin 

	-----------------------------------------------------------------
	--check if STATUS is exists
	select 
		coalesce(
			(select status_id from statuses.statuses where object_id=in_object_id and status_name=in_status_name),
			(select status_id from statuses.statuses where object_id=in_object_id and status_name='*')
			) into v_status_id;
	
	--get diagnostics v_rows = ROW_COUNT;

	return v_status_id;
end;
$$ language plpgsql;

------------------------------------------------------------------------------------------------
------------------------------------------fnSetStatus-------------------------------------------
------------------------------------------------------------------------------------------------

create or replace function statuses.fnSetStatus(in_instance_token text, in_status_name text) returns bigint as $$
declare 
	v_new_status_id int;
	v_object_id int;
    v_event_id int;
	v_object_name text;
	v_rows int;
	v_instance_id bigint;
begin
	
	-----------------------------------------------------------------
	--check if object_name is exists
	select i.instance_id, i.object_id, o.object_name into v_instance_id, v_object_id, v_object_name
	from statuses.instances i,
		 statuses.objects o 
	where 
		i.instance_token::text = in_instance_token and 
		i.object_id=o.object_id;

	get diagnostics v_rows = ROW_COUNT;
	
	if v_rows=0 then
		raise exception 'Instance token (%) doesn''t exists', in_instance_token using hint='Try use correct instance token';
	end if;
	
	
	-----------------------------------------------------------------
	--get status_id or generate exception if it is not exist
	select statuses.fnGetStatusId(v_object_id, in_status_name) into v_new_status_id;

	if v_new_status_id is null then
		raise exception 'Status (%) doesn''t exists for object (%)', in_status_name, v_object_name using hint='Try use existing object name';
	end if;

	-----------------------------------------------------------------
	--Проверка на наличие пары текущего и нового статусов в workflow-таблице
	raise info 'v_object_id=%, v_instance_id=%, v_new_status_id=%', v_object_id, v_instance_id, v_new_status_id;
	if (select statuses.fnCheckPairOfStatusesIsValid(v_object_id, statuses.fnGetLastStatusIdOfInstance(v_instance_id), v_new_status_id)) is false then
		raise exception 'New status (%) is not valid for the workflow of the object (%)', in_status_name, v_object_name using hint='Try use a valid order of statuses';
	end if;
	

	-----------------------------------------------------------------
	--creating instance of object
	insert into statuses.events (instance_id, status_id)
	values(v_instance_id, v_new_status_id)
	returning event_id into v_event_id;

	get diagnostics v_rows = ROW_COUNT;

	if v_rows=0 then
		raise exception 'Inserting stasus was failed (instance_token=%, status_name=%)', in_instance_token, in_status_name;
	end if;


	return v_event_id;

end;
$$ language plpgsql;


--
select * from statuses.fnSetStatus('15452e55-2f05-4087-827c-199f0089873e', 'RUN');
select * from statuses.fnSetStatus('171aa53b-722e-42b6-ab6a-ddd47db96a04', 'RUN');
select * from statuses.fnSetStatus('171aa53b-722e-42b6-ab6a-ddd47db96a04', 'FINISHED');
select * from statuses.fnSetStatus('171aa53b-722e-42b6-ab6a-ddd47db96a04', 'FINISHED');
select statuses.fnCreateInstance('FIXED_TASK');
select statuses.fnCheckPairOfStatusesIsValid(2,  statuses.fnGetLastStatusIdOfInstance(3), 7)
--select * from statuses.v_workflows w where w.workflow_id=statuses.fnGetWorkflowId(1, 4)
--select * from statuses.instances i
--select * from statuses.v_events
--select * from statuses.statuses s  where s.status_id=statuses.fnGetLastStatusIdOfInstance(1)
--select * from statuses.v_statuses
	
--check if new status is first in workflow & what old status is last in workflow
 
