#include "rtc_config.h"
#include <stdio.h> // 如果需要 printf 调试

// 此函数配置RTC。在实际应用中，你需要确保LSE/LSI时钟已稳定。
void RTC_Configuration(void) {
    // 1. 使能电源和备份域时钟
    RCC_APB1PeriphClockCmd(RCC_APB1Periph_PWR | RCC_APB1Periph_BKP, ENABLE);

    // 2. 使能对备份域的访问
    PWR_BackupAccessCmd(ENABLE);

    // 3. 复位备份域 (可选，有时用于解决RTC启动问题)
    // BKP_DeInit();

    // 4. 选择RTC时钟源
    // 检查是否已配置过RTC时钟 (例如通过BKP_DR1寄存器的一个标志位)
    // 如果没有，则进行配置。这里简化为每次都尝试配置。

    // 使用LSE (外部32.768kHz晶振)
    RCC_LSEConfig(RCC_LSE_ON);
    /* 等待LSE稳定 */
    while (RCC_GetFlagStatus(RCC_FLAG_LSERDY) == RESET) {}
    RCC_RTCCLKConfig(RCC_RTCCLKSource_LSE);


    // 或者使用LSI (内部RC振荡器 - 精度较低)
    /*
    RCC_LSICmd(ENABLE);
    while (RCC_GetFlagStatus(RCC_FLAG_LSIRDY) == RESET) {}
    RCC_RTCCLKConfig(RCC_RTCCLKSource_LSI);
    */

    // 5. 使能RTC时钟
    RCC_RTCCLKCmd(ENABLE);

    // 6. 等待RTC寄存器同步
    RTC_WaitForSynchro();

    // 7. 等待RTC上一次写操作完成
    RTC_WaitForLastTask();

    // 8. 设置RTC预分频器，使RTC周期为1秒
    // 对于LSE (32.768 kHz): (32768-1) -> (PrescalerLSB = 32767)
    // 对于LSI (约40 kHz): (40000-1) -> (PrescalerLSB = 39999)
    // 根据你选择的时钟源调整预分频值
    RTC_SetPrescaler(32767); // 假设使用LSE
    // RTC_SetPrescaler(39999); // 假设使用LSI

    // 9. 等待RTC上一次写操作完成
    RTC_WaitForLastTask();

    // (可选) 设置一个初始时间，例如：16:45:00
    // 只有在第一次配置或需要重置时间时调用
    // if (BKP_ReadBackupRegister(BKP_DR1) != 0xA5A5) // 检查一个标志位，看是否已设置过时间
    // {
    //    RTC_Set_Time(16, 45, 00);
    //    BKP_WriteBackupRegister(BKP_DR1, 0xA5A5);
    // }
}

// 设置RTC时间 (小时，分钟，秒)
ErrorStatus RTC_Set_Time(uint8_t hour, uint8_t minute, uint8_t second) {
    uint32_t counter_time;

    if (hour > 23 || minute > 59 || second > 59) {
        return ERROR; // 无效时间
    }

    counter_time = (uint32_t)hour * 3600 + (uint32_t)minute * 60 + (uint32_t)second;

    RTC_WaitForLastTask(); // 等待上一次写操作完成
    RTC_SetCounter(counter_time); // 设置RTC计数器的值
    RTC_WaitForLastTask(); // 等待写操作完成
    return SUCCESS;
}

// 获取格式化的RTC时间
void RTC_Get_Formatted_Time(uint8_t* pHour, uint8_t* pMinute, uint8_t* pSecond) {
    uint32_t counter_time;
    counter_time = RTC_GetCounter(); // 读取RTC计数器的值

    *pHour = (counter_time / 3600) % 24;
    *pMinute = (counter_time % 3600) / 60;
    *pSecond = (counter_time % 3600) % 60;
}
