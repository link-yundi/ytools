# CoolPipeline

流水线：并发、池化

### 特性

- 池化，复用
- 流水线定制
- 多条流水线并发

### 示例

#### 多条流水线并发：

```go

// 定制 3 台手机, 两套流水线
// 步骤： 采购零件(耗时1s) -> 组装(耗时5s) -> 打包(耗时3s), 一套完整的流程耗时 9s

func main() {
    start := time.Now()
    // 采购
    buy := func(in any) (out any) {
        time.Sleep(1 * time.Second)
        i := in.(int)
        out = fmt.Sprint("零件", i)
        ylog.Info(out)
        return
    }
	// 组装
    build := func(in any) (out any) {
        time.Sleep(5 * time.Second)
        out = "组装(" + in.(string) + ")"
        ylog.Info(out)
        return
    }
    // 打包
    pack := func(in any) (out any) {
        time.Sleep(3 * time.Second)
        out = "打包(" + in.(string) + ")"
        ylog.Info(out)
        return
    }
    // 工作流定义: 2个并发, 两套流水线生产3台手机理论上需要18s
    pipeline := NewPipelines(2, buy, build, pack)
    // 订购3台
    var ins []any
    for i := 1; i <= 3; i++ {
        ins = append(ins, i)
    }
    pipeline.AddTask(ins...)
    end := time.Now()
    duration := end.Sub(start).Seconds()
    ylog.Info("耗时: ", duration, "s")
}

// Output:
// 2022/11/05 18:03:21 [INFO] [零件1]
// 2022/11/05 18:03:21 [INFO] [零件2]
// 2022/11/05 18:03:26 [INFO] [组装(零件1)]
// 2022/11/05 18:03:26 [INFO] [组装(零件2)]
// 2022/11/05 18:03:29 [INFO] [打包(组装(零件2))]
// 2022/11/05 18:03:29 [INFO] [打包(组装(零件1))]
// 2022/11/05 18:03:30 [INFO] [零件3]
// 2022/11/05 18:03:35 [INFO] [组装(零件3)]
// 2022/11/05 18:03:38 [INFO] [打包(组装(零件3))]
// 2022/11/05 18:03:38 [INFO] [耗时:  18.008715458 s]
```

