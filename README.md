# Guuid

UUID 的生成通常基于时间和机器的物理或MAC 地址等参数，以确保在同一时空中的所有机器上生成的 UUID 是唯一的。UUID 的标准形式是 32 个十六进制数字，分为五部分，形如 8-4-4-4-12 的 36 个字符的字符串，

`例如：d75f30db-659c-262c-0000-00013c598b3d`

UUID 的主要用途包括数据库记录的唯一标识、事务处理、分布式系统、软件安装和配置、网络会话等，其可以保证在不同平台和应用间的数据一致性和无冲突性。

另外值得注意的是，虽然 UUID 的设计目的是全球唯一性，但由于其生成的特定方法可能在未来出现重复，UUID 设计上允许一定的重复。例如，根据不同的应用需求，有些情况下可以使用更简单的随机数生成器来生成 UUID，虽然这样可能会增加重复的可能性。因此，在需要绝对唯一性的场景下，可能需要额外的安全措施。

Guuid使用了服务器Hostname、进程PID、时间戳、随机数、时序元素等一系列元素来保证生成UUID的唯一性。

## 使用方式

### Docker

```shell
docker pull winston2024/guuid:latest

docker-compose up -d

# 默认端口为9000
docker-compose logs
docker-guuid  | [GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.
docker-guuid  | 
docker-guuid  | [GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
docker-guuid  |  - using env:	export GIN_MODE=release
docker-guuid  |  - using code:	gin.SetMode(gin.ReleaseMode)
docker-guuid  | 
docker-guuid  | [GIN-debug] GET    /get                      --> github.com/ykubernetes/GUUID/Service.CreateHandler (3 handlers)
docker-guuid  | [GIN-debug] GET    /get/simple               --> github.com/ykubernetes/GUUID/Service.CreateSimpleHandler (3 handlers)
docker-guuid  | [GIN-debug] GET    /mget/:num                --> github.com/ykubernetes/GUUID/Service.CreateMultiHandler (3 handlers)
docker-guuid  | [GIN-debug] GET    /mget/:num/simple         --> github.com/ykubernetes/GUUID/Service.CreateMultiSimpleHandler (3 handlers)
docker-guuid  | [GIN-debug] [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
docker-guuid  | Please check https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies for details.
docker-guuid  | [GIN-debug] Listening and serving HTTP on 0.0.0.0:9000
```

### Windows/Linux

下载对应平台制品，解压并运行

通过浏览器访问`http://IP:9000/get`

## RestAPI

- 获取UUID

```shell
$ curl http://127.0.0.1:9000/get

{
    "Version": "v0.1",
    "uuid": "d75f30db-659c-262c-0000-00013c598b3d"
}
```

- 批量获取UUID

```shell
$ curl http://127.0.0.1:9000/mget/200

"Data": [
        "d75f30db-659c-27de-0000-00020c7ca403",
        "d75f30db-659c-27de-0000-00034d76025a",
        "d75f30db-659c-27de-0000-00056dc9a6ee",
        "d75f30db-659c-27de-0000-00041d061bb3",
        "d75f30db-659c-27de-0000-0006750327ef",
        "d75f30db-659c-27de-0000-00074e3f36e9",
        "d75f30db-659c-27de-0000-00085da0e03d",
        "d75f30db-659c-27de-0000-00092a675a50",
        "d75f30db-659c-27de-0000-000a38d48d5d",
        "d75f30db-659c-27de-0000-000b44ec67a5",
        "d75f30db-659c-27de-0000-000c0acee0b5",
        "d75f30db-659c-27de-0000-000d3e4c645e",
        "d75f30db-659c-27de-0000-000e6565d933",
        "d75f30db-659c-27de-0000-000f48f8ac66",
        "d75f30db-659c-27de-0000-0010757e4cf7",
        "d75f30db-659c-27de-0000-00111eb9492e",
        "d75f30db-659c-27de-0000-00125b27b46b",
        ... ...
         ],
   "Version": "v0.1"
}

说明:
最多返回2000条记录
```

- 获取简版UUID

```shell
$ curl http://127.0.0.1:9000/get/simple

{
    "Version": "v0.1",
    "simple_uuid": "d75f30db659c281d000000cb46ae073d"
}
```

- 批量获取简版UUID

```shell
$ curl http://127.0.0.1:9000/mget/200/simple

{
    "Data": [
        "09cecc39659c28590000000140762aa4",
        "09cecc39659c28590000000249a14900",
        "09cecc39659c2859000000034fb08247",
        "09cecc39659c2859000000040c48d702",
        "09cecc39659c2859000000056ae64dda",
        "09cecc39659c2859000000065a5a49d0",
        "09cecc39659c28590000000750410980",
        "09cecc39659c285900000008626e1d25",
        "09cecc39659c285900000009267e21ac",
        "09cecc39659c28590000000a7b3b7b9f"
    ],
    "Version": "v0.1"
}
```

