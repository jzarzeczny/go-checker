package dotEnvParser

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ParseEnvVariables() (map[string]string, error) {
	envVariables := make(map[string]string)
	file, err := os.Open(".env")

	if err != nil {
		fmt.Println("Error during opening .env file: ", err)

		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)

		parts := strings.SplitN(line, "=", 2)

		if len(parts) != 2 {
			fmt.Println("Invalid line format: ", line)
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])
		envVariables[key] = value
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading .env file", err)
		return nil, err
	}

	return envVariables, nil
}
