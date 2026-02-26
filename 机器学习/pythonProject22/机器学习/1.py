import numpy as np
from sklearn.datasets import load_breast_cancer
from sklearn.model_selection import train_test_split
from sklearn.linear_model import LogisticRegression
from sklearn.metrics import accuracy_score, precision_score, recall_score
import pandas as pd # 用于更好地展示系数

# --- 1. 加载数据集 ---
cancer = load_breast_cancer()
X = cancer.data # 特征数据
y = cancer.target # 目标变量 (0: 恶性, 1: 良性)
feature_names = cancer.feature_names # 特征名称

print(f"数据集样本数: {X.shape[0]}")
print(f"特征数量: {X.shape[1]}")
print(f"目标类别: {cancer.target_names}") # ['malignant' 'benign']
print(f"类别分布: 恶性(0): {np.sum(y==0)}, 良性(1): {np.sum(y==1)}")
print("-" * 30)

# --- 2. 分割数据集 ---
X_train, X_test, y_train, y_test = train_test_split(
    X, y, test_size=0.3, random_state=42, stratify=y # stratify 保证训练集和测试集类别比例相似
)

print(f"训练集大小: {X_train.shape[0]}")
print(f"测试集大小: {X_test.shape[0]}")
print("-" * 30)

# --- 3. 创建和训练逻辑回归模型 ---
model = LogisticRegression(solver='liblinear', max_iter=5000, random_state=42)

# 使用训练集数据拟合模型
model.fit(X_train, y_train)

print("逻辑回归模型训练完成。")
print("-" * 30)

# --- 4. 在测试集上进行预测 ---
y_pred = model.predict(X_test)

# --- 5. 评估模型性能 ---
accuracy = accuracy_score(y_test, y_pred)
# precision_score 和 recall_score 默认计算正类（标签为1，即良性）的指标
# 如果需要计算负类（恶性）的指标，可以设置 pos_label=0
precision = precision_score(y_test, y_pred, pos_label=1)
recall = recall_score(y_test, y_pred, pos_label=1)

print("模型性能评估:")
print(f"  准确率 (Accuracy): {accuracy:.4f}")
print(f"  精确率 (Precision for Benign): {precision:.4f}")
print(f"  召回率 (Recall for Benign): {recall:.4f}")
print("-" * 30)

# --- 6. 查看回归系数 ---
# model.coef_ 是一个二维数组 (1, n_features)，我们取出第一行
coefficients = model.coef_[0]
intercept = model.intercept_[0]

print("模型系数:")
print(f"  截距 (Intercept): {intercept:.4f}")

# 将系数与特征名称对应起来，方便查看
coef_df = pd.DataFrame({'Feature': feature_names, 'Coefficient': coefficients})
# 按系数绝对值降序排列，看看哪些特征影响最大
coef_df['Abs_Coefficient'] = np.abs(coef_df['Coefficient'])
coef_df = coef_df.sort_values(by='Abs_Coefficient', ascending=False).drop('Abs_Coefficient', axis=1)

print("-" * 30)
print("任务完成。")

