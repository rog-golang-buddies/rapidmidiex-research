# List of patterns & conventions

- [http services according to mat ryer](https://pace.dev/blog/2018/05/09/how-I-write-http-services-after-eight-years.html)

## STDLib

```
stdlib/
    cmd/
        main.go
    user/
        handler.go
        service.go
    auth/
        handler.go
        service.go
    user.go
    auth.go

```

## Layers

```
layers/
    cmd/
        main.go
    models/
        user.go
        auth.go
    handlers/
        user.go
        auth.go
    services/
        user.go
        auth.go
```

## Hybrid

```
stdlib/
    cmd/
        main.go
    models/
        user.go
        auth.go
    user/
        handler.go
        service.go
        types.go
    auth/
        handler.go
        service.go
        types.go

```
