#ifndef _SEG_H_		//如果不包含LED_H
#define _SEG_H_   //那就定义LED_H

void Seg_Init(void);
void Seg_Set1(uint8_t digit,int num);
void intToArray(int num, int a[5]);
void Seg_Set2(int num);

#endif						//结束条件编译检查
