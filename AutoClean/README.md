# AutoClean

一个 go 编写的自动清理脚本


## 用法

只编译了 Android 版本

### Android

> 仅 root 可用，配合 Tasker 类食用

1. 下载项目[bin目录](https://github.com/SukiEva/Scripts/tree/main/AutoClean/bin)下的二进制文件，放到手机根目录 /data 目录下（只有data目录下才能授予执行权限）
2. 在用户目录 Documents 中新建 AutoClean 文件夹（想改目录自己复制源码修改编译）
3. 在上述文件夹中新建 config.json，示例如下：支持正则
```json
{
  "blackList": [
    "/storage/emulated/0/Documents/AutoClean/*.bak",
    "/storage/emulated/0/Download/.*"
  ]
}
```

## 其他

娱乐作品，随缘更新
