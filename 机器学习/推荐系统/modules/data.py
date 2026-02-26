import pandas as pd

def load_enriched_movie_data(csv_path):
    """
    从 CSV 文件加载并预处理电影数据。
    这个函数专门用于处理 'movies_enriched.csv' 或类似结构的文件。
    它会查找 'average_rating' 和 'rating_count' 并将它们标准化为 'rating' 和 'members'。
    """
    try:
        df = pd.read_csv(csv_path)
        print(f"成功从 '{csv_path}' 加载 {len(df)} 条数据。")
    except FileNotFoundError:
        print(f"错误: CSV文件 '{csv_path}' 未找到。")
        return None
    except Exception as e:
        print(f"错误: 加载CSV文件 '{csv_path}' 失败: {e}")
        return None

    # 1. 检查并处理原始的评分和计数列
    required_original_cols = ['average_rating', 'rating_count']
    for col in required_original_cols:
        if col not in df.columns:
            print(f"错误: 数据中缺少必需的列 '{col}'。")
            return None
        df[col] = pd.to_numeric(df[col], errors='coerce')
    
    df.dropna(subset=required_original_cols, inplace=True)
    
    # 2. 根据分数总和 (rating_count) 和平均分 (average_rating) 计算评分人数
    # 避免除以零的错误
    df['members'] = df.apply(
        lambda row: int(row['rating_count'] / row['average_rating']) if row['average_rating'] > 0 else 0,
        axis=1
    )
    
    # 3. 将列名标准化为内部统一使用的名称，并删除不再需要的列
    df.rename(columns={'average_rating': 'rating'}, inplace=True)
    df.drop(columns=['rating_count'], inplace=True, errors='ignore')
    
    # 4. 为了兼容性，自动重命名ID和标题列
    if 'movieId' not in df.columns and 'anime_id' in df.columns:
        print("检测到 'moive_id' 列，将重命名列以保持一致性 ('anime_id'->'movieId', 'name'->'title', 'genre'->'genres')。")
        df = df.rename(columns={'anime_id': 'movieId', 'name': 'title', 'genre': 'genres'})

    print("数据预处理完成：已正确计算 'members' (评分人数) 并标准化列名。")
    return df

def load_user_ratings(db_connection):
    """
    从数据库加载所有用户评分数据。
    """
    try:
        query = "SELECT user_id, anime_id, rating FROM user_ratings"
        ratings_df = pd.read_sql(query, db_connection)
        print(f"成功从数据库加载 {len(ratings_df)} 条用户评分。")
        return ratings_df
    except Exception as e:
        print(f"错误: 从数据库加载用户评分失败: {e}")
        return pd.DataFrame()
