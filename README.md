# A SumUp interview

## Overview
This application is designed according to a SumUp interview specification this including
three endpoints responsible for handling an incoming `POST` requests. The `endpoints` are
channel oriented meaning we have a single `endpoint` per channel. In this implementation we have:
1. an email channel for notification
2. a sms channel for notification
3. a slack channel for notification

All channels are configurable via `.env.local` file which is in `base64` encoded format.
Every channel - `email`, `sms` and `slack` has support for different providers.
For example we have `gmail` and `yahoo` for an email channel.

Here is it how email configuration looks as `JSON`:

```json
[
   {
      "server":"mail.google.com",
      "port":443,
      "ssl":true,
      "name":"gmail"
   },
   {
      "server":"mail.yahoo.com",
      "port":443,
      "ssl":true,
      "name":"yahoo"
   }
]
```

The whole application is a configuration depended on and currently there are three factories for crafting the
required number of providers.

The application is build around `FX` a dependency injection framework from `Uber`. All `SOLID` principles are kept 
in mind and the software design follow as well the so-called hexagon architecture. Also, it is designed to be a part 
from a bigger microservice architecture.

### NOTE
For this demo the actual requests to third party APIs are not performed instead we just have logged the results. 

## Requirements
1. Operating system that support `docker`
2. `Docker` and `docker-compoer`
3. `git`

## Testing details
Here is it the current setup on which the application was written and tested:

```bash
$ uname -a
Linux base9 5.15.0-39-generic #42-Ubuntu SMP Thu Jun 9 23:42:32 UTC 2022 x86_64 x86_64 x86_64 GNU/Linux
```

```bash
$ docker version
Client: Docker Engine - Community
Version:           20.10.17
API version:       1.41
Go version:        go1.17.11
Git commit:        100c701
Built:             Mon Jun  6 23:02:46 2022
OS/Arch:           linux/amd64
Context:           default
Experimental:      true

Server: Docker Engine - Community
Engine:
Version:          20.10.17
API version:      1.41 (minimum version 1.12)
Go version:       go1.17.11
Git commit:       a89b842
Built:            Mon Jun  6 23:00:51 2022
OS/Arch:          linux/amd64
Experimental:     false
containerd:
Version:          1.6.6
GitCommit:        10c12954828e7c7c9b6e0ea9b0c02b01407d3ae1
runc:
Version:          1.1.2
GitCommit:        v1.1.2-0-ga916309
docker-init:
Version:          0.19.0
GitCommit:        de40ad0
```

```bash
$ docker compose version
Docker Compose version v2.6.0
```

## Start the application

```bash
$ docker compose up
```

## Sending notification
To send a notification we are going to use `curl`. Before sending any requests here is it some details.
First let's imagine that this application is a part from something bigger, as microservice architecture and our 
responsibilities are to send notification for some providers and different channels. As such part we have implemented
a idempotency middleware, which is going to allow every message only once. If some errors occur during the processing the
application will return error and the sender ( `kafka` or  other microservice ) will have a possibility to re-send it
later on. So keeping this in ming lets continue ahead.

1. Sending an email notification to `yahoo`
```bash
$ curl -vvv -H "Idempotency-Key: 0b1ca7a6-fd4f-43c3-9e6b-236a02691425" -d '{"from":"da@da.com", "to":"ba@ko.com", "message": "test", "provider":"yahoo"}' localhost:8080/email-notify
*   Trying 127.0.0.1:8080...
* Connected to localhost (127.0.0.1) port 8080 (#0)
> POST /email-notify HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.81.0
> Accept: */*
> Idempotency-Key: 0b1ca7a6-fd4f-43c3-9e6b-236a02691425
> Content-Length: 77
> Content-Type: application/x-www-form-urlencoded
>
* Mark bundle as not supporting multiuse
< HTTP/1.1 200 OK
< Date: Sun, 26 Jun 2022 14:23:50 GMT
< Content-Length: 0
<
* Connection #0 to host localhost left intact
```

Pay attention on the request above, here we are providing our unique idempotency key via the `header`

Next lets try to re-send it again and test our idempotency middleware:

```bash
$ curl -vvv -H "Idempotency-Key: 0b1ca7a6-fd4f-43c3-9e6b-236a02691425" -d '{"from":"da@da.com", "to":"ba@ko.com", "message": "test", "provider":"yahoo"}' localhost:8080/email-notify
*   Trying 127.0.0.1:8080...
* Connected to localhost (127.0.0.1) port 8080 (#0)
> POST /email-notify HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.81.0
> Accept: */*
> Idempotency-Key: 0b1ca7a6-fd4f-43c3-9e6b-236a02691425
> Content-Length: 77
> Content-Type: application/x-www-form-urlencoded
>
* Mark bundle as not supporting multiuse
< HTTP/1.1 400 Bad Request
< Content-Type: text/plain; charset=utf-8
< X-Content-Type-Options: nosniff
< Date: Sun, 26 Jun 2022 14:25:11 GMT
< Content-Length: 35
<
message is already being processed
* Connection #0 to host localhost left intact
```

So as expected we get `HTTP/1.1 400 Bad Request` and message `message is already being processed`

Next we can switch to a different email provider - `gmail`

```bash
$ curl -vvv -H "Idempotency-Key: 0b1ca7a6-fd4f-43c3-9e6b-236a02691425" -d '{"from":"da@da.com", "to":"ba@ko.com", "message": "test", "provider":"gmail"}' localhost:8080/email-notify
*   Trying 127.0.0.1:8080...
* Connected to localhost (127.0.0.1) port 8080 (#0)
> POST /email-notify HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.81.0
> Accept: */*
> Idempotency-Key: 0b1ca7a6-fd4f-43c3-9e6b-236a02691428
> Content-Length: 77
> Content-Type: application/x-www-form-urlencoded
>
* Mark bundle as not supporting multiuse
< HTTP/1.1 200 OK
< Date: Sun, 26 Jun 2022 14:27:46 GMT
< Content-Length: 0
<
* Connection #0 to host localhost left intact
```

That was for the email next is the `SMS` channel:

2. Sending a `SMS` notification.
Here we have two providers as well - `smsone` and `smstwo`.

```bash
$ curl -vvv -H "Idempotency-Key: 0b1ca7a6-ed4f-43c3-9e6b-436a02691109" -d '{"from":"1234", "to":"+324", "message": "test", "provider":"smsone"}' localhost:8080/sms-notify
*   Trying 127.0.0.1:8080...
* Connected to localhost (127.0.0.1) port 8080 (#0)
> POST /sms-notify HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.81.0
> Accept: */*
> Idempotency-Key: 0b1ca7a6-ed4f-43c3-9e6b-436a02691109
> Content-Length: 68
> Content-Type: application/x-www-form-urlencoded
>
* Mark bundle as not supporting multiuse
< HTTP/1.1 200 OK
< Date: Sun, 26 Jun 2022 14:28:59 GMT
< Content-Length: 0
<
* Connection #0 to host localhost left intact
```

and the second provider:

```bash
$ curl -vvv -H "Idempotency-Key: 0b1ca7a6-ed4f-43c3-9e6b-636a02691109" -d '{"from":"1234", "to":"+324", "message": "test", "provider":"smstwo"}' localhost:8080/sms-notify
*   Trying 127.0.0.1:8080...
* Connected to localhost (127.0.0.1) port 8080 (#0)
> POST /sms-notify HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.81.0
> Accept: */*
> Idempotency-Key: 0b1ca7a6-ed4f-43c3-9e6b-636a02691109
> Content-Length: 68
> Content-Type: application/x-www-form-urlencoded
>
* Mark bundle as not supporting multiuse
< HTTP/1.1 200 OK
< Date: Sun, 26 Jun 2022 14:34:09 GMT
< Content-Length: 0
<
* Connection #0 to host localhost left intact
```

the idempotency middleware works as well:

```bash
$ curl -vvv -H "Idempotency-Key: 0b1ca7a6-ed4f-43c3-9e6b-636a02691109" -d '{"from":"1234", "to":"+324", "message": "test", "provider":"smstwo"}' localhost:8080/sms-notify
*   Trying 127.0.0.1:8080...
* Connected to localhost (127.0.0.1) port 8080 (#0)
> POST /sms-notify HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.81.0
> Accept: */*
> Idempotency-Key: 0b1ca7a6-ed4f-43c3-9e6b-636a02691109
> Content-Length: 68
> Content-Type: application/x-www-form-urlencoded
>
* Mark bundle as not supporting multiuse
< HTTP/1.1 400 Bad Request
< Content-Type: text/plain; charset=utf-8
< X-Content-Type-Options: nosniff
< Date: Sun, 26 Jun 2022 14:34:56 GMT
< Content-Length: 35
<
message is already being processed
* Connection #0 to host localhost left intact
```

and finally the `SLACK` notification. Here we also have two providers `slackone` and `slacktwo`

```bash
$ curl -vvv -H "Idempotency-Key: 0b1ca7a6-1d4f-43c3-9e6b-836a02691544" -d '{"channel":"sumup", "message": "test", "provider":"slackone"}' localhost:8080/slack-notify
*   Trying 127.0.0.1:8080...
* Connected to localhost (127.0.0.1) port 8080 (#0)
> POST /slack-notify HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.81.0
> Accept: */*
> Idempotency-Key: 0b1ca7a6-1d4f-43c3-9e6b-836a02691544
> Content-Length: 61
> Content-Type: application/x-www-form-urlencoded
>
* Mark bundle as not supporting multiuse
< HTTP/1.1 200 OK
< Date: Sun, 26 Jun 2022 14:36:35 GMT
< Content-Length: 0
<
* Connection #0 to host localhost left intact
```

and second provider:

```bash
$ curl -vvv -H "Idempotency-Key: 0b1ca7a6-1d4f-43c3-9e6b-836a02691544" -d '{"channel":"sumup", "message": "test", "provider":"slacktwo"}' localhost:8080/slack-notify
*   Trying 127.0.0.1:8080...
* Connected to localhost (127.0.0.1) port 8080 (#0)
> POST /slack-notify HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.81.0
> Accept: */*
> Idempotency-Key: 0b1ca7a6-1d4f-43c3-9e6b-836a02691544
> Content-Length: 61
> Content-Type: application/x-www-form-urlencoded
>
* Mark bundle as not supporting multiuse
< HTTP/1.1 400 Bad Request
< Content-Type: text/plain; charset=utf-8
< X-Content-Type-Options: nosniff
< Date: Sun, 26 Jun 2022 14:37:09 GMT
< Content-Length: 35
<
message is already being processed
* Connection #0 to host localhost left intact
```

In addition, all events has been logged from our backend.

## What is missing
1. Unit tests
2. Metrics
3. Tracer

## Production notes
For production deployment instead of docker we are going to use kubernetes. 
The orchestration of the services are:
1. A Database cluster which can scale vertically.
2. A Load Balancer and a fleet of application instances waiting for the requests. If the load become bigger, more instances 
are scheduled for running. Here we have a horizontal scaling, thanks to our stateless implementation.

## Final notes

Thanks for reading!

Dimitar Dimitrov
