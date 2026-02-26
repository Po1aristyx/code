// Seg.h
#ifndef __SEG_H
#define __SEG_H

#include "stm32f10x.h"
// #include <stdbool.h> // 如果在.c中用了bool，这里也可能需要，或者用int代替

void Seg_Init(void);
void Seg_OFF(void);
void Seg_Set1(uint8_t digit, int num); // 移除了bool show_dp, 因为DP由PB4独立控制
void Seg_Set2(int num); // 不控制DP的版本
void Seg_Set2_With_DP(int num, uint8_t dp_position); // 控制DP的版本
void intToArray(int num, int f[5]); // 如果需要导出

#endif