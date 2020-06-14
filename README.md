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

## Statuses workflow models
***
Each of object has own workflow (or status model).

For example,

2-POINT LINE TASK has workflow:<br>
- RUN -> FINISHED

