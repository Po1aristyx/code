import re
import logging
from modules.database import connect_db

# 设置日志
logging.basicConfig(level=logging.INFO, format='%(asctime)s - %(levelname)s - %(message)s')


def validate_email(email):
    """验证电子邮件是否符合规范"""
    return re.match(r'^[^@]+@[^@]+\.[^@]+$', email) is not None

def register_user(username, password, email, preferences):
    """注册新用户"""
    if not validate_email(email):
        logging.error("电子邮件格式不正确")
        return False

    conn = connect_db()
    cursor = conn.cursor()
    try:
        cursor.execute('INSERT INTO users (username, password, email, preferences) VALUES (%s, %s, %s, %s)',
                       (username, password, email, ','.join(preferences)))
        conn.commit()
        logging.info("用户注册成功")
        return True
    except Exception as e:
        logging.error(f"注册用户时出错: {e}")
        return False
    finally:
        conn.close()

def login_user(username, password):
    """用户登录"""
    conn = connect_db()
    cursor = conn.cursor()
    try:
        cursor.execute('SELECT user_id FROM users WHERE username=%s AND password=%s', (username, password))
        user = cursor.fetchone()
        return user[0] if user else None
    except Exception as e:
        logging.error(f"用户登录时出错: {e}")
        return None
    finally:
        conn.close()

def get_user_preferences(username):
    """获取用户偏好"""
    conn = connect_db()
    cursor = conn.cursor()
    try:
        cursor.execute('SELECT preferences FROM users WHERE username=%s', (username,))
        preferences = cursor.fetchone()
        return preferences[0].split(',') if preferences else []
    except Exception as e:
        logging.error(f"获取用户偏好时出错: {e}")
        return []
    finally:
        conn.close()
