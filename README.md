# http

### [http请求中加号被替换为空格](escape)

在 RFC 标准中会将数据字段中的空格编码成加号，而加号是不变的

所以在获取字段时无法区分，只会将加号都解码成空格，造成错误

解决方法:

1. 提前编码好，加号转成`%2B`
2. 直接拿到`r.URL.RawQuery`，自行处理

Ref:

* [http请求中加号被替换为空格](https://www.cnblogs.com/thisiswhy/p/12119126.html)
* [URL encoding the space character: + or %20?](https://stackoverflow.com/questions/1634271/url-encoding-the-space-character-or-20)

