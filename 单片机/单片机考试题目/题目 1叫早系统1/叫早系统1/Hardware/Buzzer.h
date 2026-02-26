#ifndef             _P_BUZZER_H_
#define				_P_BUZZER_H_
void Buzzer_Init(void);
 
void  Buzzer_ON(void);
 
void  Buzzer_OFF(void);
 
void Buzzer_Turn(void);
uint8_t Buzzer_GetStatus(void);
void Buzzer_Work(uint16_t Freq,uint16_t Timeout,uint8_t Volume);
#endif
