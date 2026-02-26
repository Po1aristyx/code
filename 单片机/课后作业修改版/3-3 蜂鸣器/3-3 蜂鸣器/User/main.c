#include <REGX52.H>
#include "Delay.h"
#include "timer.h"
#include "LCD1602.h"
sbit buzzer =P2^5;   //sbit 用于定义可位寻址变量
春日影
以八分音符为基准为一拍，一拍为60/90大约600MS,其他音符为八分音符的倍数关系 
宏定义各个音符在数组中对应的索引号
#define Speed 600   
#define P  0
#define L1  1
#define L1_ 2 
#define L2  3
#define L2_ 4 
#define L3  5  
#define L4  6
#define L4_ 7 
#define L5  8
#define L5_ 9 
#define L6  10
#define L6_ 11
#define L7  12
 
#define M1  13
#define M1_ 14
#define M2  15
#define M2_ 16 
#define M3  17 
#define M4  18
#define M4_ 19
#define M5  20
#define M5_  21
#define M6  22
#define M6_ 23
#define M7  24
 
#define H1  25
#define H1_ 26
#define H2  27
#define H2_ 28
#define H3  29 
#define H4  30
#define H4_ 31
#define H5  32
#define H5_ 33
#define H6  34
#define H6_ 35
#define H7  36
以各个音符的C调频率算出11.5092MHz频率单片机中定时器的重装载值
unsigned int fre[]={
0, //简谱的休止符
63778,63873,63968,64054,64140,64215,64291,64360,64425,64489,64547,64603,  
64655,64704,64751,64795,64837,64876,64913,64948,64981,65013,65042,65070,
65095,65120,65143,65166,65187,65206,65224,65242,65259,65274,65289,65303
};
unsigned char freselect,musicsel;
unsigned char code music[]=                  //用code表示存在Flash中
{                                            //每个音符之后的数字为对应的时值，时值表达式在之
H3,8,                                        //后delay函数中
H2,4,
H1,8,
H2,4,	
	
H3,6,
H4,2,
H3,4,
H2,12,
 
H3,8,
H2,4,
H1,8,
H2,4,
 
H3,6,
H4,2,
H3,4,
H2,12,
 
H3,8,
H2,4,
H1,8,
H2,4,
 
H3,6,
H4,2,
H3,4,
H2,12,
 
H3,8,
H2,4,
H1,8,
H2,4,
 
H3,6,
H4,2,
H3,4,
H2,8,
H1,2,
H2,2,
 
	
//*
//M5,2,
//M5,2,
/
M3,4,
M3,4,
M2,4,
M4,4,
M3,4,
M2,4,
 
M2,4,
M2,4,
M1,2,
M1,2,
M4,4,
M3,4,
M2,4,
 
M2,4,
M1,4,
M1,2,
M2,2,
M3,8+4,
 
P,12,
M3,4,
M5,4,
H1,2,
H1,2,
 
M7,8,
H1,4,
M7,8,
H1,4,
 
M7,2,
M6,2,
M5,4,
P,4,
M5,4,
M2,4,
M4,4,
 
M4,8,
M3,4,
M3,4,
P,4,
L5,4,
 
M4,4,
M3,4,
M2,4,
M3,8,
M5,4,
 
M1,12,
P,4,
P,4,
M1,4,
 
 
M2,4,
M1,6,
M1,2,
M1,4,
M5,4,
M1,4,
 
M4,8,
M3,4,
M2,4,
M1,4,
M1,4,
 
M1,12,
P,4,
P,4,
M1,2,
M2,2,
 
M3,4,
M3,4,
M2,4,
M4,4,
M3,4,
M2,4,
 
M2,4,
M2,4,
M1,4,
M4,4,
M3,4,
M2,4,
 
M2,4,
M1,4,
M1,2,
M2,2,
M3,12+8,
 
P,4,
M3,4,
M5,4,
M1,4,
 
M7,8,
H1,4,
M7,8,
H1,4,
 
M7,2,
M6,2,
M5,4,
P,4,
M5,4,
M2,4,
M4,4,
 
M4,4,
M3,4,
M3,4,
M3,4,
P,4,
L5,4,
 
M4,4,
M3,4,
M2,4,
M3,8,
M5,4,
 
M1,12+8,
M1,2,
M1,2,
 
M2,4,
M1,4,
P,4,
M1,4,
M5,4,
M1,4,
 
M4,4,
M4,2,
M4,2,
M3,2,
M2,2,
M2,4,
M1,4,
M1,4,
 
M1,8+4,
P,12,
 
H6,4,
H5,4,
H5,4,
H5,4,
H4,4,
H4,4,
 
H3,4,
H2,4,
H2,4,
H2,8,
H5,4,
 
M5,4,
M4,2,
M4,2,
M4,4,
M4,4,
M3,4,
M2,4,
 
M2,4,
M1,4,
M1,2,
L7,2,
M1,12,
 
H6,4,
H5,4,
H5,4,
H5,4,
H4,4,
H4,4,
 
H3,4,
H2,4,
H2,4,
H2,8,
M3,4,
 
M3,4,
M3,2,
M3,2,
M3,2,
M3,2,
M3,4,
M2,4,
M3,4,
 
H2,8,
H1,4,
H1,4,
P,4,
H1,4,
 
M7,8,
M6,4,
M6,12,
 
P,4,
P,4,
M6,4,
M6,4,
M5,4,
M4,2,
M4,2,
 
M4,8,
M3,4,
M3,2,
M4,2,
M5,8,
 
M5,12,
P,12,
 
M3,2,
M2,2,
M3,2,
M2,2,
M3,2,
M4,2,
M5,4,
P,4,
M4,2,
M5,2,
 
M6,4,
P,4,
M6,2,
M7,2,
H1,4,
P,4,
H2,2,
H1,2,
 
M5,8,
P,2,
L5,2,
M5,4,
M4,4,
M4,4,
 
M3,8,
M3,2,
M4,2,
M5,8,
P,4,
 
M3,2,
M2,2,
M3,2,
M2,2,
M3,2,
M4,2,
M5,4,
P,4,
M4,2,
M5,2,
 
M6,4,
P,4,
M5,2,
M6,2,
M7,4,
P,4,
P,2,
M3,2,
 
H3,4,
H3,4,
P,2,
M3,2,
H4,4,
H3,4,
H2,4,
 
H2,4+2,
H1,2,
M1,2,
M7,2,
H1,4,
P,4,
M5,2,
H1,2,
 
H2,8,
H1,4,
H1,4,
P,4,
M5,4,
 
H2,8,
H1,4,
H1,4,
P,4,
M5,2,
H1,2,
 
H2,4,
H1,4,
H1,8,
P,4,
M5,2,
H1,2,
 
H2,4+2,
H3,2,
H2,4,
H1,4,
P,4,
H1,4,
 
M7,4,
M6,4,
M6,4,
M6,4,
P,4,
M5,4,
 
M5,8,
M4,4,
M4,4,
M3,4,
M2,4,
 
M3,12,
P,12,
 
M3,4,
M4,4,
M3,4,
M4,4,
M3,4,
M2,4,
 
M1,12,
M1,8,
P,4,
 
P,12,
P,4,
P,4,
M1,2,
 M2,2,
 
M3,4,
M3,4,
M2,4,
M4,4,
M3,4,
M2,4,
 
M2,4,
M2,4,
M1,2,
M1,2,
M4,4,
M3,4,
M2,4,
 
M2,6,
M2,2,
M1,2,
M2,2,
M3,12,
 
P,12,
M3,4,
M5,4,
H1,4,
 
M7,8,
H1,4,
M7,8,
H1,4,
 
M7,2,
M6,2,
M5,4,
P,4,
M5,4,
M2,4,
M4,4,
 
M4,4,
M3,4,
M3,4,
M3,4,
P,4,
P,4,
 
M4,4,
M3,4,
M2,4,
M3,8,
M5,4,
 
M1,8,
M1,8,
P,4,
M1,2,
M1,2,
 
M2,4,
M1,4,
P,4,
M1,4,
M5,4,
M1,4,
 
M4,8,
M3,4,
M2,8,
M1,4,
 
M1,12,
P,12,
 
H6,4,
H5,4,
H5,4,
H5,4,
H4,4,
H4,4,
 
H3,4,
H2,4,
H2,4,
H2,8,
M5,4,
 
 
0xFF
};
void main()
{
	LCD_Init();
	LCD_ShowString(1,1,"It's Mygo!");
  Timer0_Init( );
  while(1)
  { if(music[musicsel]!=0xFF)
		{
		freselect=music[musicsel];
		musicsel++; ///选择时值
		Delay1ms(Speed/8*music[musicsel]);  ///以八分音符为基准的时值表达式
		musicsel++;  //选择下一位音符
        TR0=0;  //定时器关掉 Delay函数不影响
		Delay1ms(5);///音符之间停顿
		TR0=1;
        }  
		else
		{
		 TR0=0;  //歌曲播放完之后关闭
		 while(1);
		}
 
}
}
 
 
void  Timer0_Routine()  interrupt 1  //中断函数，函数中定义重装载值，同时蜂鸣器IO口翻转
{ if(fre[freselect])                 //产生对应的频率
	{
	TL0 = fre[freselect]%256;				 
	TH0 = fre[freselect]/256;		
	buzzer=!buzzer;
  }
}