import pandas as pd
import redis

# 初始化 Redis 缓存
cache = redis.StrictRedis(host='localhost', port=6379, db=0)

def recommend_based_on_survey(preferences, anime_data, num_recommendations=10):
    """根据用户偏好推荐动漫"""
    recommendations = anime_data[anime_data['genre'].apply(lambda x: any(tag in x.split(', ') for tag in preferences))]
    if recommendations.empty:
        return anime_data.head(num_recommendations)
    return recommendations.sample(num_recommendations)[['anime_id', 'name', 'genre', 'rating', 'members']]


def recommend_user_ratings(rating_data, anime_data, user_id, num_recommendations=10):
    """根据用户评分与整体评分对比推荐动漫"""
    user_ratings = rating_data[rating_data['user_id'] == user_id]
    if user_ratings.empty:
        return pd.DataFrame()

    anime_ids = user_ratings['anime_id'].unique()
    recommendations = anime_data[anime_data['anime_id'].isin(anime_ids)]

    for index, row in recommendations.iterrows():
        anime_id = row['anime_id']
        user_rating = user_ratings[user_ratings['anime_id'] == anime_id]['rating'].mean()
        overall_rating = row['rating']
        recommendations.at[index, 'rating_difference'] = (user_rating - overall_rating) ** 2

    recommendations['popularity'] = recommendations['members']
    recommendations = recommendations.sort_values(by=['rating_difference', 'popularity'], ascending=[False, False])

    return recommendations.head(num_recommendations)[['anime_id', 'name', 'genre', 'rating', 'members']]


def recommend_popular_anime(anime_data, num_recommendations=5):
    """推荐最受欢迎的动漫"""
    popular_anime = anime_data.sort_values(by='members', ascending=False).head(num_recommendations)
    return popular_anime[['anime_id', 'name', 'genre', 'rating', 'members']]


def recommend_hybrid(rating_data, anime_data, preferences, user_id, num_recommendations=10):
    """混合推荐"""
    # 基于用户偏好的推荐 (权重最高)
    survey_recommendations = recommend_based_on_survey(preferences, anime_data, num_recommendations * 3)

    # 基于用户评分与整体评分对比的推荐 (次高权重)
    user_rating_recommendations = recommend_user_ratings(rating_data, anime_data, user_id, num_recommendations * 2)

    # 最受欢迎的动漫推荐 (权重最低)
    popular_recommendations = recommend_popular_anime(anime_data, num_recommendations)

    # 合并结果，并去重
    combined_recommendations = pd.concat(
        [survey_recommendations, user_rating_recommendations, popular_recommendations]).drop_duplicates()

    # 按照权重重新排序：首先是偏好标签，其次是评分差异，最后是受欢迎程度
    combined_recommendations['weight'] = combined_recommendations.apply(
        lambda row: (3 if any(tag in row['genre'].split(', ') for tag in preferences) else 0) +
                    row.get('rating_difference', 0) + row['members'] / combined_recommendations['members'].max(), axis=1
    )
    combined_recommendations = combined_recommendations.sort_values(by='weight', ascending=False)

    return combined_recommendations.head(num_recommendations)[['anime_id', 'name', 'genre', 'rating', 'members']]


def get_recommendations(user_id, rating_data, anime_data, preferences, num_recommendations=5, use_cache=True):
    cache_key = f"user_recommendations_{user_id}"
    if use_cache:
        cached_recommendations = cache.get(cache_key)
        if cached_recommendations:
            cached_recs = pd.read_json(cached_recommendations)
            # 检查缓存的推荐是否与新的推荐相同，如果相同则返回缓存的推荐
            new_recommendations = _get_recommendations_without_cache(user_id, rating_data, anime_data, preferences, num_recommendations)
            if new_recommendations.equals(cached_recs):
                return cached_recs
            else:
                # 更新缓存
                cache.set(cache_key, new_recommendations.to_json(), ex=3600)  # 缓存一小时
                return new_recommendations
    else:
        return _get_recommendations_without_cache(user_id, rating_data, anime_data, preferences, num_recommendations)


def _get_recommendations_without_cache(user_id, rating_data, anime_data, preferences, num_recommendations):
    if rating_data[rating_data['user_id'] == user_id].empty:
        recommendations = recommend_based_on_survey(preferences, anime_data, num_recommendations)
    else:
        recommendations = recommend_hybrid(rating_data, anime_data, preferences, user_id, num_recommendations)
    return recommendations