import pandas as pd

def load_data(anime_path, rating_path):
    """加载动漫和评分数据"""
    anime_data = pd.read_csv('D:\\PYTHON-project\\mssb-course_design\dataes\\anime.csv')
    rating_data = pd.read_csv('D:\\PYTHON-project\\mssb-course_design\dataes\\rating.csv')
    return anime_data, rating_data

def preprocess_data(anime_data, rating_data):
    """预处理数据，处理缺失值和重复数据"""
    # 处理缺失值
    anime_data.dropna(inplace=True)
    rating_data.dropna(inplace=True)

    # 处理重复数据
    anime_data.drop_duplicates(inplace=True)
    rating_data.drop_duplicates(inplace=True)

    # 保留每个用户对每个物品的最高评分
    rating_data = rating_data.groupby(['user_id', 'anime_id']).rating.max().reset_index()

    return anime_data, rating_data
