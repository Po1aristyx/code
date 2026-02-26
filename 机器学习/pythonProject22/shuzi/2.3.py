import numpy as np
import matplotlib.pyplot as plt
import librosa
import librosa.display

# 设置 matplotlib 支持中文显示
plt.rcParams['font.sans-serif'] = ['SimHei']  # 指定默认字体为 SimHei
plt.rcParams['axes.unicode_minus'] = False  # 解决保存图像时负号'-'显示为方块的问题


def draw():

    # 加载示例语音信号
    # 'mtlb' 是 MATLAB 特有的信号，这里我们用一个 librosa 的示例语音代替
    try:
        # 尝试加载一个常见的语音样本
        y, sr = librosa.load(librosa.ex('libri1'), duration=1.5, sr=None)
    except Exception as e_libri:
        print(f"加载 librosa.ex('libri1') 失败: {e_libri}")
        print("尝试加载 librosa.ex('trumpet') 作为备选...")
        try:
            y, sr = librosa.load(librosa.ex('trumpet'), duration=1.5, sr=None)
            print("警告: 使用 'trumpet' 示例音频，非语音信号。")
        except Exception as e_trumpet:
            print(f"加载 librosa.ex('trumpet') 也失败: {e_trumpet}")
            print("将生成一个简单的合成信号代替。")
            sr = 8000  # 假设采样率
            t_synth = np.linspace(0, 1.5, int(1.5 * sr), endpoint=False)
            y = 0.5 * np.sin(2 * np.pi * 220 * t_synth) + 0.3 * np.sin(2 * np.pi * 440 * t_synth)
            y += 0.1 * np.random.randn(len(y))  # 添加少量噪声

    # 1. 原始语音信号
    plt.figure(figsize=(14, 10))  # 调整图形大小以容纳标题和标签

    plt.subplot(2, 2, 1)
    librosa.display.waveshow(y, sr=sr, ax=plt.gca())
    plt.title('1. 原始信号 (时域)')
    plt.xlabel('时间 (s)')
    plt.ylabel('幅度')

    # 2. FFT变换
    Y_fft = np.fft.fft(y)
    N_fft = len(Y_fft)
    freq_axis_fft = np.fft.fftfreq(N_fft, d=1 / sr)

    # 取正频率部分用于绘图
    positive_freq_mask = freq_axis_fft >= 0
    freq_positive = freq_axis_fft[positive_freq_mask]
    Y_fft_magnitude_positive = np.abs(Y_fft[positive_freq_mask])

    plt.subplot(2, 2, 2)
    plt.plot(freq_positive, Y_fft_magnitude_positive)
    plt.title('2. 原始信号FFT变换 (幅度谱)')
    plt.xlabel('频率 (Hz)')
    plt.ylabel('幅度 |X(f)|')
    plt.xlim(0, sr / 2)  # 显示到奈奎斯特频率
    plt.grid(True)

    # 3. 去掉幅值小于阈值的FFT变换值
    Y_fft_filtered = Y_fft.copy()
    threshold = 1.0
    print(
        f"FFT 幅度范围: min={np.min(np.abs(Y_fft)):.2f}, max={np.max(np.abs(Y_fft)):.2f}, mean={np.mean(np.abs(Y_fft)):.2f}")
    print(f"使用的阈值: {threshold}")

    Y_fft_filtered[np.abs(Y_fft_filtered) < threshold] = 0

    Y_fft_filtered_magnitude_positive = np.abs(Y_fft_filtered[positive_freq_mask])

    plt.subplot(2, 2, 3)
    plt.plot(freq_positive, Y_fft_filtered_magnitude_positive)
    plt.title(f'3. 滤波后FFT (幅度 < {threshold:.1f} 置零)')
    plt.xlabel('频率 (Hz)')
    plt.ylabel('幅度 |X_filtered(f)|')
    plt.xlim(0, sr / 2)
    plt.grid(True)

    # 4. 重构语音信号
    y_reconstructed = np.fft.ifft(Y_fft_filtered)
    y_reconstructed = np.real(y_reconstructed)  # 取实部，因为原始信号是实数

    plt.subplot(2, 2, 4)
    librosa.display.waveshow(y_reconstructed, sr=sr, ax=plt.gca())
    plt.title('4. 重构信号 (时域)')
    plt.xlabel('时间 (s)')
    plt.ylabel('幅度')

    plt.tight_layout(rect=[0, 0, 1, 0.95])  # Adjust layout to make space for suptitle
    plt.suptitle("题三：语音信号处理与重构", fontsize=16)
    plt.show()


if __name__ == '__main__':
    draw()