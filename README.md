# Statusek

Universal tool for reading &amp; saving statuses of task, documents &amp; other

## General

![Alt text](images/SMSGeneral.png?raw=true "General points")

## Main concept

***
It's a service has given functionality to manage statuses across different services or systems

Service defines contract of status model some one object (for example, task or document)

Imagine, we have the chain of two applications which execute some process:
**Service1 -> Service2**

Workflow between them looks like (async interaction):

1. Service1 push task to Service2
2. Service2 executes task and writes back statuses of execution
3. Service1 see statuses and waits when task will be fully executed

Conclusions:

- We should develop status model
- We should develop methods to set statuses
- We should develop method to share statuses
- Service2 can change statuses it retuns and does not inform Service1 about this
- We should foresee when task is hold on Service2 (timed out) and finish status never will be returned

Statusek realizes all of these points.

Statusek has six restAPI methods for organize interaction between services:

1. **instance/create** - create instance (process) of some object (model) and get its token
2. **status/setStatus** - set status of instance (by instance token & status name)
3. **status/checkStatusIsSet** - check certain status is set
4. **status/checkStatusIsReadyToSet** - check certain status can be set
5. **instance/checkIsFinished** - checks instance is finished or not
6. **event/getEvents** - gets all setted statuses

So, interaction is changing to:

1. Service1 calls *instance/create* and gets instanceToken
2. Service1 push task to Service2 and transmit instanceToken to it
3. Service2 executes task and writes certain statuses by call *status/setStatus*
4. Service1 call *instance/checkIsFinished* and understands when task will be fully executed
5. If Service2 hasn't returned any of its statuses then *instance/checkIsFinished* is set to True

## Statuses

***
Statuses represents states some process, action or object.

![Alt text](images/SMSModelsExample.png?raw=true "Examples of statuses models")

## Rest API

### Creating instance of process

***

First, service-initiator of process creates instance of process inside SMS (statuses management service)

It calls:
<http://hostname:8080/instance/create>

with raw json in the body:

```json
{
 "object_name":"2-POINT LINE TASK",
 "instance_timeout":600
}
```

- timeout in seconds

and get something like this:

```json
{
    "status": true,
    "message": "",
    "data": {
        "instance_token": "1f13cd00-f6fc-49fc-918b-2915bc05908f"
    }
}
```

So, field 'status' contains true if request has executed  successfully. If errors occurred, the 'status' field will be set to 'false,' and you will find an accompanying message in the 'message' field.
For example:

```json
{
    "status": false,
    "message": "ERROR: Object name '2-POINT LINE TASK1' is not found",
    "data": {}
}
```

### Setting statuses

***

Services that possess the instance token (process) can update statuses by making a call to the following endpoint:
<http://hostname:8080/status/setStatus>

with raw json in the body:

```json
{
 "instance_token":"1f13cd00-f6fc-49fc-918b-2915bc05908f",
 "status_name": "STARTED"
}
```

and get something like this:

```json
{
    "status": true,
    "message": "",
    "data": {}
}
```

If there's an error, 'status' will be 'false,' and 'message' will contain an error description.

### Checking some status is set

***
For specific logic we can to want check some status was set or not

We call:
<http://hostname:8080/status/checkStatusIsSet>

with raw json in the body:

```json
{
 "instance_token":"1f13cd00-f6fc-49fc-918b-2915bc05908f",
 "status_name": "STARTED"
}
```

and get something like this:

```json
{
    "status": true,
    "message": "",
    "data": {
        "status_is_set": true
    }
}
```

or "status_is_set": false,  if it is not set

### Checking some status is ready to be set

***
For specific logic we can to want check some status was set or not

We call:
<http://hostname:8080/status/checkStatusIsReadyToSet>

with raw json in the body:

```json
{
 "instance_token":"1f13cd00-f6fc-49fc-918b-2915bc05908f",
 "status_name": "FINISHED"
}
```

and get something like this:

```json
{
    "status": true,
    "message": "",
    "data": {
        "status_is_ready_to_set": false,
        "status_is_ready_to_set_description": "Not all previos mandatory statuses are set for status 'FINISHED'"
    }
}
```

or a positive answer if all limits to set this status have completed:

```json
{
    "status": true,
    "message": "",
    "data": {
        "status_is_ready_to_set": true,
        "status_is_ready_to_set_description": ""
    }
}
```

### Checking process is finished

***
Service-initiator want to know process is finished yet or not

We call:
<http://hostname:8080/instance/checkIsFinished>

with raw json in the body:

```json
{
 "instance_token":"1f13cd00-f6fc-49fc-918b-2915bc05908f"
}
```

and get something like this:

```json
{
    "status": true,
    "message": "",
    "data": {
        "instance_is_finished": false,
        "instanse_is_finished_description": ""
    }
}
```

or status: false and message: Instance is not finished

### Getting all statuses of process

***

We call:
<http://hostname:8080/event/getEvents>

with raw json in the body:

```json
{
    "instance_token": "1f13cd00-f6fc-49fc-918b-2915bc05908f"
}
```

and get something like this:

```json
{
    "status": true,
    "message": "",
    "data": {
        "events": [
            {
                "Status": {
                    "ID": 1,
                    "CreatedAt": "2022-01-24T00:22:27.278227+03:00",
                    "UpdatedAt": "2022-01-24T00:22:27.278227+03:00",
                    "DeletedAt": null,
                    "ObjectID": 1,
                    "StatusName": "STARTED",
                    "StatusDesc": "",
                    "StatusType": "MANDATORY"
                },
                "EventCreationDt": "2023-05-10T15:44:49.879293+03:00"
            }
        ]
    }
}
```
