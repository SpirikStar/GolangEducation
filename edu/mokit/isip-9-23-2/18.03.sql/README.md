```cmd
go mod init lesson
```
MAC OS, Linux
```cmd
go get github.com/mattn/go-sqlite3
```
```go
import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)
```

windows 10 VS 11
```cmd
go get "modernc.org/sqlite"
```

```go
import (
	"database/sql"
	_ "modernc.org/sqlite"
)
```