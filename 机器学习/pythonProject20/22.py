import pandas as pd
import numpy as np
from sklearn.model_selection import train_test_split
from sklearn.neighbors import KNeighborsRegressor
from sklearn.linear_model import LinearRegression
from sklearn.metrics import mean_squared_error, r2_score
import matplotlib.pyplot as plt

# 1. 加载数据
df = pd.read_excel("电视剧播放数据.xlsx")

# 2. 数据预处理
# 查看数据基本信息
print(df.head())
print(df.info())
print(df.describe())

# 检查缺失值
print("缺失值统计：")
print(df.isnull().sum())

# 3. 特征选择 - 使用点赞数和差评数作为特征
X = df[['点赞', '差评']]
y = df['得分']

# 4. 划分训练集和测试集
X_train, X_test, y_train, y_test = train_test_split(X, y, test_size=0.2, random_state=42)

# 5. 构建模型
# K近邻回归模型 (K=20)
knn_model = KNeighborsRegressor(n_neighbors=20)
knn_model.fit(X_train, y_train)

# 线性回归模型
lr_model = LinearRegression()
lr_model.fit(X_train, y_train)

# 6. 模型预测
y_pred_knn = knn_model.predict(X_test)
y_pred_lr = lr_model.predict(X_test)

# 7. 模型评估
# KNN模型评估
knn_mse = mean_squared_error(y_test, y_pred_knn)
knn_rmse = np.sqrt(knn_mse)
knn_r2 = r2_score(y_test, y_pred_knn)

# 线性回归模型评估
lr_mse = mean_squared_error(y_test, y_pred_lr)
lr_rmse = np.sqrt(lr_mse)
lr_r2 = r2_score(y_test, y_pred_lr)

# 8. 输出结果
print("\n模型评估结果对比：")
print(f"{'模型':<15}{'MSE':<15}{'RMSE':<15}{'R²':<15}")
print(f"{'-'*60}")
print(f"{'KNN (K=20)':<15}{knn_mse:<15.4f}{knn_rmse:<15.4f}{knn_r2:<15.4f}")
print(f"{'线性回归':<15}{lr_mse:<15.4f}{lr_rmse:<15.4f}{lr_r2:<15.4f}")

# 9. 可视化结果
plt.figure(figsize=(12, 5))

# 实际值与KNN预测值对比
plt.subplot(1, 2, 1)
plt.scatter(y_test, y_pred_knn)
plt.plot([y_test.min(), y_test.max()], [y_test.min(), y_test.max()], 'k--')
plt.xlabel('实际评分')
plt.ylabel('KNN预测评分')
plt.title('KNN回归模型 (K=20)')

# 实际值与线性回归预测值对比
plt.subplot(1, 2, 2)
plt.scatter(y_test, y_pred_lr)
plt.plot([y_test.min(), y_test.max()], [y_test.min(), y_test.max()], 'k--')
plt.xlabel('实际评分')
plt.ylabel('线性回归预测评分')
plt.title('线性回归模型')

plt.tight_layout()
plt.savefig('模型预测对比.png')
plt.show()

# 10. 线性回归系数分析
print("\n线性回归模型系数：")
print(f"截距: {lr_model.intercept_:.4f}")
print(f"点赞系数: {lr_model.coef_[0]:.4f}")
print(f"差评系数: {lr_model.coef_[1]:.4f}")
print(f"线性回归方程: 得分 = {lr_model.intercept_:.4f} + {lr_model.coef_[0]:.4f} × 点赞 + {lr_model.coef_[1]:.4f} × 差评")