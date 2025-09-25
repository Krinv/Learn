# 基础概念-蓝图-主应用

from flask import Flask
# 导入自定义蓝图
from blueprints.auth import auth_bp
from blueprints.blog import blog_bp

app=Flask(__name__)

# 注册蓝图并设置url前缀
app.register_blueprint(auth_bp,url_prefix='/auth')
app.register_blueprint(blog_bp,url_prefix='/blog')

# 主应用路由
@app.route('/')
def  home():
    return '''
    <h1>WELCOME!</h1>
    <p><a href="/auth/login">登录</a></p>
    <p><a href="/auth/register">注册</a></p>
    <p><a href="/blog/posts">查看文章</a></p>
    <p><a href="/blog/post/1">查看文章1</a></p>
    '''

if __name__=='__main__':
    app.run(debug=True)