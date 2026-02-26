# 导入所需的库
from sklearn.datasets import load_digits
from sklearn.model_selection import train_test_split, GridSearchCV
from sklearn.svm import SVC
from sklearn.metrics import classification_report, confusion_matrix
import matplotlib.pyplot as plt
import numpy as np

# 加载digits数据集
digits = load_digits()
X = digits.data
y = digits.target

# 显示数据集的一些样本
plt.figure(figsize=(14, 4))
for i in range(10):
    plt.subplot(2, 5, i + 1)
    plt.imshow(digits.images[i], cmap='gray')
    plt.title(f'Label: {digits.target[i]}')
    plt.axis('off')
plt.tight_layout()
plt.show()

# 数据集基本信息
print(f"数据集形状: {X.shape}")
print(f"标签数量: {len(np.unique(y))}")
print(f"每个类别的样本数量:")
for i in range(10):
    print(f"数字 {i}: {np.sum(y == i)}")

# 将数据集分为训练集和测试集 (80% 训练, 20% 测试)
X_train, X_test, y_train, y_test = train_test_split(X, y, test_size=0.2, random_state=42)
print(f"训练集大小: {X_train.shape}")
print(f"测试集大小: {X_test.shape}")

# 使用网格搜索找到最优参数
param_grid = {
    'C': [0.1, 1, 10, 100],
    'gamma': [0.001, 0.01, 0.1, 1],
    'kernel': ['rbf']  # 使用高斯核函数
}

# 创建网格搜索对象
grid_search = GridSearchCV(
    SVC(),
    param_grid,
    cv=5,  # 5折交叉验证
    scoring='accuracy',
    verbose=1,
    n_jobs=-1  # 使用所有可用的CPU核心
)

# 在训练集上执行网格搜索
print("开始网格搜索最优参数...")
grid_search.fit(X_train, y_train)

# 输出最优参数
print(f"最优参数: {grid_search.best_params_}")
print(f"交叉验证最佳得分: {grid_search.best_score_:.4f}")

# 使用最优参数创建SVM分类器
best_svm = SVC(
    C=grid_search.best_params_['C'],
    gamma=grid_search.best_params_['gamma'],
    kernel='rbf'
)

# 在整个训练集上训练最终模型
best_svm.fit(X_train, y_train)

# 在测试集上进行预测
y_pred = best_svm.predict(X_test)

# 计算准确率
accuracy = metrics.accuracy_score(y_test, y_pred)
print(f"测试集准确率: {accuracy:.4f}")

# 输出混淆矩阵
conf_matrix = confusion_matrix(y_test, y_pred)
print("混淆矩阵:")
print(conf_matrix)

# 可视化混淆矩阵
plt.figure(figsize=(10, 8))
plt.imshow(conf_matrix, interpolation='nearest', cmap=plt.cm.Blues)
plt.title('Confusion Matrix')
plt.colorbar()
tick_marks = np.arange(10)
plt.xticks(tick_marks, digits.target_names, rotation=45)
plt.yticks(tick_marks, digits.target_names)
plt.xlabel('Predicted Label')
plt.ylabel('True Label')

# 在混淆矩阵中添加数值标签
thresh = conf_matrix.max() / 2
for i in range(conf_matrix.shape[0]):
    for j in range(conf_matrix.shape[1]):
        plt.text(j, i, format(conf_matrix[i, j], 'd'),
                 horizontalalignment="center",
                 color="white" if conf_matrix[i, j] > thresh else "black")
plt.tight_layout()
plt.show()

# 输出分类报告
print("分类报告:")
print(classification_report(y_test, y_pred))

# 可视化一些预测结果
plt.figure(figsize=(15, 6))
for i in range(15):
    plt.subplot(3, 5, i + 1)
    plt.imshow(X_test[i].reshape(8, 8), cmap='gray')
    pred = y_pred[i]
    true = y_test[i]
    plt.title(f'Pred: {pred}, True: {true}')
    plt.axis('off')
plt.tight_layout()
plt.show()