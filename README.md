### 服务实现
go-zero框架

### 下载镜像
```
sudo docker pull aiotpocket/file_server:latest
```

### 创建临时容器
```
sudo docker run -it \
--name file_server \
--restart=always \
--privileged=true \
-p 8888:8888 \
-d aiotpocket/file_server:latest
```

### 复制服务配置文件
```
sudo mkdir /file
sudo docker cp file_server:/app/etc /file/etc
sudo docker rm -f file_server
```

### 创建正式容器
```
sudo docker run -it \
--name file_server \
--restart=always \
--privileged=true \
-p 8888:8888 \
-v /file:/file:rw \
-v /file/logs:/app/logs:rw \
-v /file/etc:/app/etc:rw \
-d aiotpocket/file_server:latest
```

### 赋予挂载卷权限
```
sudo chmod -R 777 /file
```

### http post 请求
#### 上传需要 CRC16-Modbus 校验的 bin 文件
请求接口 `http://127.0.0.1:8888/api/v1/file/upload/bin`

请求json
```json
{
    "file_name": "test_file.bin",
    "file_context": "010300120010E403",
    "pass_key": "43ZJJWpCZsosAzUr6KiPNz4Ek5iZlZVE"
}
```
file_name: 选择参数，为空时使用 file_context 十六进制倒数第4、3位为文件名。

file_context: 必填参数，十六进制数字符串，最后两位为 CRC16-Modbus 校验码。

pass_key: 必填参数，上传密钥。

请求返回json
```json
{
	"code": 1,
	"type": "BinFileUpload",
	"msg": "SUCCESS",
	"data": {
		"file_path": "/file/test_file.bin"
	}
}
```
