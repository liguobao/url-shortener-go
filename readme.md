# 短链服务 Golang 版本

- GET /{shortCode} 直接走 302 跳转；可选 QueryString s=T, 代表是否需要跳过分享链接事件，默认不跳过
- POST /share-link/ 创建短链链接, 传入 OpenId(类似于用户 Id 或者其他的用户凭证) 和需要生成短链的原始链接

## POST /share-link/

### request json

```json
{
  "openId": "123",
  "url": "https://zhihu.com"
}
```

### response json

```json
{
  "code": 0,
  "data": {
    "shareLink": "localhost:8080/15FUgA",
    "shortCode": "15FUgA"
  }
}
```

### curl

```sh
curl --location --request POST 'localhost:8080/share-link' \
--header 'Content-Type: application/json' \
--data-raw '{
	"openId": "123",
	"url":"https://zhihu.com"
}'
```

## 开发指南

- golang 1.14
- vs code
- 直接运行 main.go 即可
- 依赖的包文件在 go.mod

## 项目原理说明

- 使用 MySQL 存储所有的短链数据
- 使用 MySQL 自增主键生成短链 Id 后，转换成 62 进制的字符串作为短链路径
- 短链访问记录依赖 cookie 中的
- link.sql 为 link 表初始化脚本

## 环境变量说明

- LINK_CODES_PREFIX 每次访问会记录一下当前用户访问过哪些短链，在此 cookie 存储
- URL_SHORTER_DOMAIN 短链服务域名
- MYSQL_CONF 数据库配置
- MQ_CONF rabbitMQ 配置
