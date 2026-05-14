# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## 构建与运行

```bash
# 构建
go build -o bin/server/myapp ./cmd/server

# 运行（configDir 默认为 ./config）
go run ./cmd/server --configDir ./config

# 指定配置文件
go run ./cmd/server --configDir ./config --configFile config.yaml
```

## 技术栈

- **Go 1.26.2**，模块名: `app`
- **Gin**（HTTP 框架）— `gin.Default()` 内置 Logger + Recovery 中间件
- **GORM**（ORM），MySQL 驱动
- **Viper**（配置管理）+ pflag 解析命令行参数
- **golang-jwt/jwt/v5**，JWT 令牌，HS256 签名
- **bwmarrin/snowflake**，雪花 ID 生成器
- **go-redis/redis/v9**（配置已加载，但应用中尚未初始化连接）

## 架构

分层架构 + 依赖注入，在 `router/user.go` 中完成组装：

```
Handler（HTTP 层，处理 gin.Context） → Service（业务逻辑） → Repository（GORM 数据访问）
```

- **`cmd/server/main.go`** — 入口：加载配置 → 连接 MySQL → 启动服务
- **`config/`** — 基于 Viper 的配置加载；`config.yaml` 存放运行时配置
- **`internal/app/server.go`** — 创建 Gin 引擎，注册全局中间件
- **`router/`** — 路由注册，每个路由分组一个文件
- **`middleware/exception.go`** — 全局异常中间件（目前为桩代码，仅打印错误类型）
- **`middleware/auth.go`** — JWT 鉴权中间件，解析 Bearer token，将 claims 注入 context
- **`internal/model/`** — GORM 实体（`BaseModel` 包含 ID/CreatedAt/UpdatedAt）
- **`internal/dao/`** — 请求 DTO（在此绑定 JSON 请求体）
- **`internal/repository/`** — 通过 GORM 进行数据访问
- **`internal/service/`** — 业务逻辑
- **`internal/handler/`** — HTTP 处理器（绑定请求 → 调用 service → 写入 `resp.Output()`）
- **`pkg/resp/`** — 统一 JSON 响应格式：`{"code": int, "message": string, "data": ...}`
- **`pkg/jwt/`** — JWT 令牌生成（access token + refresh token，HS256）

## 路由分组

| 分组 | 路径 | 接口 |
|------|------|------|
| User（公开） | `/user` | `POST /login`, `POST /register` |
| User（鉴权） | `/user` | `GET /info` |

全局中间件：`ExceptionMiddleware`。`/user/info` 使用 `AuthMiddleware` 鉴权。

## 本地开发环境

```bash
# 启动 MySQL + Redis
docker compose -f scripts/docker-compose.yaml up -d
```

MySQL 8.4 端口 3306，Redis 端口 6379。连接信息与 `config/config.yaml` 一致。

## 已知问题

- **密码明文存储**：密码未经过哈希处理直接存储。
- **无数据库迁移**：表结构需手动创建，代码中未调用 `AutoMigrate`。
- **Redis 未使用**：配置已加载但 Redis 从未建立连接或使用。
- **无测试**：仓库中没有任何测试文件。
