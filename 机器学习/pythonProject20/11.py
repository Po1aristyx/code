import pandas as pd
import numpy as np
from sklearn.linear_model import LinearRegression
import matplotlib.pyplot as plt

# 假设df是您的数据框
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