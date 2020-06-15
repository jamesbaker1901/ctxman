package cmd

import (
	"fmt"

	"github.com/spf13/viper"
)

type Environment struct {
	Kubernetes struct {
		Cluster   string `yaml:"cluster"`
		Namespace string `yaml:"namespace"`
	} `yaml:"kubernetes"`
	Env []struct {
		Variable struct {
			Key   string `yaml:"key"`
			Value string `yaml:"value"`
		} `yaml:"variable"`
	} `yaml:"env"`
	Openvpn struct {
		ConfigFile string `yaml:"configFile"`
	} `yaml:"openvpn"`
	Shell struct {
		Command string `yaml:"command"`
	} `yaml:"shell"`
	Path struct {
		Include string `yaml:"include"`
		Exclude string `yaml:"exclude"`
	} `yaml:"path"`
	Symlink struct {
		Source string `yaml:"source"`
		Dest   string `yaml:"dest"`
	} `yaml:"symlink"`
}

func Swap(env string, namespace string) error {
	//fmt.Println(viper.Get(env))

	var e Environment
	err := viper.UnmarshalKey(env, &e)
	if err != nil {
		return err
	}

	if len(e.Env) != 0 {
		err = e.SetEnv()
		if err != nil {
			return err
		}
	}

	if e.Kubernetes.Cluster != "" {
		err = e.SetKubernetes(namespace)
		if err != nil {
			return err
		}
	}

	if e.Symlink.Source != "" {
		err = e.SetSymlink()
		if err != nil {
			return err
		}
	}

	if e.Openvpn.ConfigFile != "" {
		err = e.SetOpenvpn()
		if err != nil {
			return err
		}
	}

	if e.Path.Include != "" || e.Path.Exclude != "" {
		err = e.SetPath()
		if err != nil {
			return err
		}
	}

	if e.Shell.Command != "" {
		err = e.ExceShell()
		if err != nil {
			return err
		}
	}

	return nil
}

func (e *Environment) SetPath() error {
	if e.Path.Include != "" {
		fmt.Println("adding to $PATH:", e.Path.Include)
	}

	if e.Path.Exclude != "" {
		fmt.Println("removing from $PATH:", e.Path.Exclude)
	}

	return nil
}

func (e *Environment) SetOpenvpn() error {
	fmt.Println("Starting openvpn connection with profile:", e.Openvpn.ConfigFile)

	return nil
}

func (e *Environment) SetSymlink() error {
	fmt.Println("setting symlink:")
	fmt.Println(" > ln -s", e.Symlink.Source, e.Symlink.Dest)
	return nil
}
