## 插件调度插件


### 目录

```TXT
plugin-sche/
 ├── main.go
 ├── plugins/   插件目录
 │    ├── plugin-a/
 │    │    ├── plugin.yaml 插件配置yaml
 │    │    └── plugin-a.exe 插件的可执行文件
 │    ├── plugin-b/
 │
 ├── runtime/  
 │    └── plugin_state.json   ← 关键文件 存储运行状态
 ├── cmd/    可执行命令
 ├── schema/    基本结构体
 ├── core/    核心调度器
 │    ├── plugin_man.go  插件的调度器
 │    ├── config_manager.go 插件配置调度器
      └── status_manager.go 插件状态调度器

```


### 依赖插件

go install github.com/spf13/cobra-cli@latest
* 引入理由：成熟的cli开发框架，高效易用。

gopkg.in/yaml.v3
* 引入理由： yaml 数据管理


### 使用说明

```BASH


$ ./plugin-sche 
My application description

Usage:
  plugin-sche [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  disable     禁用指定插件
  enable      启用指定插件
  genPlugin   生成插件的配置文件和Go文件
  genPlugin   生成插件的配置文件和Go文件
  help        Help about any command
  ls          查看当前插件列表
  pipline     数据按照顺序依次执行
  reload      重新生成status.json状态数据
  reload      重新生成status.json状态数据
  run         执行指定插件

Flags:
  -h, --help                help for plugin-sche
  -p, --plugins string      配置文件路径 (default "plugins")
  -t, --timeout int         超时时间(秒) (default 60)
  -v, --verbose             详细输出
  -w, --wrok-spcae string   运行状态文件 (default "runtime")
```


### 扩展
 1.  status 的管理器 已经抽象出 interface 把status.json 的 管理迁移到sqlite 或者etcd等
 2.  插件运行和针对版本进行统一抽出和管理
 3.  写入文件.lock空文件  实现多命令同时执行的不占用 status.json文件



 ### 执行步骤



```BASH

$ go build -o plugin-sche 
$ ./plugin-sche genPlugin plu1 
成功为 "plu1" 插件生成文件到 plugins\plu1
$ ./plugin-sche genPlugin plu2
成功为 "plu2" 插件生成文件到 plugins\plu2
$ ./plugin-sche genPlugin plu3
成功为 "plu3" 插件生成文件到 plugins\plu3
$ ./plugin-sche ls
NAME  VERSION ENTRY     ENABLED TIMEOUT STATUS
demo1 1.0     demo1.exe true    3s      idle
plu1  1.0     plu1.exe  true    3s      idle
plu2  1.0     plu2.exe  true    3s      idle
plu3  1.0     plu3.exe  true    3s      idle
$ ./plugin-sche disable plu2
$ ./plugin-sche run plu2
<nil>Error: 插件 plu2 已禁用
Error: 插件 plu2 已禁用
$ ./plugin-sche run plu1 '{"data":{"foo":"bar"}}'
&{success map[data:Hello from plu1! msg map[foo:bar]!] }
$ ./plugin-sche pipline plu1 plu3  -d '{"data":{"foo":"bar"}}'
output &{success map[data:Hello from plu1! msg map[foo:bar]!] } 
output &{success map[data:Hello from plu3! msg map[data:map[data:Hello from plu1! msg map[foo:bar]!]]!] } 
```
