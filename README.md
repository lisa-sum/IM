# IM

## 前置条件
1. 运行环境存在`GO`开发环境
2. 本地或远程存在`MongoDB`服务器
3. 本地或远程存在`Redis`服务器

## 设置环境
在项目根路径的`config`目录下的`db.yaml`配置, 改成你的数据库的相关信息

## MongoDB 配置表内容

### 消息表
用于读取历史消息列表
```mongodb
use im
db.message_basic.insertMany(
[
  {
    "created_at": 1,
    "data": "发送的数据",
    "room_identity": "房间的唯一标识",
    "updated_at": 1,
    "user_identity": "用户的唯一标识"
  }
])
```

### 用户表
存储用户信息
```mongodb
use im
db.user_basic.insertMany(
[
  {
    "account": "账号",
    "avatar": "头像",
    "created_at": 1,
    "email": "邮箱",
    "nickname": "昵称",
    "password": "密码",
    "sex": 1,
    "updated_at": 1
  },
])
```

### 群聊/房间表
存储 群聊/房间 信息
```mongodb
use im
db.user_room.insertMany(
[
  {
    "created_at": 1,
    "info": "房间简介",
    "name": "房间名称",
    "number": "房间号",
    "updated_at": 1,
    "user_identity": "房间创建者的唯一标识"
  }
])
```

### 用户群聊表
存储用户的加入的 群聊/房间 信息
```mongodb
use im
db.user_room.insertMany(
[
  {
    "created_at": 1,
    "info": "房间简介",
    "message_identity": "123123",
    "name": "aa",
    "number": "1",
    "room_identity": "1",
    "updated_at": 1,
    "user_identity": "d22e9e82fab041cc8f8d57d51d6b802d"
  },
])
```
## 测试
按照上述内容进行修改插入
1. 进入项目根路径的`test`目录下
2. 运行`getPath_test.go`文件校验项目配置文件的路径设置是否正确
3. 运行`pingDB_test.go`文件校验`mongodb`的设置是否正确
