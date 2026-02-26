#include "stm32f10x.h"   // Device header
#include "Delay.h"
#include "Seg.h"		//数码管头文件


int i;
int main(void)
{
	
	Seg_Init();				//数码管初始化函数
	while(1)
	{
		Seg_Set2(8888);
	}
}


