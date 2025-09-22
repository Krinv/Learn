# 基础概念-视图函数
from flask import Flask

app=Flask(__name__)

#实现取URL上的变量值
@app.route('/greet/<name>')
def greet(name):
    return f'Hello,{name}!'

if __name__=='__main__':
    app.run()