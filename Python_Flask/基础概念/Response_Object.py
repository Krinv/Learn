# 基础概念-响应对象
from flask import Flask,make_response
import urllib.parse

app=Flask(__name__)

# 首页
@app.route('/')
def index():
    return '''
    <h1>Flask响应对象演示</h1>
    <p><a href="/custom_response">自定义响应</a></p>
    <p><a href="/normal_response">普通响应</a></p>
    '''

# 普通响应页
@app.route('/normal_response')
def normal_response():
    return '这是普通响应，Flask自动创建响应对象'

# 自定义响应页
@app.route('/custom_response')
def custom_response():

    response=make_response("这是自定义响应！")

    # 自适应响应头部
    response.headers["Custom-Header"]='Value of Custom Header'
    response.headers['Content-Type']='text/plain;charset=utf-8'

    # 可选：包括中文时，使用URL编码实现
    encoded_value=urllib.parse.quote("中文具体值")
    response.headers['Encoded-Header']=encoded_value

    # 设置状态码
    response.status_code=200
    
    return  response


if __name__=='__main__':
    app.run()