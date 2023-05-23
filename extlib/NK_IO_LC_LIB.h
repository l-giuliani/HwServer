#ifndef _NK_IO_LC_LIB_H_
#define _NK_IO_LC_LIB_H_

//#define __cplusplus

#if defined(_MSC_VER) || defined(WIN64) || defined(_WIN64) || defined(__WIN64__) || defined(WIN32) || defined(_WIN32) || defined(__WIN32__) || defined(__NT__)
#include <Windows.h>
   #ifdef NK_EXPORTS
	#define NK_API __declspec(dllexport)
   #else
	#define NK_API __declspec(dllimport)
   #endif
#else
   #define NK_API
#endif


#ifndef VOID
#define VOID void
#endif

#ifndef BOOL
#define BOOL int
#endif

#ifndef DWORD
#define DWORD unsigned int
#endif

#ifndef WORD
#define WORD unsigned short
#endif

#ifndef BYTE
#define BYTE unsigned char
#endif

#ifdef __cplusplus
extern "C" {
#endif
	typedef enum
	{
		E_NOERR = 0,              /*!< No error. */
		E_NOREG = 1,              /*!< Illegal register address. */
		E_INVAL = 2,              /*!< Illegal argument. */
		E_PORTERR = 3,            /*!< Porting layer error. */
		E_NORES = 4,              /*!< Insufficient resources. */
		E_IO = 5,                 /*!< I/O error. */
		E_TIMEDOUT = 6           /*!< Timeout error occurred. */
	} eErrorCode;

	typedef struct _OPENCOM_CALLBACK_ARG_T
	{
		unsigned short portNum;
		unsigned char devId;
		unsigned char hardwareMajorVer;
		unsigned char hardwareMinorVer;
		unsigned char hardwareRevVer;
		unsigned char firmwareMajorVer;
		unsigned char firmwareMinorVer;
		unsigned char firmwareRevVer;
		unsigned char fillup_1;
		unsigned char fillup_2;
		unsigned char error;
		unsigned int errorId;

	}OPENCOM_CALLBACK_ARG_T;

	typedef struct _CLOSECOM_CALLBACK_ARG_T
	{
		unsigned int fill_up;
		unsigned char fill_up_1;
		unsigned char fill_up_2;
		unsigned char fill_up_3;
		unsigned char error;
		unsigned int errorId;
		unsigned int fill_up_4;
	}CLOSECOM_CALLBACK_ARG_T;

	typedef struct _GET_DEVICE_VER_CALLBACK_ARG_T
	{
		unsigned char hardwareMajorVer;
		unsigned char hardwareMinorVer;
		unsigned char hardwareRevVer;
		unsigned char firmwareMajorVer;
		unsigned char firmwareMinorVer;
		unsigned char firmwareRevVer;
		unsigned char fill_up;
		unsigned char error;
		unsigned int errorId;
		unsigned int fill_up_2;
	}GET_DEVICE_VER_CALLBACK_ARG_T;

	typedef struct _SET_PWM_PARAMS_CALLBACK_ARG_T
	{
		unsigned int fill_up;
		unsigned char chIdx;
		unsigned char fill_up_1;
		unsigned char fill_up_2;
		unsigned char error;
		unsigned int errorId;
		unsigned int fill_up_3;
	}SET_PWM_PARAMS_CALLBACK_ARG_T;

	typedef struct _GET_PWM_PARAMS_CALLBACK_ARG_T
	{
		unsigned char chIdx;
		unsigned char pwmMode;
		unsigned char pwmValue;
		unsigned char pwmHoldingTime;
		unsigned char pwmOnOff;
		unsigned char fill_up_2;
		unsigned char fill_up_3;
		unsigned char error;
		unsigned int errorId;
		unsigned int fill_up_4;
	}GET_PWM_PARAMS_CALLBACK_ARG_T;

	// for the general parameters settings
	
	typedef struct _SET_GENERAL_PARAM_CALLBACK_ARG_T
	{
		unsigned char devId;
		unsigned char cmdId;
		unsigned char paramId;
		unsigned char paramLen;
		unsigned int paramValue;
		unsigned char fill_up_1;
		unsigned char fill_up_2;
		unsigned char fill_up_3;
		unsigned char error;
		unsigned int errorId;
	}SET_GENERAL_PARAM_CALLBACK_ARG_T;

	typedef struct _GET_GENERAL_PARAM_CALLBACK_ARG_T
	{
		unsigned char devId;
		unsigned char cmdId;
		unsigned char paramId;
		unsigned char paramLen;
		unsigned int paramValue;
		unsigned char fill_up_1;
		unsigned char fill_up_2;
		unsigned char fill_up_3;
		unsigned char error;
		unsigned int errorId;
	}GET_GENERAL_PARAM_CALLBACK_ARG_T;

	typedef struct _IAP_CALLBACK_ARG_T
	{
		unsigned char devId;
		unsigned char error;
		unsigned char fill_up_1;
		unsigned char fill_up_2;
		unsigned int errorId;
		unsigned int totalFileSize;
		unsigned int downloadedSize;

	}IAP_CALLBACK_ARG_T;

	

	// for the older version 
	
	typedef union
	{
		OPENCOM_CALLBACK_ARG_T openComCallbackArg;
		CLOSECOM_CALLBACK_ARG_T closeComCallbackArg;
		GET_DEVICE_VER_CALLBACK_ARG_T getDeviceVerCallbackArg;
		SET_PWM_PARAMS_CALLBACK_ARG_T setPwmParamsCallbackArg;
		GET_PWM_PARAMS_CALLBACK_ARG_T getPwmParamsCallbackArg;
		
		// for the general params settings and gettings
		GET_GENERAL_PARAM_CALLBACK_ARG_T getGeneralParamCallbackArg;
		SET_GENERAL_PARAM_CALLBACK_ARG_T setGeneralParamCallbackArg;		
		
		IAP_CALLBACK_ARG_T iapUpdateCallbackArg;
		
	}LC_CALLBACK_ARG_T;

	typedef void(*pLcCallbackFunc)(LC_CALLBACK_ARG_T arg);

	/*************************************************************************************************/
	/*                                                                                               */
	/*    Nodka DIO and Light Control library                                                        */
	/*                                                                                               */
	/*************************************************************************************************/
	NK_API int DIOLC_LibraryBaseInit(const char *configFile);
	NK_API int  DIOLC_OpenDevice(unsigned short port, pLcCallbackFunc pCallBackFun);
	NK_API int  DIOLC_CloseDevice(unsigned short port, pLcCallbackFunc pCallBackFun);
	NK_API int  DIOLC_IsDeviceOpened(unsigned short port, unsigned int devId);
	NK_API int  DIOLC_Process(void);
	NK_API int DIOLC_LibraryBaseDeinit();
	/*************************************************************************************************/
	/*                                                                                               */
	/*    Nodka DIO                                                                                  */
	/*                                                                                               */
	/*************************************************************************************************/
	// Read and Write DIO in polling mode
	/*-Polling Mode API------------------------------------------------------------------------------*/
	NK_API BOOL DIO_PollingReadDiBit(BYTE diByteIndex, BYTE diBitIndex);
	NK_API BYTE DIO_PollingReadDiByte(BYTE diByteIndex);
	NK_API WORD DIO_PollingReadDiWord(BYTE diWordIndex);
	NK_API VOID DIO_PollingWriteDoBit(BYTE doByteIndex, BYTE doBitIndex, BYTE doBitValue);
	NK_API VOID DIO_PollingWriteDoByte(BYTE doByteIndex, BYTE doByteValue);
	NK_API VOID DIO_PollingWriteDoWord(BYTE doWordIndex, WORD doWordValue);

	NK_API 	BOOL DIO_PollingReadDoBit(BYTE doByteIndex, BYTE doBitIndex);
	NK_API 	BYTE DIO_PollingReadDoByte(BYTE doByteIndex);
	NK_API 	WORD DIO_PollingReadDoWord(BYTE doWordIndex);
	
	

	/*************************************************************************************************/
    /*                                                                                               */
    /*    Nodka Light Control                                                                        */
    /*                                                                                               */
    /*************************************************************************************************/
	

	//NK_API int LC_NotifyProcess(void);

	//NK_API int LC_OpenPort(unsigned int devId, unsigned short port, pLcCallbackFunc pCallBackFun);
	//NK_API int LC_ClosePort(unsigned int devId, unsigned short port, pLcCallbackFunc pCallBackFun);
	//NK_API int LC_IsDeviceOpened(unsigned int devId);

	NK_API int LC_GetVerInfo(unsigned int devId, pLcCallbackFunc pCallBackFun);

	NK_API int LC_SetPwmParams(unsigned int devId,
		unsigned char ucChIdx,
		unsigned char ucPwmMode,
		unsigned char ucPwmValue,
		unsigned char ucPwmHoldingTime,
		unsigned char ucPwmOnOff,
		pLcCallbackFunc pCallBackFun);

	NK_API int LC_GetPwmParams(unsigned int devId,
		unsigned char ucChIdx,
		pLcCallbackFunc pCallBackFun);


	// for the general params settings and gettings
	NK_API int DIOLC_SetGeneralParam(unsigned int devId,
		unsigned char ucParamId,
		unsigned char ucParamLen,
		unsigned int uiParamValue,
		pLcCallbackFunc pCallBackFun);

	NK_API int DIOLC_GetGeneralParam(unsigned int devId,
		unsigned char ucParamId,
		unsigned char ucParamLen,
		pLcCallbackFunc pCallBackFun);	

	// for the firmware download
	NK_API int LC_IAPDownload(unsigned int devId,
		const char *imageFile,
		pLcCallbackFunc pCallBackIAPDownload);


#ifdef __cplusplus
}
#endif


#endif // _NK_IO_LC_LIB_H_
