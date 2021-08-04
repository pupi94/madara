## DB 迁移脚本
```go
// 生成 sql 迁移文件
go run main.go db:generate_migration create_products
// db/migrations/20201207105022_create_products.down.sql
// db/20201207105022_create_products.up.sql  用于回滚 DB

// 执行迁移
go run main.go db:migrate up

// DB 回滚
go run main.go db:migrate down
```

## Grpc [=>](https://grpc.io/docs/languages/go/quickstart/)