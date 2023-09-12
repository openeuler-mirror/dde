# Changelog生成工具

- [项目简述](#项目简述)
- [编译](#编译)
- [使用](#使用)
- [编辑器](#编辑器)
- [许可证](#许可证)

## 项目简述

系统打包的日常工作中，需要根据git commit信息生成changelog。这往往需要开发输入多条命令，手动编写changelog格式，浪费了时间。  
本工具可以实现一键根据commit 信息生成rpm spec格式的changelog。

**自定义**

可以指定要生成changelog的commit条目数量，选择将某几条commit合并为一个changelog或忽略某些commit，在生成最终结果前进行自由调整

**简单上手**

实现的功能逻辑与执行 `git rebase -i` 逻辑相同，经过简单的文档查看即可迅速上手

## 编译
编译环境，Go>=1.18

```shell
git clone https://gitee.com/openeuler/dde.git
cd dde/develop/changelog
go build
```

## 使用
```shell
Usage of changelog:
  -c int
        num of changelog need to create (default 1)
  -e    useDefaultEditor, Windows: notepad; Linux: vi
  -g string
        choose the git path (default ".")
  -o string
        output file
  -s    use short entry
  -v string
        set changelog version

```

## 编辑器
使用的编辑器的优先级如下
1. 命令行参数 `-e`
2. 环境变量 `GIT_EDITOR`
3. Git设置：`core.editor`
4. 环境变量 `VISUAL`
5. 环境变量 `EDITOR`
6. 默认编辑器

## 许可证

[MulanPSL-2.0](http://license.coscl.org.cn/MulanPSL2/)


## 参考

[Git编辑器](https://mirrors.edge.kernel.org/pub/software/scm/git/docs/git-var.html)