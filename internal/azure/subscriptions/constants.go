// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package subscriptions

const (
	moduleName    = "armcostmanagement"
	moduleVersion = "v2.0.0"
)

// AccumulatedType - Show costs accumulated over time.
type AccumulatedType string

const (
	AccumulatedTypeFalse AccumulatedType = "false"
	AccumulatedTypeTrue  AccumulatedType = "true"
)

// PossibleAccumulatedTypeValues returns the possible values for the AccumulatedType const type.
func PossibleAccumulatedTypeValues() []AccumulatedType {
	return []AccumulatedType{
		AccumulatedTypeFalse,
		AccumulatedTypeTrue,
	}
}

// ActionType - Enum. Indicates the action type. "Internal" refers to actions that are for internal only APIs.
type ActionType string

const (
	ActionTypeInternal ActionType = "Internal"
)

// PossibleActionTypeValues returns the possible values for the ActionType const type.
func PossibleActionTypeValues() []ActionType {
	return []ActionType{
		ActionTypeInternal,
	}
}

// AlertCategory - Alert category
type AlertCategory string

const (
	AlertCategoryBilling AlertCategory = "Billing"
	AlertCategoryCost    AlertCategory = "Cost"
	AlertCategorySystem  AlertCategory = "System"
	AlertCategoryUsage   AlertCategory = "Usage"
)

// PossibleAlertCategoryValues returns the possible values for the AlertCategory const type.
func PossibleAlertCategoryValues() []AlertCategory {
	return []AlertCategory{
		AlertCategoryBilling,
		AlertCategoryCost,
		AlertCategorySystem,
		AlertCategoryUsage,
	}
}

// AlertCriteria - Criteria that triggered alert
type AlertCriteria string

const (
	AlertCriteriaCostThresholdExceeded          AlertCriteria = "CostThresholdExceeded"
	AlertCriteriaCreditThresholdApproaching     AlertCriteria = "CreditThresholdApproaching"
	AlertCriteriaCreditThresholdReached         AlertCriteria = "CreditThresholdReached"
	AlertCriteriaCrossCloudCollectionError      AlertCriteria = "CrossCloudCollectionError"
	AlertCriteriaCrossCloudNewDataAvailable     AlertCriteria = "CrossCloudNewDataAvailable"
	AlertCriteriaForecastCostThresholdExceeded  AlertCriteria = "ForecastCostThresholdExceeded"
	AlertCriteriaForecastUsageThresholdExceeded AlertCriteria = "ForecastUsageThresholdExceeded"
	AlertCriteriaGeneralThresholdError          AlertCriteria = "GeneralThresholdError"
	AlertCriteriaInvoiceDueDateApproaching      AlertCriteria = "InvoiceDueDateApproaching"
	AlertCriteriaInvoiceDueDateReached          AlertCriteria = "InvoiceDueDateReached"
	AlertCriteriaMultiCurrency                  AlertCriteria = "MultiCurrency"
	AlertCriteriaQuotaThresholdApproaching      AlertCriteria = "QuotaThresholdApproaching"
	AlertCriteriaQuotaThresholdReached          AlertCriteria = "QuotaThresholdReached"
	AlertCriteriaUsageThresholdExceeded         AlertCriteria = "UsageThresholdExceeded"
)

// PossibleAlertCriteriaValues returns the possible values for the AlertCriteria const type.
func PossibleAlertCriteriaValues() []AlertCriteria {
	return []AlertCriteria{
		AlertCriteriaCostThresholdExceeded,
		AlertCriteriaCreditThresholdApproaching,
		AlertCriteriaCreditThresholdReached,
		AlertCriteriaCrossCloudCollectionError,
		AlertCriteriaCrossCloudNewDataAvailable,
		AlertCriteriaForecastCostThresholdExceeded,
		AlertCriteriaForecastUsageThresholdExceeded,
		AlertCriteriaGeneralThresholdError,
		AlertCriteriaInvoiceDueDateApproaching,
		AlertCriteriaInvoiceDueDateReached,
		AlertCriteriaMultiCurrency,
		AlertCriteriaQuotaThresholdApproaching,
		AlertCriteriaQuotaThresholdReached,
		AlertCriteriaUsageThresholdExceeded,
	}
}

// AlertOperator - operator used to compare currentSpend with amount
type AlertOperator string

const (
	AlertOperatorEqualTo              AlertOperator = "EqualTo"
	AlertOperatorGreaterThan          AlertOperator = "GreaterThan"
	AlertOperatorGreaterThanOrEqualTo AlertOperator = "GreaterThanOrEqualTo"
	AlertOperatorLessThan             AlertOperator = "LessThan"
	AlertOperatorLessThanOrEqualTo    AlertOperator = "LessThanOrEqualTo"
	AlertOperatorNone                 AlertOperator = "None"
)

// PossibleAlertOperatorValues returns the possible values for the AlertOperator const type.
func PossibleAlertOperatorValues() []AlertOperator {
	return []AlertOperator{
		AlertOperatorEqualTo,
		AlertOperatorGreaterThan,
		AlertOperatorGreaterThanOrEqualTo,
		AlertOperatorLessThan,
		AlertOperatorLessThanOrEqualTo,
		AlertOperatorNone,
	}
}

// AlertSource - Source of alert
type AlertSource string

const (
	AlertSourcePreset AlertSource = "Preset"
	AlertSourceUser   AlertSource = "User"
)

// PossibleAlertSourceValues returns the possible values for the AlertSource const type.
func PossibleAlertSourceValues() []AlertSource {
	return []AlertSource{
		AlertSourcePreset,
		AlertSourceUser,
	}
}

// AlertStatus - alert status
type AlertStatus string

const (
	AlertStatusActive     AlertStatus = "Active"
	AlertStatusDismissed  AlertStatus = "Dismissed"
	AlertStatusNone       AlertStatus = "None"
	AlertStatusOverridden AlertStatus = "Overridden"
	AlertStatusResolved   AlertStatus = "Resolved"
)

// PossibleAlertStatusValues returns the possible values for the AlertStatus const type.
func PossibleAlertStatusValues() []AlertStatus {
	return []AlertStatus{
		AlertStatusActive,
		AlertStatusDismissed,
		AlertStatusNone,
		AlertStatusOverridden,
		AlertStatusResolved,
	}
}

// AlertTimeGrainType - Type of timegrain cadence
type AlertTimeGrainType string

const (
	AlertTimeGrainTypeAnnually       AlertTimeGrainType = "Annually"
	AlertTimeGrainTypeBillingAnnual  AlertTimeGrainType = "BillingAnnual"
	AlertTimeGrainTypeBillingMonth   AlertTimeGrainType = "BillingMonth"
	AlertTimeGrainTypeBillingQuarter AlertTimeGrainType = "BillingQuarter"
	AlertTimeGrainTypeMonthly        AlertTimeGrainType = "Monthly"
	AlertTimeGrainTypeNone           AlertTimeGrainType = "None"
	AlertTimeGrainTypeQuarterly      AlertTimeGrainType = "Quarterly"
)

// PossibleAlertTimeGrainTypeValues returns the possible values for the AlertTimeGrainType const type.
func PossibleAlertTimeGrainTypeValues() []AlertTimeGrainType {
	return []AlertTimeGrainType{
		AlertTimeGrainTypeAnnually,
		AlertTimeGrainTypeBillingAnnual,
		AlertTimeGrainTypeBillingMonth,
		AlertTimeGrainTypeBillingQuarter,
		AlertTimeGrainTypeMonthly,
		AlertTimeGrainTypeNone,
		AlertTimeGrainTypeQuarterly,
	}
}

// AlertType - type of alert
type AlertType string

const (
	AlertTypeBudget         AlertType = "Budget"
	AlertTypeBudgetForecast AlertType = "BudgetForecast"
	AlertTypeCredit         AlertType = "Credit"
	AlertTypeGeneral        AlertType = "General"
	AlertTypeInvoice        AlertType = "Invoice"
	AlertTypeQuota          AlertType = "Quota"
	AlertTypeXCloud         AlertType = "xCloud"
)

// PossibleAlertTypeValues returns the possible values for the AlertType const type.
func PossibleAlertTypeValues() []AlertType {
	return []AlertType{
		AlertTypeBudget,
		AlertTypeBudgetForecast,
		AlertTypeCredit,
		AlertTypeGeneral,
		AlertTypeInvoice,
		AlertTypeQuota,
		AlertTypeXCloud,
	}
}

// BenefitKind - Kind/type of the benefit.
type BenefitKind string

const (
	// BenefitKindIncludedQuantity - Benefit is IncludedQuantity.
	BenefitKindIncludedQuantity BenefitKind = "IncludedQuantity"
	// BenefitKindReservation - Benefit is Reservation.
	BenefitKindReservation BenefitKind = "Reservation"
	// BenefitKindSavingsPlan - Benefit is SavingsPlan.
	BenefitKindSavingsPlan BenefitKind = "SavingsPlan"
)

// PossibleBenefitKindValues returns the possible values for the BenefitKind const type.
func PossibleBenefitKindValues() []BenefitKind {
	return []BenefitKind{
		BenefitKindIncludedQuantity,
		BenefitKindReservation,
		BenefitKindSavingsPlan,
	}
}

// ChartType - Chart type of the main view in Cost Analysis. Required.
type ChartType string

const (
	ChartTypeArea          ChartType = "Area"
	ChartTypeGroupedColumn ChartType = "GroupedColumn"
	ChartTypeLine          ChartType = "Line"
	ChartTypeStackedColumn ChartType = "StackedColumn"
	ChartTypeTable         ChartType = "Table"
)

// PossibleChartTypeValues returns the possible values for the ChartType const type.
func PossibleChartTypeValues() []ChartType {
	return []ChartType{
		ChartTypeArea,
		ChartTypeGroupedColumn,
		ChartTypeLine,
		ChartTypeStackedColumn,
		ChartTypeTable,
	}
}

// CheckNameAvailabilityReason - The reason why the given name is not available.
type CheckNameAvailabilityReason string

const (
	CheckNameAvailabilityReasonAlreadyExists CheckNameAvailabilityReason = "AlreadyExists"
	CheckNameAvailabilityReasonInvalid       CheckNameAvailabilityReason = "Invalid"
)

// PossibleCheckNameAvailabilityReasonValues returns the possible values for the CheckNameAvailabilityReason const type.
func PossibleCheckNameAvailabilityReasonValues() []CheckNameAvailabilityReason {
	return []CheckNameAvailabilityReason{
		CheckNameAvailabilityReasonAlreadyExists,
		CheckNameAvailabilityReasonInvalid,
	}
}

// CostDetailsDataFormat - The data format of the report
type CostDetailsDataFormat string

const (
	// CostDetailsDataFormatCSVCostDetailsDataFormat - Csv data format.
	CostDetailsDataFormatCSVCostDetailsDataFormat CostDetailsDataFormat = "Csv"
)

// PossibleCostDetailsDataFormatValues returns the possible values for the CostDetailsDataFormat const type.
func PossibleCostDetailsDataFormatValues() []CostDetailsDataFormat {
	return []CostDetailsDataFormat{
		CostDetailsDataFormatCSVCostDetailsDataFormat,
	}
}

// CostDetailsMetricType - The type of the detailed report. By default ActualCost is provided
type CostDetailsMetricType string

const (
	// CostDetailsMetricTypeActualCostCostDetailsMetricType - Actual cost data.
	CostDetailsMetricTypeActualCostCostDetailsMetricType CostDetailsMetricType = "ActualCost"
	// CostDetailsMetricTypeAmortizedCostCostDetailsMetricType - Amortized cost data.
	CostDetailsMetricTypeAmortizedCostCostDetailsMetricType CostDetailsMetricType = "AmortizedCost"
)

// PossibleCostDetailsMetricTypeValues returns the possible values for the CostDetailsMetricType const type.
func PossibleCostDetailsMetricTypeValues() []CostDetailsMetricType {
	return []CostDetailsMetricType{
		CostDetailsMetricTypeActualCostCostDetailsMetricType,
		CostDetailsMetricTypeAmortizedCostCostDetailsMetricType,
	}
}

// CostDetailsStatusType - The status of the cost details operation
type CostDetailsStatusType string

const (
	// CostDetailsStatusTypeCompletedCostDetailsStatusType - Operation is Completed.
	CostDetailsStatusTypeCompletedCostDetailsStatusType CostDetailsStatusType = "Completed"
	// CostDetailsStatusTypeFailedCostDetailsStatusType - Operation Failed.
	CostDetailsStatusTypeFailedCostDetailsStatusType CostDetailsStatusType = "Failed"
	// CostDetailsStatusTypeNoDataFoundCostDetailsStatusType - Operation is Completed and no cost data found.
	CostDetailsStatusTypeNoDataFoundCostDetailsStatusType CostDetailsStatusType = "NoDataFound"
)

// PossibleCostDetailsStatusTypeValues returns the possible values for the CostDetailsStatusType const type.
func PossibleCostDetailsStatusTypeValues() []CostDetailsStatusType {
	return []CostDetailsStatusType{
		CostDetailsStatusTypeCompletedCostDetailsStatusType,
		CostDetailsStatusTypeFailedCostDetailsStatusType,
		CostDetailsStatusTypeNoDataFoundCostDetailsStatusType,
	}
}

// CreatedByType - The type of identity that created the resource.
type CreatedByType string

const (
	CreatedByTypeApplication     CreatedByType = "Application"
	CreatedByTypeKey             CreatedByType = "Key"
	CreatedByTypeManagedIdentity CreatedByType = "ManagedIdentity"
	CreatedByTypeUser            CreatedByType = "User"
)

// PossibleCreatedByTypeValues returns the possible values for the CreatedByType const type.
func PossibleCreatedByTypeValues() []CreatedByType {
	return []CreatedByType{
		CreatedByTypeApplication,
		CreatedByTypeKey,
		CreatedByTypeManagedIdentity,
		CreatedByTypeUser,
	}
}

// DaysOfWeek - Days of Week.
type DaysOfWeek string

const (
	DaysOfWeekFriday    DaysOfWeek = "Friday"
	DaysOfWeekMonday    DaysOfWeek = "Monday"
	DaysOfWeekSaturday  DaysOfWeek = "Saturday"
	DaysOfWeekSunday    DaysOfWeek = "Sunday"
	DaysOfWeekThursday  DaysOfWeek = "Thursday"
	DaysOfWeekTuesday   DaysOfWeek = "Tuesday"
	DaysOfWeekWednesday DaysOfWeek = "Wednesday"
)

// PossibleDaysOfWeekValues returns the possible values for the DaysOfWeek const type.
func PossibleDaysOfWeekValues() []DaysOfWeek {
	return []DaysOfWeek{
		DaysOfWeekFriday,
		DaysOfWeekMonday,
		DaysOfWeekSaturday,
		DaysOfWeekSunday,
		DaysOfWeekThursday,
		DaysOfWeekTuesday,
		DaysOfWeekWednesday,
	}
}

// ExecutionStatus - The last known status of the export run.
type ExecutionStatus string

const (
	ExecutionStatusCompleted           ExecutionStatus = "Completed"
	ExecutionStatusDataNotAvailable    ExecutionStatus = "DataNotAvailable"
	ExecutionStatusFailed              ExecutionStatus = "Failed"
	ExecutionStatusInProgress          ExecutionStatus = "InProgress"
	ExecutionStatusNewDataNotAvailable ExecutionStatus = "NewDataNotAvailable"
	ExecutionStatusQueued              ExecutionStatus = "Queued"
	ExecutionStatusTimeout             ExecutionStatus = "Timeout"
)

// PossibleExecutionStatusValues returns the possible values for the ExecutionStatus const type.
func PossibleExecutionStatusValues() []ExecutionStatus {
	return []ExecutionStatus{
		ExecutionStatusCompleted,
		ExecutionStatusDataNotAvailable,
		ExecutionStatusFailed,
		ExecutionStatusInProgress,
		ExecutionStatusNewDataNotAvailable,
		ExecutionStatusQueued,
		ExecutionStatusTimeout,
	}
}

// ExecutionType - The type of the export run.
type ExecutionType string

const (
	ExecutionTypeOnDemand  ExecutionType = "OnDemand"
	ExecutionTypeScheduled ExecutionType = "Scheduled"
)

// PossibleExecutionTypeValues returns the possible values for the ExecutionType const type.
func PossibleExecutionTypeValues() []ExecutionType {
	return []ExecutionType{
		ExecutionTypeOnDemand,
		ExecutionTypeScheduled,
	}
}

// ExportType - The type of the query.
type ExportType string

const (
	ExportTypeActualCost    ExportType = "ActualCost"
	ExportTypeAmortizedCost ExportType = "AmortizedCost"
	ExportTypeUsage         ExportType = "Usage"
)

// PossibleExportTypeValues returns the possible values for the ExportType const type.
func PossibleExportTypeValues() []ExportType {
	return []ExportType{
		ExportTypeActualCost,
		ExportTypeAmortizedCost,
		ExportTypeUsage,
	}
}

type ExternalCloudProviderType string

const (
	ExternalCloudProviderTypeExternalBillingAccounts ExternalCloudProviderType = "externalBillingAccounts"
	ExternalCloudProviderTypeExternalSubscriptions   ExternalCloudProviderType = "externalSubscriptions"
)

// PossibleExternalCloudProviderTypeValues returns the possible values for the ExternalCloudProviderType const type.
func PossibleExternalCloudProviderTypeValues() []ExternalCloudProviderType {
	return []ExternalCloudProviderType{
		ExternalCloudProviderTypeExternalBillingAccounts,
		ExternalCloudProviderTypeExternalSubscriptions,
	}
}

// FileFormat - Destination of the view data. Currently only CSV format is supported.
type FileFormat string

const (
	FileFormatCSV FileFormat = "Csv"
)

// PossibleFileFormatValues returns the possible values for the FileFormat const type.
func PossibleFileFormatValues() []FileFormat {
	return []FileFormat{
		FileFormatCSV,
	}
}

// ForecastOperatorType - The operator to use for comparison.
type ForecastOperatorType string

const (
	ForecastOperatorTypeIn ForecastOperatorType = "In"
)

// PossibleForecastOperatorTypeValues returns the possible values for the ForecastOperatorType const type.
func PossibleForecastOperatorTypeValues() []ForecastOperatorType {
	return []ForecastOperatorType{
		ForecastOperatorTypeIn,
	}
}

// ForecastTimeframe - The time frame for pulling data for the forecast.
type ForecastTimeframe string

const (
	ForecastTimeframeCustom ForecastTimeframe = "Custom"
)

// PossibleForecastTimeframeValues returns the possible values for the ForecastTimeframe const type.
func PossibleForecastTimeframeValues() []ForecastTimeframe {
	return []ForecastTimeframe{
		ForecastTimeframeCustom,
	}
}

// ForecastType - The type of the forecast.
type ForecastType string

const (
	ForecastTypeActualCost    ForecastType = "ActualCost"
	ForecastTypeAmortizedCost ForecastType = "AmortizedCost"
	ForecastTypeUsage         ForecastType = "Usage"
)

// PossibleForecastTypeValues returns the possible values for the ForecastType const type.
func PossibleForecastTypeValues() []ForecastType {
	return []ForecastType{
		ForecastTypeActualCost,
		ForecastTypeAmortizedCost,
		ForecastTypeUsage,
	}
}

// FormatType - The format of the export being delivered. Currently only 'Csv' is supported.
type FormatType string

const (
	FormatTypeCSV FormatType = "Csv"
)

// PossibleFormatTypeValues returns the possible values for the FormatType const type.
func PossibleFormatTypeValues() []FormatType {
	return []FormatType{
		FormatTypeCSV,
	}
}

// FunctionName - The name of the column to aggregate.
type FunctionName string

const (
	FunctionNameCost          FunctionName = "Cost"
	FunctionNameCostUSD       FunctionName = "CostUSD"
	FunctionNamePreTaxCost    FunctionName = "PreTaxCost"
	FunctionNamePreTaxCostUSD FunctionName = "PreTaxCostUSD"
)

// PossibleFunctionNameValues returns the possible values for the FunctionName const type.
func PossibleFunctionNameValues() []FunctionName {
	return []FunctionName{
		FunctionNameCost,
		FunctionNameCostUSD,
		FunctionNamePreTaxCost,
		FunctionNamePreTaxCostUSD,
	}
}

// FunctionType - The name of the aggregation function to use.
type FunctionType string

const (
	FunctionTypeSum FunctionType = "Sum"
)

// PossibleFunctionTypeValues returns the possible values for the FunctionType const type.
func PossibleFunctionTypeValues() []FunctionType {
	return []FunctionType{
		FunctionTypeSum,
	}
}

// GenerateDetailedCostReportMetricType - The type of the detailed report. By default ActualCost is provided
type GenerateDetailedCostReportMetricType string

const (
	GenerateDetailedCostReportMetricTypeActualCost    GenerateDetailedCostReportMetricType = "ActualCost"
	GenerateDetailedCostReportMetricTypeAmortizedCost GenerateDetailedCostReportMetricType = "AmortizedCost"
)

// PossibleGenerateDetailedCostReportMetricTypeValues returns the possible values for the GenerateDetailedCostReportMetricType const type.
func PossibleGenerateDetailedCostReportMetricTypeValues() []GenerateDetailedCostReportMetricType {
	return []GenerateDetailedCostReportMetricType{
		GenerateDetailedCostReportMetricTypeActualCost,
		GenerateDetailedCostReportMetricTypeAmortizedCost,
	}
}

// Grain - Grain which corresponds to value.
type Grain string

const (
	// GrainDaily - Hourly grain corresponds to value per day.
	GrainDaily Grain = "Daily"
	// GrainHourly - Hourly grain corresponds to value per hour.
	GrainHourly Grain = "Hourly"
	// GrainMonthly - Hourly grain corresponds to value per month.
	GrainMonthly Grain = "Monthly"
)

// PossibleGrainValues returns the possible values for the Grain const type.
func PossibleGrainValues() []Grain {
	return []Grain{
		GrainDaily,
		GrainHourly,
		GrainMonthly,
	}
}

type GrainParameter string

const (
	// GrainParameterDaily - Hourly grain corresponds to value per day.
	GrainParameterDaily GrainParameter = "Daily"
	// GrainParameterHourly - Hourly grain corresponds to value per hour.
	GrainParameterHourly GrainParameter = "Hourly"
	// GrainParameterMonthly - Hourly grain corresponds to value per month.
	GrainParameterMonthly GrainParameter = "Monthly"
)

// PossibleGrainParameterValues returns the possible values for the GrainParameter const type.
func PossibleGrainParameterValues() []GrainParameter {
	return []GrainParameter{
		GrainParameterDaily,
		GrainParameterHourly,
		GrainParameterMonthly,
	}
}

// GranularityType - The granularity of rows in the forecast.
type GranularityType string

const (
	GranularityTypeDaily GranularityType = "Daily"
)

// PossibleGranularityTypeValues returns the possible values for the GranularityType const type.
func PossibleGranularityTypeValues() []GranularityType {
	return []GranularityType{
		GranularityTypeDaily,
	}
}

// KpiType - KPI type (Forecast, Budget).
type KpiType string

const (
	KpiTypeBudget   KpiType = "Budget"
	KpiTypeForecast KpiType = "Forecast"
)

// PossibleKpiTypeValues returns the possible values for the KpiType const type.
func PossibleKpiTypeValues() []KpiType {
	return []KpiType{
		KpiTypeBudget,
		KpiTypeForecast,
	}
}

// LookBackPeriod - The number of days used to look back.
type LookBackPeriod string

const (
	// LookBackPeriodLast30Days - 30 days used to look back.
	LookBackPeriodLast30Days LookBackPeriod = "Last30Days"
	// LookBackPeriodLast60Days - 60 days used to look back.
	LookBackPeriodLast60Days LookBackPeriod = "Last60Days"
	// LookBackPeriodLast7Days - 7 days used to look back.
	LookBackPeriodLast7Days LookBackPeriod = "Last7Days"
)

// PossibleLookBackPeriodValues returns the possible values for the LookBackPeriod const type.
func PossibleLookBackPeriodValues() []LookBackPeriod {
	return []LookBackPeriod{
		LookBackPeriodLast30Days,
		LookBackPeriodLast60Days,
		LookBackPeriodLast7Days,
	}
}

// MetricType - Metric to use when displaying costs.
type MetricType string

const (
	MetricTypeAHUB          MetricType = "AHUB"
	MetricTypeActualCost    MetricType = "ActualCost"
	MetricTypeAmortizedCost MetricType = "AmortizedCost"
)

// PossibleMetricTypeValues returns the possible values for the MetricType const type.
func PossibleMetricTypeValues() []MetricType {
	return []MetricType{
		MetricTypeAHUB,
		MetricTypeActualCost,
		MetricTypeAmortizedCost,
	}
}

// OperationStatusType - The status of the long running operation.
type OperationStatusType string

const (
	OperationStatusTypeCompleted OperationStatusType = "Completed"
	OperationStatusTypeFailed    OperationStatusType = "Failed"
	OperationStatusTypeRunning   OperationStatusType = "Running"
)

// PossibleOperationStatusTypeValues returns the possible values for the OperationStatusType const type.
func PossibleOperationStatusTypeValues() []OperationStatusType {
	return []OperationStatusType{
		OperationStatusTypeCompleted,
		OperationStatusTypeFailed,
		OperationStatusTypeRunning,
	}
}

// OperatorType - The operator to use for comparison.
type OperatorType string

const (
	OperatorTypeContains OperatorType = "Contains"
	OperatorTypeIn       OperatorType = "In"
)

// PossibleOperatorTypeValues returns the possible values for the OperatorType const type.
func PossibleOperatorTypeValues() []OperatorType {
	return []OperatorType{
		OperatorTypeContains,
		OperatorTypeIn,
	}
}

// Origin - The intended executor of the operation; as in Resource Based Access Control (RBAC) and audit logs UX. Default
// value is "user,system"
type Origin string

const (
	OriginSystem     Origin = "system"
	OriginUser       Origin = "user"
	OriginUserSystem Origin = "user,system"
)

// PossibleOriginValues returns the possible values for the Origin const type.
func PossibleOriginValues() []Origin {
	return []Origin{
		OriginSystem,
		OriginUser,
		OriginUserSystem,
	}
}

// PivotType - Data type to show in view.
type PivotType string

const (
	PivotTypeDimension PivotType = "Dimension"
	PivotTypeTagKey    PivotType = "TagKey"
)

// PossiblePivotTypeValues returns the possible values for the PivotType const type.
func PossiblePivotTypeValues() []PivotType {
	return []PivotType{
		PivotTypeDimension,
		PivotTypeTagKey,
	}
}

// QueryColumnType - The type of the column in the report.
type QueryColumnType string

const (
	// QueryColumnTypeDimension - The dimension of cost data.
	QueryColumnTypeDimension QueryColumnType = "Dimension"
	// QueryColumnTypeTagKey - The tag associated with the cost data.
	QueryColumnTypeTagKey QueryColumnType = "TagKey"
)

// PossibleQueryColumnTypeValues returns the possible values for the QueryColumnType const type.
func PossibleQueryColumnTypeValues() []QueryColumnType {
	return []QueryColumnType{
		QueryColumnTypeDimension,
		QueryColumnTypeTagKey,
	}
}

// QueryOperatorType - The operator to use for comparison.
type QueryOperatorType string

const (
	QueryOperatorTypeIn QueryOperatorType = "In"
)

// PossibleQueryOperatorTypeValues returns the possible values for the QueryOperatorType const type.
func PossibleQueryOperatorTypeValues() []QueryOperatorType {
	return []QueryOperatorType{
		QueryOperatorTypeIn,
	}
}

// RecurrenceType - The schedule recurrence.
type RecurrenceType string

const (
	RecurrenceTypeAnnually RecurrenceType = "Annually"
	RecurrenceTypeDaily    RecurrenceType = "Daily"
	RecurrenceTypeMonthly  RecurrenceType = "Monthly"
	RecurrenceTypeWeekly   RecurrenceType = "Weekly"
)

// PossibleRecurrenceTypeValues returns the possible values for the RecurrenceType const type.
func PossibleRecurrenceTypeValues() []RecurrenceType {
	return []RecurrenceType{
		RecurrenceTypeAnnually,
		RecurrenceTypeDaily,
		RecurrenceTypeMonthly,
		RecurrenceTypeWeekly,
	}
}

// ReportConfigSortingType - Direction of sort.
type ReportConfigSortingType string

const (
	ReportConfigSortingTypeAscending  ReportConfigSortingType = "Ascending"
	ReportConfigSortingTypeDescending ReportConfigSortingType = "Descending"
)

// PossibleReportConfigSortingTypeValues returns the possible values for the ReportConfigSortingType const type.
func PossibleReportConfigSortingTypeValues() []ReportConfigSortingType {
	return []ReportConfigSortingType{
		ReportConfigSortingTypeAscending,
		ReportConfigSortingTypeDescending,
	}
}

// ReportGranularityType - The granularity of rows in the report.
type ReportGranularityType string

const (
	ReportGranularityTypeDaily   ReportGranularityType = "Daily"
	ReportGranularityTypeMonthly ReportGranularityType = "Monthly"
)

// PossibleReportGranularityTypeValues returns the possible values for the ReportGranularityType const type.
func PossibleReportGranularityTypeValues() []ReportGranularityType {
	return []ReportGranularityType{
		ReportGranularityTypeDaily,
		ReportGranularityTypeMonthly,
	}
}

// ReportOperationStatusType - The status of the long running operation.
type ReportOperationStatusType string

const (
	ReportOperationStatusTypeCompleted       ReportOperationStatusType = "Completed"
	ReportOperationStatusTypeFailed          ReportOperationStatusType = "Failed"
	ReportOperationStatusTypeInProgress      ReportOperationStatusType = "InProgress"
	ReportOperationStatusTypeNoDataFound     ReportOperationStatusType = "NoDataFound"
	ReportOperationStatusTypeQueued          ReportOperationStatusType = "Queued"
	ReportOperationStatusTypeReadyToDownload ReportOperationStatusType = "ReadyToDownload"
	ReportOperationStatusTypeTimedOut        ReportOperationStatusType = "TimedOut"
)

// PossibleReportOperationStatusTypeValues returns the possible values for the ReportOperationStatusType const type.
func PossibleReportOperationStatusTypeValues() []ReportOperationStatusType {
	return []ReportOperationStatusType{
		ReportOperationStatusTypeCompleted,
		ReportOperationStatusTypeFailed,
		ReportOperationStatusTypeInProgress,
		ReportOperationStatusTypeNoDataFound,
		ReportOperationStatusTypeQueued,
		ReportOperationStatusTypeReadyToDownload,
		ReportOperationStatusTypeTimedOut,
	}
}

// ReportTimeframeType - The time frame for pulling data for the report. If custom, then a specific time period must be provided.
type ReportTimeframeType string

const (
	ReportTimeframeTypeCustom      ReportTimeframeType = "Custom"
	ReportTimeframeTypeMonthToDate ReportTimeframeType = "MonthToDate"
	ReportTimeframeTypeWeekToDate  ReportTimeframeType = "WeekToDate"
	ReportTimeframeTypeYearToDate  ReportTimeframeType = "YearToDate"
)

// PossibleReportTimeframeTypeValues returns the possible values for the ReportTimeframeType const type.
func PossibleReportTimeframeTypeValues() []ReportTimeframeType {
	return []ReportTimeframeType{
		ReportTimeframeTypeCustom,
		ReportTimeframeTypeMonthToDate,
		ReportTimeframeTypeWeekToDate,
		ReportTimeframeTypeYearToDate,
	}
}

// ReportType - The type of the report. Usage represents actual usage, forecast represents forecasted data and UsageAndForecast
// represents both usage and forecasted data. Actual usage and forecasted data can be
// differentiated based on dates.
type ReportType string

const (
	ReportTypeUsage ReportType = "Usage"
)

// PossibleReportTypeValues returns the possible values for the ReportType const type.
func PossibleReportTypeValues() []ReportType {
	return []ReportType{
		ReportTypeUsage,
	}
}

// ReservationReportSchema - The CSV file from the reportUrl blob link consists of reservation usage data with the following
// schema at daily granularity
type ReservationReportSchema string

const (
	ReservationReportSchemaInstanceFlexibilityGroup ReservationReportSchema = "InstanceFlexibilityGroup"
	ReservationReportSchemaInstanceFlexibilityRatio ReservationReportSchema = "InstanceFlexibilityRatio"
	ReservationReportSchemaInstanceID               ReservationReportSchema = "InstanceId"
	ReservationReportSchemaKind                     ReservationReportSchema = "Kind"
	ReservationReportSchemaReservationID            ReservationReportSchema = "ReservationId"
	ReservationReportSchemaReservationOrderID       ReservationReportSchema = "ReservationOrderId"
	ReservationReportSchemaReservedHours            ReservationReportSchema = "ReservedHours"
	ReservationReportSchemaSKUName                  ReservationReportSchema = "SkuName"
	ReservationReportSchemaTotalReservedQuantity    ReservationReportSchema = "TotalReservedQuantity"
	ReservationReportSchemaUsageDate                ReservationReportSchema = "UsageDate"
	ReservationReportSchemaUsedHours                ReservationReportSchema = "UsedHours"
)

// PossibleReservationReportSchemaValues returns the possible values for the ReservationReportSchema const type.
func PossibleReservationReportSchemaValues() []ReservationReportSchema {
	return []ReservationReportSchema{
		ReservationReportSchemaInstanceFlexibilityGroup,
		ReservationReportSchemaInstanceFlexibilityRatio,
		ReservationReportSchemaInstanceID,
		ReservationReportSchemaKind,
		ReservationReportSchemaReservationID,
		ReservationReportSchemaReservationOrderID,
		ReservationReportSchemaReservedHours,
		ReservationReportSchemaSKUName,
		ReservationReportSchemaTotalReservedQuantity,
		ReservationReportSchemaUsageDate,
		ReservationReportSchemaUsedHours,
	}
}

// ScheduleFrequency - Frequency of the schedule.
type ScheduleFrequency string

const (
	// ScheduleFrequencyDaily - Cost analysis data will be emailed every day.
	ScheduleFrequencyDaily ScheduleFrequency = "Daily"
	// ScheduleFrequencyMonthly - Cost analysis data will be emailed every month.
	ScheduleFrequencyMonthly ScheduleFrequency = "Monthly"
	// ScheduleFrequencyWeekly - Cost analysis data will be emailed every week.
	ScheduleFrequencyWeekly ScheduleFrequency = "Weekly"
)

// PossibleScheduleFrequencyValues returns the possible values for the ScheduleFrequency const type.
func PossibleScheduleFrequencyValues() []ScheduleFrequency {
	return []ScheduleFrequency{
		ScheduleFrequencyDaily,
		ScheduleFrequencyMonthly,
		ScheduleFrequencyWeekly,
	}
}

// ScheduledActionKind - Kind of the scheduled action.
type ScheduledActionKind string

const (
	// ScheduledActionKindEmail - Cost analysis data will be emailed.
	ScheduledActionKindEmail ScheduledActionKind = "Email"
	// ScheduledActionKindInsightAlert - Cost anomaly information will be emailed. Available only on subscription scope at daily
	// frequency. If no anomaly is detected on the resource, an email won't be sent.
	ScheduledActionKindInsightAlert ScheduledActionKind = "InsightAlert"
)

// PossibleScheduledActionKindValues returns the possible values for the ScheduledActionKind const type.
func PossibleScheduledActionKindValues() []ScheduledActionKind {
	return []ScheduledActionKind{
		ScheduledActionKindEmail,
		ScheduledActionKindInsightAlert,
	}
}

// ScheduledActionStatus - Status of the scheduled action.
type ScheduledActionStatus string

const (
	// ScheduledActionStatusDisabled - Scheduled action is saved but will not be run.
	ScheduledActionStatusDisabled ScheduledActionStatus = "Disabled"
	// ScheduledActionStatusEnabled - Scheduled action is saved and will be run.
	ScheduledActionStatusEnabled ScheduledActionStatus = "Enabled"
	// ScheduledActionStatusExpired - Scheduled action is expired.
	ScheduledActionStatusExpired ScheduledActionStatus = "Expired"
)

// PossibleScheduledActionStatusValues returns the possible values for the ScheduledActionStatus const type.
func PossibleScheduledActionStatusValues() []ScheduledActionStatus {
	return []ScheduledActionStatus{
		ScheduledActionStatusDisabled,
		ScheduledActionStatusEnabled,
		ScheduledActionStatusExpired,
	}
}

// Scope - Kind of the recommendation scope.
type Scope string

const (
	// ScopeShared - Shared scope recommendation.
	ScopeShared Scope = "Shared"
	// ScopeSingle - Single scope recommendation.
	ScopeSingle Scope = "Single"
)

// PossibleScopeValues returns the possible values for the Scope const type.
func PossibleScopeValues() []Scope {
	return []Scope{
		ScopeShared,
		ScopeSingle,
	}
}

// StatusType - The status of the export's schedule. If 'Inactive', the export's schedule is paused.
type StatusType string

const (
	StatusTypeActive   StatusType = "Active"
	StatusTypeInactive StatusType = "Inactive"
)

// PossibleStatusTypeValues returns the possible values for the StatusType const type.
func PossibleStatusTypeValues() []StatusType {
	return []StatusType{
		StatusTypeActive,
		StatusTypeInactive,
	}
}

// Term - Grain which corresponds to value.
type Term string

const (
	// TermP1Y - Benefit term is 1 year.
	TermP1Y Term = "P1Y"
	// TermP3Y - Benefit term is 3 years.
	TermP3Y Term = "P3Y"
)

// PossibleTermValues returns the possible values for the Term const type.
func PossibleTermValues() []Term {
	return []Term{
		TermP1Y,
		TermP3Y,
	}
}

// TimeframeType - The time frame for pulling data for the query. If custom, then a specific time period must be provided.
type TimeframeType string

const (
	TimeframeTypeBillingMonthToDate  TimeframeType = "BillingMonthToDate"
	TimeframeTypeCustom              TimeframeType = "Custom"
	TimeframeTypeMonthToDate         TimeframeType = "MonthToDate"
	TimeframeTypeTheLastBillingMonth TimeframeType = "TheLastBillingMonth"
	TimeframeTypeTheLastMonth        TimeframeType = "TheLastMonth"
	TimeframeTypeWeekToDate          TimeframeType = "WeekToDate"
)

// PossibleTimeframeTypeValues returns the possible values for the TimeframeType const type.
func PossibleTimeframeTypeValues() []TimeframeType {
	return []TimeframeType{
		TimeframeTypeBillingMonthToDate,
		TimeframeTypeCustom,
		TimeframeTypeMonthToDate,
		TimeframeTypeTheLastBillingMonth,
		TimeframeTypeTheLastMonth,
		TimeframeTypeWeekToDate,
	}
}

// WeeksOfMonth - Weeks of month.
type WeeksOfMonth string

const (
	WeeksOfMonthFirst  WeeksOfMonth = "First"
	WeeksOfMonthFourth WeeksOfMonth = "Fourth"
	WeeksOfMonthLast   WeeksOfMonth = "Last"
	WeeksOfMonthSecond WeeksOfMonth = "Second"
	WeeksOfMonthThird  WeeksOfMonth = "Third"
)

// PossibleWeeksOfMonthValues returns the possible values for the WeeksOfMonth const type.
func PossibleWeeksOfMonthValues() []WeeksOfMonth {
	return []WeeksOfMonth{
		WeeksOfMonthFirst,
		WeeksOfMonthFourth,
		WeeksOfMonthLast,
		WeeksOfMonthSecond,
		WeeksOfMonthThird,
	}
}
