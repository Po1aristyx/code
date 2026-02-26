#ifndef __RTC_CONFIG_H
#define __RTC_CONFIG_H

#include "stm32f10x.h" // 根据你的MCU型号调整

void RTC_Configuration(void);
// 可选：设置时间的函数
ErrorStatus RTC_Set_Time(uint8_t hour, uint8_t minute, uint8_t second);
// 读取时间的函数（可以直接用库函数，但封装一下有时更方便）
void RTC_Get_Formatted_Time(uint8_t* pHour, uint8_t* pMinute, uint8_t* pSecond);

#endif

