package ezpq

import (
	"github.com/lib/pq"
)

// Postgres responce codes

var (
	ErrorsPQ = map[pq.ErrorCode]error{
		// Class 00 — Successful Completion
		"00000": nil,

		// Class 01 — Warning
		"01000": ErrWarning,
		"0100C": ErrDynamicResultSetsReturned,
		"01008": ErrImplicitZeroBitPadding,
		"01003": ErrNullValueEliminatedInSetFunction,
		"01007": ErrPrivilegeNotGranted,
		"01006": ErrPrivilegeNotRevoked,
		"01004": ErrStringDataRightTruncation,
		"01P01": ErrDeprecatedFeature,

		// Class 02 — No Data (this is also a Erring class per the SQL standard)

		"2000": ErrNoData,
		"2001": ErrNoAdditionalDynamicResultSetsReturned,

		// Class 03 — SQL Statement Not Yet Complete
		"3000": ErrSqlStatementNotYetComplete,

		// Class 08 — Connection Exception
		"8000":  ErrConnectionException,
		"8003":  ErrConnectionDoesNotExist,
		"8006":  ErrConnectionFailure,
		"8001":  ErrSqlclientUnableToEstablishSqlconnection,
		"8004":  ErrSqlserverRejectedEstablishmentOfSqlconnection,
		"8007":  ErrTransactionResolutionUnknown,
		"08P01": ErrProtocolViolation,

		// Class 09 — Triggered Action Exception
		"9000": ErrTriggeredActionException,

		// Class 0A — Feature Not Supported
		"0A000": ErrFeatureNotSupported,

		// Class 0B — Invalid Transaction Initiation
		"0B000": ErrInvalidTransactionInitiation,

		// Class 0F — Locator Exception
		"0F000": ErrLocatorException,
		"0F001": ErrInvalidLocatorSpecification,

		// Class 0L — Invalid Grantor
		"0L000": ErrInvalidGrantor,
		"0LP01": ErrInvalidGrantOperation,

		// Class 0P — Invalid Role Specification
		"0P000": ErrInvalidRoleSpecification,

		// Class 0Z — Diagnostics Exception
		"0Z000": ErrDiagnosticsException,
		"0Z002": ErrStackedDiagnosticsAccessedWithoutActiveHandler,

		// Class 20 — Case Not Found
		"20000": ErrCaseNotFound,

		// Class 21 — Cardinality Violation
		"21000": ErrCardinalityViolation,

		// Class 22 — Data Exception
		"22000": ErrDataException,
		"2202E": ErrArraySubscriptError,
		"22021": ErrCharacterNotInRepertoire,
		"22008": ErrDatetimeFieldOverflow,
		"22012": ErrDivisionByZero,
		"22005": ErrErrorInAssignment,
		"2200B": ErrEscapeCharacterConflict,
		"22022": ErrIndicatorOverflow,
		"22015": ErrIntervalFieldOverflow,
		"2201E": ErrInvalidArgumentForLogarithm,
		"22014": ErrInvalidArgumentForNtileFunction,
		"22016": ErrInvalidArgumentForNthValueFunction,
		"2201F": ErrInvalidArgumentForPowerFunction,
		"2201G": ErrInvalidArgumentForWidthBucketFunction,
		"22018": ErrInvalidCharacterValueForCast,
		"22007": ErrInvalidDatetimeFormat,
		"22019": ErrInvalidEscapeCharacter,
		"2200D": ErrInvalidEscapeOctet,
		"22025": ErrInvalidEscapeSequence,
		"22P06": ErrNonstandardUseOfEscapeCharacter,
		"22010": ErrInvalidIndicatorParameterValue,
		"22023": ErrInvalidParameterValue,
		"2201B": ErrInvalidRegularExpression,
		"2201W": ErrInvalidRowCountInLimitClause,
		"2201X": ErrInvalidRowCountInResultOffsetClause,
		"22009": ErrInvalidTimeZoneDisplacementValue,
		"2200C": ErrInvalidUseOfEscapeCharacter,
		"2200G": ErrMostSpecificTypeMismatch,
		"22004": ErrDataExceptoinNullValueNotAllowed,
		"22002": ErrNullValueNoIndicatorParameter,
		"22003": ErrNumericValueOutOfRange,
		"22026": ErrStringDataLengthMismatch,
		"22001": ErrStringDataRightTruncation,
		"22011": ErrSubstringError,
		"22027": ErrTrimError,
		"22024": ErrUnterminatedCString,
		"2200F": ErrZeroLengthCharacterString,
		"22P01": ErrFloatingPointException,
		"22P02": ErrInvalidTextRepresentation,
		"22P03": ErrInvalidBinaryRepresentation,
		"22P04": ErrBadCopyFileFormat,
		"22P05": ErrUntranslatableCharacter,
		"2200L": ErrNotAnXmlDocument,
		"2200M": ErrInvalidXmlDocument,
		"2200N": ErrInvalidXmlContent,
		"2200S": ErrInvalidXmlComment,
		"2200T": ErrInvalidXmlProcessingInstruction,

		// Class 23 — Integrity Constraint Violation
		"23000": ErrIntegrityConstraintViolation,
		"23001": ErrRestrictViolation,
		"23502": ErrNotNullViolation,
		"23503": ErrForeignKeyViolation,
		"23505": ErrUniqueViolation,
		"23514": ErrCheckViolation,
		"23P01": ErrExclusionViolation,

		// Class 24 — Invalid Cursor State
		"24000": ErrInvalidCursorState,

		// Class 25 — Invalid Transaction State
		"25000": ErrInvalidTransactionState,
		"25001": ErrActiveSqlTransaction,
		"25002": ErrBranchTransactionAlreadyActive,
		"25008": ErrHeldCursorRequiresSameIsolationLevel,
		"25003": ErrInappropriateAccessModeForBranchTransaction,
		"25004": ErrInappropriateIsolationLevelForBranchTransaction,
		"25005": ErrNoActiveSqlTransactionForBranchTransaction,
		"25006": ErrReadOnlySqlTransaction,
		"25007": ErrSchemaAndDataStatementMixingNotSupported,
		"25P01": ErrNoActiveSqlTransaction,
		"25P02": ErrInFailedSqlTransaction,

		// Class 26 — Invalid SQL Statement Name
		"26000": ErrInvalidSqlStatementName,

		// Class 27 — Triggered Data Change Violation
		"27000": ErrTriggeredDataChangeViolation,

		// Class 28 — Invalid Authorization Specification
		"28000": ErrInvalidAuthorizationSpecification,
		"28P01": ErrInvalidPassword,

		// Class 2B — Dependent Privilege Descriptors Still Exist
		"2B000": ErrDependentPrivilegeDescriptorsStillExist,
		"2BP01": ErrDependentObjectsStillExist,

		// Class 2D — Invalid Transaction Termination
		"2D000": ErrInvalidTransactionTermination,

		// Class 2F — SQL Routine Exception
		"2F000": ErrSqlRoutineException,
		"2F005": ErrFunctionExecutedNoReturnStatement,
		"2F002": ErrSQLRoutineModifyingSqlDataNotPermitted,
		"2F003": ErrSQLRoutineProhibitedSqlStatementAttempted,
		"2F004": ErrSQLRoutineReadingSqlDataNotPermitted,

		// Class 34 — Invalid Cursor Name
		"34000": ErrInvalidCursorName,

		// Class 38 — External Routine Exception
		"38000": ErrExternalRoutineException,
		"38001": ErrContainingSqlNotPermitted,
		"38002": ErrExternalRoutineModifyingSqlDataNotPermitted,
		"38003": ErrExternalRoutineProhibitedSqlStatementAttempted,
		"38004": ErrExternalRoutineReadingSqlDataNotPermitted,

		// Class 39 — External Routine Invocation Exception
		"39000": ErrExternalRoutineInvocationException,
		"39001": ErrInvalidSqlstateReturned,
		"39004": ErrExternalRoutineNullValueNotAllowed,
		"39P01": ErrTriggerProtocolViolated,
		"39P02": ErrSrfProtocolViolated,

		// Class 3B — Savepoint Exception
		"3B000": ErrSavepointException,
		"3B001": ErrInvalidSavepointSpecification,

		// Class 3D — Invalid Catalog Name
		"3D000": ErrInvalidCatalogName,

		// Class 3F — Invalid Schema Name
		"3F000": ErrInvalidSchemaName,

		// Class 40 — Transaction Rollback
		"40000": ErrTransactionRollback,
		"40002": ErrTransactionIntegrityConstraintViolation,
		"40001": ErrSerializationFailure,
		"40003": ErrStatementCompletionUnknown,
		"40P01": ErrDeadlockDetected,

		// Class 42 — Syntax Error or Access Rule Violation
		"42000": ErrSyntaxErrorOrAccessRuleViolation,
		"42601": ErrSyntaxError,
		"42501": ErrInsufficientPrivilege,
		"42846": ErrCannotCoerce,
		"42803": ErrGroupingError,
		"42P20": ErrWindowingError,
		"42P19": ErrInvalidRecursion,
		"42830": ErrInvalidForeignKey,
		"42602": ErrInvalidName,
		"42622": ErrNameTooLong,
		"42939": ErrReservedName,
		"42804": ErrDatatypeMismatch,
		"42P18": ErrIndeterminateDatatype,
		"42P21": ErrCollationMismatch,
		"42P22": ErrIndeterminateCollation,
		"42809": ErrWrongObjectType,
		"42703": ErrUndefinedColumn,
		"42883": ErrUndefinedFunction,
		"42P01": ErrUndefinedTable,
		"42P02": ErrUndefinedParameter,
		"42704": ErrUndefinedObject,
		"42701": ErrDuplicateColumn,
		"42P03": ErrDuplicateCursor,
		"42P04": ErrDuplicateDatabase,
		"42723": ErrDuplicateFunction,
		"42P05": ErrDuplicatePreparedStatement,
		"42P06": ErrDuplicateSchema,
		"42P07": ErrDuplicateTable,
		"42712": ErrDuplicateAlias,
		"42710": ErrDuplicateObject,
		"42702": ErrAmbiguousColumn,
		"42725": ErrAmbiguousFunction,
		"42P08": ErrAmbiguousParameter,
		"42P09": ErrAmbiguousAlias,
		"42P10": ErrInvalidColumnReference,
		"42611": ErrInvalidColumnDefinition,
		"42P11": ErrInvalidCursorDefinition,
		"42P12": ErrInvalidDatabaseDefinition,
		"42P13": ErrInvalidFunctionDefinition,
		"42P14": ErrInvalidPreparedStatementDefinition,
		"42P15": ErrInvalidSchemaDefinition,
		"42P16": ErrInvalidTableDefinition,
		"42P17": ErrInvalidObjectDefinition,

		// Class 44 — WITH CHECK OPTION Violation
		"44000": ErrWithCheckOptionViolation,

		// Class 53 — Insufficient Resources
		"53000": ErrInsufficientResources,
		"53100": ErrDiskFull,
		"53200": ErrOutOfMemory,
		"53300": ErrTooManyConnections,
		"53400": ErrConfigurationLimitExceeded,

		// Class 54 — Program Limit Exceeded
		"54000": ErrProgramLimitExceeded,
		"54001": ErrStatementTooComplex,
		"54011": ErrTooManyColumns,
		"54023": ErrTooManyArguments,

		// Class 55 — Object Not In Prerequisite State
		"55000": ErrObjectNotInPrerequisiteState,
		"55006": ErrObjectInUse,
		"55P02": ErrCantChangeRuntimeParam,
		"55P03": ErrLockNotAvailable,

		// Class 57 — Operator Intervention
		"57000": ErrOperatorIntervention,
		"57014": ErrQueryCanceled,
		"57P01": ErrAdminShutdown,
		"57P02": ErrCrashShutdown,
		"57P03": ErrCannotConnectNow,
		"57P04": ErrDatabaseDropped,

		// Class 58 — System Error (errors external to PostgreSQL itself)
		"58000": ErrSystemError,
		"58030": ErrIoError,
		"58P01": ErrUndefinedFile,
		"58P02": ErrDuplicateFile,

		// Class F0 — Configuration File Error
		"F0000": ErrConfigFileError,
		"F0001": ErrLockFileExists,

		// Class HV — Foreign Data Wrapper Error (SQL/MED)
		"HV000": ErrFdwError,
		"HV005": ErrFdwColumnNameNotFound,
		"HV002": ErrFdwDynamicParameterValueNeeded,
		"HV010": ErrFdwFunctionSequenceError,
		"HV021": ErrFdwInconsistentDescriptorInformation,
		"HV024": ErrFdwInvalidAttributeValue,
		"HV007": ErrFdwInvalidColumnName,
		"HV008": ErrFdwInvalidColumnNumber,
		"HV004": ErrFdwInvalidDataType,
		"HV006": ErrFdwInvalidDataTypeDescriptors,
		"HV091": ErrFdwInvalidDescriptorFieldIdentifier,
		"HV00B": ErrFdwInvalidHandle,
		"HV00C": ErrFdwInvalidOptionIndex,
		"HV00D": ErrFdwInvalidOptionName,
		"HV090": ErrFdwInvalidStringLengthOrBufferLength,
		"HV00A": ErrFdwInvalidStringFormat,
		"HV009": ErrFdwInvalidUseOfNullPointer,
		"HV014": ErrFdwTooManyHandles,
		"HV001": ErrFdwOutOfMemory,
		"HV00P": ErrFdwNoSchemas,
		"HV00J": ErrFdwOptionNameNotFound,
		"HV00K": ErrFdwReplyHandle,
		"HV00Q": ErrFdwSchemaNotFound,
		"HV00R": ErrFdwTableNotFound,
		"HV00L": ErrFdwUnableToCreateExecution,
		"HV00M": ErrFdwUnableToCreateReply,
		"HV00N": ErrFdwUnableToEstablishConnection,

		// Class P0 — PL/pgSQL Error
		"P0000": ErrPlpgsqlError,
		"P0001": ErrRaiseException,
		"P0002": ErrNoDataFound,
		"P0003": ErrTooManyRows,

		// Class XX — Internal Error
		"XX000": ErrInternalError,
		"XX001": ErrDataCorrupted,
		"XX002": ErrIndexCorrupted,
	}
)
