# ylog

### 特性

- 级别
- 控制台+文件输出
- `error stack trace`

### 示例

```go
import log "ytools/ylog"

// 设置log level
ylog.SetLogLevel(log.LevelDebug)

// 通过string设置
ylog.SetLogLevelFromStr("info")

// 设置log file
logFilePath := "log/to/path"
log.SetLogFile(logFilePath)
```

#### panic

```go
func main() {
    panicFunc := func() {
        defer ylog.Panic() // 打印 panic ,不中断 运行
        a := make([]int, 0)
        fmt.Println(a[0])
    
    }
    panicFunc()
    fmt.Println("go on")
}

// Output:
// [PANIC] original err:runtime error: index out of range [0] with length 0
// [PANIC] stack trace: 
// go on 恢复运行
```

