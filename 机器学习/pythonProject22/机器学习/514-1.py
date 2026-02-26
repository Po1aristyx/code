import numpy as np
from sklearn.datasets import make_classification
from sklearn.model_selection import train_test_split
from sklearn.tree import DecisionTreeClassifier
from sklearn.ensemble import RandomForestClassifier
from sklearn.metrics import accuracy_score

print("实验二：随机森林与决策树对比")
print("=" * 40)

# 1. 随机生成数据
# 生成一个包含1000个样本，20个特征的二分类问题数据集
# n_informative: 有效特征的数量
# n_redundant: 冗余特征的数量 (由有效特征线性组合而成)
# n_clusters_per_class: 每个类别的簇数，影响类别分离的难度
# random_state: 保证每次生成的数据相同，便于结果复现
X, y = make_classification(
    n_samples=1000,
    n_features=20,
    n_informative=15,
    n_redundant=3,
    n_clusters_per_class=1, # 简化问题，每个类别一个簇
    random_state=42
)
print(f"随机数据已生成：{X.shape[0]} 个样本，{X.shape[1]} 个特征。")

# 2. 划分训练集和测试集
# test_size=0.3 表示测试集占30%
X_train, X_test, y_train, y_test = train_test_split(
    X, y,
    test_size=0.3,
    random_state=42
)
print(f"数据已划分为训练集 ({X_train.shape[0]} 样本) 和测试集 ({X_test.shape[0]} 样本)。")
print("-" * 40)

# 3. 采用决策树算法
print("开始训练决策树模型...")
dt_classifier = DecisionTreeClassifier(random_state=42)
dt_classifier.fit(X_train, y_train)
print("决策树模型训练完成。")

# 在测试集上进行预测
y_pred_dt = dt_classifier.predict(X_test)

# 计算决策树的准确率
accuracy_dt = accuracy_score(y_test, y_pred_dt)
print(f"决策树算法在测试集上的准确率: {accuracy_dt:.4f}")
print("-" * 40)

# 4. 采用随机森林算法
print("开始训练随机森林模型...")
# n_estimators: 森林中决策树的数量。通常数量越多，模型越稳定，但训练时间也越长。
# n_jobs=-1: 使用所有可用的CPU核心进行并行计算，加快训练速度。
rf_classifier = RandomForestClassifier(n_estimators=100, random_state=42, n_jobs=-1)
rf_classifier.fit(X_train, y_train)
print("随机森林模型训练完成。")

# 在测试集上进行预测
y_pred_rf = rf_classifier.predict(X_test)

# 计算随机森林的准确率
accuracy_rf = accuracy_score(y_test, y_pred_rf)
print(f"随机森林算法在测试集上的准确率: {accuracy_rf:.4f}")
print("-" * 40)

# 5. 对比两种算法的准确率
print("\n实验二结论：")
print(f"决策树算法准确率: {accuracy_dt:.4f}")
print(f"随机森林算法准确率: {accuracy_rf:.4f}")

if accuracy_rf > accuracy_dt:
    print("结论：随机森林算法的准确率更高。")
elif accuracy_dt > accuracy_rf:
    print("结论：决策树算法的准确率更高。")
else:
    print("结论：两种算法的准确率相同。")

print("\n" + "=" * 40 + "\n")