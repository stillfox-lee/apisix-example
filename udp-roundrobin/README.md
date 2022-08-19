# 对于 Apisix 的 UDP roundrobin 测试

Build udp server:

```bash
docker build -f Dockerfile.server -t lsy/udp-echo-server .
```

Startup:

```bash
docker-compose up
```

Run Client:

```bash
go run client.go
```

Apisix 对于 UDP 的 roundrobin 处理逻辑，不是以请求为粒度来分发的。所以，测试的时候是每次建立一个新的 UDP，再发送请求。
