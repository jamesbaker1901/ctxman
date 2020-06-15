package cmd

import "os"

func (e *Environment) SetEnv() error {
	file, err := os.Create("/home/jay/.config/ctxman/env")
	if err != nil {
		return err
	}
	defer file.Close()

	for _, v := range e.Env {

		key := "\"" + v.Variable.Key + "\""
		value := "\"" + v.Variable.Value + "\""

		_, err = file.WriteString("export " + key + "=" + value + "\n")
		if err != nil {
			return err
		}
	}

	return nil
}
