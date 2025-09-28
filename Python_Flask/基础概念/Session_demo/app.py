# 基础概念-会话
from flask import Flask,session,request,render_template

app=Flask(__name__)
app.secret_key='secret_key'

@app.route('/',methods=['GET','POST'])
def index():
    # 初始化会话数据:在会话中创建visit_count项
    if 'visit_count' not in session:
        session['visit_count']=0

    # 处理表单提交
    message=""

    # 当用户名输入框填写完毕点击完成会发送HTTP POST请求
    if request.method=='POST':
        username=request.form.get('username')

        if username:
            session['username']=username
            message=f'用户名已经设置为：{username}'
            message_type='success'
        else:
            message = '请输入用户名'
            message_type='error'
    else:
        message_type=''

    # 默认值为0
    session['visit_count']=session.get('visit_count',0)+1

    # 获取当前会话数据
    username=session.get('username','未设置')
    visit_count=session.get('visit_count',0)

    return render_template('index.html',
                        username=username,
                        visit_count=visit_count,
                        message=message,
                        message_type=message_type
                        )

@app.route('/clear')
def clear_session():
    session.clear()

    return render_template('index.html',
                        username='未设置',
                        visit_count=0,
                        message='所有会话数据已经清除',
                        message_type='success'
                        ) 

if __name__=='__main__':
    app.run(debug=True)