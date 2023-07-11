package config

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/pelletier/go-toml/v2"
	"github.com/spf13/viper"
	"github.com/yb7/alilog"
	"reflect"
	"strings"
)

var C = Configuration{}

// Configuration Config example config
type Configuration struct {
	Ports *PortsConfig `validate:"required"`
	Db    *DbConfig    `validate:"required"`
}

type PortsConfig struct {
	Http string
}
type DbConfig struct {
	Url string
}

func init() {
	// viper.AutomaticEnv()
	// viper.SetEnvPrefix("C")
	// viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	// viper.BindEnv("redis.password")
	viper.SetConfigName("config")                            // name of config file (without extension)
	viper.SetConfigType("toml")                              // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("/etc/gangkou56.com/truck-on-time/") // path to look for the config file in
	viper.AddConfigPath("$HOME/.tsp")                        // call multiple times to add many search paths
	viper.AddConfigPath(".")                                 // optionally look for config in the working directory

	bindAllKeys(reflect.TypeOf(C))

	// for _, k := range viper.AllKeys() {
	// 	alilog.Infof("key >> %s", k)
	// }

	// alilog.Warnf("REDIS.PASSWORD >> [%s]", os.Getenv("REDIS.PASSWORD"))

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		alilog.Warnf("\nFatal error config file: %v", err)
	}
	if err := viper.Unmarshal(&C); err != nil {
		alilog.Fatalf("Fatal error marshal config file: %w", err)
	}

	validate := validator.New()
	if err := validate.Struct(C); err != nil {
		printValidateError(err)
	}
	b, err := toml.Marshal(C)
	if err != nil {
		alilog.Fatal(err)
	}

	alilog.Infof("config as following\n%s", string(b))
}

// https://github.com/spf13/viper/issues/584
func bindAllKeys(t reflect.Type) {
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		walkThroughFields(envKeyName(f), t.Field(i).Type)
	}
}
func envKeyName(f reflect.StructField) string {
	envKey, ok := f.Tag.Lookup("env")
	if ok {
		return envKey
	}
	return strings.ToLower(f.Name[0:1]) + f.Name[1:] //lowerFirstChar(f.Name)
}
func walkThroughFields(fieldKey string, t reflect.Type) {
	if reflect.Ptr == t.Kind() {
		walkThroughFields(fieldKey, t.Elem())
	} else if reflect.Struct == t.Kind() {
		for i := 0; i < t.NumField(); i++ {
			f := t.Field(i)
			walkThroughFields(fieldKey+"."+envKeyName(f), f.Type)
		}
	} else {
		// alilog.Infof("bind key[%s]", fieldKey)
		alterKey := strings.ReplaceAll(fieldKey, ".", "_")
		upperAlterKey := strings.ToUpper(alterKey)
		// alilog.Infof("bind >> %s, %s, %s, %s", fieldKey, fieldKey, alterKey, upperAlterKey)
		viper.BindEnv(fieldKey, fieldKey, alterKey, upperAlterKey)
	}
}

func printValidateError(err error) {

	// this check is only needed when your code could produce
	// an invalid value for validation such as interface with nil
	// value most including myself do not usually have code like this.
	if _, ok := err.(*validator.InvalidValidationError); ok {
		alilog.Fatal(err)
	}
	var builder strings.Builder
	builder.WriteString("validate errors:\n")
	for _, err := range err.(validator.ValidationErrors) {
		builder.WriteString("  ")
		builder.WriteString(err.Namespace())
		builder.WriteString(" ")
		builder.WriteString(err.Field())
		builder.WriteString(" ")
		builder.WriteString(err.StructNamespace())
		builder.WriteString(" ")
		builder.WriteString(err.StructField())
		builder.WriteString(" ")
		builder.WriteString(err.Tag())
		builder.WriteString(" ")
		builder.WriteString(err.ActualTag())
		builder.WriteString(" ")
		builder.WriteString(err.Kind().String())
		builder.WriteString(" ")
		builder.WriteString(err.Type().String())
		builder.WriteString(" ")
		builder.WriteString(fmt.Sprintf("%v", err.Value()))
		builder.WriteString(" ")
		builder.WriteString(err.Param())
		builder.WriteString("\n")
	}
	alilog.Fatalf(builder.String())
}
