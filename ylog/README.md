# ylog

### 安装

```sh
go get -u github.com/link-yundi/ytools
```

### 特性

- 级别
- 控制台+文件输出
- `error stack trace`

### 示例

```go
import log "github.com/link-yundi/ytools/ylog"

// 设置log level
ylog.SetLogLevel(log.LevelDebug)

// 通过string设置
ylog.SetLogLevelFromStr("info")

// 设置log file
logFilePath := "log/to/path"
log.SetLogFile(logFilePath)
```

