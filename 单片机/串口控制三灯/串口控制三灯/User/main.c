#include "stm32f10x.h"     // Device header
#include "Delay.h"
#include "OLED.h"
#include "Serial.h"
#include "LED.h"           // 新增：LED 控制头文件

uint8_t RxData;

int main(void)
{
    // 初始化
    OLED_Init();
    OLED_ShowString(1, 1, "RxData:");
    
    Serial_Init();
    LED_Init();            // 新增：LED 初始化

    while (1)
    {
        // 串口有数据到来
        if (Serial_GetRxFlag() == 1)
        {
            RxData = Serial_GetRxData();    // 取出收到的字节
            Serial_SendByte(RxData);        // 回显到串口
            OLED_ShowHexNum(1, 8, RxData, 2);// OLED 上以 16 进制显示

            // 根据收到字符进行 LED 切换
            switch (RxData)
            {
                case '1':
                    LED1_Turn();
                    break;
                case '2':
                    LED2_Turn();
                    break;
                case '3':
                    LED3_Turn();
                    break;
                default:
                    // 非 1/2/3 的输入不做处理
                    break;
            }
        }
    }
}
