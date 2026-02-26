import numpy as np
import matplotlib.pyplot as plt
from sklearn.svm import SVC
from sklearn.model_selection import train_test_split
from sklearn import metrics
import pandas as pd

# 定义数据集
X = np.array([[-2.23, -1.], [-2.06, -1.58], [0.22, -2.12], [-1.18, -2.82], [-1.44, -1.55],
              [-0.5, -1.76], [-0.92, -1.47], [-0.71, -4.54], [-1.52, -2.22], [-2.45, -2.62],
              [-1.31, -2.07], [-3.81, -1.72], [-2.93, -1.88], [-1.19, -2.51], [-0.76, -0.79],
              [-4., -1.89], [-3.08, -1.99], [-1.96, -1.37], [-0.78, -2.87], [-2.87, -2.07],
              [1.62, 3.05], [3.15, 2.29], [1.58, 2.45], [2.93, 3.04], [1.13, 2.5],
              [0.73, 1.06], [1.22, 2.91], [2.18, 1.41], [2.12, 2.2], [1.04, 1.9],
              [-0.64, 3.57], [1.6, 3.34], [0.52, 0.73], [4.25, 1.86], [0.37, 2.02],
              [1.95, 1.74], [3.56, 0.85], [3.19, 1.37], [2.93, 3.32], [1.82, 1.68]])

y = np.array([0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
              1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1])

# 创建SVM分类器，使用线性核函数
svm_classifier = SVC(kernel='linear', C=1.0)

# 训练SVM模型
svm_classifier.fit(X, y)

# 获取超平面参数
w = svm_classifier.coef_[0]  # 权重向量
b = svm_classifier.intercept_[0]  # 偏置项
print(f"超平面参数 - 权重向量w: {w}, 偏置项b: {b}")

# 获取支持向量
support_vectors = svm_classifier.support_vectors_
print(f"支持向量数量: {len(support_vectors)}")

# 创建网格以绘制决策边界
x_min, x_max = X[:, 0].min() - 1, X[:, 0].max() + 1
y_min, y_max = X[:, 1].min() - 1, X[:, 1].max() + 1
xx, yy = np.meshgrid(np.arange(x_min, x_max, 0.02),
                     np.arange(y_min, y_max, 0.02))

# 预测网格中每个点的类别
Z = svm_classifier.predict(np.c_[xx.ravel(), yy.ravel()])
Z = Z.reshape(xx.shape)

# 绘制决策边界和数据点
plt.figure(figsize=(10, 8))
plt.contourf(xx, yy, Z, alpha=0.3, cmap=plt.cm.Paired)

# 绘制超平面
plt.contour(xx, yy, Z, colors='k', levels=[-1, 0, 1], alpha=0.5, linestyles=['--', '-', '--'])

# 绘制支持向量
plt.scatter(support_vectors[:, 0], support_vectors[:, 1], s=100,
            linewidth=1, facecolors='none', edgecolors='k', label='Support Vectors')

# 绘制数据点
plt.scatter(X[y == 0][:, 0], X[y == 0][:, 1], c='b', marker='o', label='Class 0')
plt.scatter(X[y == 1][:, 0], X[y == 1][:, 1], c='r', marker='^', label='Class 1')

# 绘制超平面方程
# 计算超平面上的两个点
x_points = np.array([x_min, x_max])
y_points = (-b - w[0] * x_points) / w[1]

# 计算间隔边界上的点
margin_plus = (-b + 1 - w[0] * x_points) / w[1]  # 正间隔边界
margin_minus = (-b - 1 - w[0] * x_points) / w[1]  # 负间隔边界

# 绘制超平面和间隔边界
plt.plot(x_points, y_points, 'k-')
plt.plot(x_points, margin_plus, 'k--')
plt.plot(x_points, margin_minus, 'k--')

plt.xlabel('Feature 1')
plt.ylabel('Feature 2')
plt.title('SVM Classification with Decision Boundary and Support Vectors')
plt.legend()
plt.grid(True)
plt.show()
