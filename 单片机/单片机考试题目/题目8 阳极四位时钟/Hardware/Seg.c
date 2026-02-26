#include "stm32f10x.h"                  // Device header
#include "Delay.h"

// 引脚分配 (假设你的硬件连接保持不变):
// 段选 a-g, dp: GPIOA 引脚 0-7
// 位选 1-4: GPIOB 引脚 0, 1, 5, 6


// 关闭所有段并取消选中所有位 (共阳驱动)
void Seg_OFF(void)
{
    // 段熄灭：所有段选引脚输出高电平
    GPIO_SetBits(GPIOA, GPIO_Pin_0 | GPIO_Pin_1 | GPIO_Pin_2 | GPIO_Pin_3 | GPIO_Pin_4 | GPIO_Pin_5 | GPIO_Pin_6 | GPIO_Pin_7);
    // 位不选：所有位选引脚输出低电平
    GPIO_ResetBits(GPIOB, GPIO_Pin_0 | GPIO_Pin_1 | GPIO_Pin_5 | GPIO_Pin_6);
}


void Seg_Init(void)
{
    RCC_APB2PeriphClockCmd(RCC_APB2Periph_GPIOA, ENABLE);
    RCC_APB2PeriphClockCmd(RCC_APB2Periph_GPIOB, ENABLE);

    GPIO_InitTypeDef GPIO_Initstructure;
    GPIO_Initstructure.GPIO_Mode = GPIO_Mode_Out_PP;
    GPIO_Initstructure.GPIO_Pin = GPIO_Pin_0 | GPIO_Pin_1 | GPIO_Pin_2 | GPIO_Pin_3 | GPIO_Pin_4 | GPIO_Pin_5 | GPIO_Pin_6 | GPIO_Pin_7; // 段选引脚
    GPIO_Initstructure.GPIO_Speed = GPIO_Speed_50MHz;
    GPIO_Init(GPIOA, &GPIO_Initstructure);

    // 对于共阳数码管，位选引脚连接到每个数码管的公共阳极。
    // 要选中一个数码管，其对应的位选引脚将输出高电平。
    GPIO_Initstructure.GPIO_Pin = GPIO_Pin_0 | GPIO_Pin_1 | GPIO_Pin_5 | GPIO_Pin_6;
	GPIO_Init(GPIOB, &GPIO_Initstructure);

    // 初始化显示为关闭状态
    Seg_OFF();
}

// 函数：在特定位上显示一个数字 (共阳驱动)
// digit: 1, 2, 3, or 4 (数码管位号)
// num: 0-9 (要显示的数字)
void Seg_Set1(uint8_t digit, int num)
{
    // 位选 - 对于共阳，高电平选中数码管
    switch (digit)
    {
        case 1: 
            GPIO_SetBits(GPIOB, GPIO_Pin_0);    // 选中第1位
            GPIO_ResetBits(GPIOB, GPIO_Pin_1 | GPIO_Pin_5 | GPIO_Pin_6); // 取消选中其他位
            break;
        case 2: 
            GPIO_SetBits(GPIOB, GPIO_Pin_1);    // 选中第2位
            GPIO_ResetBits(GPIOB, GPIO_Pin_0 | GPIO_Pin_5 | GPIO_Pin_6); // 取消选中其他位
            break;
        case 3:
            GPIO_SetBits(GPIOB, GPIO_Pin_5);    // 选中第3位
            GPIO_ResetBits(GPIOB, GPIO_Pin_0 | GPIO_Pin_1 | GPIO_Pin_6); // 取消选中其他位
            break;
        case 4:
			GPIO_SetBits(GPIOB, GPIO_Pin_6);    // 选中第4位 (输出高电平)
			GPIO_ResetBits(GPIOB, GPIO_Pin_0 | GPIO_Pin_1 | GPIO_Pin_5); // 取消选中其
            break;
        default: // 如果位号无效，关闭所有位
            GPIO_ResetBits(GPIOB, GPIO_Pin_0 | GPIO_Pin_1 | GPIO_Pin_5 | GPIO_Pin_6);
            break;
    }

    switch (num)
    {
        case 0: // 段 a,b,c,d,e,f 亮 (低电平); g 灭 (高电平)
            GPIO_ResetBits(GPIOA, GPIO_Pin_2 | GPIO_Pin_3 | GPIO_Pin_4 | GPIO_Pin_5 | GPIO_Pin_6 | GPIO_Pin_7);
            GPIO_SetBits(GPIOA, GPIO_Pin_1 | GPIO_Pin_0);
            break;
        case 1: // 段 b,c 亮 (低电平)
            GPIO_ResetBits(GPIOA, GPIO_Pin_5 | GPIO_Pin_6);
            GPIO_SetBits(GPIOA, GPIO_Pin_0 | GPIO_Pin_1 | GPIO_Pin_2 | GPIO_Pin_3 | GPIO_Pin_4 | GPIO_Pin_7);
            break;
        case 2: // 段 a,b,d,e,g 亮 (低电平)
            GPIO_ResetBits(GPIOA, GPIO_Pin_1 | GPIO_Pin_3 | GPIO_Pin_4 | GPIO_Pin_6 | GPIO_Pin_7);
            GPIO_SetBits(GPIOA, GPIO_Pin_0 | GPIO_Pin_2 | GPIO_Pin_5);
            break;
        case 3: // 段 a,b,c,d,g 亮 (低电平)
            GPIO_ResetBits(GPIOA, GPIO_Pin_1 | GPIO_Pin_4 | GPIO_Pin_5 | GPIO_Pin_6 | GPIO_Pin_7);
            GPIO_SetBits(GPIOA, GPIO_Pin_0 | GPIO_Pin_2 | GPIO_Pin_3);
            break;
        case 4: // 段 b,c,f,g 亮 (低电平)
            GPIO_ResetBits(GPIOA, GPIO_Pin_1 | GPIO_Pin_2 | GPIO_Pin_5 | GPIO_Pin_6);
            GPIO_SetBits(GPIOA, GPIO_Pin_0 | GPIO_Pin_3 | GPIO_Pin_4 | GPIO_Pin_7);
            break;
        case 5: // 段 a,c,d,f,g 亮 (低电平)
            GPIO_ResetBits(GPIOA, GPIO_Pin_1 | GPIO_Pin_2 | GPIO_Pin_4 | GPIO_Pin_5 | GPIO_Pin_7);
            GPIO_SetBits(GPIOA, GPIO_Pin_0 | GPIO_Pin_3 | GPIO_Pin_6);
            break;
        case 6: // 段 a,c,d,e,f,g 亮 (低电平)
            GPIO_ResetBits(GPIOA, GPIO_Pin_1 | GPIO_Pin_2 | GPIO_Pin_3 | GPIO_Pin_4 | GPIO_Pin_5 | GPIO_Pin_7);
            GPIO_SetBits(GPIOA, GPIO_Pin_0 | GPIO_Pin_6);
            break;
        case 7: // 段 a,b,c 亮 (低电平)
            GPIO_ResetBits(GPIOA, GPIO_Pin_5 | GPIO_Pin_6 | GPIO_Pin_7);
            GPIO_SetBits(GPIOA, GPIO_Pin_0 | GPIO_Pin_1 | GPIO_Pin_2 | GPIO_Pin_3 | GPIO_Pin_4);
            break;
        case 8: // 段 a,b,c,d,e,f,g 亮 (低电平)
            GPIO_ResetBits(GPIOA, GPIO_Pin_1 | GPIO_Pin_2 | GPIO_Pin_3 | GPIO_Pin_4 | GPIO_Pin_5 | GPIO_Pin_6 | GPIO_Pin_7);
            GPIO_SetBits(GPIOA, GPIO_Pin_0);
            break;
        case 9: // 段 a,b,c,d,f,g 亮 (低电平)
            GPIO_ResetBits(GPIOA,  GPIO_Pin_1 | GPIO_Pin_2 | GPIO_Pin_4 | GPIO_Pin_5 | GPIO_Pin_6 | GPIO_Pin_7);
            GPIO_SetBits(GPIOA, GPIO_Pin_0 | GPIO_Pin_3);
            break;
        default: // 如果数字无效，关闭所有段
            GPIO_SetBits(GPIOA, GPIO_Pin_0 | GPIO_Pin_1 | GPIO_Pin_2 | GPIO_Pin_3 | GPIO_Pin_4 | GPIO_Pin_5 | GPIO_Pin_6 | GPIO_Pin_7);
            break;
    }
}

void intToArray(int num, int f[5]) {
    int digits = 0;
    int temp = num;

    if (temp == 0) { // 处理数字0的情况
        digits = 1;
    } else {
        while (temp != 0) {
            temp /= 10;
            digits++;
        }
    }
    f[0] = digits > 4 ? 4 : digits; // 最多4位

    temp = num;
    for (int i = f[0]; i >= 1; i--) {
        f[5-i] = temp % 10; // 注意这里数组索引的对应关系，确保与 Seg_Set2 中的使用一致
        temp /= 10;
    }
     // 如果位数不足4位，用0填充高位 (如果Seg_Set2总是扫描4位的话)
    for (int i = 1; i < 5 - f[0]; i++) {
        f[i] = 0;
    }
}



void Seg_Set2(int num)
{
    int i;
    int display_data[4]; // 存储要显示的4个数字

    // 确保 num 在可显示范围内
    if (num < 0) num = 0;
    if (num > 9999) num = 9999; // 最多显示9999

    display_data[0] = (num / 1000) % 10; // 千位 (假设为第1位 - 最高位)
    display_data[1] = (num / 100) % 10;  // 百位 (第2位)
    display_data[2] = (num / 10) % 10;   // 十位 (第3位)
    display_data[3] = num % 10;          // 个位 (第4位 - 最低位)


    for (i = 0; i < 4; i++)
    {
        
        GPIO_SetBits(GPIOA, GPIO_Pin_0 | GPIO_Pin_1 | GPIO_Pin_2 | GPIO_Pin_3 | GPIO_Pin_4 | GPIO_Pin_5 | GPIO_Pin_6 | GPIO_Pin_7); // 所有段熄灭

        Seg_Set1(i+1 , display_data[i]); 
        Delay_ms(2); 
    }
    
}

