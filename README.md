# CacheService(缓存服务)

一个基于`redis`的缓存服务。



### 配置信息

在该程序的所在文件夹，创建两个文件夹，分别是`configs`、`data`。

在`configs`中，创建服务的配置信息文件`config.json`。内容如下：

```json
{
  "port": 4930,
  "redis": {
    "host": "172.17.0.1",
    "port": 6379,
    "password": "",
    "db": 1
  }
}
```

`data`文件夹创建好即可，此文件夹由服务维护，用来存放缓存数据的相关信息（`cfn`）。



### 关于Cfn

`cfn`是此服务为每个缓存对象设计的一个结构体，结构如下:

```go
type Fn struct {
	Method string      `form:"method" binding:"required" json:"method"`
	Uri    string      `form:"uri" binding:"required" json:"uri"`
	Param  interface{} `form:"param" binding:"" json:"param,omitempty"`
}

type Cfn struct {
	Name   string `form:"name" binding:"required" json:"name"` // 缓存对象名称
	Expire int64  `form:"expire" binding:"" json:"expire"` // 缓存存活时间 单位 秒
	Fn     *Fn    `form:"fn" binding:"required" json:"fn"` // 当缓存失效时，需要获取缓存数据的地址
}
```



### Api

// TODO