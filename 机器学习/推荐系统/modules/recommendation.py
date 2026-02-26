import pandas as pd
from sklearn.feature_extraction.text import TfidfVectorizer
from sklearn.metrics.pairwise import cosine_similarity
from modules.database import get_db_engine
import re

def recommend_top_rated_global(movie_df, min_ratings_threshold=1000, num_recommendations=10):
    """
    推荐全局评分最高、且评分人数达到阈值的电影。
    """
    qualified_movies = movie_df[movie_df['members'] >= min_ratings_threshold].copy()
    if qualified_movies.empty:
        return pd.DataFrame()
    
    # 使用加权评分或其他更复杂的评分机制可能会更好，但这里为了简单直接用平均分
    # 此处假设 'rating' 列是电影的平均分
    top_movies = qualified_movies.sort_values(by='rating', ascending=False).head(num_recommendations)
    return top_movies

def recommend_content_based(liked_movies, movie_df, num_recommendations=10):
    """
    基于内容的推荐。
    liked_movies: 一个包含用户喜欢电影名称的列表。
    """
    if not liked_movies or movie_df.empty:
        return pd.DataFrame()

    # 使用更灵活的、不区分大小写的搜索来查找匹配的电影
    # 将用户输入的电影列表转换为一个正则表达式，每个电影名用 | 分隔
    # 正则表达式中的特殊字符需要转义
    search_pattern = '|'.join([re.escape(s) for s in liked_movies])
    
    # 查找包含任何一个喜欢电影名字的电影 (case-insensitive)
    liked_movie_mask = movie_df['name'].str.contains(search_pattern, case=False, na=False)
    liked_movie_indices = movie_df[liked_movie_mask].index

    if len(liked_movie_indices) == 0:
        # 如果找不到完全匹配的，可以返回一个空结果或者一个提示
        return pd.DataFrame()

    movie_df['genre'] = movie_df['genre'].fillna('')
    tfidf = TfidfVectorizer(stop_words='english')
    tfidf_matrix = tfidf.fit_transform(movie_df['genre'])
    
    liked_movie_matrix = tfidf_matrix[liked_movie_indices]
    
    cosine_sim = cosine_similarity(liked_movie_matrix, tfidf_matrix)
    
    # 计算平均相似度
    sim_scores = cosine_sim.mean(axis=0)
    
    sim_scores = list(enumerate(sim_scores))
    sim_scores = sorted(sim_scores, key=lambda x: x[1], reverse=True)
    
    # 获取最相似的电影索引，排除掉用户已经喜欢的
    top_indices = [i[0] for i in sim_scores if i[0] not in liked_movie_indices][:num_recommendations]
    
    return movie_df.iloc[top_indices]

def recommend_user_based_cf(user_id, movie_df, ratings_df, num_recommendations=10):
    if ratings_df.empty or user_id not in ratings_df['user_id'].values:
        return pd.DataFrame()

    user_item_matrix = ratings_df.pivot_table(index='user_id', columns='movie_id', values='rating').fillna(0)
    
    user_similarity = cosine_similarity(user_item_matrix)
    user_sim_df = pd.DataFrame(user_similarity, index=user_item_matrix.index, columns=user_item_matrix.index)

    similar_users = user_sim_df[user_id][user_sim_df[user_id] > 0].sort_values(ascending=False)
    
    if similar_users.empty:
        return pd.DataFrame()
        
    similar_users = similar_users.head(50)

    similar_users_watched = user_item_matrix.loc[similar_users.index]
    
    recommendations = similar_users_watched.apply(lambda row: row * similar_users, axis=1).sum()
    
    user_watched = user_item_matrix.loc[user_id][user_item_matrix.loc[user_id] > 0].index
    recommendations = recommendations.drop(user_watched, errors='ignore')
    
    top_recs_ids = recommendations.sort_values(ascending=False).index[:num_recommendations]
    return movie_df[movie_df['movie_id'].isin(top_recs_ids)]

def recommend_item_based_cf(user_id, movie_df, ratings_df, num_recommendations=10):
    if ratings_df.empty or user_id not in ratings_df['user_id'].values:
        return pd.DataFrame()

    user_ratings = ratings_df[ratings_df['user_id'] == user_id]
    if user_ratings.empty:
        return pd.DataFrame()
        
    item_user_matrix = ratings_df.pivot_table(index='movie_id', columns='user_id', values='rating').fillna(0)
    
    item_similarity = cosine_similarity(item_user_matrix)
    item_sim_df = pd.DataFrame(item_similarity, index=item_user_matrix.index, columns=item_user_matrix.index)

    all_recs = pd.Series(dtype=float)
    for _, row in user_ratings.iterrows():
        movie_id, rating = row['movie_id'], row['rating']
        if movie_id in item_sim_df:
            sim_scores = item_sim_df[movie_id].drop(movie_id, errors='ignore')
            sim_scores = sim_scores * rating
            all_recs = all_recs.add(sim_scores, fill_value=0)

    user_watched = user_ratings['movie_id'].tolist()
    all_recs = all_recs.drop(user_watched, errors='ignore')

    top_recs_ids = all_recs.sort_values(ascending=False).index[:num_recommendations]
    return movie_df[movie_df['movie_id'].isin(top_recs_ids)]


def get_recommendations(user_id, recommendation_type, liked_movies=None, num_recommendations=10):
    try:
        movie_df = pd.read_csv('dataes/movies_enriched.csv')
        
        # [修复] 对加载的DataFrame进行与数据库初始化时相同的转换
        movie_df.rename(columns={
            'movieId': 'movie_id', 'title': 'name', 'genres': 'genre', 'average_rating': 'rating'
        }, inplace=True)
        
        # 检查 'rating_count' 列是否存在并计算 'members'
        if 'rating_count' in movie_df.columns and 'rating' in movie_df.columns:
            movie_df['members'] = movie_df.apply(
                lambda row: int(row['rating_count'] / row['rating']) if row['rating'] > 0 else 0, axis=1
            )
        # 如果关键列不存在，则创建一个默认的 'members' 列以避免崩溃
        elif 'members' not in movie_df.columns:
            movie_df['members'] = 0

        engine = get_db_engine()
        all_ratings_df = pd.DataFrame()
        user_ratings_df = pd.DataFrame()
        if engine:
            # 加载所有用户的评分用于CF算法
            all_ratings_df = pd.read_sql('SELECT user_id, movie_id, rating FROM user_ratings', engine)
            print(f"从数据库加载了 {len(all_ratings_df)} 条评分记录。")
            
            # 单独获取当前用户的评分记录用于过滤
            user_ratings_df = pd.read_sql(f'SELECT movie_id FROM user_ratings WHERE user_id = {user_id}', engine)

    except FileNotFoundError:
        return pd.DataFrame(), "错误：找不到电影数据文件(movies_enriched.csv)。"
    except Exception as e:
        return pd.DataFrame(), f"加载数据时出错: {e}"

    message = ""
    recs_df = pd.DataFrame()

    # 获取用户已评分的电影ID列表
    user_rated_movie_ids = user_ratings_df['movie_id'].tolist() if not user_ratings_df.empty else []

    if recommendation_type == 'top_rated_global':
        # 先获取一个更大的列表，以便过滤后仍有足够数量的推荐
        popular_movies = recommend_top_rated_global(movie_df, min_ratings_threshold=1000, num_recommendations=num_recommendations * 5)
        # 过滤掉用户已经评分过的电影
        recs_df = popular_movies[~popular_movies['movie_id'].isin(user_rated_movie_ids)]
        # 取最终数量的推荐
        recs_df = recs_df.head(num_recommendations)
        message = "为您展示大众热门高分电影！"
    
    elif recommendation_type == 'content_based':
        if liked_movies:
            recs_df = recommend_content_based(liked_movies, movie_df, num_recommendations)
            if recs_df.empty:
                 message = f"无法根据您输入的电影 '{', '.join(liked_movies)}' 找到相似推荐。请检查电影名称是否正确，或尝试其他电影。"
            else:
                message = f"为您找到与 '{', '.join(liked_movies)}' 相似的电影！"
        else:
            message = "请输入您喜欢的电影以获得基于内容的推荐。"

    elif recommendation_type == 'user_cf':
        recs_df = recommend_user_based_cf(user_id, movie_df, all_ratings_df, num_recommendations)
        if recs_df.empty:
            message = '您的评分记录太少，无法生成"基于用户"的推荐。系统为您展示全局热门电影，请多多评分！'

    elif recommendation_type == 'item_cf':
        recs_df = recommend_item_based_cf(user_id, movie_df, all_ratings_df, num_recommendations)
        if recs_df.empty:
            message = '您的评分记录太少，无法生成"基于物品"的推荐。系统为您展示全局热门电影，请多多评分！'

    return recs_df, message