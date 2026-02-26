#include "stm32f10x.h"                  // Device header

  int main(void)
{
	RCC->APB2RSTR = 00000010;
	while(1)
	{
		HAL_GPIO_TogglePin(LED_GPIO_Port,LED_Pin);
		HAL_Delay(500);
	}
}
