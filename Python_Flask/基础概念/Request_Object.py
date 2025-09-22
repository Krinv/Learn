# 基础概念-请求对象

from flask import Flask
from flask import request

app=Flask(__name__)

@app.route('/submit',methods=['GET','POST'])
def submit():

    method=request.method

    #获取查询参数
    query_params=request.args.to_dict()

    #获取表单数据
    form_data=request.form.to_dict()

    #简单响应
    if method=='GET':
        return f'GET请求参数：{query_params}'
    else:
        username=form_data.get('username','未知用户')
        return f'你好,{username}! 你的表单数据：{form_data}'

if __name__=='__main__':
    app.run()