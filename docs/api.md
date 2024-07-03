## api接口

### 说明
- ${host}表示访问域
- ${explain_url}地址前缀

##### 文件上传接口

- 访问地址：/file/upload
- 方法：PUT
- 参数方式： formdata
- 参数列表：

|名称|类型|是否可空|提交方式|说明|
|:--:|:--:|:--:|:--:|:--:|
|file|文件|否|formdata|上传文件体|
|app_id|int|否|formdata|应用id|
|version|string|否|formdata|应用版本号|
|sha|string|否|formdata|64位sha256校验码|
- 响应格式：json
- 响应内容示例：
```
{
    "code": 0,
    "message": "success",
    "data": null
}
```

##### 文件下载接口

- 地址：/file/download/{app_id}/{version}
- 方法：GET
- 参数列表：

|名称|类型|是否可空|提交方式|说明|
|:--:|:--:|:--:|:--:|:--:|
|app_id|int|否|urlpath|应用id|
|version|string|否|urlpath|应用版本|
|token|string|否|param|验证串|
|file_name|string|是|param|指定下载文件名,不填使用源文件名|

- 响应格式：binary
- 响应说明：文件体

##### 导出仓库接口

- 地址：/repo/export
- 方法：POST
- 参数列表：

|名称|类型|是否可空|提交方式|说明|
|:--:|:--:|:--:|:--:|:--:|
|app_list|json|是|request_body|要导出的app清单，空为全量导出|
- 请求示例：
```
{
    "app_list":[
        {
            "id": 1,
            "versions":["1.0","1.1","1.2"]
        },
        {
            "id": 2,
            "versions":["1.0","1.1","1.5"]
        }
    ]
}
```
- 响应格式：binary
- 响应说明：导出文件包

##### 导入仓库
- 地址：/repo/import
- 方法：POST
- 参数列表：

|名称|类型|是否可空|提交方式|说明|
|:--:|:--:|:--:|:--:|:--:|
|file|文件|否|formdata|导入文件|
|sha|string|否|formdata|64位sha256校验码|

- 响应格式：json
- 响应内容示例：
```
{
    "code": 0,
    "message": "success",
    "data": null
}
```

## 数据结构

#### 3、 apps
应用信息表

| 序号 | 名称 | 描述 | 类型 | 键 | 为空 | 额外 | 默认值 |
| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| 1 | `id` |  | bigint | PRI | NO | auto_increment |  |
| 2 | `create_time` |  | timestamp |  | NO | DEFAULT_GENERATED | CURRENT_TIMESTAMP |
| 3 | `update_time` |  | timestamp |  | YES |  |  |
| 4 | `user_id` | 所属用户 | bigint | MUL | NO |  |  |
| 5 | `name` | 应用名称 | varchar(45) |  | NO |  |  |
| 6 | `dis_name` | 应用显示名称 | varchar(45) |  | YES |  |  |
| 7 | `volume_limit` | 容量上限 | int |  | YES |  |  |
| 8 | `version_limit` | 版本数上限 | int |  | YES |  |  |

#### 4、 app_package
应用包文件表

| 序号 | 名称 | 描述 | 类型 | 键 | 为空 | 额外 | 默认值 |
| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| 1 | `id` |  | bigint | PRI | NO | auto_increment |  |
| 2 | `create_time` |  | timestamp |  | NO | DEFAULT_GENERATED | CURRENT_TIMESTAMP |
| 3 | `update_time` |  | timestamp |  | YES |  |  |
| 4 | `file_id` | 文件id | bigint | MUL | YES |  |  |
| 5 | `volume` | 容量 | int |  | YES |  |  |


#### 5、 file
文件表

| 序号 | 名称 | 描述 | 类型 | 键 | 为空 | 额外 | 默认值 |
| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| 1 | `id` |  | bigint | PRI | NO | auto_increment |  |
| 2 | `create_time` |  | timestamp |  | NO | DEFAULT_GENERATED | CURRENT_TIMESTAMP |
| 3 | `update_time` |  | timestamp |  | YES |  |  |
| 4 | `file_name` | 存储的文件名 | varchar(254) |  | YES |  |  |
| 5 | `length` | 文件大小 | int |  | YES |  |  |
| 6 | `sha256` | 文件的数字摘要信息 | varchar(254) |  | YES |  |  |
| 7 | `original_name` | 上传时的原文件名 | varchar(254) |  | YES |  |  |
| 8 | `path` | 文件保存地址 | varchar(254) |  | YES |  |  |