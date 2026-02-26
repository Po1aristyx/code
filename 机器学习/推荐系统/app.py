from flask import Flask, render_template, request, redirect, url_for, session, flash
from modules.database import connect_db, create_user_table, create_user_ratings_table, add_user_rating
from modules.recommendation import get_recommendations
import pandas as pd
import hashlib
from io import StringIO
import csv

app = Flask(__name__)
app.secret_key = 'your_very_secret_key'  # 生产环境中请使用更安全的密钥

def get_db_connection():
    """获取数据库连接"""
    return connect_db()

@app.route('/')
def index():
    """主页，如果已登录则重定向到推荐页面，否则重定向到登录页面。"""
    if 'user_id' in session:
        return redirect(url_for('recommendations_page'))
    return redirect(url_for('login'))

@app.route('/login', methods=['GET', 'POST'])
def login():
    """用户登录"""
    if request.method == 'POST':
        username = request.form['username']
        password = request.form['password']
        
        conn = get_db_connection()
        cursor = conn.cursor(dictionary=True)
        
        # 简单的密码哈希处理（生产环境建议使用更安全的库如 passlib）
        hashed_password = hashlib.sha256(password.encode()).hexdigest()
        
        cursor.execute("SELECT * FROM users WHERE username = %s AND password = %s", (username, hashed_password))
        user = cursor.fetchone()
        
        cursor.close()
        conn.close()
        
        if user:
            session['user_id'] = user['user_id']
            session['username'] = user['username']
            session['preferences'] = user['preferences'].split(',') if user['preferences'] else []
            flash('登录成功!', 'success')
            return redirect(url_for('recommendations_page'))
        else:
            flash('用户名或密码错误', 'danger')
            
    return render_template('login.html')

@app.route('/register', methods=['GET', 'POST'])
def register():
    """用户注册"""
    if request.method == 'POST':
        username = request.form['username']
        password = request.form['password']
        email = request.form['email']
        preferences = request.form.getlist('preferences')
        
        conn = get_db_connection()
        cursor = conn.cursor()
        
        # 检查用户名是否已存在
        cursor.execute("SELECT * FROM users WHERE username = %s", (username,))
        if cursor.fetchone():
            flash('用户名已存在!', 'danger')
            return redirect(url_for('register'))
            
        # 密码哈希
        hashed_password = hashlib.sha256(password.encode()).hexdigest()
        preferences_str = ','.join(preferences)
        
        try:
            cursor.execute(
                "INSERT INTO users (username, password, email, preferences) VALUES (%s, %s, %s, %s)",
                (username, hashed_password, email, preferences_str)
            )
            conn.commit()
            flash('注册成功，请登录!', 'success')
            return redirect(url_for('login'))
        except Exception as e:
            conn.rollback()
            flash(f'注册失败: {e}', 'danger')
        finally:
            cursor.close()
            conn.close()
            
    return render_template('register.html')

@app.route('/recommendations', methods=['GET', 'POST'])
def recommendations_page():
    if 'user_id' not in session:
        return redirect(url_for('login'))

    user_id = session['user_id']
    recommendation_type = 'top_rated_global'
    liked_movies = None

    if request.method == 'POST':
        recommendation_type = request.form.get('recommendation_type', 'top_rated_global')
        if recommendation_type == 'content_based':
            liked_movies_str = request.form.get('liked_movies')
            if not liked_movies_str or not liked_movies_str.strip():
                flash('进行"基于内容的推荐"时，请输入至少一部您喜欢的电影名称。', 'warning')
                return redirect(url_for('recommendations_page'))

            # 使用更健壮的CSV解析器来处理带引号的标题
            string_io = StringIO(liked_movies_str)
            reader = csv.reader(string_io, skipinitialspace=True)
            # 将解析出的列表扁平化
            liked_movies = [item for sublist in reader for item in sublist if item]

    recs_df, message = get_recommendations(
        user_id=user_id,
        recommendation_type=recommendation_type,
        liked_movies=liked_movies
    )

    if message:
        flash(message, 'info')

    # 从数据库获取用户已评分的电影，以在前端展示评分
    db_conn = get_db_connection()
    user_ratings = {} # 默认为空字典
    if db_conn:
        try:
            cursor = db_conn.cursor(dictionary=True)
            query = "SELECT movie_id, rating FROM user_ratings WHERE user_id = %s"
            cursor.execute(query, (user_id,))
            user_ratings_list = cursor.fetchall() # 使用 fetchall() 获取所有结果
            user_ratings = {item['movie_id']: item['rating'] for item in user_ratings_list}
        except Exception as e:
            flash(f"获取用户评分时出错: {e}", "error")
        finally:
            db_conn.close()

    recommendations = recs_df.to_dict('records') if not recs_df.empty else []
    
    return render_template('recommendations.html', 
                           recommendations=recommendations, 
                           user_ratings=user_ratings)

@app.route('/rate_movie', methods=['POST'])
def rate_movie():
    """处理用户提交的电影评分"""
    if 'user_id' not in session:
        return redirect(url_for('login'))

    user_id = session['user_id']
    movie_id = request.form.get('movie_id')
    rating = request.form.get('rating')

    if not movie_id or not rating:
        flash('评分提交失败，缺少信息。', 'danger')
        return redirect(url_for('recommendations_page'))

    try:
        conn = connect_db()
        if conn:
            if add_user_rating(conn, user_id, int(movie_id), int(rating)):
                flash('评分成功！', 'success')
            else:
                flash('评分失败，请稍后重试。', 'danger')
            conn.close()
        else:
            flash('无法连接到数据库。', 'danger')
    except ValueError:
        flash('无效的评分或电影ID。', 'danger')
    
    return redirect(url_for('recommendations_page'))

@app.route('/logout')
def logout():
    """用户登出"""
    session.clear()
    flash('您已成功登出。', 'info')
    return redirect(url_for('login'))

if __name__ == '__main__':
    conn = get_db_connection()
    if conn:
        # 确保启动时表已创建
        create_user_table(conn)
        create_user_ratings_table(conn)
        conn.close()
    app.run(debug=True) 