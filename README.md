# domainQuery

`domainQuery` 是一个用于对域名进行批量解析查询的工具。它可以解析域名的CNAME和A记录，并使用QQWry数据库查询解析出来的IP地址归属地。最终结果会输出到CSV文件中。

## 功能

- 解析域名的CNAME记录
- 解析域名的A记录
- 使用QQWry数据库查询IP地址归属地
- 将结果输出到CSV文件中

## 安装

**克隆仓库**

```bash
git clone https://github.com/your_username/domainQuery.git
cd domainQuery
```
**编译**

确保你已经安装了make环境，然后运行：

```
make
```

编译好的文件保存在：`./release/`目录中。

## 使用

不修改dns服务器和线程时，只`-f`指定域名列表文件即可。

```
Usage of ./domainQuery_1.0.0:
  -d string
    	输入dns服务器 (default "223.6.6.6")
  -f string
    	输入域名文件路径
  -t int
    	输入线程数 (default 10)
```

结果输出到当前目录下的csv文件中。