
// 该文件用于 beep驱动 
#include "beep.h"

// 引用外部变量
extern TIM_HandleTypeDef htim2;

// 初始化
void beep_init(void)
{
	HAL_TIM_PWM_Start(&htim2,TIM_CHANNEL_1);
	return;
}

// 给定频率 驱动蜂鸣器响 
void beep(uint16_t HZ)
{
	uint16_t b = 0,c = 0;
	// 入参检查
	if(HZ < 20 )
	{// 不响
		__HAL_TIM_SET_COMPARE(&htim2,TIM_CHANNEL_1,0);
		return;
	}
	c = 1000000 / HZ;  // 计算重载值 
	b = c / 2; // 设置占空比 50%
	
	__HAL_TIM_SET_AUTORELOAD(&htim2,c);
	__HAL_TIM_SET_COMPARE(&htim2,TIM_CHANNEL_1,b);
	return;
}


