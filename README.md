## 实现selpg

课程链接：[CLI 命令行实用程序开发基础](https://pmlpml.github.io/ServiceComputingOnCloud/ex-cli-basic)

开发 Linux 命令行实用程序，Linux命令行程序设计

### 要求

实现一个selpg命令行程序，能够输入命令行或文件的起始页码和结束页码，将内容输出到命令行或文件中。

### 输入

标准输入：
```
[-s start_page] [-e end_page] [-l page_len ] [-f page_type] [-d destination ] [filename], program_name
```
例如：
```
selpg -s 1 -e 1 test.txt
```

### 输出
例如：
```
test1
test2
test3
...
```

对于
```
selpg -s 1 -e 1 test.txt > out.txt
```
已将out.txt上传至selpg文件夹中

### 测试
![image](http://wx2.sinaimg.cn/mw690/932e8e0cgy1fzc0kmxxrvj20g004ta9u.jpg)