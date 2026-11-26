package app

// pmSaved names the recovery that saved this run's package manager
// (e.g. apt_hook_retried); empty when none was needed. Kept for the logs
// and the tests.
var pmSaved string
