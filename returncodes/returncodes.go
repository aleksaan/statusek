package returncodes

import "golang.org/x/text/language"

type ReturnCode int

const (
	SUCCESS ReturnCode = (iota + 1) * 5
	INSTANCE_TOKEN_IS_NOT_FOUND
	OBJECT_NAME_IS_NOT_FOUND
	STATUS_NAME_IS_NOT_FOUND_FOR_OBJECT
	NOT_ALL_PREVIOS_MANDATORY_STATUSES_ARE_SET
	NO_ONE_PREVIOS_OPTIONAL_STATUSES_ARE_SET
	ALL_PREVIOS_STATUSES_ARE_SET
	INSTANCE_IS_NOT_FINISHED
	INSTANCE_IS_FINISHED
	NEXT_STATUSES_IS_NOT_SET
	AT_LEAST_ONE_NEXT_STATUS_IS_SET
	STATUS_IS_ACCORDING_TO_INSTANCE
	STATUS_IS_NOT_ACCORDING_TO_INSTANCE
	INSTANCE_IS_IN_TIMEOUT
	STATUS_ID_IS_NOT_FOUND
	ALL_MANDATORY_ARE_SET
	NOT_ALL_MANDATORY_ARE_SET
	STATUS_IS_SET
	STATUS_IS_NOT_SET
	AT_LEAST_ONE_OF_PREVIOS_STATUSES_IS_SET_FOR_STOP_STATUS
	NO_ONE_PREVIOS_STATUSES_ARE_SET_FOR_STOP_STATUS
	DATABASE_ERROR
	DB_IS_NOT_UPDATED
	PARAMETERS_PARSING_ERROR
)

var ReturnCodes map[ReturnCode]string = make(map[ReturnCode]string)

func init() {
	InitReturnCodes(language.Russian)

}

func InitReturnCodes(lang language.Tag) {
	ReturnCodes[SUCCESS] = "Success"
	ReturnCodes[INSTANCE_TOKEN_IS_NOT_FOUND] = "ERROR: Instance token '%s' is not found"
	ReturnCodes[OBJECT_NAME_IS_NOT_FOUND] = "ERROR: Object name '%s' is not found"
	ReturnCodes[STATUS_NAME_IS_NOT_FOUND_FOR_OBJECT] = "ERROR: Status name '%s' is not found for object '%s'"
	ReturnCodes[NOT_ALL_PREVIOS_MANDATORY_STATUSES_ARE_SET] = "Not all previos mandatory statuses are set for status '%s'"
	ReturnCodes[NO_ONE_PREVIOS_OPTIONAL_STATUSES_ARE_SET] = "No one previos optional statuses are sets for status '%s'"
	ReturnCodes[ALL_PREVIOS_STATUSES_ARE_SET] = "All previos statuses are set for status '%s'"
	ReturnCodes[INSTANCE_IS_NOT_FINISHED] = "Instance '%s' is not finished"
	ReturnCodes[INSTANCE_IS_FINISHED] = "Instance '%s' is finished by '%s'"
	ReturnCodes[NEXT_STATUSES_IS_NOT_SET] = "Next statuses are not set for status '%s'"
	ReturnCodes[AT_LEAST_ONE_NEXT_STATUS_IS_SET] = "At least one next status is set for status '%s'"
	ReturnCodes[STATUS_IS_ACCORDING_TO_INSTANCE] = "Status '%s' is according to instance '%s'"
	ReturnCodes[STATUS_IS_NOT_ACCORDING_TO_INSTANCE] = "ERROR: Status '%s' is not according to instance '%s'"
	ReturnCodes[INSTANCE_IS_IN_TIMEOUT] = "Instance '%s' is timed out"
	ReturnCodes[STATUS_ID_IS_NOT_FOUND] = "StatusId '%d' is not found"
	ReturnCodes[ALL_MANDATORY_ARE_SET] = "All mandatory statuses are set for instance '%s'"
	ReturnCodes[NOT_ALL_MANDATORY_ARE_SET] = "Not all mandatory statuses are set for instance '%s'"
	ReturnCodes[STATUS_IS_SET] = "Status '%s' is set"
	ReturnCodes[STATUS_IS_NOT_SET] = "Status '%s' is not set"
	ReturnCodes[AT_LEAST_ONE_OF_PREVIOS_STATUSES_IS_SET_FOR_STOP_STATUS] = "At least one of previos statuses is set for a stop-status"
	ReturnCodes[NO_ONE_PREVIOS_STATUSES_ARE_SET_FOR_STOP_STATUS] = "No one previos statuses are set for a stop-status"
	ReturnCodes[DATABASE_ERROR] = "Database error: '%s'"
	ReturnCodes[DB_IS_NOT_UPDATED] = "Database was not updated because db_update_if_older_version is false"

	if lang == language.Russian {
		ReturnCodes[SUCCESS] = "Успех"
	}
}
