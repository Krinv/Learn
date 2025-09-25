# 基础概念-应用工厂

from flask import Flask

def create_app():
    app=Flask(__name__)

    @app.route('/')
    def home():
        return '<h1>简单应用工厂示例</h1>'
    
    @app.route('/hello')
    def  hello():
        return '<h1>Hello World!</h1>'
    
    return


# 通过该函数，则可以批量实例化具体的app，并且可以通过定义入参的方式，实现批量定义、修改配置的效果