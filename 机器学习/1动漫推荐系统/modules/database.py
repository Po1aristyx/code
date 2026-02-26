import pandas as pd
import mysql.connector
from mysql.connector import errorcode

def connect_db():
    """连接到 MySQL 数据库"""
    try:
        conn = mysql.connector.connect(
            host='localhost',  # 请替换为你的 MySQL 主机名
            user='root',  # 请替换为你的 MySQL 用户名
            password='18858652154y',  # 请替换为你的 MySQL 密码
            database='anime_recommendation'  # 请替换为你的 MySQL 数据库名
        )
        return conn
    except mysql.connector.Error as err:
        if err.errno == errorcode.ER_ACCESS_DENIED_ERROR:
            print("用户名或密码错误")
        elif err.errno == errorcode.ER_BAD_DB_ERROR:
            print("数据库不存在")
        else:
            print(err)

def create_anime_table(conn):
    """创建动漫表"""
    cursor = conn.cursor()
    cursor.execute('''
    CREATE TABLE IF NOT EXISTS anime (
        anime_id INT PRIMARY KEY,
        name VARCHAR(255) NOT NULL,
        genre VARCHAR(255),
        type VARCHAR(255),  # 确保包含 type 列
        episodes INT,
        rating FLOAT,
        members INT
    )
    ''')
    conn.commit()

def create_user_table(conn):
    """创建用户表"""
    cursor = conn.cursor()
    cursor.execute('''
    CREATE TABLE IF NOT EXISTS users (
        user_id INT AUTO_INCREMENT PRIMARY KEY,
        username VARCHAR(255) NOT NULL UNIQUE,
        password VARCHAR(255) NOT NULL,
        email VARCHAR(255) NOT NULL,
        preferences VARCHAR(255) NOT NULL  # 存储用户喜爱的类型
    )
    ''')
    conn.commit()

def create_user_ratings_table(conn):
    """创建用户评分表"""
    cursor = conn.cursor()
    cursor.execute('''
    CREATE TABLE IF NOT EXISTS user_ratings (
        rating_id INT AUTO_INCREMENT PRIMARY KEY,
        user_id INT NOT NULL,
        anime_id INT NOT NULL,
        rating INT NOT NULL
    )
    ''')
    conn.commit()


def load_csv_to_db(conn, anime_csv, rating_csv):
    """加载CSV文件数据到数据库"""
    anime_data = pd.read_csv(anime_csv)
    rating_data = pd.read_csv(rating_csv)

    # 处理缺失值，使用适当的默认值
    anime_data['name'].fillna('Unknown', inplace=True)
    anime_data['genre'].fillna('Unknown', inplace=True)
    anime_data['type'].fillna('Unknown', inplace=True)
    anime_data['episodes'].fillna(0, inplace=True)
    anime_data['rating'].fillna(0, inplace=True)
    anime_data['members'].fillna(0, inplace=True)

    # 将 'Unknown' 转换为适当的默认值，确保数值列是正确的格式
    anime_data['episodes'] = pd.to_numeric(anime_data['episodes'], errors='coerce').fillna(0).astype(int)
    anime_data['rating'] = pd.to_numeric(anime_data['rating'], errors='coerce').fillna(0)
    anime_data['members'] = pd.to_numeric(anime_data['members'], errors='coerce').fillna(0).astype(int)

    # 将评分中的 -1 值转换为 0（未评分）
    rating_data['rating'] = rating_data['rating'].replace(-1, 0)

    cursor = conn.cursor()

    # 插入动漫数据
    for _, row in anime_data.iterrows():
        cursor.execute('''
        INSERT INTO anime (anime_id, name, genre, type, episodes, rating, members)
        VALUES (%s, %s, %s, %s, %s, %s, %s)
        ON DUPLICATE KEY UPDATE name=%s, genre=%s, type=%s, episodes=%s, rating=%s, members=%s
        ''', (int(row['anime_id']), row['name'], row['genre'], row['type'], int(row['episodes']), float(row['rating']),
              int(row['members']),
              row['name'], row['genre'], row['type'], int(row['episodes']), float(row['rating']), int(row['members'])))

    # 插入评分数据到 ratings 表
    for _, row in rating_data.iterrows():
        cursor.execute('''
        INSERT INTO ratings (user_id, anime_id, rating)
        VALUES (%s, %s, %s)
        ON DUPLICATE KEY UPDATE rating=%s
        ''', (int(row['user_id']), int(row['anime_id']), int(row['rating']), int(row['rating'])))

    conn.commit()

if __name__ == '__main__':
    conn = connect_db()
    if conn:
        create_anime_table(conn)
        create_user_table(conn)
        create_user_ratings_table(conn)
        # 使用实际的CSV文件路径替换下面的路径
        load_csv_to_db(conn, 'D:/PYTHON-project/mssb-course_design/dataes/anime.csv',
                       'D:/PYTHON-project/mssb-course_design/dataes/rating.csv')
        conn.close()
        print("数据加载成功！")
