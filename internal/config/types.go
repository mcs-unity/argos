package config

const (
	AllowOfflineTxForUnknownId              key = "AllowOfflineTxForUnknownId"
	AuthorizationCacheEnabled               key = "AuthorizationCacheEnabled"
	AuthorizeRemoteTxRequests               key = "AuthorizeRemoteTxRequests"
	BlinkRepeat                             key = "BlinkRepeat"
	ClockAlignedDataInterval                key = "ClockAlignedDataInterval"
	ConnectionTimeOut                       key = "ConnectionTimeOut"
	ConnectorPhaseRotation                  key = "ConnectorPhaseRotation"
	ConnectorPhaseRotationMaxLength         key = "ConnectorPhaseRotationMaxLength"
	GetConfigurationMaxKeys                 key = "GetConfigurationMaxKeys"
	HeartbeatInterval                       key = "HeartbeatInterval"
	LightIntensity                          key = "LightIntensity"
	LocalAuthorizeOffline                   key = "LocalAuthorizeOffline"
	LocalPreAuthorize                       key = "LocalPreAuthorize"
	MaxEnergyOnInvalidId                    key = "MaxEnergyOnInvalidId"
	MeterValuesAlignedData                  key = "MeterValuesAlignedData"
	MeterValuesAlignedDataMaxLength         key = "MeterValuesAlignedDataMaxLength"
	MeterValuesSampledData                  key = "MeterValuesSampledData"
	MeterValuesSampledDataMaxLength         key = "MeterValuesSampledDataMaxLength"
	MeterValueSampleInterval                key = "MeterValueSampleInterval"
	MinimumStatusDuration                   key = "MinimumStatusDuration"
	NumberOfConnectors                      key = "NumberOfConnectors"
	ResetRetries                            key = "ResetRetries"
	StopTransactionOnEVSideDisconnect       key = "StopTransactionOnEVSideDisconnect"
	StopTransactionOnInvalidId              key = "StopTransactionOnInvalidId"
	StopTxnAlignedData                      key = "StopTxnAlignedData"
	StopTxnAlignedDataMaxLength             key = "StopTxnAlignedDataMaxLength"
	StopTxnSampledData                      key = "StopTxnSampledData"
	StopTxnSampledDataMaxLength             key = "StopTxnSampledDataMaxLength"
	SupportedFeatureProfiles                key = "SupportedFeatureProfiles"
	SupportedFeatureProfilesMaxLength       key = "SupportedFeatureProfilesMaxLength"
	TransactionMessageAttempts              key = "TransactionMessageAttempts"
	TransactionMessageRetryInterval         key = "TransactionMessageRetryInterval"
	UnlockConnectorOnEVSideDisconnect       key = "UnlockConnectorOnEVSideDisconnect"
	WebSocketPingInterval                   key = "WebSocketPingInterval"
	LocalAuthListEnabled                    key = "LocalAuthListEnabled"
	LocalAuthListMaxLength                  key = "LocalAuthListMaxLength"
	SendLocalListMaxLength                  key = "SendLocalListMaxLength"
	ReserveConnectorZeroSupported           key = "ReserveConnectorZeroSupported"
	ChargeProfileMaxStackLevel              key = "ChargeProfileMaxStackLevel"
	ChargingScheduleAllowedChargingRateUnit key = "ChargingScheduleAllowedChargingRateUnit"
	ChargingScheduleMaxPeriods              key = "ChargingScheduleMaxPeriods"
	ConnectorSwitch3to1PhaseSupported       key = "ConnectorSwitch3to1PhaseSupported"
	MaxChargingProfilesInstalled            key = "MaxChargingProfilesInstalled"
)

const (
	UINT8  = "float64"
	UINT16 = "float64"
	UINT32 = "float64"
	STRING = "string"
	BOOL   = "bool"
)

type key string
type list = map[key]Variable

type Variable struct {
	ReadOnly bool   `json:"readOnly"`
	Value    any    `json:"value"`
	Type     string `json:"type"`
}
