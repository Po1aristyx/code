import numpy as np
import matplotlib.pyplot as plt
plt.rcParams['font.sans-serif'] = ['SimHei']  # 黑体
plt.rcParams['axes.unicode_minus'] = False     # 解决负号显示问题

def dconv(f1, f2, k1, k2):
    f = np.convolve(f1, f2)
    k = np.arange(k1[0] + k2[0], k1[-1] + k2[-1] + 1)

    plt.figure(figsize=(8, 6))

    plt.subplot(3, 1, 1)
    plt.stem(k1, f1)
    plt.title('f1(k)')
    plt.grid(True)

    plt.subplot(3, 1, 2)
    plt.stem(k2, f2)
    plt.title('f2(k)')
    plt.grid(True)

    plt.subplot(3, 1, 3)
    plt.stem(k, f)
    plt.title('f(k) = f1(k) * f2(k)')
    plt.grid(True)

    plt.tight_layout()
    plt.show()

    return f, k

# 题三的序列
f1 = [1, 2, 1];      k1 = np.arange(-1, 2)    # k = -1, 0, 1
f2 = [1, 1, 1, 1, 1]; k2 = np.arange(-2, 3)    # k = -2, -1, 0, 1, 2
# 使用上面写好的 dconv 函数
dconv(f1, f2, k1, k2)
