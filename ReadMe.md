# web脚手架项目

## 准备
1.设置环境变量
```bash
# jwt secret key
export SECRET_KEY="123456"

# service port. don't set use default 8080
export PORT=":8888"

export LOG_LEVEL="info"

# jwt token过期时间, 单位: 秒. 默认为86400
export JWT_EXPIRE_TIME=3600
```