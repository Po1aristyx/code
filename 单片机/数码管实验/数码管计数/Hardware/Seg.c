#include "stm32f10x.h"                  // Device header
#include "Delay.h"
#include <stdbool.h> // 为了使用 bool 类型 (如果你的编译器支持 C99 或更高标准)
                     // 如果不支持，可以用 typedef char bool; #define true 1 #define false 0;

// 引脚分配:
// 段选 a-g: GPIOA 引脚 0-6 (假设PA7不是DP的一部分，或者DP是一个独立的灯)
// DP (小数点/冒号): GPIOB, GPIO_Pin_4 (由PB4独立控制，低电平点亮)
// 位选 1-4: GPIOB 引脚 0, 1, 5, 6

// 定义DP控制引脚
#define DP_CONTROL_PIN GPIO_Pin_0
#define DP_CONTROL_PORT GPIOA

// 关闭所有段并取消选中所有位 (共阳驱动)
void Seg_OFF(void)
{
    // 段熄灭：所有段选引脚(GPIOA)输出高电平
    GPIO_SetBits(GPIOA, GPIO_Pin_0 | GPIO_Pin_1 | GPIO_Pin_2 | GPIO_Pin_3 | GPIO_Pin_4 | GPIO_Pin_5 | GPIO_Pin_6 | GPIO_Pin_7); // 假设PA7也可能是某个段
    // DP熄灭：DP控制引脚(PB4)输出高电平 (因为假设低电平点亮)
    GPIO_SetBits(DP_CONTROL_PORT, DP_CONTROL_PIN);
    // 位不选：所有位选引脚(GPIOB)输出低电平
    GPIO_ResetBits(GPIOB, GPIO_Pin_0 | GPIO_Pin_1 | GPIO_Pin_5 | GPIO_Pin_6);
}


void Seg_Init(void)
{
    RCC_APB2PeriphClockCmd(RCC_APB2Periph_GPIOA, ENABLE);
    RCC_APB2PeriphClockCmd(RCC_APB2Periph_GPIOB, ENABLE);

    GPIO_InitTypeDef GPIO_Initstructure;
    GPIO_Initstructure.GPIO_Mode = GPIO_Mode_Out_PP;
    // 初始化GPIOA (PA0-PA7 用于段 a-g，或者如果你的数码管DP也是GPIOA上的，也包含在这里)
    GPIO_Initstructure.GPIO_Pin = GPIO_Pin_0 | GPIO_Pin_1 | GPIO_Pin_2 | GPIO_Pin_3 | GPIO_Pin_4 | GPIO_Pin_5 | GPIO_Pin_6 | GPIO_Pin_7;
    GPIO_Initstructure.GPIO_Speed = GPIO_Speed_50MHz;
    GPIO_Init(GPIOA, &GPIO_Initstructure);

    // 初始化GPIOB (PB0, PB1, PB5, PB6 用于位选; PB4 用于独立的DP控制)
    GPIO_Initstructure.GPIO_Pin = GPIO_Pin_0 | GPIO_Pin_1 | GPIO_Pin_5 | GPIO_Pin_6 ; // 已包含 DP_CONTROL_PIN (PB4)
    GPIO_Init(GPIOB, &GPIO_Initstructure);

    // 初始化显示为关闭状态 (包括DP也应该熄灭)
    Seg_OFF();
}

// 函数：在特定位上显示一个数字0-9的a-g段 (共阳驱动)
// digit: 1, 2, 3, or 4 (数码管位号)
// num: 0-9 (要显示的数字)
// 这个函数不再直接控制独立的DP (PB4)
void Seg_Set1(uint8_t digit, int num)
{
    // 位选 - 对于共阳，高电平选中数码管
    switch (digit)
    {
        case 1:
            GPIO_SetBits(GPIOB, GPIO_Pin_0);
            GPIO_ResetBits(GPIOB, GPIO_Pin_1 | GPIO_Pin_5 | GPIO_Pin_6);
            break;
        case 2:
            GPIO_SetBits(GPIOB, GPIO_Pin_1);
            GPIO_ResetBits(GPIOB, GPIO_Pin_0 | GPIO_Pin_5 | GPIO_Pin_6);
            break;
        case 3:
            GPIO_SetBits(GPIOB, GPIO_Pin_5);
            GPIO_ResetBits(GPIOB, GPIO_Pin_0 | GPIO_Pin_1 | GPIO_Pin_6);
            break;
        case 4:
            GPIO_SetBits(GPIOB, GPIO_Pin_6);
            GPIO_ResetBits(GPIOB, GPIO_Pin_0 | GPIO_Pin_1 | GPIO_Pin_5);
            break;
        default: // 如果位号无效，关闭所有位，并熄灭所有段
            GPIO_ResetBits(GPIOB, GPIO_Pin_0 | GPIO_Pin_1 | GPIO_Pin_5 | GPIO_Pin_6);
            GPIO_SetBits(GPIOA, GPIO_Pin_0 | GPIO_Pin_1 | GPIO_Pin_2 | GPIO_Pin_3 | GPIO_Pin_4 | GPIO_Pin_5 | GPIO_Pin_6 | GPIO_Pin_7);
            return; // 直接返回
    }

    // 段选 (a-g段，不包括独立的DP)
    // !!! 你必须根据你的硬件连接（GPIOA的哪个引脚对应a-g段）来验证或修改这里的段码 !!!
    // 假设 PA0=a, PA1=b, PA2=c, PA3=d, PA4=e, PA5=f, PA6=g. PA7不在这里控制（除非它是g的一部分）
    // 低电平点亮
    switch (num)
    {
        // 以下段码基于你原始代码的逻辑，请仔细核对是否与你的硬件一致
        // 并且确保它们只控制 a-g 段，不涉及独立的DP (PB4)
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

// intToArray 函数保持不变 (如果需要)
void intToArray(int num, int f[5]) {
    // ... (你的 intToArray 代码) ...
}


// 修改 Seg_Set2 以控制独立的 DP (PB4)
// dp_position: 希望DP在哪一位数字 *之后* 显示 (传入1则在第1位后, 2在第2位后, ...)
//              传入0则不显示DP
void Seg_Set2_With_DP(int num, uint8_t dp_position)
{
    int i;
    uint8_t display_data[4];

    if (num < 0) num = 0;
    if (num > 9999) num = 9999;

    display_data[0] = (num / 1000) % 10;
    display_data[1] = (num / 100) % 10;
    display_data[2] = (num / 10) % 10;
    display_data[3] = num % 10;

    for (i = 0; i < 4; i++) // i 从 0 到 3, 代表数码管位 1 到 4
    {
        // 1. 消隐所有段 (a-g)
        GPIO_SetBits(GPIOA, GPIO_Pin_0 | GPIO_Pin_1 | GPIO_Pin_2 | GPIO_Pin_3 | GPIO_Pin_4 | GPIO_Pin_5 | GPIO_Pin_6 | GPIO_Pin_7);
        // 2. 消隐独立的 DP (PB4) - 拉高DP引脚
        GPIO_SetBits(DP_CONTROL_PORT, DP_CONTROL_PIN);

        // 3. 设置当前位要显示的数字 (a-g段)
        Seg_Set1((uint8_t)(i + 1), display_data[i]);

        // 4. 如果当前位是需要显示DP的那一位，则点亮DP
        // 例如，如果 dp_position = 2, 意味着在第二个数码管（i=1）显示时，DP亮
        if (dp_position != 0 && (i + 1) == dp_position) {
            GPIO_ResetBits(DP_CONTROL_PORT, DP_CONTROL_PIN); // 低电平点亮DP
        }

        Delay_ms(2); // 延时
    }
    // 扫描完成后，可以选择是否关闭DP
    // GPIO_SetBits(DP_CONTROL_PORT, DP_CONTROL_PIN); // 如果不希望DP在扫描间隙一直亮
}

// 你原来的 Seg_Set2 函数可以保留，或者修改为调用新的函数
// 如果保留，它将不控制DP
void Seg_Set2(int num)
{
    // 这个版本不控制DP，如果想控制DP，调用 Seg_Set2_With_DP
    int i;
    uint8_t display_data[4];

    if (num < 0) num = 0;
    if (num > 9999) num = 9999;

    display_data[0] = (num / 1000) % 10;
    display_data[1] = (num / 100) % 10;
    display_data[2] = (num / 10) % 10;
    display_data[3] = num % 10;

    for (i = 0; i < 4; i++)
    {
        GPIO_SetBits(GPIOA, GPIO_Pin_0 | GPIO_Pin_1 | GPIO_Pin_2 | GPIO_Pin_3 | GPIO_Pin_4 | GPIO_Pin_5 | GPIO_Pin_6 | GPIO_Pin_7);
        // 确保DP是灭的，如果这个函数不应该控制DP
        GPIO_SetBits(DP_CONTROL_PORT, DP_CONTROL_PIN);
        Seg_Set1((uint8_t)(i + 1) , display_data[i]);
        Delay_ms(2);
    }
}