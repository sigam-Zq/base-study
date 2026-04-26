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
  pipline     数据按照顺序依次执行插件[数据串联]
  reload      重新生成status.json状态数据
  run         执行指定插件

Flags:
  -h, --help                help for plugin-sche
  -p, --plugins string      配置文件路径 (default "plugins")
  -t, --timeout int         超时时间(秒) (default 60)
  -v, --verbose             详细输出
  -w, --wrok-spcae string   运行状态文件 (default "runtime")
```

### 选择实现
以下内容不要求全部实现，可选择性完成或仅给出设计说明:
1. 插件热加载/热卸载机制 √
> 直接读取plugin目录进行展示等价于热加载，但是依赖配置的plugin.yaml记录元数据

2. 插件执行超时控制
 > -t 参数可传入限制运行时间，不传以配置文件 plugin.yaml 超时时间为准
3. 插件隔离方案(如进程隔离、资源隔离思路)
 > 二进制运行本身就是进程隔离 可以通过 -w 参数实现资源隔离 -p 指定不同插件目录进行二进制资源隔离
4. 插件依赖关系与版本约束
 > 未涉及，只依赖yaml记录最基本的版本号
5. 插件执行结果的失败隔离与降级策略
 > 二进制运行本身为单次运行，不存在守护进程 失败本身就是隔离开的
6. 不同类型插件支持(只要满足插件规范，可以对接Js、Python、Golang等任意语言任意形式的插件)
 > 当前插件基于os/exec运行，只要可以打包二进制均可以执行，后期可以考虑识别后缀调用宿主机本身的python，node，goland执行其他语言
- 插件规范需要满足json格式  具体规范详见schema\base_schema.go的文件结构

### 扩展
 1.  status 的管理器 已经抽象出 interface 把status.json 的 管理迁移到sqlite 或者etcd等
 2.  插件运行和针对版本进行统一抽出和管理
 3.  写入文件.lock空文件  实现多命令同时执行的不占用 status.json文件



 ### 执行步骤

 **当前二进制均基于windows进行打包，更换运行系统需要重新进行打包**

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


$ ./plugin-sche run timeOut '{"data":{"foo":"bar"}}' -t 4
<nil>Error: 插件 timeOut 执行超时 (4s)
Error: 插件 timeOut 执行超时 (4s)
```
