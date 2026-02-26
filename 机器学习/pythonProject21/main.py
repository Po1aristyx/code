import numpy as np
import pandas as pd
import matplotlib.pyplot as plt
from scipy import stats

# 第一类：高发展水平国家
class1 = np.array([
    [76, 99, 5374],    # 美国
    [79.5, 99, 5359],  # 日本
    [78, 99, 5372],    # 瑞士
    [72.1, 95.9, 5242], # 阿根廷
    [73.8, 77.7, 5370]  # 阿联酋
])

# 第二类：中等发展水平国家
class2 = np.array([
    [71.2, 93, 4250],   # 保加利亚
    [75.3, 94.9, 3412], # 古巴
    [70, 91.2, 3390],   # 巴拉圭
    [72.8, 99, 2300],   # 格鲁吉亚
    [62.9, 80.6, 3799]  # 南非
])

# 待判样本
unknown = np.array([
    [68.5, 79.3, 1950],  # 中国
    [69.9, 96.9, 2840],  # 罗马尼亚
    [77.6, 93.8, 5233],  # 希腊
    [69.3, 90.3, 5158]   # 哥伦比亚
])

# 国家名称
country_names = ["中国", "罗马尼亚", "希腊", "哥伦比亚"]

# 计算两类样本的均值向量
mean1 = np.mean(class1, axis=0)
mean2 = np.mean(class2, axis=0)

print("第一类(高发展水平国家)均值向量:", mean1)
print("第二类(中等发展水平国家)均值向量:", mean2)

# 计算类内散度矩阵Sw
n1, n2 = class1.shape[0], class2.shape[0]
p = class1.shape[1]  # 特征数量

# 计算每个类别的协方差矩阵
S1 = np.zeros((p, p))
for i in range(n1):
    diff = (class1[i] - mean1).reshape(p, 1)
    S1 += np.dot(diff, diff.T)

S2 = np.zeros((p, p))
for i in range(n2):
    diff = (class2[i] - mean2).reshape(p, 1)
    S2 += np.dot(diff, diff.T)

# 类内散度矩阵
Sw = S1 + S2
print("\n协方差矩阵S1:")
print(S1)
print("\n协方差矩阵S2:")
print(S2)

# 计算Fisher判别向量w
mean_diff = (mean1 - mean2).reshape(p, 1)
try:
    Sw_inv = np.linalg.inv(Sw)
    w = np.dot(Sw_inv, mean_diff)
    # 归一化w
    w = w / np.linalg.norm(w)
    print("\nFisher判别向量w:")
    print(w.flatten())
except np.linalg.LinAlgError:
    print("警告：协方差矩阵不可逆，使用伪逆")
    Sw_inv = np.linalg.pinv(Sw)
    w = np.dot(Sw_inv, mean_diff)
    # 归一化w
    w = w / np.linalg.norm(w)
    print("\nFisher判别向量w:")
    print(w.flatten())

# 计算投影后的均值
m1 = np.dot(w.T, mean1.reshape(p, 1))[0, 0]
m2 = np.dot(w.T, mean2.reshape(p, 1))[0, 0]

# 计算临界值
c = (m1 + m2) / 2
print("\n投影后第一类均值:", m1)
print("投影后第二类均值:", m2)
print("判别临界值:", c)

# 对已知样本进行判别
y1 = np.zeros(n1)
for i in range(n1):
    y1[i] = np.dot(w.T, class1[i].reshape(p, 1))[0, 0]

y2 = np.zeros(n2)
for i in range(n2):
    y2[i] = np.dot(w.T, class2[i].reshape(p, 1))[0, 0]

print("\n第一类样本的判别函数值:")
print(y1)
print("第二类样本的判别函数值:")
print(y2)

# 检查分类准确性
class1_correct = np.sum(y1 > c)
class2_correct = np.sum(y2 < c)
accuracy = (class1_correct + class2_correct) / (n1 + n2)

print(f"\n训练集分类准确率: {accuracy*100:.2f}%")
print(f"第一类正确分类: {class1_correct}/{n1}")
print(f"第二类正确分类: {class2_correct}/{n2}")

# 对待测样本进行判别
y_unknown = np.zeros(len(unknown))
for i in range(len(unknown)):
    y_unknown[i] = np.dot(w.T, unknown[i].reshape(p, 1))[0, 0]

print("\n待测样本的判别函数值:")
for i in range(len(unknown)):
    print(f"{country_names[i]}: {y_unknown[i]:.4f}")

# 判断类别
predictions = []
for i in range(len(unknown)):
    if y_unknown[i] > c:
        predictions.append("第一类(高发展水平国家)")
    else:
        predictions.append("第二类(中等发展水平国家)")

print("\n待测样本的分类结果:")
for i in range(len(unknown)):
    print(f"{country_names[i]}: {predictions[i]}")

# 可视化结果
plt.figure(figsize=(10, 6))

# 绘制已知样本的投影
plt.scatter(np.ones(n1), y1, c='red', marker='o', label='高发展水平国家')
plt.scatter(np.ones(n2), y2, c='blue', marker='s', label='中等发展水平国家')

# 绘制待测样本的投影
for i in range(len(unknown)):
    if y_unknown[i] > c:
        color = 'red'
    else:
        color = 'blue'
    plt.scatter(1.1, y_unknown[i], c=color, marker='^')
    plt.text(1.12, y_unknown[i], country_names[i])

# 绘制判别临界线
plt.axhline(y=c, color='green', linestyle='--', label='判别临界值')

plt.xlim(0.9, 1.2)
plt.title('Fisher线性判别分析结果')
plt.legend()
plt.grid(True)
plt.show()