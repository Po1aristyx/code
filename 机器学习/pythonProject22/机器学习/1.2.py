import pandas as pd
import numpy as np
import matplotlib.pyplot as plt
import seaborn as sns
from sklearn.model_selection import train_test_split
from sklearn.linear_model import LinearRegression
from sklearn.metrics import mean_squared_error, r2_score, mean_absolute_error
import statsmodels.api as sm

# 1. 加载数据
df = pd.read_excel('北京市空气质量数据.xlsx')

# 2. 数据预处理
# 查看数据基本信息
print("原始数据信息:")
print(df.info())
print("\n原始数据前5行:")
print(df.head())

# 将0值替换为NaN
df.replace(0, np.nan, inplace=True)

# 剔除缺失值
df_clean = df.dropna()

# 筛选PM2.5浓度低于200且SO2浓度低于20的数据
df_filtered = df_clean[(df_clean['PM2.5'] < 200) & (df_clean['SO2'] < 20)]

print(f"\n预处理后的数据量: {len(df_filtered)}")
print("\n预处理后的数据统计:")
print(df_filtered.describe())

# 3. 数据可视化
plt.figure(figsize=(12, 5))

# CO与PM2.5的散点图
plt.subplot(1, 2, 1)
plt.scatter(df_filtered['CO'], df_filtered['PM2.5'], alpha=0.5)
plt.title('CO vs PM2.5')
plt.xlabel('CO')
plt.ylabel('PM2.5')

# SO2与PM2.5的散点图
plt.subplot(1, 2, 2)
plt.scatter(df_filtered['SO2'], df_filtered['PM2.5'], alpha=0.5)
plt.title('SO2 vs PM2.5')
plt.xlabel('SO2')
plt.ylabel('PM2.5')

plt.tight_layout()
plt.savefig('scatter_plots.png')
plt.show()

# 相关性热图
plt.figure(figsize=(8, 6))
correlation = df_filtered[['PM2.5', 'CO', 'SO2']].corr()
sns.heatmap(correlation, annot=True, cmap='coolwarm', fmt='.2f')
plt.title('相关性热图')
plt.savefig('correlation_heatmap.png')
plt.show()

# 4. 一元线性回归模型 (PM2.5 ~ CO)
X_single = df_filtered[['CO']]
y = df_filtered['PM2.5']

# 划分训练集和测试集
X_single_train, X_single_test, y_train, y_test = train_test_split(
    X_single, y, test_size=0.2, random_state=42)

# 建立一元线性回归模型
model_single = LinearRegression()
model_single.fit(X_single_train, y_train)

# 使用statsmodels获取详细统计信息
X_single_sm = sm.add_constant(X_single_train)
model_single_sm = sm.OLS(y_train, X_single_sm).fit()
print("\n一元线性回归模型详细统计:")
print(model_single_sm.summary())

# 预测
y_pred_single = model_single.predict(X_single_test)

# 评估模型
mse_single = mean_squared_error(y_test, y_pred_single)
rmse_single = np.sqrt(mse_single)
mae_single = mean_absolute_error(y_test, y_pred_single)
r2_single = r2_score(y_test, y_pred_single)

print("\n一元线性回归模型 (PM2.5 ~ CO):")
print(f"回归方程: PM2.5 = {model_single.intercept_:.4f} + {model_single.coef_[0]:.4f} * CO")
print(f"均方误差 (MSE): {mse_single:.4f}")
print(f"均方根误差 (RMSE): {rmse_single:.4f}")
print(f"平均绝对误差 (MAE): {mae_single:.4f}")
print(f"决定系数 (R²): {r2_single:.4f}")

# 可视化一元回归模型
plt.figure(figsize=(10, 6))
plt.scatter(X_single_test, y_test, color='blue', alpha=0.5, label='实际值')
plt.plot(X_single_test, y_pred_single, color='red', linewidth=2, label='预测值')
plt.title('一元线性回归: CO vs PM2.5')
plt.xlabel('CO')
plt.ylabel('PM2.5')
plt.legend()
plt.grid(True)
plt.savefig('single_regression_model.png')
plt.show()

# 5. 多元线性回归模型 (PM2.5 ~ CO + SO2)
X_multi = df_filtered[['CO', 'SO2']]

# 划分训练集和测试集
X_multi_train, X_multi_test, y_train, y_test = train_test_split(
    X_multi, y, test_size=0.2, random_state=42)

# 建立多元线性回归模型
model_multi = LinearRegression()
model_multi.fit(X_multi_train, y_train)

# 使用statsmodels获取详细统计信息
X_multi_sm = sm.add_constant(X_multi_train)
model_multi_sm = sm.OLS(y_train, X_multi_sm).fit()
print("\n多元线性回归模型详细统计:")
print(model_multi_sm.summary())

# 预测
y_pred_multi = model_multi.predict(X_multi_test)

# 评估模型
mse_multi = mean_squared_error(y_test, y_pred_multi)
rmse_multi = np.sqrt(mse_multi)
mae_multi = mean_absolute_error(y_test, y_pred_multi)
r2_multi = r2_score(y_test, y_pred_multi)

print("\n多元线性回归模型 (PM2.5 ~ CO + SO2):")
print(f"回归方程: PM2.5 = {model_multi.intercept_:.4f} + {model_multi.coef_[0]:.4f} * CO + {model_multi.coef_[1]:.4f} * SO2")
print(f"均方误差 (MSE): {mse_multi:.4f}")
print(f"均方根误差 (RMSE): {rmse_multi:.4f}")
print(f"平均绝对误差 (MAE): {mae_multi:.4f}")
print(f"决定系数 (R²): {r2_multi:.4f}")

# 可视化多元回归模型预测结果
plt.figure(figsize=(10, 6))
plt.scatter(y_test, y_pred_multi, alpha=0.5)
plt.plot([y_test.min(), y_test.max()], [y_test.min(), y_test.max()], 'r--')
plt.title('多元线性回归: 预测值 vs 实际值')
plt.xlabel('实际PM2.5')
plt.ylabel('预测PM2.5')
plt.grid(True)
plt.savefig('multi_regression_prediction.png')
plt.show()

# 残差分析
residuals = y_test - y_pred_multi
plt.figure(figsize=(12, 5))

# 残差散点图
plt.subplot(1, 2, 1)
plt.scatter(y_pred_multi, residuals, alpha=0.5)
plt.axhline(y=0, color='r', linestyle='-')
plt.title('残差图')
plt.xlabel('预测PM2.5')
plt.ylabel('残差')
plt.grid(True)

# 残差直方图
plt.subplot(1, 2, 2)
plt.hist(residuals, bins=20, alpha=0.7, color='skyblue', edgecolor='black')
plt.title('残差分布')
plt.xlabel('残差')
plt.ylabel('频数')
plt.grid(True)

plt.tight_layout()
plt.savefig('residual_analysis.png')
plt.show()

# 6. 模型比较
print("\n模型比较:")
print(f"一元线性回归 (PM2.5 ~ CO) - R²: {r2_single:.4f}, RMSE: {rmse_single:.4f}")
print(f"多元线性回归 (PM2.5 ~ CO + SO2) - R²: {r2_multi:.4f}, RMSE: {rmse_multi:.4f}")
print(f"R²提升: {(r2_multi - r2_single):.4f}")
print(f"RMSE降低: {(rmse_single - rmse_multi):.4f}")

# 3D可视化多元回归模型
fig = plt.figure(figsize=(12, 8))
ax = fig.add_subplot(111, projection='3d')

# 创建网格
x_surf = np.linspace(X_multi_test['CO'].min(), X_multi_test['CO'].max(), 100)
y_surf = np.linspace(X_multi_test['SO2'].min(), X_multi_test['SO2'].max(), 100)
x_surf, y_surf = np.meshgrid(x_surf, y_surf)

# 计算预测值
z_surf = model_multi.intercept_ + model_multi.coef_[0] * x_surf + model_multi.coef_[1] * y_surf

# 绘制3D表面
ax.plot_surface(x_surf, y_surf, z_surf, alpha=0.5, cmap='viridis')

# 绘制实际数据点
ax.scatter(X_multi_test['CO'], X_multi_test['SO2'], y_test, color='blue', alpha=0.7)

ax.set_xlabel('CO')
ax.set_ylabel('SO2')
ax.set_zlabel('PM2.5')
ax.set_title('多元线性回归模型的3D可视化')

plt.savefig('3d_regression_visualization.png')
plt.show()