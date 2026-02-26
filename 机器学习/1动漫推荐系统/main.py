from flask import Flask, render_template, request, redirect, url_for, session, jsonify, logging
from modules.database import connect_db
from modules.user import register_user, login_user, get_user_preferences
from modules.data import load_data, preprocess_data
from modules.recommendation import (recommend_based_on_survey, get_recommendations)

app = Flask(__name__)
app.secret_key = 'your_secret_key'


@app.route('/')
def home():
    return redirect(url_for('login'))


@app.route('/login', methods=['GET', 'POST'])
def login():
    if request.method == 'POST':
        username = request.form['username']
        password = request.form['password']
        user_id = login_user(username, password)
        if user_id:
            session['user_id'] = user_id
            session['username'] = username
            session['preferences'] = get_user_preferences(username)  # 获取用户偏好
            return redirect(url_for('recommendations'))
        return '登录失败，用户名或密码错误'
    return render_template('login.html')



@app.route('/register', methods=['GET', 'POST'])
def register():
    if request.method == 'POST':
        username = request.form['username']
        password = request.form['password']
        email = request.form['email']
        preferences = request.form.getlist('preferences')
        register_user(username, password, email, preferences)
        return redirect(url_for('login'))
    return render_template('register.html')


@app.route('/recommendations', methods=['GET'])
def recommendations():
    if 'user_id' in session:
        user_id = session['user_id']
        preferences = session.get('preferences', ['Action', 'Comedy'])

        anime_path = r'D:\PYTHON-project\mssb-course_design\dataes\anime.csv'
        rating_path = r'D:\PYTHON-project\mssb-course_design\dataes\rating.csv'
        anime_data, rating_data = load_data(anime_path, rating_path)
        anime_data, rating_data = preprocess_data(anime_data, rating_data)

        # 推荐结果
        recommendations = get_recommendations(user_id, rating_data, anime_data, preferences, use_cache=False)

        return render_template('recommendations.html', recommendations=recommendations.to_dict(orient='records'))
    return redirect(url_for('login'))


@app.route('/rate_anime', methods=['POST'])
def rate_anime():
    if 'user_id' in session:
        user_id = session['user_id']
        data = request.json
        anime_id = data['anime_id']
        rating = data['rating']
        try:
            conn = connect_db()
            cursor = conn.cursor()
            cursor.execute('''
            INSERT INTO user_ratings (user_id, anime_id, rating) VALUES (%s, %s, %s)
            ON DUPLICATE KEY UPDATE rating=%s
            ''', (user_id, anime_id, rating, rating))
            conn.commit()
            return jsonify(success=True, message='评分成功')
        except Exception as e:
            return jsonify(success=False, message=f'评分失败: {e}')
        finally:
            conn.close()
    return jsonify(success=False, message='用户未登录')



@app.route('/logout')
def logout():
    session.pop('user_id', None)
    session.pop('username', None)
    session.pop('preferences', None)
    return redirect(url_for('login'))


if __name__ == '__main__':
    app.run(debug=True)