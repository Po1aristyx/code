#include "stm32f10x.h"                  // Device header
#include "Delay.h"
#include "Buzzer.h"

// 新增静态状态变量
static uint8_t buzzer_status = 0;

void Buzzer_Init(void)
{
    RCC_APB2PeriphClockCmd(RCC_APB2Periph_GPIOA,ENABLE);
    
    GPIO_InitTypeDef GPIO_InitStructure;
    GPIO_InitStructure.GPIO_Mode=GPIO_Mode_Out_PP;
    GPIO_InitStructure.GPIO_Pin=GPIO_Pin_12;
    GPIO_InitStructure.GPIO_Speed=GPIO_Speed_50MHz;
    GPIO_Init(GPIOA,&GPIO_InitStructure);
    
    GPIO_SetBits(GPIOA,GPIO_Pin_12);
    buzzer_status = 0;  // 初始化状态为关闭
}

void Buzzer_ON(void)
{
    GPIO_ResetBits (GPIOA,GPIO_Pin_12);
    buzzer_status = 1;  // 更新状态为开启
}

void Buzzer_OFF(void)
{
    GPIO_SetBits (GPIOA,GPIO_Pin_12);
    buzzer_status = 0;  // 更新状态为关闭
}

void Buzzer_Turn(void)
{
    if(GPIO_ReadOutputDataBit (GPIOA,GPIO_Pin_12)==0)
    {
        GPIO_SetBits(GPIOA,GPIO_Pin_12);
        buzzer_status = 0;
    }
    else
    {
        GPIO_ResetBits(GPIOA,GPIO_Pin_12);
        buzzer_status = 1;
    }
}

void Buzzer_Work(uint16_t Freq,uint16_t Timeout,uint8_t Volume)
{
    uint16_t i;
    uint16_t Cycle=1000000.0/Freq;                //一个周期的时间  单位us  范围(2000~5000)us
    uint16_t H_time=Volume/100.0*Cycle;        //高电平时间     单位us
    uint16_t L_time=Cycle-H_time;                    //低电平时间     单位us
    uint16_t n=Timeout*1000/Cycle;                            //多少个周期
    
    buzzer_status = 1;  // 工作状态视为开启
    for(i=0;i<n;i++)
    {
        GPIO_ResetBits(GPIOA,GPIO_Pin_12);  // 直接操作寄存器保持原有性能
        Delay_us(H_time);
        GPIO_SetBits(GPIOA,GPIO_Pin_12);
        Delay_us(L_time);
    }
    buzzer_status = 0;  // 工作结束恢复关闭状态
}

// 新增函数：获取蜂鸣器状态
uint8_t Buzzer_GetStatus(void)
{
    return buzzer_status;
}