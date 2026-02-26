import pandas as pd
import numpy as np
import matplotlib.pyplot as plt
import seaborn as sns
from sklearn.model_selection import train_test_split
from sklearn.linear_model import LinearRegression
from sklearn.metrics import mean_squared_error, r2_score
from mpl_toolkits.mplot3d import Axes3D
# 加载数据
df = pd.read_csv('housing_price.csv')


# 检查缺失值

# 准备数据
X_single = df[['size']]
y = df['Price']

# 划分训练集和测试集
X_single_train, X_single_test, y_train, y_test = train_test_split(X_single, y, test_size=0.2, random_state=42)

# 建立单因子模型
model_single = LinearRegression()
model_single.fit(X_single_train, y_train)

# 预测
y_pred_single = model_single.predict(X_single_test)

# 评估模型
mse_single = mean_squared_error(y_test, y_pred_single)
r2_single = r2_score(y_test, y_pred_single)

print(f"单因子模型 - 面积")
print(f"截距: {model_single.intercept_:.2f}")
print(f"系数: {model_single.coef_[0]:.2f}")
print(f"均方误差 (MSE): {mse_single:.2f}")
print(f"决定系数 (R²): {r2_single:.4f}")
# 准备数据
X_multi = df[['Avg. Area Income', 'Avg. Area House Age', 'Avg. Area Number of Rooms', 'Area Population', 'size']]

# 划分训练集和测试集
X_multi_train, X_multi_test, y_train, y_test = train_test_split(X_multi, y, test_size=0.2, random_state=42)

# 建立多因子模型
model_multi = LinearRegression()
model_multi.fit(X_multi_train, y_train)

# 预测
y_pred_multi = model_multi.predict(X_multi_test)

# 评估模型
mse_multi = mean_squared_error(y_test, y_pred_multi)
r2_multi = r2_score(y_test, y_pred_multi)

print(f"\n多因子模型")
print(f"截距: {model_multi.intercept_:.2f}")
for i, col in enumerate(X_multi.columns):
    print(f"{col} 系数: {model_multi.coef_[i]:.2f}")
print(f"均方误差 (MSE): {mse_multi:.2f}")
print(f"决定系数 (R²): {r2_multi:.4f}")
# 准备数据
X_multi = df[['Avg. Area Income', 'Avg. Area House Age', 'Avg. Area Number of Rooms', 'Area Population', 'size']]

# 划分训练集和测试集
X_multi_train, X_multi_test, y_train, y_test = train_test_split(X_multi, y, test_size=0.2, random_state=42)

# 建立多因子模型
model_multi = LinearRegression()
model_multi.fit(X_multi_train, y_train)

# 预测
y_pred_multi = model_multi.predict(X_multi_test)

# 评估模型
mse_multi = mean_squared_error(y_test, y_pred_multi)
r2_multi = r2_score(y_test, y_pred_multi)

print(f"\n多因子模型")
print(f"截距: {model_multi.intercept_:.2f}")
for i, col in enumerate(X_multi.columns):
    print(f"{col} 系数: {model_multi.coef_[i]:.2f}")
print(f"均方误差 (MSE): {mse_multi:.2f}")
print(f"决定系数 (R²): {r2_multi:.4f}")
plt.figure(figsize=(10, 6))
plt.scatter(X_single_test, y_test, color='blue', alpha=0.5, label='实际值')
plt.plot(X_single_test, y_pred_single, color='red', linewidth=2, label='预测值')
plt.title('单因子模型：面积与房价的关系')
plt.xlabel('面积 (size)')
plt.ylabel('房价 (Price)')
plt.legend()
plt.grid(True)
plt.tight_layout()
plt.savefig('single_factor_model.png')
plt.show()
# 预测值与实际值对比
plt.figure(figsize=(10, 6))
plt.scatter(y_test, y_pred_multi, alpha=0.5)
plt.plot([y_test.min(), y_test.max()], [y_test.min(), y_test.max()], 'r--')
plt.title('多因子模型：预测值 vs 实际值')
plt.xlabel('实际房价')
plt.ylabel('预测房价')
plt.grid(True)
plt.tight_layout()
plt.show()

# 残差图
residuals = y_test - y_pred_multi
plt.figure(figsize=(10, 6))
plt.scatter(y_pred_multi, residuals, alpha=0.5)
plt.axhline(y=0, color='r', linestyle='-')
plt.title('多因子模型：残差图')
plt.xlabel('预测房价')
plt.ylabel('残差')
plt.grid(True)
plt.tight_layout()
plt.show()

# 特征重要性
importance = pd.DataFrame({
    'Feature': X_multi.columns,
    'Coefficient': model_multi.coef_
})
importance['Abs_Coefficient'] = abs(importance['Coefficient'])
importance = importance.sort_values('Abs_Coefficient', ascending=False)

plt.figure(figsize=(10, 6))
sns.barplot(x='Abs_Coefficient', y='Feature', data=importance)
plt.title('多因子模型：特征重要性')
plt.xlabel('系数绝对值')
plt.ylabel('特征')
plt.tight_layout()
plt.show()
models = ['单因子模型', '多因子模型']
r2_scores = [r2_single, r2_multi]
mse_scores = [mse_single, mse_multi]

plt.figure(figsize=(12, 5))

plt.subplot(1, 2, 1)
plt.bar(models, r2_scores, color=['skyblue', 'orange'])
plt.title('R² 比较')
plt.ylim(0, 1)
for i, v in enumerate(r2_scores):
    plt.text(i, v + 0.01, f'{v:.4f}', ha='center')

plt.subplot(1, 2, 2)
plt.bar(models, mse_scores, color=['skyblue', 'orange'])
plt.title('MSE 比较')
for i, v in enumerate(mse_scores):
    plt.text(i, v + 0.01, f'{v:.2f}', ha='center')

plt.tight_layout()
plt.show()




# 选择两个最重要的特征（假设是面积和收入）
feature1 = 'Avg. Area House Age'
feature2 = 'Avg. Area Number of Rooms'

# 创建3D图
fig = plt.figure(figsize=(12, 10))
ax = fig.add_subplot(111, projection='3d')

# 绘制散点图
scatter = ax.scatter(df[feature1], df[feature2], df['Price'],
                    c=df['Price'], cmap='viridis', alpha=0.6, s=50)

# 创建预测平面的网格点
x_min, x_max = df[feature1].min(), df[feature1].max()
y_min, y_max = df[feature2].min(), df[feature2].max()
x_grid, y_grid = np.meshgrid(np.linspace(x_min, x_max, 20),
                            np.linspace(y_min, y_max, 20))

# 使用多因子模型预测网格点的值
model_subset = LinearRegression()
model_subset.fit(df[[feature1, feature2]], df['Price'])
z_grid = model_subset.predict(np.c_[x_grid.ravel(), y_grid.ravel()]).reshape(x_grid.shape)

# 绘制预测平面
surface = ax.plot_surface(x_grid, y_grid, z_grid, alpha=0.3, cmap='viridis')

# 添加标题和标签
ax.set_title(f'3D可视化: {feature1}, {feature2} 与房价的关系', fontsize=15)
ax.set_xlabel(feature1, fontsize=12)
ax.set_ylabel(feature2, fontsize=12)
ax.set_zlabel('Price', fontsize=12)

# 添加颜色条
fig.colorbar(scatter, ax=ax, shrink=0.5, aspect=5, label='房价')

plt.tight_layout()
plt.show()



X = df.drop('Price', axis=1)  # 所有特征
y = df['Price']  # 目标变量

# 训练模型
model = LinearRegression()
model.fit(X, y)

# 查看所有特征的系数
feature_coefficients = pd.DataFrame({
    '特征': X.columns,
    '系数': model.coef_
})

print("所有特征的系数:")
print(feature_coefficients)

# 按系数绝对值排序并可视化
feature_coefficients = feature_coefficients.reindex(
    feature_coefficients['系数'].abs().sort_values(ascending=False).index
)

plt.figure(figsize=(10, 6))
plt.barh(feature_coefficients['特征'], feature_coefficients['系数'])
plt.xlabel('系数值')
plt.ylabel('特征')
plt.title('所有特征的重要性系数')
plt.grid(True, linestyle='--', alpha=0.7, axis='x')
plt.tight_layout()
plt.show()