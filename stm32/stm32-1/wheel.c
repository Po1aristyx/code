#include "wheel.h"

void wheel_Init(){
	RCC_APB2PeriphClockCmd(RCC_APB2Periph_GPIOA, ENABLE);
	
	GPIO_InitTypeDef GPIO_InitStructure;
	GPIO_InitStructure.GPIO_Mode = GPIO_Mode_Out_PP;
	GPIO_InitStructure.GPIO_Pin = IN1 | IN2 | IN3 | IN4;
	GPIO_InitStructure.GPIO_Speed = GPIO_Speed_50MHz;
	GPIO_Init(GPIOA, &GPIO_InitStructure);
}


void Car_Go(){
	//左轮向前
	GPIO_SetBits(GPIOA, IN1);
	GPIO_ResetBits(GPIOA, IN2);
	//右轮向前
	GPIO_SetBits(GPIOA, IN3);
	GPIO_ResetBits(GPIOA, IN4);
	
	//左轮速度
	PWM_SetCompare2(100);
	//右轮速度
	PWM_SetCompare3(100);
	
}

void Stop(){
	//左轮速度
	PWM_SetCompare2(0);
	//右轮速度
	PWM_SetCompare3(0);
}

void Car_GoLeft(){
	//左轮向前
	GPIO_SetBits(GPIOA, IN1);
	GPIO_ResetBits(GPIOA, IN2);
	//右轮向前
	GPIO_SetBits(GPIOA, IN3);
	GPIO_ResetBits(GPIOA, IN4);
	
	//左轮速度
	PWM_SetCompare2(50);
	//右轮速度
	PWM_SetCompare3(55);
	
}


void Car_GoRight(){
	//左轮向前
	GPIO_SetBits(GPIOA, IN1);
	GPIO_ResetBits(GPIOA, IN2);
	//右轮向前
	GPIO_SetBits(GPIOA, IN3);
	GPIO_ResetBits(GPIOA, IN4);
	
	//左轮速度
	PWM_SetCompare2(55);
	//右轮速度
	PWM_SetCompare3(50);
	
}


void TurnLeft_Slow(){
	//左轮向后
	GPIO_ResetBits(GPIOA, IN1);
	GPIO_SetBits(GPIOA, IN2);
	//右轮向前
	GPIO_SetBits(GPIOA, IN3);
	GPIO_ResetBits(GPIOA, IN4);
	
	//左轮速度
	PWM_SetCompare2(75);
	//右轮速度
	PWM_SetCompare3(75);
}

void TurnRight_Slow(){
	//左轮向前
	GPIO_SetBits(GPIOA, IN1);
	GPIO_ResetBits(GPIOA, IN2);
	//右轮向后
	GPIO_ResetBits(GPIOA, IN3);
	GPIO_SetBits(GPIOA, IN4);
	
	//左轮速度
	PWM_SetCompare2(75);
	//右轮速度
	PWM_SetCompare3(75);
}

void TurnLeft_Fast(){
	//左轮向后
	GPIO_ResetBits(GPIOA, IN1);
	GPIO_SetBits(GPIOA, IN2);
	//右轮向前
	GPIO_SetBits(GPIOA, IN3);
	GPIO_ResetBits(GPIOA, IN4);
	
	//左轮速度
	PWM_SetCompare2(70);
	//右轮速度
	PWM_SetCompare3(70);
}

void TurnRight_Fast(){
	//左轮向前
	GPIO_SetBits(GPIOA, IN1);
	GPIO_ResetBits(GPIOA, IN2);
	//右轮向后
	GPIO_ResetBits(GPIOA, IN3);
	GPIO_SetBits(GPIOA, IN4);
	
	//左轮速度
	PWM_SetCompare2(70);
	//右轮速度
	PWM_SetCompare3(70);
}


void Angle_Left(){
	//左轮向后
	GPIO_ResetBits(GPIOA, IN1);
	GPIO_SetBits(GPIOA, IN2);
	//右轮向前
	GPIO_SetBits(GPIOA, IN3);
	GPIO_ResetBits(GPIOA, IN4);
	
	//左轮速度
	PWM_SetCompare2(90);
	//右轮速度
	PWM_SetCompare3(90);
}

void Angle_Right(){
	//左轮向前
	GPIO_SetBits(GPIOA, IN1);
	GPIO_ResetBits(GPIOA, IN2);
	//右轮向后
	GPIO_ResetBits(GPIOA, IN3);
	GPIO_SetBits(GPIOA, IN4);
	
	//左轮速度
	PWM_SetCompare2(90);
	//右轮速度
	PWM_SetCompare3(90);
}
