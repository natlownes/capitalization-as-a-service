# capitalization_as_a_service

http://strings.microservice.narf.io/

Microservices:  Aw shit.

### dev server

```
cd cli && goapp serve -admin_port 8082
```

### tests

```
gom -test install && go test
```

### app engine deploy

making note of this because I will forget otherwise

```
cd cli && appcfg.py --oauth2 update .
```
