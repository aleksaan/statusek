# Statusek
Universal tool for reading &amp; saving statuses of task, documents &amp; other

## General

![Alt text](images/SMSGeneral.png?raw=true "General points")

## Main concept
***
It's a service has given functionality to manage statuses across different services or systems

Service defines contract of status model some one object (for example, task or document) 

Imagine, we have the chain of two application:<br>
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

Statusek has three restAPI methods for organize interaction between services:
1. **instance/create** - create instance of some object and get its token
2. **status/setStatus** - set status of instance (by instance token & status name)
3. **instance/checkIsFinished** - checks instance is finished or not

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
---

First, service-initiator of process creates instance of process inside SMS (statuses management service)

It call:
http://hostname:8080/instance/create

with raw json in the body:
```json
{
	"object_name":"2-POINT LINE TASK",
	"instance_timeout":600
}
```
* timeout in seconds

and get something like this:
```json
{
    "instance_token": "1d36bcdf-d776-4e7a-97a6-4431079d2b2e",
    "message": "Success",
    "status": true
}
```

### Setting statuses
---

Services, who knows token of instance (process) can set statuses by calling:
http://hostname:8080/status/setStatus

with raw json in the body:
```json
{
	"instance_token":"1d36bcdf-d776-4e7a-97a6-4431079d2b2e",
	"status_name": "RUN"
}
```

and get something like this:
```json
{
    "message": "Success",
    "status": true
}
```

or status: false and message with error text

### Check some status is set
---
For specific logic we can to want check some status was set or not

We call:
http://hostname:8080/status/checkStatusIsSet

with raw json in the body:
```json
{
	"instance_token":"1d36bcdf-d776-4e7a-97a6-4431079d2b2e",
	"status_name": "FINISHED"
}
```

and get something like this:
```json
{
    "message": "Status is not set",
    "status": false
}
```
or status: true if checked status is set

### Checking process is finished
---
Service-initiator want to know process is finished yet or not

We call:
http://hostname:8080/instance/checkIsFinished

with raw json in the body:
```json
{
	"instance_token":"1d36bcdf-d776-4e7a-97a6-4431079d2b2e"
}
```

and get something like this:
```json
{
    "message": "Instance is finished",
    "status": true
}
```
or status: false and message: Instance is not finished 

### Getting all statuses of process
---

We call:
http://hostname:8080/instance/checkIsFinished

with raw json in the body:
```json
{
	"instance_token":"1d36bcdf-d776-4e7a-97a6-4431079d2b2e"
}
```

and get something like this:
```json
{
    "events": [
        {
            "Status": {
                "ObjectID": 2,
                "StatusID": 2,
                "StatusName": "RUN",
                "StatusDesc": "Task is running",
                "StatusType": "MANDATORY"
            },
            "EventCreationDt": "2020-06-15T02:22:05.38288+03:00"
        },
        {
            "Status": {
                "ObjectID": 2,
                "StatusID": 3,
                "StatusName": "FINISHED",
                "StatusDesc": "Task is finished",
                "StatusType": "MANDATORY"
            },
            "EventCreationDt": "2020-06-15T02:22:08.998913+03:00"
        }
    ],
    "message": "Success",
    "status": true
}
```