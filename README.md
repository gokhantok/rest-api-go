# webserver

## Installing Operator For Docker

```bash
$ make all
```

## Installing Operator For Local

```bash
$ make build
$ make run
```
## Installing Operator For Docker-compose

```bash
$ docker-compose up    
```

Dockerfile2 is not used. It shows how it happened before becoming a multi-stage docker.
 

## API

The APi provides a single endpoint `/events` that will respond to `GET` and `POST` methods.

`GET /events?from=1234&to=1334&type=my_event` will return a number of aggregated events in the time between `1234` and `1334` (timestamps).
Sample response:
```json
{
  "count": 52,
  "type": "my_event"
}
```

`POST /events` will receive an event with a number such as:
```json
{
  "count": 3,
  "type": "my_event",
  "timestamp": 1238
}
```
