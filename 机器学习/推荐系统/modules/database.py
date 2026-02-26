import mysql.connector
import pandas as pd
from sqlalchemy import create_engine
from urllib.parse import quote_plus

# 请将这里的配置替换成您自己的数据库连接信息
DB_CONFIG = {
    'host': 'localhost',
    'user': 'root',
    'password': 'Tyx@050529',
    'database': 'movies_recommendation'

}

def connect_db():
    """连接到MySQL数据库"""
    try:
        conn = mysql.connector.connect(**DB_CONFIG)
        return conn
    except mysql.connector.Error as err:
        print(f"数据库连接失败: {err}")
        return None

def get_db_engine():
    """为Pandas创建一个 SQLAlchemy 引擎"""
    try:
        # 对密码进行URL编码，以安全处理特殊字符
        encoded_password = quote_plus(DB_CONFIG['password'])
        
        db_uri = (
            f"mysql+mysqlconnector://{DB_CONFIG['user']}:{encoded_password}"
            f"@{DB_CONFIG['host']}/{DB_CONFIG['database']}"
        )
        engine = create_engine(db_uri)
        return engine
    except Exception as e:
        print(f"创建 SQLAlchemy 引擎失败: {e}")
        return None

def create_movie_table(conn):
    """根据 movies_enriched.csv 的结构创建电影表 (movie)"""
    cursor = conn.cursor()
    cursor.execute('''
    CREATE TABLE IF NOT EXISTS movie (
        movie_id INT PRIMARY KEY,
        name VARCHAR(255),
        genre VARCHAR(255),
        rating FLOAT,
        members INT
    )
    ''')
    conn.commit()

def create_user_table(conn):
    """创建用户表 (users)"""
    cursor = conn.cursor()
    cursor.execute('''
    CREATE TABLE IF NOT EXISTS users (
        user_id INT AUTO_INCREMENT PRIMARY KEY,
        username VARCHAR(255) NOT NULL UNIQUE,
        password VARCHAR(255) NOT NULL,
        email VARCHAR(255) NOT NULL,
        preferences TEXT
    )
    ''')
    conn.commit()

def create_user_ratings_table(conn):
    """创建用户评分表 (user_ratings)"""
    cursor = conn.cursor()
    cursor.execute('''
    CREATE TABLE IF NOT EXISTS user_ratings (
        rating_id INT AUTO_INCREMENT PRIMARY KEY,
        user_id INT NOT NULL,
        movie_id INT NOT NULL,
        rating INT NOT NULL,
        UNIQUE KEY `user_movie_unique` (`user_id`, `movie_id`)
    )
    ''')
    conn.commit()

def add_user_rating(conn, user_id, movie_id, rating):
    """向 user_ratings 表中添加或更新一条评分记录"""
    try:
        cursor = conn.cursor()
        query = """
        INSERT INTO user_ratings (user_id, movie_id, rating)
        VALUES (%s, %s, %s)
        ON DUPLICATE KEY UPDATE rating=VALUES(rating)
        """
        cursor.execute(query, (user_id, movie_id, rating))
        conn.commit()
        cursor.close()
        return True
    except Exception as e:
        print(f"添加评分时出错: {e}")
        conn.rollback()
        return False

def load_csv_to_db(conn, enriched_movie_path, rating_csv_path):
    """加载电影元数据和10000条外部评分数据到数据库。"""
    try:
        movie_data = pd.read_csv(enriched_movie_path)
        movie_data.rename(columns={
            'movieId': 'movie_id', 'title': 'name', 'genres': 'genre', 'average_rating': 'rating'
        }, inplace=True)
        
        # 修复 'members' 列的计算逻辑
        movie_data['members'] = movie_data.apply(
            lambda row: int(row['rating_count'] / row['rating']) if row['rating'] > 0 else 0, axis=1
        )
        
        final_movie_data = movie_data[['movie_id', 'name', 'genre', 'rating', 'members']]

    except KeyError as e:
        print(f"处理CSV文件时出错: 找不到列 {e}。请检查 'movies_enriched.csv' 的列名是否正确。")
        return
    except Exception as e:
        print(f"处理CSV文件时出错: {e}")
        return

    cursor = conn.cursor()

    insert_movie_query = """
    INSERT INTO movie (movie_id, name, genre, rating, members)
    VALUES (%s, %s, %s, %s, %s)
    ON DUPLICATE KEY UPDATE 
        name=VALUES(name), genre=VALUES(genre), rating=VALUES(rating), members=VALUES(members)
    """
    movie_records = [tuple(row) for row in final_movie_data.to_numpy()]
    
    try:
        cursor.executemany(insert_movie_query, movie_records)
        conn.commit()
        print(f"成功插入或更新 {len(movie_records)} 条电影数据到 'movie' 表。")
    except Exception as e:
        print(f"插入电影数据时出错: {e}")
        conn.rollback()
        return # 如果电影数据插入失败，则不继续

    # 2. 加载并插入前10000条评分数据
    try:
        rating_data = pd.read_csv(rating_csv_path, nrows=10000)
        rating_data.rename(columns={'userId': 'user_id', 'movieId': 'movie_id'}, inplace=True)
        final_rating_data = rating_data[['user_id', 'movie_id', 'rating']]

        insert_rating_query = """
        INSERT INTO user_ratings (user_id, movie_id, rating)
        VALUES (%s, %s, %s)
        ON DUPLICATE KEY UPDATE rating=VALUES(rating)
        """
        rating_records = [tuple(row) for row in final_rating_data.to_numpy()]
        
        cursor.executemany(insert_rating_query, rating_records)
        conn.commit()
        print(f"成功插入 {len(rating_records)} 条外部评分数据到 'user_ratings' 表。")

    except FileNotFoundError:
        print(f"警告: 未找到评分文件 '{rating_csv_path}'。将跳过评分数据加载。")
    except Exception as e:
        print(f"加载或插入评分数据时出错: {e}")
        conn.rollback()

if __name__ == '__main__':
    db_conn = connect_db()
    if db_conn:
        print("成功连接到数据库。")
        
        print("正在检查/创建数据表...")
        create_movie_table(db_conn)
        create_user_table(db_conn)
        create_user_ratings_table(db_conn)
        print("数据表创建/检查完成。")
        
        print("正在从CSV文件加载电影元数据和10000条评分(采用ON DUPLICATE KEY UPDATE来避免重复)...")
        import os
        script_dir = os.path.dirname(os.path.abspath(__file__))
        project_root = os.path.dirname(script_dir)
        movie_file_path = os.path.join(project_root, 'dataes', 'movies_enriched.csv')
        rating_file_path = os.path.join(project_root, 'dataes', 'ratings.csv')
        
        load_csv_to_db(db_conn, movie_file_path, rating_file_path)
        
        db_conn.close()
        print("\n数据库数据补充/更新完成！")
    else:
        print("无法连接到数据库，初始化失败。")
