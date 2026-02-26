#include "stm32f10x.h"
#include "Delay.h"
#include "Buzzer.h"
#include "beep.h"
#include "PWM.h"
#include "LightSensor.h"
#include "LED.h"

uint8_t lastLightState = 0;

void BEEP_Init(void)
{   
    GPIO_InitTypeDef  GPIO_InitStructure;
    RCC_APB2PeriphClockCmd(RCC_APB2Periph_GPIOA, ENABLE);
    GPIO_InitStructure.GPIO_Pin = GPIO_Pin_12;     
    GPIO_InitStructure.GPIO_Mode = GPIO_Mode_Out_PP;
    GPIO_InitStructure.GPIO_Speed = GPIO_Speed_50MHz;
    GPIO_Init(GPIOB, &GPIO_InitStructure);
    GPIO_SetBits(GPIOA,GPIO_Pin_12);    
}


void Sound(uint16_t frq, uint16_t duration)
{
    uint32_t time;
    uint32_t cycles = duration * 1000UL / (1000000UL / frq); // 计算完整周期数
    
    for(uint32_t i = 0; i < cycles; i++){
        if(frq != 1000)
        {
            Buzzer_ON();
            PWM_SetCompare1(100); // LED全亮
            time = 1000000 / frq / 2;
            Delay_us(time);
            
            Buzzer_OFF();
            PWM_SetCompare1(0);   // LED全灭
            Delay_us(time);
        }
        else
        {
            PWM_SetCompare1(0);   // 不发音时LED灭
            Delay_us(1000);
        }
    }
}

void play_music(void)
{
    uint16_t tone[] = {247,262,294,330,349,392,440,494,523,587,659,698,784,1000};
    uint8_t music[] = {
        1,1,5,5,6,6,5,4,4,3,3,2,2,1,
        5,5,4,4,3,3,2,5,5,4,4,3,3,2,
        1,1,5,5,6,6,5,4,4,3,3,2,2,1
    };

    uint8_t time[] = {
        2,2,2,2,2,2,4,2,2,2,2,2,2,4,
        2,2,2,2,2,2,4,2,2,2,2,2,2,4,
        2,2,2,2,2,2,4,2,2,2,2,2,2,4
    };

    for(uint8_t i = 0; i < sizeof(music)/sizeof(music[0]); i++){
        // 根据音符时长设置不同亮度变化速度
        uint16_t note_duration = time[i] * 300; // 基础时长300ms
        Sound(tone[music[i]], note_duration);
        
        // 音符间隔时LED渐灭
        for(int j=100; j>=0; j-=20){
            PWM_SetCompare1(j);
            Delay_ms(note_duration/50);
        }
    }
}

int main(void)
{
    PWM_Init();
    Buzzer_Init();
    LightSensor_Init();
    BEEP_Init();
    LED_Init();

    while(1)
    {
        uint8_t currentLightState = LightSensor_Get();

        if(lastLightState && !currentLightState){
            play_music();
            PWM_SetCompare1(0); // 音乐结束后关闭LED
        }
        
        lastLightState = currentLightState;

        // 空闲状态呼吸灯（仅在无音乐时运行）
        if(Buzzer_GetStatus() == 0){
            static uint8_t dir = 0;
            static uint8_t brightness = 0;
            
            if(dir == 0){
                if(++brightness >= 100) dir = 1;
            }else{
                if(--brightness <= 0) dir = 0;
            }
            PWM_SetCompare1(brightness);
            Delay_ms(10);
        }
    }
}
