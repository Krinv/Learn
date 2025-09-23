# 基础概念-模板
from flask import render_template,Flask

app=Flask(__name__)

@app.route('/hello/<name>')
def hello(name):
    # 默认会寻找当前目录下的templates文件夹，在里面寻文档
    # 如果要改默认文件夹，要在app实例化时添加入参
    return render_template('hello.html',name=name)

if __name__=='__main__':
    app.run()