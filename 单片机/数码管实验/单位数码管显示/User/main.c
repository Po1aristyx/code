#include "stm32f10x.h"   // Device header
#include "Delay.h"
#include "Seg.h"		//数码管头文件

int i;
int main(void)
{
	
	Seg_Init();				//数码管初始化函数
	
	while (i < 10)
	{
		i++;
		if(i == 10)  i = 0;
		Seg_Set1(1,i);
		Delay_ms(250);
		Seg_Set1(2,i);
		Delay_ms(250);
		Seg_Set1(3,i);
		Delay_ms(250);
		Seg_Set1(4,i);
		Delay_ms(250);
	}
}


