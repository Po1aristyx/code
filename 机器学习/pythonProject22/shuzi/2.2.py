import numpy as np
import matplotlib.pyplot as plt

# 设置 matplotlib 支持中文显示
plt.rcParams['font.sans-serif'] = ['SimHei']  # 指定默认字体为 SimHei
plt.rcParams['axes.unicode_minus'] = False  # 解决保存图像时负号'-'显示为方块的问题


def draw():
    fs = 1000  # 采样率 (Hz)
    T = 1.0  # 信号持续时间 (秒)
    t = np.linspace(0, T, int(fs * T), endpoint=False)  # 时间向量

    # 原始信号
    f1 = 50  # Hz
    f2 = 120  # Hz
    signal_clean = np.sin(2 * np.pi * f1 * t) + np.sin(2 * np.pi * f2 * t)

    # 添加均值为0的随机噪声
    # np.random.seed(0) # 为了结果可复现
    noise_power = 0.5  # 调整噪声强度
    # 噪声标准差设为信号标准差的一定比例，使得噪声与信号幅度相关
    noise = noise_power * np.random.normal(loc=0, scale=np.std(signal_clean), size=t.shape)
    signal_noisy = signal_clean + noise

    # 进行傅立叶变换
    N = len(signal_noisy)
    Y_fft = np.fft.fft(signal_noisy)

    # 计算频率轴
    freq_axis = np.fft.fftfreq(N, d=1 / fs)

    # 只取正频率部分进行绘图
    positive_freq_indices = np.where(freq_axis >= 0)
    freq_positive = freq_axis[positive_freq_indices]
    # 幅度谱需要取绝对值并进行归一化 (除以N，有时除以N/2得到单边谱的实际幅度)
    # 对于功率谱，通常是 |X(k)|^2 / N^2
    # 这里我们显示幅度谱，所以是 |X(k)| / N
    Y_fft_magnitude_positive = np.abs(Y_fft[positive_freq_indices]) / N

    # 绘图
    plt.figure(figsize=(12, 6))

    plt.subplot(2, 1, 1)
    plt.plot(t[:200], signal_noisy[:200])  # 只显示前200个点以看清波形
    plt.title('含噪声信号 (时域)')
    plt.xlabel('时间 (s)')
    plt.ylabel('幅度')
    plt.grid(True)

    plt.subplot(2, 1, 2)
    plt.plot(freq_positive, Y_fft_magnitude_positive)
    plt.title('含噪声信号的傅立叶变换 (频域)')
    plt.xlabel('频率 (Hz)')
    plt.ylabel('归一化幅度')
    plt.xlim(0, fs / 2)  # 通常只显示到奈奎斯特频率
    plt.xticks(np.arange(0, fs / 2 + 1, 20))
    plt.grid(True)

    plt.tight_layout()
    plt.suptitle("题二：噪声污染信号的傅立叶变换", fontsize=16)
    plt.subplots_adjust(top=0.92)
    plt.show()
if __name__ == '__main__':
    draw()