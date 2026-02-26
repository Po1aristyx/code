#ifndef MOTOR_H_
#define MOTOR_H_

#include "main.h"

#define AIN1_RESET  HAL_GPIO_WritePin(AIN1_GPIO_Port,AIN1_Pin,GPIO_PIN_RESET)//设置AIN1 PB13为 低电平
#define AIN1_SET    HAL_GPIO_WritePin(AIN1_GPIO_Port,AIN1_Pin,GPIO_PIN_SET)//设置AIN1 PB13为 高电平

#define BIN1_RESET 	HAL_GPIO_WritePin(BIN1_GPIO_Port,BIN1_Pin,GPIO_PIN_RESET)  //设置BIN1 PB3为低电平
#define BIN1_SET    HAL_GPIO_WritePin(BIN1_GPIO_Port,BIN1_Pin,GPIO_PIN_SET)//设置AIN1 PB13为 高电平


void Motoer_Set (int motor1,int motor2);
/*******************
*  @brief  设置两个电机转速和方向
*  @param  motor1:电机B设置参数、motor2:设置参数
*  @param  motor1: 输入1~100 对应控制B电机正方向速度在1%-100%、输入-1~-100 对应控制B电机反方向速度在1%-100%、motor2同理
*  @return  无
*
*******************/


#endif
