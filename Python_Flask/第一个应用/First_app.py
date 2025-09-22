# 第一个Flask应用
from flask import Flask

# 创建Flask实例
# __name__特殊变量:
# 在模块被直接运行时，__name__='__main__'
# 在被其他模块导入时，是本模块的名字
app=Flask(__name__) 

#装饰器指定Flask哪一个URL触发下面的函数，这里指定网页根目录（主页）
@app.route('/')
def hello_world():
    return 'Hello World!'

#条件判断，本模块是否被直接执行，直接执行则运行代码块内容
if __name__=='__main__':
    #Flask实例的run方法：启动内置的开发服务器，入参为启动调试模式
    app.run(debug=True)