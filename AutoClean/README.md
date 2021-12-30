# AutoClean

一个 go 编写的自动清理脚本


## 用法

只编译了 Android 版本

### Android

> 仅 root 可用，配合 Tasker 类食用

1. 下载项目[bin目录](https://github.com/SukiEva/Scripts/tree/main/AutoClean/bin)下的二进制文件（[蓝奏云](https://suki.lanzoui.com/iZNoNvht42d)），放到手机根目录 /data 目录下（只有data目录下才能授予执行权限）
2. 授予 autoclean 可执行权限(以下二选一)：
    - MT管理器等手动授权
    - 命令行获取su后输入 chmod 770 /data/autoclean
3. 在用户目录 Documents 中新建 AutoClean 文件夹（想改目录自己复制源码修改编译）
4. 在上述文件夹中新建 config.prop，每行一个文件/文件夹路径，支持通配符正则，示例如下：
```shell
# 详细说明见
# https://github.com/SukiEva/Scripts/tree/main/AutoClean
# 路径支持正则
# 一行一个

/storage/emulated/0/Documents/AutoClean/*.bak
```
5. 在 Tasker 中新建shell 命令 `/data/autoclean`，自行配置运行条件


## 其他

自带白名单保护：忽略以下的路径

```go
var (
	workDir    = "/data/media/0/Documents/AutoClean/"
	ignoreList = []string{
		"Android",
		"Android/data",
		"Android/media",
		"Android/obb",
		"DCIM",
		"Documents",
		"Download",
		"Movies",
		"Music",
		"Pictures",
	}
	sdcardList = []string{
		"/sdcard/",
		"/storage/emulated/0/",
		"/data/media/0/",
	}
)
```

