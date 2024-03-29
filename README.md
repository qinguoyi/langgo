# LangGo
基于Gin封装的业务脚手架
```shell
 _                       _____       
| |                     |  __ \      
| |     __ _ _ __   __ _| |  \/ ___  
| |    / _` | '_ \ / _` | | __ / _ \ 
| |___| (_| | | | | (_| | |_\ \ (_) |
\_____/\__,_|_| |_|\__, |\____/\___/ 
                    __/ |            
                   |___/             
```


## 功能
* 面向对象编程，避免全局变量满天飞
* 支持插件化管理，DB/Redis/Minio等基础组件
* 支持TraceId
* 支持请求/响应参数打印
* 支持日志输出到文件和控制台
* 支持服务优雅重启
* DB支持多个实例切换
* 支持swag api文档
* 自定义GORM封装及操作示例(包括事务)
* 完整的CURD接口示例

## TODO
* 静态资源api测试
* 单元测试
* redis分布式锁使用示例
* 雪花算法生成分布式ID
* 增加es基础组件
* Makefile
* Dockerfile

## 测试
* 生成swag文档
```shell
swag init -g .\cmd\main.go
```