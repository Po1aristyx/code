#include "stm32f10x.h"                  // Device header
#include "Delay.h"
#include "OLED.h"
#include "Timer.h"
#include "Key.h"

uint16_t Seconds; // 记录秒数
uint8_t Minutes; // 记录分钟
uint16_t Milliseconds; // 记录毫秒
uint8_t IsRunning = 0; // 计时状态标志位，0为停止，1为运行

int main(void) {
    OLED_Init();
    Timer_Init();
    Key_Init();

    OLED_ShowString(1, 1, "Time:");

    while (1) {
        OLED_ShowString(1, 7, "");
        OLED_ShowNum(1, 7, Minutes, 2); // 显示分钟
        OLED_ShowChar(1, 9, ':');
        OLED_ShowNum(1, 10, Seconds % 60, 2); // 显示秒数
        OLED_ShowChar(1, 12, '.');
        OLED_ShowNum(1, 13, Milliseconds % 1000, 3); // 显示毫秒

        uint8_t Key = Key_GetNum();
        if (Key == 1) {
            IsRunning = !IsRunning; // 切换计时状态
        } else if (Key == 2) {
            Seconds = 0; // 复位
            Minutes = 0;
            Milliseconds = 0;
            IsRunning = 0;
        }
    }
}

void TIM2_IRQHandler(void) {
    if (TIM_GetITStatus(TIM2, TIM_IT_Update) == SET) {
        if (IsRunning) {
            Milliseconds++;
            if (Milliseconds >= 1000) {
                Seconds++;
                Milliseconds = 0;
                if (Seconds >= 60) {
                    Minutes++;
                    Seconds = 0;
                }
            }
        }
        TIM_ClearITPendingBit(TIM2, TIM_IT_Update);
    }
}