package envVars

import (
	"fmt"
	"os"
)

func GetEnvValue(envVarName string) (string, error) {
	envVar, envVarDefined := os.LookupEnv(envVarName)
	if !envVarDefined {
		return "", fmt.Errorf("could not find '%v' environment variable", envVarName)
	}

	return envVar, nil
}
