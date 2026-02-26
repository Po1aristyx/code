import pandas as pd


def extract_genres(file_path):
    """从动漫数据集提取所有独特的类型"""
    # 读取CSV文件
    data = pd.read_csv(file_path)

    # 提取所有类型
    all_genres = data['genre'].dropna().str.split(',').explode().str.strip().unique()

    return sorted(all_genres)


if __name__ == '__main__':
    file_path = 'D:/PYTHON-project/mssb-course_design/dataes/anime.csv'  # 替换为实际的CSV文件路径
    genres = extract_genres(file_path)
    print("所有的动漫类型:")
    for genre in genres:
        print(genre)
