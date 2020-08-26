package constants

const (
	// ErrorMissingRecord is an error for missing DNS record
	ErrorMissingRecord = "Missing DNS Record! Must be specified with --record flag."

	// ErrorIPLookupFailed is an error for when IP lookup fails
	ErrorIPLookupFailed = "Public IP lookup failed! Check your internet connection."

	// ErrorRecordSetFailed is an error for when setting DNS records fails
	ErrorRecordSetFailed = "Failed to set DNS records! Check API token is valid and correct."
)

// ErrorMissingToken is an error for a missing auth token
func ErrorMissingToken(envName string) string {
	return "Missing API token! Must be specified with --token flag or " + envName + " environment variable."
}
