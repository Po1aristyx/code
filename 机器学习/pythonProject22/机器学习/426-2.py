import pandas as pd
from sklearn.tree import DecisionTreeClassifier, plot_tree
from sklearn.preprocessing import LabelEncoder
import matplotlib.pyplot as plt
import matplotlib

# 对于Windows:
plt.rcParams['font.sans-serif'] = ['SimHei']


# 1. 准备数据
data = {
    '厨师': ['Sita', 'Sita', 'Asha', 'Asha', 'Usha', 'Usha', 'Asha', 'Asha', 'Usha', 'Usha', 'Sita', 'Sita'],
    '状态': ['不好', '好', '不好', '好', '不好', '不好', '不好', '好', '好', '好', '好', '不好'],
    '菜系': ['印度', '欧式', '印度', '印度', '印度', '欧式', '欧式', '欧式', '印度', '欧式', '印度', '欧式'],
    '味道': ['好', '好', '不好', '好', '好', '不好', '不好', '好', '好', '不好', '好', '好']
}
df = pd.DataFrame(data)


# 2. 数据预处理
le_dict = {} # 用于存储每个特征的LabelEncoder，方便后续理解编码
features = ['厨师', '状态', '菜系']
target = '味道'

# 对特征进行编码
df_encoded = pd.DataFrame()
for col in features:
    le = LabelEncoder()
    df_encoded[col] = le.fit_transform(df[col])
    le_dict[col] = le
    print(f"\n{col} 编码映射:")
    print(dict(zip(le.classes_, le.transform(le.classes_))))


# 对目标变量进行编码
le_target = LabelEncoder()
df_encoded[target] = le_target.fit_transform(df[target])
le_dict[target] = le_target
print(f"\n{target} 编码映射:")
print(dict(zip(le_target.classes_, le_target.transform(le_target.classes_))))

# print("\n编码后的数据:")
# print(df_encoded)

X = df_encoded[features]
y = df_encoded[target]

# 3. 训练模型

dt_classifier = DecisionTreeClassifier(criterion='entropy', random_state=42)
dt_classifier.fit(X, y)

# 4. 可视化
plt.figure(figsize=(15, 10)) # 设置图形大小
plot_tree(dt_classifier,
          feature_names=features, # 使用原始特征名
          class_names=le_target.classes_, # 使用原始类别名
          filled=True, # 填充颜色以表示类别
          rounded=True, # 节点使用圆角框
          fontsize=10) # 设置字体大小

plt.title("决策树可视化 (味道预测)", fontsize=16)
plt.show() # 显示图形 (在脚本运行时会弹出窗口)
