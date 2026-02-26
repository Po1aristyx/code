#ifndef __BEEP_H
#define __BEEP_H

#include "stm32f1xx_hal.h"
// 初始化
void beep_init(void);

// 给定频率 驱动蜂鸣器响 
void beep(uint16_t HZ);

#endif
