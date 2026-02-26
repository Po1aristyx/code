#include "stm32f10x.h"                  // Device header
#include "Delay.h"
#include "OLED.h"
#include "MyRTC.h"
#include "Timer.h"
#include "time.h"
#include "Serial.h"
#include "LightSensor.h"

void Light_On(void)
{
	GPIO_ResetBits(GPIOA, GPIO_Pin_0);
}

void Light_Off(void){
	GPIO_SetBits(GPIOA, GPIO_Pin_0);
}

void Buzzer_On(void)
{
    GPIO_ResetBits(GPIOA, GPIO_Pin_1);
}
 
void Buzzer_Off(void)
{
    GPIO_SetBits(GPIOA, GPIO_Pin_1);
}
 
 
void BEEP_Init(void)
{   
    
		RCC_APB2PeriphClockCmd(RCC_APB2Periph_GPIOA, ENABLE);
	
		GPIO_InitTypeDef  GPIO_InitStructure;
	
    GPIO_InitStructure.GPIO_Pin = GPIO_Pin_0 | GPIO_Pin_1;	 
    GPIO_InitStructure.GPIO_Mode = GPIO_Mode_Out_PP;
    GPIO_InitStructure.GPIO_Speed = GPIO_Speed_50MHz;
    GPIO_Init(GPIOA, &GPIO_InitStructure);
	
    GPIO_SetBits(GPIOA,GPIO_Pin_1);	
}

int flag;

time_t cur;
 
void Sound(u16 frq)
{
	
	if(RTC_GetCounter()-cur>=1){
		//MyRTC_ReadTime();
		//cur=RTC_GetCounter();
	}
	u32 time;
	if(frq != 1000)
	{
		time = 100000/((u32)frq);
		if (LightSensor_Get() == 1){
			flag=1;
			return ;
		}
		Buzzer_On();
		Light_On();
		Delay_ms(time);
		Buzzer_Off();
		Light_Off();
		Delay_ms(time);
	}else
		Delay_ms(100);
}

void play_music(void)
{
	uc16 tone[] = {247,262,294,330,349,392,440,494,523,587,659,698,784,1000};
 
		u8 music[]={3,5,8,6,5,13,
	                3,5,6,8,5,13,
	                8,10,9,8,9,8,6,8,5,13,
					3,5,6,5,6,8,9,5,6,13,
					3,2,1,2,13,
					2,2,3,5,5,8,2,3,5,13};
		u8 time[] ={2,2,2,2,6,4, 
				2,2,2,2,6,4,
                6,2,4,4,2,2,2,2,6,4,
				6,2,4,2,2,4,2,2,6,4,
				2,2,4,6,4,
				4,2,2,4,4,4,2,2,6,4};
 
	u32 yanshi;
	u16 i,e;
	yanshi = 2;
	for(i=0;i<sizeof(music)/sizeof(music[0]);i++){
		for(e=0;e<((u16)time[i])*tone[music[i]]/yanshi;e++){
			if (LightSensor_Get() == 1){
				flag=1;
				return ;
			}
			Sound((u32)tone[music[i]]);
			if(RTC_GetCounter()-cur>=2){
				//MyRTC_ReadTime();
				cur=RTC_GetCounter();
			}
		}
	}
}

int main(void)
{
		OLED_Init();
		MyRTC_Init();
		Timer_Init();
		Serial_Init();
		BEEP_Init();
		LightSensor_Init();
	
		MyRTC_SetTime();
		time_t time_cnt;
	
		while (1)
		{
			//MyRTC_ReadTime();
			//OLED_Clear();
			OLED_ShowString(1, 1, "2022337621031");
			OLED_ShowChinese(2, 1, 0);
			OLED_ShowChinese(2, 3, 1);
			OLED_ShowChinese(2, 5, 2);
			MyRTC_ReadTime();
			cur=RTC_GetCounter();
			if(flag==1){
				if(RTC_GetCounter()>=time_cnt+60) flag=0;
			}
			if(flag==0){
				play_music();
				if(flag==1) time_cnt = RTC_GetCounter();
			}
		}
	
}

void TIM2_IRQHandler(void)
{
	if (TIM_GetITStatus(TIM2, TIM_IT_Update) == SET)
	{
		TIM_ClearITPendingBit(TIM2, TIM_IT_Update);
	}
}
