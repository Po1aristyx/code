#include "stm32f10x.h"                  // Device header
#include "Delay.h"
#include "OLED.h"
#include "MyRTC.h"	
#include "LightSensor.h"
#include "PWM.h"
void Beep(uint16_t freq, uint16_t duration_ms) {
    RCC_APB2PeriphClockCmd(RCC_APB2Periph_GPIOA, ENABLE);
    GPIO_InitTypeDef GPIO_InitStructure;
    GPIO_InitStructure.GPIO_Mode = GPIO_Mode_Out_PP;
    GPIO_InitStructure.GPIO_Pin = GPIO_Pin_5;
    GPIO_InitStructure.GPIO_Speed = GPIO_Speed_50MHz;
    GPIO_Init(GPIOA, &GPIO_InitStructure);

    uint32_t period_us = 1000000 / freq;
    uint32_t half_period_us = period_us / 2;
    uint32_t cycles = duration_ms * 1000 / period_us;

    for (uint32_t i = 0; i < cycles; i++) {
        GPIO_SetBits(GPIOA, GPIO_Pin_5);
        Delay_us(half_period_us);
        GPIO_ResetBits(GPIOA, GPIO_Pin_5);
        Delay_us(half_period_us);
    }
}
uint16_t MyRTC_Time[] = {2025,4,7,16,0,0};		

int main(void)
{
	OLED_Init();
	MyRTC_Init();		
	LightSensor_Init();
	PWM_Init();
	OLED_ShowString(3, 1, "Date XXXX XX XX");
	OLED_ShowString(4, 1, "Time XX:XX:XX");

	OLED_ShowNum(2, 1, 202333, 6);
	OLED_ShowNum(2, 7, 7621117, 7);
	OLED_ShowChinese(1,1,0);
	OLED_ShowChinese(1,2,1);
	OLED_ShowChinese(1,3,2);
	uint16_t frequencies[] = {262, 262, 392, 392, 440, 440, 392, 349, 349, 330, 330, 294, 294, 262,
                               392, 392, 349, 349, 330, 330, 294, 392, 392, 349, 349, 330, 330, 294};
    uint16_t durations[] = {500, 500, 500, 500, 500, 500, 1000, 500, 500, 500, 500, 500, 500, 1000,
                             500, 500, 500, 500, 500, 500, 1000, 500, 500, 500, 500, 500, 500, 1000};
	while (1)
	{
		if (LightSensor_Get() == 1){
        uint8_t i;		
		uint16_t c;
		c=0;
		for (i = 0; i <sizeof(frequencies)/sizeof(frequencies[0]); i++)
		{
			if(LightSensor_Get() == 1){
			Beep(frequencies[i], durations[i]);
			if(c==1){
			    PWM_SetCompare1(10);
			}
		    else{
				PWM_SetCompare1(100);
			}
		    MyRTC_ReadTime();													
		    OLED_ShowNum(3, 6, MyRTC_Time[0], 4);		
		    OLED_ShowNum(3, 11, MyRTC_Time[1], 2);		
		    OLED_ShowNum(3, 14, MyRTC_Time[2], 2);		
		    OLED_ShowNum(4, 6, MyRTC_Time[3], 2);		
		    OLED_ShowNum(4, 9, MyRTC_Time[4], 2);		
		    OLED_ShowNum(4, 12, MyRTC_Time[5], 2);
			Delay_ms(100);
			c=(c+1)%2;
		    }
			else{
			break;
			}
	    }
		}
        else{
			Delay_s(60);
		}
	}
}
