# RESTful架构

Representational State Transfer	表征性状态转移，一组架构约束条件和原则



随着互联网和移动设备发展，Web应用的使用需求增加，传统的动态页面因低效率逐渐被HTML+JS的前后端分离替代，且Android、IOS、小程序等形式客户端层出不穷，客户端种类多元，要求一套结构清晰、符合标准、易于理解、扩展方便的接口风格。因而RESTful被实践流行

![img](https://pica.zhimg.com/v2-5016de164cff86f8aa48e9c108dc20bc_1440w.jpg)



### RESTful架构特征

1. 以资源为基础
2. 统一接口：对资源包括获取（GET）、创建（POST）、修改（PUT）、删除（DELETE），与HTTP协议提供的各方法对应。使用RESTful风格接口，你能够定位资源，但无法知晓资源进行了什么操作





## 资源与URL

表述性状态转移，指的是资源的表述。

Web中，资源的**唯一标识**是URI（Uniform Resource Identifier），可视作资源的名称/地址。

没有URI表示的信息，不能被视为资源，只能算是资源的一些信息。

### 可寻址性原则

URI设计应遵循可寻址性原则，拥有自描述性

1. 使用_或-提升可读性

   ```
   http://www.oschina.net/news/38119/oschina-translate-reward-plan
   ```

2. 使用/表示资源的层级关系

   ```
   /git/git/commit/e3af72cdafab5993d18fae056f87e1d675913d08
   表示一个多级的资源，这里是一个用户的commit记录
   ```

3. 使用？来过滤资源

   ```
   /git/git/pulls		git项目的所有推入请求
   /pulls?state=closed		git项目已经关闭的推入请求
   ```

4. 用,或；来表示同级资源关系

   ```
   /git/git /block-sha1/sha1.h/compare/5913d08;febca	假如要表示两个同级commit
   ```

   





## 统一资源接口

RESTful框架遵循统一接口原则

统一接口包含一组受限的预定义操作，不管什么资源都通过使用相同的接口进行资源的访问

按HTTP方法的语义在暴露资源时，接口拥有安全性（无论请求多少次，都不会改变服务器状态）和幂等性（无论对资源操作多少次，结果总是一样，后面的请求不会产生比第一次更多的影响）的特性



### GET

安全、幂等；

获取表示；变更获取表示（缓存）

```
典型用法
200（OK）		表示已在响应中发出
204（无内容）	资源有空表示
301（Moved Permanently）		资源URI已被更新
303（See Other）		其他（如负载均衡等）
304（Not Modified）		资源未更改（缓存）
400（Bad Request）		指代坏请求（如参数错误等）
404（Not Found）		资源不存在
406（Not Acceptable）		服务端不支持所需表示
500（Internal Server Error）		通用错误响应
503（Server Unavailable）		服务端当前无法处理请求
```



### POST

不安全、不幂等；

使用服务端管理（自动产生）的实例号创建资源；创建子资源；部分更新资源；若没有被修改，则不过更新资源（乐观锁）

```
典型用法
200（OK）		现有资源已被更改
201（Created） 	新资源被创建
202（Accepted）		已接受处理请求但尚未完成（异步处理）
301（Moved Permanently）		资源URI已被更新
303（See Other）		其他（如负载均衡等）
400（Bad Request）		指代坏请求（如参数错误等）
404（Not Found）		资源不存在
406（Not Acceptable）		服务端不支持所需表示
409（Conflict）		通用冲突
412（Precondition Failed）		前置条件失败（如执行条件更新时的冲突）
415（Unsupported Media Type）		接受到的表示不受支持
500（Internal Server Error）		通用错误响应
503（Server Unavailable）		服务端当前无法处理请求
```



### PUT

不安全、幂等；

通过客户端管理的实例号创建一个资源；通过替换的方式更新资源；如果未被修改，则更新资源（乐观锁）

```
200（OK）		现有资源已被更改
201（Created） 	新资源被创建
301（Moved Permanently）		资源URI已被更新
303（See Other）		其他（如负载均衡等）
400（Bad Request）		指代坏请求（如参数错误等）
404（Not Found）		资源不存在
406（Not Acceptable）		服务端不支持所需表示
409（Conflict）		通用冲突
412（Precondition Failed）		前置条件失败（如执行条件更新时的冲突）
415（Unsupported Media Type）		接受到的表示不受支持
500（Internal Server Error）		通用错误响应
503（Server Unavailable）		服务端当前无法处理请求
```



### DELETE

不安全、幂等；

删除资源

```
200（OK）		现有资源已被删除
301（Moved Permanently）		资源URI已被更改
303（See Other）		其他（如负载均衡等）
400（Bad Request）		指代坏请求（如参数错误等）
404（Not Found）		资源不存在
409（Conflict）		通用冲突
500（Internal Server Error）		通用错误响应
503（Server Unavailable）		服务端当前无法处理请求
```



### 常见问题

1. POST与PUT创建资源时的区别：
   创建资源的名称URI是否由客户端决定

2. 统一接口是否意味着不能扩展带特殊语义的方法：

   不阻止扩展方法，只要方法对资源有具体、可识别的语义，并能保持整个接口的统一性

3. 统一资源接口对URI的指导意义：

   统一资源接口要求使用标准HTTP方法对资源操作，因而URI只能表示资源的名称，而不该包括资源的操作（不使用动作描述）