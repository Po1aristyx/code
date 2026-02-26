#include "stm32f10x.h"                  // Device header
#include "Delay.h"
#include "OLED.h"
#include "PWM.h"
#include "AD.h"

uint16_t ADValue;         // 定义AD值变量
float Voltage;            // 定义电压变量
float PWMValue;// 定义PWM值变量
float PWMV;
int main(void)
{
    // 初始化系统
    OLED_Init();             // OLED初始化
    OLED_ShowString(1, 1, "AD Value:");
    OLED_ShowString(2, 1, "Voltage:");
    OLED_ShowString(3, 1, "PWM:");
    // 初始化ADC和PWM
    AD_Init();
    PWM_Init();
    // 主循环
    while(1)
    {
        // 获取当前光强值
        uint16_t lightValue = AD_GetValue();
        ADValue = lightValue;
        Voltage = (float)ADValue / 4095 * 3.3; // 将AD值转换为电压值
        // 显示AD值和电压值
        OLED_ShowNum(1, 11, ADValue, 4);          // 显示AD值
        OLED_ShowNum(2, 11, (uint16_t)(Voltage * 100), 4); // 显示电压值，保留两位小数
        // 计算PWM占空比
        PWMValue = (ADValue - 800) * (20000 - 0) / (2000 - 800);  // 将AD值映射到100-0%的占空比，与光强互补
        // 设置PWM比较寄存器
		PWMV=(float)ADValue/4095*100;
        PWM_SetCompare2(PWMValue);
        // 显示占空比
        OLED_ShowNum(3, 11, (uint16_t)(PWMV), 4); // 显示占空比
        // 添加适当的延迟
        Delay_ms(50);
    }
}