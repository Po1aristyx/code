#include "stm32f10x.h"
#include "Delay.h"
#include "Seg.h"        // 数码管驱动
#include "rtc_config.h" // RTC配置 (如果单独放文件)
#include <stdbool.h>

uint8_t current_hour = 0;
uint8_t current_minute = 0;
uint8_t current_second = 0; // 虽然不显示秒，但RTC会更新它

// 如果不把RTC配置放单独文件，可以直接在这里实现RTC_Configuration等函数
// void RTC_Configuration(void) { ... }
// ErrorStatus RTC_Set_Time(uint8_t hour, uint8_t minute, uint8_t second) { ... }
// void RTC_Get_Formatted_Time(uint8_t* pHour, uint8_t* pMinute, uint8_t* pSecond) { ... }


int main(void)
{
    // SystemInit(); // 通常在 startup_stm32f10x_xx.s 中已调用

    //Delay_ms();   // 初始化延时函数
    Seg_Init();     // 初始化数码管GPIO

    RTC_Configuration(); // 配置RTC

    // 检查RTC是否是第一次配置，如果是，可以设置一个初始时间
    // 例如，通过读取备份寄存器的一个特定值来判断
    if (BKP_ReadBackupRegister(BKP_DR1) != 0xA5A5) { // 0xA5A5 是一个任意选择的魔法数字
        // 第一次配置RTC，设置初始时间为 16:45:00
        if (RTC_Set_Time(16, 45, 00) == SUCCESS) {
            BKP_WriteBackupRegister(BKP_DR1, 0xA5A5); // 写入标志，表示已设置过
        }
    }


    while(1)
    {
        // 1. 从RTC获取当前时间
        RTC_Get_Formatted_Time(&current_hour, &current_minute, &current_second);
		

        // 2. 在数码管上显示时间 HH.MM
        Seg_Set2_With_DP(current_hour*100+current_minute,2);

        
    }
}
