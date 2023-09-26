#ifndef NP6118BASELIB_H
#define NP6118BASELIB_H


#if defined _WIN32

#include <Windows.h>

#ifdef NKIOLIB_EXPORTS
#define NK_API __declspec(dllexport)
#else
#define NK_API __declspec(dllimport)
#endif

#elif defined __linux__
#define NK_API
#else
#define NK_API
#endif



// Error Codes
#define NKIO_ENOERR				0		/*!< No error. */
#define NKIO_EBUSY				1		/*!< Bus is Busy. */
#define NKIO_EINUSED			2		/*!< Bus is in used. */
#define NKIO_EIO				3		/*!< Bus is error. */
#define NKIO_ETIMEOUT			4		/*!< Timeout error occurred. */
#define NKIO_EDEVERR			5		/*!< IO Device is not accessed. */
#define NKIO_EINVAL				6		/*!< Illegal argument. */
#define NKIO_EFILE				7		/*!< File is read error. */
#define NKIO_ELIB				8		/*!< load dll error. */


#ifdef __cplusplus
extern "C" {
#endif

	/**
	* @brief: NKIO library and driver initialized by the configure file.
	* @param: configFile - The path of the IO board configure file.
	*
	* @return: return the NKIO_ENOERR if succeed, otherwise, return the errorcode, please see the error codes.
	*
	* @note: Must be called once firstly before calling the read and write DIO functions.
	*/
	NK_API int NKBASE_LibraryInit();



	/**
	* @brief: deinitialize the driver and release the resources allocated in the NKIO library.
	*
	*/
	NK_API void NKBASE_LibraryDeinit(void);

	
	NK_API  int NKBASE_PollingWriteR_STS(unsigned char value);

	NK_API  int NKBASE_PollingReadR_STS(unsigned char *pValue);
	/**
	* @brief: Read DI P_OK status in the polling mode.
	* @param: pValue - the address to store the read value, the according bit is 0 if on, otherwise the according bit is 1.
	* @return: return the NKIO_ENOERR if succeed, otherwise, return the errorcode, please see the error codes.
	* @note: P_OK is reserved to connect the external UPS module to check the power status, which can also be used for the general DI.
	*/
	NK_API  int NKBASE_PollingReadP_OK(unsigned char *pValue);

	/**
	* @brief: Read Internal UPS module power status in the polling mode.
	* @param: pValue - the address to store the read value, the according bit is 0 if on, otherwise the according bit is 1.
	* @return: return the NKIO_ENOERR if succeed, otherwise, return the errorcode, please see the error codes.
	* @note: This signal can only be used when the internal UPS module is installed.
	*/
	NK_API  int NKBASE_PollingReadIntUpsPwrSts(unsigned char *pValue);

	/************************************************************************************************************/
	/*           Watchdog timer                                                                                 */
	/************************************************************************************************************/
	/**
	* @brief: Configure and start the watchdog timer.
	* @param: unit: time value unit, 0: 1sec, 1:60sec.
	*		  time: the value of the timer count, 1~255
	*
	* @return: return the NKIO_ENOERR if succeed, otherwise, return the errorcode, please see the error codes.
	*
	*/
	NK_API int NKWDT_Start(unsigned char unit,unsigned char time);

	/**
	* @brief: Reset the count value of the watchdog timer.
	* @param: unit: time: the value of the timer count, 1~255
	*
	* @return: return the NKIO_ENOERR if succeed, otherwise, return the errorcode, please see the error codes.
	*
	*/
	NK_API int NKWDT_Reset(unsigned char time);

	/**
	* @brief: Stop the timer count decrease of the watchdog timer.
	* @param: null
	*
	* @return: return the NKIO_ENOERR if succeed, otherwise, return the errorcode, please see the error codes.
	*
	*/
	NK_API int NKWDT_Stop(void);

#ifdef __cplusplus
}
#endif
#endif // NP_6118SDK_H
