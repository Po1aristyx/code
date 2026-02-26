import numpy as np
from sklearn.datasets import load_breast_cancer
from sklearn.model_selection import cross_val_score, KFold
from sklearn.tree import DecisionTreeClassifier
from sklearn.ensemble import AdaBoostClassifier # 使用 AdaBoost 作为提升算法的代表

print("实验三：提升法在乳腺癌数据集上的应用")
print("=" * 40)

# 1. 加载乳腺癌数据集
cancer = load_breast_cancer()
X, y = cancer.data, cancer.target
print(f"乳腺癌数据集已加载：{X.shape[0]} 个样本，{X.shape[1]} 个特征。")

# 2. 定义K-折交叉验证
# n_splits: 折数。常用的有5或10。
# shuffle=True: 在划分前打乱数据，推荐使用。
# random_state: 保证每次划分相同，便于结果复现。
n_splits = 5
kf = KFold(n_splits=n_splits, shuffle=True, random_state=42)
print(f"将使用 {n_splits}-折交叉验证。")
print("-" * 40)

# 3. 采用决策树算法 (CART树)
print("开始评估决策树 (CART) 模型...")
# DecisionTreeClassifier 默认使用 CART 算法 (criterion='gini' for classification)
cart_tree_classifier = DecisionTreeClassifier(random_state=42)

# 使用 cross_val_score 进行交叉验证
# scoring='accuracy': 评估指标为准确率
scores_cart = cross_val_score(cart_tree_classifier, X, y, cv=kf, scoring='accuracy')
mean_accuracy_cart = scores_cart.mean()
std_accuracy_cart = scores_cart.std()

print(f"决策树 (CART) 的 {n_splits}-折交叉验证准确率列表: {scores_cart}")
print(f"决策树 (CART) 的平均准确率: {mean_accuracy_cart:.4f} (标准差: {std_accuracy_cart:.4f})")
print("-" * 40)

# 4. 采用提升算法分类器 (AdaBoost)
print("开始评估 AdaBoost 提升算法模型...")
# AdaBoostClassifier
# estimator: 基学习器。如果不指定，默认为 DecisionTreeClassifier(max_depth=1)，即决策树桩。
#   视频中提到CART树，提升法通常使用弱的CART树作为基学习器。
# n_estimators: 基学习器的数量（即提升的轮数）。
boosting_classifier = AdaBoostClassifier(
    estimator=DecisionTreeClassifier(max_depth=1, random_state=42), # 使用决策树桩作为基学习器
    n_estimators=100, # 提升100轮
    random_state=42
)

# 使用 cross_val_score 进行交叉验证
scores_boosting = cross_val_score(boosting_classifier, X, y, cv=kf, scoring='accuracy')
mean_accuracy_boosting = scores_boosting.mean()
std_accuracy_boosting = scores_boosting.std()

print(f"AdaBoost 提升算法的 {n_splits}-折交叉验证准确率列表: {scores_boosting}")
print(f"AdaBoost 提升算法的平均准确率: {mean_accuracy_boosting:.4f} (标准差: {std_accuracy_boosting:.4f})")
print("-" * 40)

# 5. 检验结果
print(f"决策树 (CART) 平均准确率 (K-折交叉验证): {mean_accuracy_cart:.4f}")
print(f"AdaBoost 提升算法平均准确率 (K-折交叉验证): {mean_accuracy_boosting:.4f}")

if mean_accuracy_boosting > mean_accuracy_cart:
    print("结论：AdaBoost 提升算法在乳腺癌数据集上的平均准确率更高。")
elif mean_accuracy_cart > mean_accuracy_boosting:
    print("结论：决策树 (CART) 算法在乳腺癌数据集上的平均准确率更高。")
else:
    print("结论：两种算法在乳腺癌数据集上的平均准确率表现相似。")

print("\n" + "=" * 40)