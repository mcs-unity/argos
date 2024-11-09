package connector

const (
	CONNECTOR_LOCK_FAILURE ErrorCode = "ConnectorLockFailure"
	EV_COMMUNICATION_ERROR ErrorCode = "EVCommunicationError"
	GROUND_FAILURE         ErrorCode = "GroundFailure"
	HIGH_TEMPERATURE       ErrorCode = "HighTemperature"
	INTERNAL_ERROR         ErrorCode = "InternalError"
	LOCAL_LIST_CONFLICT    ErrorCode = "LocalListConflict"
	NO_ERROR               ErrorCode = "NoError"
	OTHER_ERROR            ErrorCode = "OtherError"
	OVER_CURRENT_FAILURE   ErrorCode = "OverCurrentFailure"
	POWER_METER_FAILURE    ErrorCode = "PowerMeterFailure"
	POWER_SWITCH_FAILURE   ErrorCode = "PowerSwitchFailure"
	READER_FAILURE         ErrorCode = "ReaderFailure"
	RESET_FAILURE          ErrorCode = "ResetFailure"
	UNDER_VOLTAGE          ErrorCode = "UnderVoltage"
	OVER_VOLTAGE           ErrorCode = "OverVoltage"
	WEAK_SIGNAL            ErrorCode = "WeakSignal"
)
