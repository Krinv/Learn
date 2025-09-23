# Flask

使用python编写的轻量级Web应用框架

基于WSGI（Web Server Gateway Interface）和Jinja2模板引擎

微框架，使用简单核心，用扩展增加功能

## 特点

1. 轻量级、简洁
2. 灵活：提供基本框架，无强制项目布局与组件
3. 可扩展性
4. 内置开发服务器
5. RESTful支持

## 组成

1. Flask应用实例：核心，通过创建Flask对象初始化应用
2. 路由与视图函数：
   路由将URL映射到视图函数
   视图函数处理请求并返回响应
3. 模板系统：使用Jinja2模板引擎渲染HTML页面，将数据动态插入页面
4. 请求和响应：处理HTTP请求，生成响应，支持多种HTTP方法


## Flask相关基本概念

### 路由 Routing

路由是URL到Python函数的映射，当访问特定URL时，Flask调用对应的视图函数处理请求

```python
from flask import Flask

app=Flask(__name__)

@app.route('/')		#将URL：/ 映射到home()函数
def home():
    return 'Welcome to the Home Page'

@app.route('/about')	#将URL：/about映射到about()函数
def about():
    return 'This is the About Page'
```

### 视图函数 View Functions

视图函数是处理请求并返回响应的Python函数。通常接收请求对象作为参数，并返回响应对象，或直接返回字符串、html等内容

```python
@app.route('/greet/<name>')
def greet(name):
    return f'Hello,{name}!'
```

### 请求对象 Request Object

请求对象包括客户端发送的请求信息，包括：请求方法、URL、请求头、表单数据等，Flask 用request对象来访问这些信息

```python
from flask import request

@app.route('/submit',methods=['POST'])
def submit():
    username=request.form.get('username')
    return f'Hello,{username}!'
```

### 响应对象 Response Object

响应对象包含发送给客户端的响应信息。响应信息（状态码、响应头、响应体）。Flask默认将字符串、HTML直接作为响应体

```python
from flask import make_response

@app.route('/custom_response')
def custom_response():
    response=make_response('THis is a custom response!')		//创建自定义响应对象
    response.headers['X-Custom-Header']='Value'		//设置响应头X-Custom-Header
    return response
```

### 模板

Flask使用Jinja2模板引擎渲染HTML，允许将python代码嵌入HTML，动态生成网页内容

```html
<!--hello.html内容-->
<html>
    <head>
        <title>Hello</title>
    </head>
    <body>
        <h1>
            Hello,{{name}}!		<!--在一些模板引擎中表示变量替换-->
        </h1>
    </body>
</html>    
```



```python
# 已有hello.html模板文件的前提下
from flask import render_template

@app.route('/hello/<name>')
def hello(name)
	return render_template('hello.html',name=name)
```

### 应用工厂 Application Factory

是一个py函数，允许你创建和配置多个Flask应用实例，或在不同配置下初始化应用

```python
def create_app(config_name)
	app=Flask(__name__)
	app.config.from_object(config_name)

	from . import routes
	app.register_blueprint(routes.bp)

	return app
```

