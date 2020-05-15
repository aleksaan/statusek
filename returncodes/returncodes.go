package rc

type ReturnCode int

const (
	SUCCESS ReturnCode = (iota + 1) * 5
	INSTANCE_TOKEN_IS_NOT_FOUND
	OBJECT_NAME_IS_NOT_FOUND
	STATUS_NAME_IS_NOT_FOUND_FOR_OBJECT
	NOT_ALL_PREVIOS_MANDATORY_STATUSES_IS_SET
	NO_ONE_PREVIOS_OPTIONAL_STATUSES_IS_SET
	ALL_PREVIOS_STATUSES_IS_SET
	INSTANCE_IS_NOT_FINISHED
	INSTANCE_IS_FINISHED
	NEXT_STATUSES_IS_NOT_SET
	AT_LEAST_ONE_NEXT_STATUS_IS_SET
	CURRENT_STATUS_IS_SET
	CURRENT_STATUS_IS_NOT_SET
	STATUS_IS_ACCORDING_TO_INSTANCE
	STATUS_IS_NOT_ACCORDING_TO_INSTANCE
	SET_STATUS_VALIDATION_IS_FAIL
	SET_STATUS_DB_ERROR
	INSTANCE_IS_IN_TIMEOUT
	STATUS_ID_IS_NOT_FOUND
	ALL_MANDATORY_ARE_SET
	NOT_ALL_MANDATORY_ARE_SET
	STATUS_IS_SET
	STATUS_IS_NOT_SET
)

var ReturnCodes map[ReturnCode]string = make(map[ReturnCode]string)

func init() {
	ReturnCodes[SUCCESS] = "Success"
	ReturnCodes[INSTANCE_TOKEN_IS_NOT_FOUND] = "ERROR: Instance token is not found"
	ReturnCodes[OBJECT_NAME_IS_NOT_FOUND] = "ERROR: Object name is not found"
	ReturnCodes[STATUS_NAME_IS_NOT_FOUND_FOR_OBJECT] = "ERROR: Status name is not found for object"
	ReturnCodes[NOT_ALL_PREVIOS_MANDATORY_STATUSES_IS_SET] = "Not all previos mandatory statuses are set"
	ReturnCodes[NO_ONE_PREVIOS_OPTIONAL_STATUSES_IS_SET] = "No one previos optional statuses are sets"
	ReturnCodes[ALL_PREVIOS_STATUSES_IS_SET] = "All previos statuses are set"
	ReturnCodes[INSTANCE_IS_NOT_FINISHED] = "Instance is not finished"
	ReturnCodes[INSTANCE_IS_FINISHED] = "Instance is finished"
	ReturnCodes[NEXT_STATUSES_IS_NOT_SET] = "Next statuses is not set"
	ReturnCodes[AT_LEAST_ONE_NEXT_STATUS_IS_SET] = "At least one next status is set"
	ReturnCodes[CURRENT_STATUS_IS_SET] = "Status hasn't set because is already set"
	ReturnCodes[CURRENT_STATUS_IS_NOT_SET] = "Status is not set yet"
	ReturnCodes[STATUS_IS_ACCORDING_TO_INSTANCE] = "Status is according to instance"
	ReturnCodes[STATUS_IS_NOT_ACCORDING_TO_INSTANCE] = "ERROR: Status is not according to instance"
	ReturnCodes[SET_STATUS_VALIDATION_IS_FAIL] = "ERROR: Status validation is failed"
	ReturnCodes[SET_STATUS_DB_ERROR] = "ERROR: Set status database error"
	ReturnCodes[INSTANCE_IS_IN_TIMEOUT] = "Instance has been timed out"
	ReturnCodes[STATUS_ID_IS_NOT_FOUND] = "StatusId is not found"
	ReturnCodes[ALL_MANDATORY_ARE_SET] = "All mandatory statuses are set"
	ReturnCodes[NOT_ALL_MANDATORY_ARE_SET] = "Not all mandatory statuses are set"
	ReturnCodes[STATUS_IS_SET] = "Status is set"
	ReturnCodes[STATUS_IS_NOT_SET] = "Status is not set"
}
