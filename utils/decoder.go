package utils

import (
	"fmt"
	"reflect"
	"time"

	"github.com/go-viper/mapstructure/v2"
)

const (
	defaultTagName = "json"
)

type DecoderOption func(*mapstructure.DecoderConfig)

// WithTagName is a decoder option that sets the tag name used for decoding struct fields.
// It takes a string parameter `tagName` and returns a decoderOption function.
// The returned function sets the `TagName` field of the `DecoderConfig` to the provided `tagName` value.
func WithTagName(tagName string) DecoderOption {
	return func(c *mapstructure.DecoderConfig) {
		c.TagName = tagName
	}
}

// NewDecoder creates a new mapstructure decoder with the given output and options.
// The output parameter must be a pointer to the structure where the decoded data will be stored.
// The opts parameter is a variadic argument that allows specifying additional decoder options.
// The function returns a pointer to the created mapstructure.Decoder and an error, if any.
//
// Example usage:
//
//	var output MyStruct
//	decoder, err := NewDecoder(&output, WithTagName("json"))
//	if err != nil {
//	  // handle error
//	}
//	// Use the decoder to decode data into the output structure
//
// Note: The NewDecoder function internally uses the mapstructure package.
// Make sure to import "github.com/go-viper/mapstructure/v2" to use this function.
func NewDecoder(output any, opts ...DecoderOption) (*mapstructure.Decoder, error) {

	if reflect.ValueOf(output).Kind() != reflect.Ptr {
		return nil, fmt.Errorf("output must be a pointer")
	}

	conf := &mapstructure.DecoderConfig{
		TagName: defaultTagName,
		Result:  output,
		DecodeHook: mapstructure.ComposeDecodeHookFunc(
			durationDecoderHook,
		),
	}

	for _, opt := range opts {
		opt(conf)
	}

	return mapstructure.NewDecoder(conf)
}

// durationDecoderHook is a custom decoder hook function that converts a string representation of a duration to a time.Duration value.
// It is used to decode data during the process of converting between different types using reflection.
//
// Parameters:
//   - from: The kind of the source data.
//   - to: The kind of the target data.
//   - data: The data to be decoded.
//
// Returns:
//   - any: The decoded data.
//   - error: An error if the decoding process fails.
func durationDecoderHook(from reflect.Kind, to reflect.Kind, data any) (any, error) {
	if from == reflect.String && to == reflect.TypeOf(time.Duration(0)).Kind() {
		return time.ParseDuration(data.(string))
	}
	return data, nil
}
