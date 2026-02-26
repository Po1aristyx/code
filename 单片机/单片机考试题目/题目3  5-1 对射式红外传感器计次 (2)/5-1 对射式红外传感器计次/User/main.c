#include "stm32f10x.h"
#include "Delay.h"
#include "OLED.h"
#include "CountSensor.h"

uint16_t count = 0;
uint32_t time_ms = 60000; 

void Show_Time(uint32_t ms)
{
	uint8_t sec = (ms / 1000) % 60;
	uint8_t min = (ms / 60000) % 60;
	OLED_ShowString(4, 1, "Time:");
	OLED_ShowNum(4, 6, min, 2);
	OLED_ShowChar(4, 8, ':');
	OLED_ShowNum(4, 9, sec, 2);
	OLED_ShowChar(4, 11, ':');
	OLED_ShowNum(4, 12, 0, 2); 
}

int main(void)
{
	OLED_Init();
	CountSensor_Init();
	
	OLED_ShowString(1, 1, "Count:");
	OLED_ShowString(4, 1, "Time:01:00:00");

	while (time_ms > 0)
	{
		count = CountSensor_Get(); 
		OLED_ShowNum(1, 7, count, 5); 
		
		OLED_ShowChinese(2, 1, 0);
		OLED_ShowChinese(2, 2, 1);
		OLED_ShowChinese(2, 3, 2);
		
		time_ms -= 100; 
		Show_Time(time_ms);
		
		Delay_ms(100); 
	}
	
	OLED_ShowString(3, 1, "Done!");
	while (1);
}
