## 插件调度插件


### 目录

```TXT
project/
 ├── main.go
 ├── plugins/
 │    ├── plugin-a/
 │    │    ├── plugin.yaml
 │    │    └── plugin-a.exe
 │    ├── plugin-b/
 │
 ├── runtime/
 │    └── plugin_state.json   ← 关键文件
```


### 依赖插件

go install github.com/spf13/cobra-cli@latest
* 引入理由：成熟的cli开发框架，高效易用。

gopkg.in/yaml.v3
* 引入理由： yaml 数据管理



### 扩展
 1.  PluginMan 内部拆分出 针对 plugin.YAML的 config 的管理器 和 status 的管理器
 2.  status 的管理器 抽象出 interface 把status.json 的 管理迁移到sqlite 或者etcd等