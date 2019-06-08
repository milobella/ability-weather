# Weather
Ability to get the weather.

## Features
>
> TODO: describe features
>

## Prerequisites

- Having access to [gitlab.milobella.com](https://gitlab.milobella.com/milobella)
- Having ``golang`` installed [instructions](https://golang.org/doc/install)
- Having ``go dep`` installed [instructions](https://golang.github.io/dep/docs/installation.html)

## Build

```bash
$ dep ensure
$ go build -o bin/ability cmd/ability/main.go
```

## Run

```bash
$ bin/ability
```

## Requests example

#### Get the weather of now in the default city.
```bash
$ curl -i -X POST http://localhost:4444/resolve -d '{"nlu":{"BestIntent": "GET_WEATHER"}}'
HTTP/1.1 200 OK
Date: Fri, 07 Jun 2019 17:08:46 GMT
Content-Length: 294
Content-Type: text/plain; charset=utf-8

{"nlg":{"sentence":"In {{city}} now, the temperature is {{temperature}}. {{weather_sentence}}","params":[{"name":"city","value":"Cannes","type":"string"},{"name":"temperature","value":21.78,"type":"string"},{"name":"weather_sentence","value":"","type":"inner"}]},"context":{"slot_filling":{}}}
```

## CHANGELOGS
- [Application changelog](CHANGELOG.md)
- [Helm chart changelog](weather-ability/helm/ability-cinema/CHANGELOG.md)