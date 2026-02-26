import numpy as np
import matplotlib.pyplot as plt

# 定义 Sigmoid 函数
def sigmoid(z):
  """计算 Sigmoid 函数的值"""
  return 1 / (1 + np.exp(-z))

# 生成 x 轴数据 (-10 到 10 之间取 100 个点)
z = np.linspace(-10, 10, 100)

# 计算对应的 Sigmoid 值 (y 轴数据)
y = sigmoid(z)

# --- 绘图 ---
plt.figure(figsize=(8, 6)) # 设置图形大小
plt.plot(z, y, label='Sigmoid Function', color='blue') # 绘制曲线

# 添加标题和坐标轴标签
plt.title('Sigmoid')
plt.xlabel('z')
plt.ylabel('sigmoid(z)')

# 添加网格线
plt.grid(True)

# 在 y=0.5 处添加一条水平虚线，标示 Sigmoid 函数的中心点
plt.axhline(y=0.5, color='red', linestyle='--', label='y=0.5')
# 在 x=0 处添加一条垂直虚线
plt.axvline(x=0, color='gray', linestyle=':')

# 显示图例
plt.legend()

# 显示图形
plt.show()