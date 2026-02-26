import pandas as pd
import os


def process_movie_data(
        movies_file='movies.csv',
        tags_file='tags.csv',
        ratings_file='ratings.csv',  # 新增ratings文件参数
        output_file='movies_enriched.csv'  # 修改输出文件名以反映更多信息
):
    """
    加载电影、标签和评分数据。
    1. 将标签聚合后添加到电影数据中。
    2. 计算每部电影的平均评分和评分次数并添加到电影数据中。
    3. 保存最终结果。
    """
    print(f"开始加载数据: {movies_file}, {tags_file}, {ratings_file}")
    try:
        movies_df = pd.read_csv(movies_file)
        tags_df = pd.read_csv(tags_file)
        ratings_df = pd.read_csv(ratings_file)  # 加载评分数据
    except FileNotFoundError as e:
        print(f"错误: 文件未找到 - {e}")
        return
    except Exception as e:
        print(f"加载文件时出错: {e}")
        return

    print("电影数据 (前5行):")
    print(movies_df.head())
    print("\n标签数据 (前5行):")
    print(tags_df.head())
    print("\n评分数据 (前5行):")
    print(ratings_df.head())

    # --- 1. 聚合标签 ---
    print("\n开始聚合标签...")
    tags_df['tag'] = tags_df['tag'].astype(str)
    movie_tags_aggregated = tags_df.groupby('movieId')['tag'].apply(
        lambda x: '|'.join(x.dropna().unique())
    ).reset_index()
    movie_tags_aggregated.rename(columns={'tag': 'movie_tags'}, inplace=True)

    print("\n聚合后的标签数据 (前5行):")
    print(movie_tags_aggregated.head())

    # 将标签合并到电影DataFrame中
    movies_enriched_df = pd.merge(movies_df, movie_tags_aggregated, on='movieId', how='left')
    movies_enriched_df['movie_tags'].fillna('', inplace=True)
    print("\n合并标签后的电影数据 (前5行):")
    print(movies_enriched_df.head())

    # --- 2. 计算并合并评分统计数据 ---
    print("\n开始计算评分统计...")
    if 'rating' not in ratings_df.columns or 'movieId' not in ratings_df.columns:
        print("错误: ratings.csv 文件缺少 'movieId' 或 'rating' 列。")
        return

    movie_rating_stats = ratings_df.groupby('movieId').agg(
        average_rating=('rating', 'mean'),
        rating_count=('rating', 'count')  # 或者 num_ratings
    ).reset_index()

    print("\n计算得到的评分统计 (前5行):")
    print(movie_rating_stats.head())

    # 将评分统计合并到电影DataFrame中
    # 确保movieId类型一致以进行合并
    try:
        if movies_enriched_df['movieId'].dtype != movie_rating_stats['movieId'].dtype:
            print(
                f"警告: movieId 类型不匹配。尝试转换 movies_enriched_df['movieId'] 为 {movie_rating_stats['movieId'].dtype} 类型。")
            movies_enriched_df['movieId'] = movies_enriched_df['movieId'].astype(movie_rating_stats['movieId'].dtype)
    except Exception as e:
        print(f"movieId 类型转换失败: {e}. 合并可能失败或产生意外结果。")

    movies_enriched_df = pd.merge(movies_enriched_df, movie_rating_stats, on='movieId', how='left')

    # 对于没有评分的电影，其统计列会是 NaN，可以填充为0或适当的值
    movies_enriched_df['average_rating'].fillna(0, inplace=True)  # 或 np.nan
    movies_enriched_df['rating_count'].fillna(0, inplace=True)  # 或 np.nan

    print("\n合并评分统计后的电影数据 (前5行):")
    print(movies_enriched_df.head())
    print("\n列信息:")
    print(movies_enriched_df.info())

    try:
        # 确保输出目录存在
        output_dir = os.path.dirname(output_file)
        if output_dir and not os.path.exists(output_dir):
            os.makedirs(output_dir)
            print(f"创建目录: {output_dir}")

        print(f"\n开始保存结果到: {output_file}")
        movies_enriched_df.to_csv(output_file, index=False, encoding='utf-8')
        print(f"处理完成！丰富的电影数据已保存到 {output_file}")
    except Exception as e:
        print(f"保存文件时出错: {e}")


if __name__ == '__main__':
    # 定义输入和输出文件路径
    # 假设输入文件在项目根目录
    INPUT_DIR = '.'
    OUTPUT_DIR = 'dataes'

    # 确保输出目录存在
    if not os.path.exists(OUTPUT_DIR):
        os.makedirs(OUTPUT_DIR)
        print(f"创建目录: {OUTPUT_DIR}")

    process_movie_data(
        movies_file=os.path.join(INPUT_DIR, 'movies.csv'),
        tags_file=os.path.join(INPUT_DIR, 'tags.csv'),
        ratings_file=os.path.join(INPUT_DIR, 'dataes/ratings.csv'),
        output_file=os.path.join(OUTPUT_DIR, 'movies_enriched.csv')
    )