#include "stm32f10x.h"                  // Device header
#include "Delay.h"
#include "OLED.h"
#include "MPU6050.h"
#include "LED.h"  // 包含LED灯驱动文件
#include "Servo.h"
uint8_t ID;
int16_t AX, AY, AZ, GX, GY, GZ;
#define THRESHOLD 250  // 自由下落判断阈值，根据实际测试调整

int main(void)
{
    OLED_Init();
    MPU6050_Init();
    LED_Init();  // 初始化LED灯
	Servo_Init();

    OLED_ShowString(1, 1, "ID:");
    ID = MPU6050_GetID();
    OLED_ShowHexNum(1, 4, ID, 2);

    while (1)
    {
        MPU6050_GetData(&AX, &AY, &AZ, &GX, &GY, &GZ);

        // 显示数据
        OLED_ShowSignedNum(2, 1, AX, 5);
        OLED_ShowSignedNum(3, 1, AY, 5);
        OLED_ShowSignedNum(4, 1, AZ, 5);
        OLED_ShowSignedNum(2, 8, GX, 5);
        OLED_ShowSignedNum(3, 8, GY, 5);
        OLED_ShowSignedNum(4, 8, GZ, 5);

        // 判断是否处于自由下落状态
        if (abs(GX) > THRESHOLD && abs(GY) > THRESHOLD && abs(GZ) > THRESHOLD)
        {
            LED1_ON();  // 触发防摔措施，开启LED1灯提示
			Servo_SetAngle(0); 
			Delay_ms(500);
            OLED_ShowString(5, 7, "Fall Detected!");
        }
        else
        {
            LED1_OFF();  // 正常状态，关闭LED1灯
			Servo_SetAngle(90); 
			
            OLED_ShowString(5, 7, "Normal     ");  // 显示正常状态信息
        }

        Delay_ms(100);  // 延时，避免过于频繁的检测
    }
}
