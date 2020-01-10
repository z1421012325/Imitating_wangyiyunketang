# 不使用文件配置,直接加载进入系统环境中

export IS_LOAD_CONFIG_FILE = "true"

export SESSION_SECRE = "1234567890QWER%$&$%&TYUIOPLKJH12345<>*&(GFSDAZXC45VBNMqwepoir63123rltykcmbvajhn"
export GIN_MODE = "debug"
export gin_mode = "release"

export REDIS_ADDR = "localhost:6379"
export REDIS_HOST = "localhost"
export REDIS_PORT = "6379"
export REDIS_DB = "10"
export REDIS_NAME = ""
export REDIS_PASSWORD = ""

export MYSQL_HOST = "localhost"
export MYSQL_PORT = "3306"
export MYSQL_NAME = "root"
export MYSQL_PASSWORD = "zyms90bdcs"
export MYSQL_DB = "test"

export SERVER_PORT = ":7999"

# 一些环境配置 alipay,aliyun-oss.....

# 更新系统一些配置
sudo apt install update
# 默认为宿主机安装了golang  移动到文件目录
cd /usr/home
go build main.go


# 下载build
yum install docker-io
# 启动docker服务
service docker start

# --------------------------pull ---------------------------------
# 如果以及由仓库并存在仓库中,pull下来
# docker pull NAME:version
# docker run -itd NAME:version -p 3306:3360 -p 6379:6379 -p 8888:7999 /usr/home/main

# ------------------------ 先安装再push ---------------------------
# build Dockerfile
docker build -t app:v1 .
docker images

# reids,mysql以及app端口映射
docker run -dit app:v1 -p 3306:3360 -p 6379:6379 -p 8888:7999 /usr/home/main



# login
docker login -u 账号 -p 密码
docker push 仓库名







