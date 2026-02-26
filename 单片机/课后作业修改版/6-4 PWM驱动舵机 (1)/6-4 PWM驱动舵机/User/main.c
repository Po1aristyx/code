#include "stm32f10x.h"                  // Device header
#include "Delay.h"
#include "OLED.h"
#include "Servo.h"
#include "Key.h"
#include "CountSensor.h"
uint8_t KeyNum;			//定义用于接收键码的变量
float Angle=0;			//定义角度变量
int count=0;
int main(void)
{
	/*模块初始化*/
	OLED_Init();		//OLED初始化
	Servo_Init();		//舵机初始化
	Key_Init();			//按键初始化
	CountSensor_Init();
	/*显示静态字符串*/
	OLED_ShowString(1, 1, "Angle:");	//1行1列显示字符串Angle:
	OLED_ShowString(2, 1, "count:");
	Servo_SetAngle(Angle);
	while (1)
	{
        if(CountSensor_Get()>count){
			Angle += 90;			//角度变量自增30
			Servo_SetAngle(Angle);
			count++;
			OLED_ShowNum(2, 7, count, 3);
			Delay_s(10);
			Angle = 0;				//角度变量归零
			Servo_SetAngle(Angle);
			
		}
		OLED_ShowNum(1, 7, Angle, 3);	//OLED显示角度变量
	}
}
