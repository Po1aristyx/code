import pandas as pd
from sklearn.model_selection import train_test_split
from sklearn.tree import DecisionTreeClassifier, export_graphviz
from sklearn.preprocessing import LabelEncoder
from sklearn.metrics import accuracy_score
import graphviz
import matplotlib.pyplot as plt # 虽然这个题目不直接用，但通常会一起导入
import numpy as np # 同上

# 1. 导入和准备数据
# (3) 使用 Pandas 库读取数据文件

data = pd.read_csv('play.txt', sep='\s+')



print("原始数据 (前5行):")
print(data.head())
print(f"\n数据总行数: {len(data)}")

# 数据清洗与转换：将分类特征转换为数值特征
label_encoders = {}
data_encoded = pd.DataFrame()

# 确保所有列都存在
expected_cols = ['温度', '风力', '天气', 'PM2.5', '是否运动']
if not all(col in data.columns for col in expected_cols):
    print(f"错误：数据文件列名与预期不符。期望列名: {expected_cols}, 实际列名: {data.columns.tolist()}")
    # 可以选择在此处退出或抛出异常
    # exit()


for column in data.columns:
    if data[column].dtype == 'object': # 只对对象类型（通常是字符串）进行编码
        le = LabelEncoder()
        data_encoded[column] = le.fit_transform(data[column])
        label_encoders[column] = le # 保存编码器以便后续解码或理解
    else:
        data_encoded[column] = data[column] # 如果已经是数值型则直接使用

print("\n编码后的数据 (前5行):")
print(data_encoded.head())

# 分离特征 (X) 和目标变量 (y)
# 目标变量是 '是否运动'
target_column = '是否运动'
X = data_encoded.drop(target_column, axis=1)
y = data_encoded[target_column]

feature_names = X.columns.tolist()
# 获取原始类名用于可视化，确保 target_column 在 label_encoders 中
if target_column in label_encoders:
    class_names = label_encoders[target_column].classes_.tolist()
else: # 如果目标列已经是数值0/1，手动定义
    class_names = [str(cls) for cls in sorted(y.unique())]
    print(f"警告：目标列 '{target_column}' 未使用LabelEncoder，将使用其唯一值作为类名: {class_names}")


# (通常我们会划分训练集和测试集，但对于这个小数据集，我们也可以用全部数据来训练以观察树的完整结构)
# X_train, X_test, y_train, y_test = train_test_split(X, y, test_size=0.3, random_state=42, stratify=y)
# 这里为了演示参数影响和可视化整个数据集的树，我们使用全部数据进行训练
X_train, y_train = X, y


# (2) 导入 Scikit-learn 提供的决策树模块 DecisionTreeClassifier() (已在顶部导入)
# (4) 构造决策树模型 (基础模型)
print("\n--- 构建基础决策树模型 ---")
base_clf = DecisionTreeClassifier(random_state=42)
base_clf.fit(X_train, y_train)

# (5) 模型的评估和预测 (在训练集上)
y_pred_train = base_clf.predict(X_train)
train_accuracy = accuracy_score(y_train, y_pred_train)
print(f"基础模型在训练集上的准确率: {train_accuracy:.4f}")

# (7) 决策树的可视化、用 graphviz 将决策树呈现出来。
def visualize_tree(clf, feature_names_list, class_names_list, filename_prefix="tree"):
    # 注意：Graphviz 对中文路径/文件名的支持可能依赖系统配置
    # 尝试使用 feature_names_list 和 class_names_list 直接传入
    dot_data = export_graphviz(clf, out_file=None,
                               feature_names=feature_names_list,
                               class_names=class_names_list,
                               filled=True, rounded=True,
                               special_characters=True,
                               fontname="Microsoft YaHei") # 尝试指定中文字体, 'SimHei' 或 'Microsoft YaHei'
                                                          # 如果Graphviz找不到字体，可能会报错或乱码
    try:
        graph = graphviz.Source(dot_data, encoding='utf-8') # 指定编码
        graph.render(filename_prefix, view=False, format='png') # view=False避免自动打开图片
        print(f"决策树已保存为 {filename_prefix}.png (如果Graphviz命令行工具可用)")
        return graph
    except graphviz.backend.execute.CalledProcessError as e:
        print(f"Graphviz 渲染失败: {e}")
        print("请确保Graphviz已安装并配置在系统PATH中，且支持所选字体（如 Microsoft YaHei）。")
        print("如果仍然存在问题，可以尝试将 feature_names 和 class_names 替换为英文。")
        return None
    except Exception as e_gen:
        print(f"生成决策树图像时发生其他错误: {e_gen}")
        return None


print("\n可视化基础决策树...")
# 使用原始中文特征名和类名
graph_base = visualize_tree(base_clf, feature_names, class_names, "play_decision_tree_base_new_data")

print("\n--- 参数对模型影响总结 ---")
print("""
- criterion: ('gini'/'entropy') 决定不纯度度量方式，影响分裂点的选择。
- splitter: ('best'/'random') 'best'在所有特征中找最优分裂，'random'在特征子集中找。
- max_depth: (int) 限制树的最大深度，防止过拟合。
- min_samples_split: (int or float) 节点分裂所需的最小样本数。值越大，模型越简单。
- max_features: (int, float, 'sqrt', 'log2') 每次分裂时考虑的特征数量。
""")