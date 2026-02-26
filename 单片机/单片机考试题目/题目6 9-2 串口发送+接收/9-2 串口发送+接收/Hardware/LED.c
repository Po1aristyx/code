#include "stm32f10x.h"                  // Device header

void LED_Init(void)
{
    RCC_APB2PeriphClockCmd(RCC_APB2Periph_GPIOA, ENABLE);  // 启用GPIOA时钟

    GPIO_InitTypeDef GPIO_InitStructure;
    GPIO_InitStructure.GPIO_Mode = GPIO_Mode_Out_PP;  // 推挽输出模式
    GPIO_InitStructure.GPIO_Pin = GPIO_Pin_1 | GPIO_Pin_2 | GPIO_Pin_3;  // 配置三个LED的引脚
    GPIO_InitStructure.GPIO_Speed = GPIO_Speed_50MHz;  // 设置引脚速度
    GPIO_Init(GPIOA, &GPIO_InitStructure);  // 初始化GPIOA

    // 默认设置，关闭所有LED
    GPIO_SetBits(GPIOA, GPIO_Pin_1 | GPIO_Pin_2 | GPIO_Pin_3);
}


void LED1_ON(void)
{
	GPIO_ResetBits(GPIOA, GPIO_Pin_1);
}

void LED1_OFF(void)
{
	GPIO_SetBits(GPIOA, GPIO_Pin_1);
}

void LED1_Turn(void)
{
	if (GPIO_ReadOutputDataBit(GPIOA, GPIO_Pin_1) == 0)
	{
		GPIO_SetBits(GPIOA, GPIO_Pin_1);
	}
	else
	{
		GPIO_ResetBits(GPIOA, GPIO_Pin_1);
	}
}

void LED2_ON(void)
{
	GPIO_ResetBits(GPIOA, GPIO_Pin_2);
}

void LED2_OFF(void)
{
	GPIO_SetBits(GPIOA, GPIO_Pin_2);
}

void LED2_Turn(void)
{
	if (GPIO_ReadOutputDataBit(GPIOA, GPIO_Pin_2) == 0)
	{
		GPIO_SetBits(GPIOA, GPIO_Pin_2);
	}
	else
	{
		GPIO_ResetBits(GPIOA, GPIO_Pin_2);
	}
	
}
void LED3_ON(void)
{
    GPIO_ResetBits(GPIOA, GPIO_Pin_3);
}

void LED3_OFF(void)
{
    GPIO_SetBits(GPIOA, GPIO_Pin_3);
}

void LED3_Turn(void)
{
    if (GPIO_ReadOutputDataBit(GPIOA, GPIO_Pin_3) == 0)
    {
        GPIO_SetBits(GPIOA, GPIO_Pin_3);
    }
    else
    {
        GPIO_ResetBits(GPIOA, GPIO_Pin_3);
    }
}
