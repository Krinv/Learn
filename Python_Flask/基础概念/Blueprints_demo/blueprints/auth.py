# 认证蓝图
from flask import Blueprint

# 创建认证蓝图
auth_bp=Blueprint('auth',__name__)

@auth_bp.route('/login')
def login():
    return '<h1>登录页面</h1><p>这是认证模块</p>'

@auth_bp.route('/register')
def register():
    return '<h1>注册页面</h1><p>新用户注册</p>'