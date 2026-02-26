import numpy as np
import matplotlib.pyplot as plt

# 设置 matplotlib 支持中文显示
plt.rcParams['font.sans-serif'] = ['SimHei']  # 指定默认字体为 SimHei
plt.rcParams['axes.unicode_minus'] = False  # 解决保存图像时负号'-'显示为方块的问题


def draw():

    N = 12
    n = np.arange(N)
    x_n = np.cos(n * np.pi / 6)

    # 计算 DFT
    X_k = np.fft.fft(x_n)

    # 计算幅度和相位
    magnitude = np.abs(X_k)
    phase = np.angle(X_k)

    k = np.arange(N)

    # 绘图
    plt.figure(figsize=(12, 8))

    plt.subplot(2, 1, 1)
    plt.stem(k, magnitude, basefmt=" ")
    plt.title(f'x(n)=cos(n*pi/6) (N={N}) 的DFT幅度谱')
    plt.xlabel('k (频率索引)')
    plt.ylabel('幅度 |X(k)|')
    plt.xticks(k)
    plt.grid(True)

    plt.subplot(2, 1, 2)
    plt.stem(k, phase, basefmt=" ")
    plt.title(f'x(n)=cos(n*pi/6) (N={N}) 的DFT相位谱')
    plt.xlabel('k (频率索引)')
    plt.ylabel('相位 (弧度)')
    plt.xticks(k)
    plt.yticks([-np.pi, -np.pi / 2, 0, np.pi / 2, np.pi], ['$-\pi$', '$-\pi/2$', '0', '$\pi/2$', '$\pi$'])
    plt.grid(True)

    plt.tight_layout()
    plt.suptitle("题一：x(n)=cos(n*pi/6) 的 DFT", fontsize=16)
    plt.subplots_adjust(top=0.92)
    plt.show()

if __name__ == '__main__':
    draw()