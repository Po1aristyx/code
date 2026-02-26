#include "stm32f10x.h"

uint16_t  table[]={0x3f,0x06,0x5b,0x4f,0x66,0x6d,0x7d,0x07,0x7f,0x6f};                             //定义0～9十个数字的字型码表
uint16_t disp[2];
uint16_t temp,i;

void Delay(unsigned int count)						  	//延时函数
{
	unsigned int i;
	for(;count!=0;count--)
	{
		i=5000;
		while(i--);
	}	
}

int main(void)
{
	GPIO_InitTypeDef  GPIO_InitStructure;
	RCC_APB2PeriphClockCmd(RCC_APB2Periph_GPIOC, ENABLE); 	//使能GPIOB时钟
	GPIO_InitStructure.GPIO_Pin = 0xffff;//PC0-PC15引脚配置
	//GPIO_InitStructure.GPIO_Pin = GPIO_Pin_8|GPIO_Pin_9|GPIO_Pin_10|GPIO_Pin_11;//PB8-PB11引脚配置
	GPIO_InitStructure.GPIO_Mode = GPIO_Mode_Out_PP;   	//配置PB8为推挽输出
	GPIO_InitStructure.GPIO_Speed = GPIO_Speed_50MHz;   //GPIOB速度为50MHz
    GPIO_Init(GPIOC, &GPIO_InitStructure);				//初始化PB8
	//GPIO_SetBits(GPIOB,GPIO_Pin_8); 					//PB8输出高电平，LED熄灭
	//GPIO_Write(GPIOC, 0x0FFFF);		//先熄灭所有数码管
	//while(1);
		
	while(1)
	{
		//GPIO_ResetBits(GPIOC, 0x0FFFF);		//先熄灭所有LED
//		temp = 0x0001;
		
		
		for(i=0;i<=20;i++)
		{
			disp[1]=table[i/10];      	//十位数码管显示0
       		disp[0]=table[i%10];       	//个位数码管显示0
			temp=~((disp[1]<<8)|(disp[0]&0x0ff));
			GPIO_Write(GPIOC,temp);  
			//GPIO_SetBits(GPIOC, temp);	//向GPIOB口写控制码
			Delay(100);
			//temp =( temp<<1)+1;			// temp左移一位加1获得下一个控制码
		}
	}
}
