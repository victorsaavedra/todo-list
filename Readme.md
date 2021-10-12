# Todo List

## Go + Gorm + Fiber + Docker

To run it:

```bash
make local
```

Once everything is correct, run

```bash
    go get github.com/gofiber/fiber/v2
    go get gorm.io/gorm
    go get gorm.io/driver/mysql
```

Then,

```bash
make docker-run-air
```

Connection parameters:

- 127.0.0.1:8081
- For the database: 
  - Host: 127.0.0.1
  - Username: mysql_go
  - Database: todoGo
  - Port: 4306