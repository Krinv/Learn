# 示例代码：基础概念-路由

from flask import Flask

app=Flask(__name__)

@app.route('/')
def home():
    return 'Welcome to Home Page!'

@app.route('/about')
def about():
    return 'This is About Page.'

if __name__=='__main__':
    app.run(debug=True)