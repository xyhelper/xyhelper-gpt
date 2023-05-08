# xyhelper-gpt

兼容官方界面

## 支持的环境变量

- PORT: 服务端口,不指定则随机
- API_PROXY: API 代理地址, 默认为`http://127.0.0.1:8001`

## docker 命令行模式

启动服务

```bash
docker run --name xyhelper-gpt -d -p 8000:8000 -e API_PROXY=https://personalchat.xyhelper.cn -e PORT=8000 xyhelper/xyhelper-gpt
```

更新服务

```bash
docker pull xyhelper/xyhelper-gpt
docker stop xyhelper-gpt
docker rm xyhelper-gpt
docker run --name xyhelper-gpt -d -p 8000:8000 -e API_PROXY=https://personalchat.xyhelper.cn -e PORT=8000 xyhelper/xyhelper-gpt
```


## docker-compose 模式

创建`docker-compose.yml`文件

```yaml
version: ‘3’
services:
  xyhelper-gpt:
    image: xyhelper/xyhelper-gpt
    container_name: xyhelper-gpt
    ports:
      - 8000:8000
    environment:
      - API_PROXY=https://personalchat.xyhelper.cn
      - PORT=8000
    restart: always
```

启动服务

```bash
docker-compose up -d
```

更新服务

```bash
docker-compose pull
docker-compose up -d
```

