# 博客蓝图
from flask import Blueprint

# 创建博客蓝图
blog_bp=Blueprint('blog',__name__)

@blog_bp.route('/posts')
def list_post():
    return '<h1>文章列表</h1> ' \
    '<ul>' \
    '<li>文章1</li>' \
    '<li>文章2</li>' \
    '</ul>'

@blog_bp.route('/post/<int:post_id>')
def show_post(post_id):
    return f'<h1>文章详情</h1><p>这是文章#{post_id}</p>'