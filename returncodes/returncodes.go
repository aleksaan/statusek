package returncodes

import (
	"golang.org/x/text/language"
)

type ReturnCode int

const (
	SUCCESS ReturnCode = (iota + 1) * 5
	INSTANCE_TOKEN_IS_NOT_FOUND
	OBJECT_NAME_IS_NOT_FOUND
	STATUS_NAME_IS_NOT_FOUND_FOR_OBJECT
	NOT_ALL_PREVIOS_MANDATORY_STATUSES_ARE_SET
	NO_ONE_PREVIOS_OPTIONAL_STATUSES_ARE_SET
	INSTANCE_IS_NOT_FINISHED
	INSTANCE_IS_FINISHED
	STATUS_IS_NOT_ACCORDING_TO_INSTANCE
	INSTANCE_IS_TIMEOUT
	INSTANCE_IS_NOT_TIMEOUT
	ALL_MANDATORY_ARE_SET
	NOT_ALL_MANDATORY_ARE_SET
	STATUS_IS_ALREADY_SET
	DATABASE_ERROR
	DB_IS_NOT_UPDATED
	AT_LEAST_ONE_NEXT_STATUS_IS_SET
	STATUS_IS_SET
	STATUS_IS_NOT_SET
	PARAMS_PARSING_IS_FAILED
)

var ReturnCodes map[ReturnCode]string = make(map[ReturnCode]string)

func init() {
	InitReturnCodes(language.Russian)
}

func InitReturnCodes(lang language.Tag) {
	ReturnCodes[SUCCESS] = "Success"
	ReturnCodes[INSTANCE_TOKEN_IS_NOT_FOUND] = "ERROR: Instance token '<InstanceToken>' is not found"
	ReturnCodes[OBJECT_NAME_IS_NOT_FOUND] = "ERROR: Object name '<ObjectName>' is not found"
	ReturnCodes[STATUS_NAME_IS_NOT_FOUND_FOR_OBJECT] = "ERROR: Status name '<StatusName>' is not found for instance '<InstanceToken>'"
	ReturnCodes[NOT_ALL_PREVIOS_MANDATORY_STATUSES_ARE_SET] = "Not all previos mandatory statuses are set for status '<StatusName>'"
	ReturnCodes[NO_ONE_PREVIOS_OPTIONAL_STATUSES_ARE_SET] = "No one previos optional statuses are sets for status '<StatusName>'"
	ReturnCodes[INSTANCE_IS_NOT_FINISHED] = "Instance '<InstanceToken>' is not finished"
	ReturnCodes[INSTANCE_IS_FINISHED] = "Instance '<InstanceToken>' is finished by '<InstanceIsFinishedDescription>'"
	ReturnCodes[STATUS_IS_NOT_ACCORDING_TO_INSTANCE] = "ERROR: Status '<StatusName>' is not according to instance '<InstanceToken>'"
	ReturnCodes[INSTANCE_IS_TIMEOUT] = "Instance '<InstanceToken>' is timed out"
	ReturnCodes[ALL_MANDATORY_ARE_SET] = "All mandatory statuses are set for instance '<InstanceToken>'"
	ReturnCodes[NOT_ALL_MANDATORY_ARE_SET] = "Not all mandatory statuses are set for instance '<InstanceToken>'"
	ReturnCodes[STATUS_IS_ALREADY_SET] = "Status '<StatusName>' is already set for instance '<InstanceToken>'"
	ReturnCodes[DATABASE_ERROR] = "Database error: '%s'"
	ReturnCodes[DB_IS_NOT_UPDATED] = "Database was not updated because db_update_if_older_version is false"
	ReturnCodes[AT_LEAST_ONE_NEXT_STATUS_IS_SET] = "At least one next status is set for optional status '<StatusName>'"
	ReturnCodes[PARAMS_PARSING_IS_FAILED] = "Error while parameters parsing"

	if lang == language.Russian {
		ReturnCodes[SUCCESS] = "Успех"
	}
}
