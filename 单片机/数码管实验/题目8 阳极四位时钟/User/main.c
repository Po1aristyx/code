#include "stm32f10x.h"   // Device header
#include "Delay.h"
#include "Timer.h"
#include "Seg.h"		//数码管头文件

uint8_t temp;
uint8_t minute =9;
uint8_t second =54;
void TIM2_IRQHandler(void)
{
	if (TIM_GetITStatus(TIM2, TIM_IT_Update) == SET)		//每10ms中断，显示一遍当前数值
	{
		temp++;
		if (temp == 100)
		{
			second++;
			temp = 0;
		}
		if (second == 60) 										
		{
			minute++;
			second = 0;
		}
		if (minute == 60)	minute = 0;
		Seg_Set1(1,minute/10);	
		Delay_ms(2);
		Seg_Set1(2,minute%10);	
		GPIO_ResetBits(GPIOA, GPIO_Pin_0);
		Delay_ms(2);
		Seg_Set1(3,second/10);	
		Delay_ms(2);
		Seg_Set1(4,second%10);	
		Delay_ms(2);
		TIM_ClearITPendingBit(TIM2, TIM_IT_Update);
	}
}

int main(void)
{
	Timer2_Init();
	Seg_Init();				//数码管初始化函数
	
	while (1)
	{
	}
}


