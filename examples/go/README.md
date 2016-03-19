Устанавливаем окружение и запускаем сервис

```bash
GOPATH=`pwd`                # set evnironment 
go get "github.com/lib/pq"  # install dependencies
go run src/server.go        # run server
```

Прогоняем тесты

```bash
go test server
```

Считаем coverage

```bash
go get code.google.com/p/go.tools/cmd/cover

go test -tags test -cover -coverprofile cover.out server
go tool cover -html=cover.out -o cover.html
```

