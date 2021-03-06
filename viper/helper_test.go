package viper

import (
	"fmt"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/spf13/viper"
	)

func setEnv(key, value string) func(t *testing.T) {
	return func(t *testing.T) {
		require.NoError(t, os.Setenv(key, value))
	}
}

func setViper(key, value string) func(t *testing.T) {
	return func(t *testing.T) {
		viper.Set(key, value)
	}
}

func noop(t *testing.T) {
}

func reset() {
	viper.Reset()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
}

func TestGetStringSlice(t *testing.T) {
	for k, tc := range []struct {
		d string
		p func(t *testing.T)
		e []string
		f []string
		c func(t *testing.T)
	}{
		{
			d: "Load legacy environment variable in legacy format",
			p: setEnv("VIPERX_GET_STRING_SLICE_LEGACY_LEGACY", "foo,bar"),
			c: setEnv("VIPERX_GET_STRING_SLICE_LEGACY_LEGACY", ""),
			e: []string{"foo", "bar"},
		},
		{
			d: "Load legacy environment variable in new format",
			p: setEnv("VIPERX_GET_STRING_SLICE_LEGACY_LEGACY", "foo, bar"),
			c: setEnv("VIPERX_GET_STRING_SLICE_LEGACY_LEGACY", ""),
			e: []string{"foo", "bar"},
		},
		{
			d: "Load legacy environment variable in legacy format",
			p: setEnv("VIPERX_GET_STRING_SLICE_LEGACY", "foo,bar"),
			c: setEnv("VIPERX_GET_STRING_SLICE_LEGACY", ""),
			e: []string{"foo", "bar"},
		},
		{
			d: "Load legacy environment variable in new format",
			p: setEnv("VIPERX_GET_STRING_SLICE_LEGACY", "foo, bar"),
			c: setEnv("VIPERX_GET_STRING_SLICE_LEGACY", ""),
			e: []string{"foo", "bar"},
		},
		{
			d: "Load new environment variable in legacy format",
			p: setEnv("VIPERX_GET_STRING_SLICE", "foo,bar"),
			c: setEnv("VIPERX_GET_STRING_SLICE", ""),
			e: []string{"foo", "bar"},
		},
		{
			d: "Load new environment variable in new format",
			p: setEnv("VIPERX_GET_STRING_SLICE", "foo, bar"),
			c: setEnv("VIPERX_GET_STRING_SLICE", ""),
			e: []string{"foo", "bar"},
		},
		{
			d: "Load new environment variable in legacy format",
			p: setViper("viper.get_string_slice", "foo,bar"),
			c: setViper("viper.get_string_slice", ""),
			e: []string{"foo", "bar"},
		},
		{
			d: "Load new environment variable in new format",
			p: setViper("viper.get_string_slice", "foo, bar"),
			c: setViper("viper.get_string_slice", ""),
			e: []string{"foo", "bar"},
		},
		{
			d: "Use fallback",
			p: noop,
			c: noop,
			f: []string{"foo", "bar"},
			e: []string{"foo", "bar"},
		},
	} {
		t.Run(fmt.Sprintf("case=%d/description=%s", k, tc.d), func(t *testing.T) {
			reset()

			tc.p(t)
			assert.EqualValues(t, tc.e, GetStringSlice("viper.get_string_slice", tc.f, "viper.get_string_slice_legacy", "viper.get_string_slice_legacy_legacy"))
			tc.c(t)
		})
	}
}

func TestGetString(t *testing.T) {
	for k, tc := range []struct {
		d string
		p func(t *testing.T)
		e string
		f string
		c func(t *testing.T)
	}{
		{
			d: "Load legacy environment variable",
			p: setEnv("VIPERX_GET_STRING_LEGACY_LEGACY", "foo"),
			c: setEnv("VIPERX_GET_STRING_LEGACY_LEGACY", ""),
			e: "foo",
		},
		{
			d: "Load legacy environment",
			p: setEnv("VIPERX_GET_STRING_LEGACY", "foo"),
			c: setEnv("VIPERX_GET_STRING_LEGACY", ""),
			e: "foo",
		},
		{
			d: "Load new environment variable",
			p: setEnv("VIPERX_GET_STRING", "foo"),
			c: setEnv("VIPERX_GET_STRING", ""),
			e: "foo",
		},
		{
			d: "Load new viper variable",
			p: setViper("viper.get_string", "foo"),
			c: setViper("viper.get_string", ""),
			e: "foo",
		},
		{
			d: "Use fallback",
			p: noop,
			c: noop,
			f: "foo",
			e: "foo",
		},
	} {
		t.Run(fmt.Sprintf("case=%d/description=%s", k, tc.d), func(t *testing.T) {
			reset()

			tc.p(t)
			assert.EqualValues(t, tc.e, GetString("viper.get_string", tc.f, "viper.get_string_legacy", "viper.get_string_legacy_legacy"))
			tc.c(t)
		})
	}
}

func TestGetInt(t *testing.T) {
	for k, tc := range []struct {
		d string
		p func(t *testing.T)
		e int
		f int
		c func(t *testing.T)
	}{
		{
			d: "Load legacy environment variable",
			p: setEnv("VIPERX_GET_INT_LEGACY_LEGACY", "1"),
			c: setEnv("VIPERX_GET_INT_LEGACY_LEGACY", ""),
			e: 1,
		},
		{
			d: "Load legacy environment",
			p: setEnv("VIPERX_GET_INT_LEGACY", "1"),
			c: setEnv("VIPERX_GET_INT_LEGACY", ""),
			e: 1,
		},
		{
			d: "Load new environment variable",
			p: setEnv("VIPERX_GET_INT", "1"),
			c: setEnv("VIPERX_GET_INT", ""),
			e: 1,
		},
		{
			d: "Load new viper variable",
			p: setViper("viper.get_int", "1"),
			c: setViper("viper.get_int", ""),
			e: 1,
		},
		{
			d: "Use fallback",
			p: noop,
			c: noop,
			f: 1,
			e: 1,
		},
	} {
		t.Run(fmt.Sprintf("case=%d/description=%s", k, tc.d), func(t *testing.T) {
			reset()

			tc.p(t)
			assert.EqualValues(t, tc.e, GetInt("viper.get_int", tc.f, "viper.get_int_legacy", "viper.get_int_legacy_legacy"))
			tc.c(t)
		})
	}
}

func TestGetFloat64(t *testing.T) {
	for k, tc := range []struct {
		d string
		p func(t *testing.T)
		e float64
		f float64
		c func(t *testing.T)
	}{
		{
			d: "Load legacy environment variable",
			p: setEnv("VIPERX_GET_FLOAT64_LEGACY_LEGACY", "1"),
			c: setEnv("VIPERX_GET_FLOAT64_LEGACY_LEGACY", ""),
			e: 1,
		},
		{
			d: "Load legacy environment",
			p: setEnv("VIPERX_GET_FLOAT64_LEGACY", "1"),
			c: setEnv("VIPERX_GET_FLOAT64_LEGACY", ""),
			e: 1,
		},
		{
			d: "Load new environment variable",
			p: setEnv("VIPERX_GET_FLOAT64", "1"),
			c: setEnv("VIPERX_GET_FLOAT64", ""),
			e: 1,
		},
		{
			d: "Load new viper variable",
			p: setViper("viper.get_float64", "1"),
			c: setViper("viper.get_float64", ""),
			e: 1,
		},
		{
			d: "Use fallback",
			p: noop,
			c: noop,
			f: 1,
			e: 1,
		},
	} {
		t.Run(fmt.Sprintf("case=%d/description=%s", k, tc.d), func(t *testing.T) {
			reset()

			tc.p(t)
			assert.EqualValues(t, tc.e, GetFloat64("viper.get_float64", tc.f, "viper.get_float64_legacy", "viper.get_float64_legacy_legacy"))
			tc.c(t)
		})
	}
}

func TestGetBool(t *testing.T) {
	for k, tc := range []struct {
		d string
		p func(t *testing.T)
		e bool
		f bool
		c func(t *testing.T)
	}{
		{
			d: "Use fallback",
			p: noop,
			c: noop,
			e: false,
			f: false,
		},
		{
			d: "Use fallback",
			p: noop,
			c: noop,
			e: true,
			f: true,
		},
		{
			d: "Load legacy environment variable",
			p: setEnv("VIPERX_GET_BOOL_LEGACY_LEGACY", "1"),
			c: setEnv("VIPERX_GET_BOOL_LEGACY_LEGACY", ""),
			e: true,
		},
		{
			d: "Load legacy environment",
			p: setEnv("VIPERX_GET_BOOL_LEGACY", "1"),
			c: setEnv("VIPERX_GET_BOOL_LEGACY", ""),
			e: true,
		},
		{
			d: "Load new environment variable",
			p: setEnv("VIPERX_GET_BOOL", "1"),
			c: setEnv("VIPERX_GET_BOOL", ""),
			e: true,
		},
		{
			d: "Load new viper variable",
			p: setViper("viper.get_bool", "1"),
			c: setViper("viper.get_bool", ""),
			e: true,
		},
	} {
		t.Run(fmt.Sprintf("case=%d/description=%s", k, tc.d), func(t *testing.T) {
			reset()

			tc.p(t)
			assert.EqualValues(t, tc.e, GetBool("viper.get_bool", tc.f, "viper.get_bool_legacy", "viper.get_bool_legacy_legacy"))
			tc.c(t)
		})
	}
}

func TestGetDuration(t *testing.T) {
	for k, tc := range []struct {
		d string
		p func(t *testing.T)
		e time.Duration
		f time.Duration
		c func(t *testing.T)
	}{
		{
			d: "Load legacy environment variable",
			p: setEnv("VIPERX_GET_DURATION_LEGACY_LEGACY", "1m"),
			c: setEnv("VIPERX_GET_DURATION_LEGACY_LEGACY", ""),
			e: time.Minute,
		},
		{
			d: "Load legacy environment",
			p: setEnv("VIPERX_GET_DURATION_LEGACY", "1m"),
			c: setEnv("VIPERX_GET_DURATION_LEGACY", ""),
			e: time.Minute,
		},
		{
			d: "Load new environment variable",
			p: setEnv("VIPERX_GET_DURATION", "1m"),
			c: setEnv("VIPERX_GET_DURATION", ""),
			e: time.Minute,
		},
		{
			d: "Load new viper variable",
			p: setViper("viper.get_duration", "1m"),
			c: setViper("viper.get_duration", ""),
			e: time.Minute,
		},
		{
			d: "Use fallback",
			p: noop,
			c: noop,
			f: 1,
			e: 1,
		},
	} {
		t.Run(fmt.Sprintf("case=%d/description=%s", k, tc.d), func(t *testing.T) {
			reset()

			tc.p(t)
			assert.EqualValues(t, tc.e, GetDuration("viper.get_duration", tc.f, "viper.get_duration_legacy", "viper.get_duration_legacy_legacy"))
			tc.c(t)
		})
	}
}

func TestGetStringMapConfig(t *testing.T) {

	t.Run("suite=with-config-path", func(t *testing.T) {
		for k, tc := range []struct {
			f      string
			fatals bool
		}{
			{f: "./stub/json/.project-stub-name.json"},
			{f: "./stub/toml/.project-stub-name.toml"},
			{f: "./stub/yaml/.project-stub-name.yaml"},
			{f: "./stub/yml/.project-stub-name.yml"},
			{f: "./stub/does-not-exist/foo.yml", fatals: true},
		} {
			t.Run(fmt.Sprintf("case=%d/path=%s", k, tc.f), func(t *testing.T) {
				viper.Reset()

				cfgFile = tc.f
				if tc.fatals {
					assert.Panics(t, func() {
						InitializeConfig("project-stub-name", "")
					})
				} else {
					InitializeConfig("project-stub-name", "")
					config := GetStringMapConfig("authenticators", "oauth2_introspection", "config")

					assert.Equal(t, "http://myurl", config["introspection_url"])
				}

				cfgFile = ""
			})
		}
	})

	t.Run("suite=with-env-var", func(t *testing.T) {
		for k, tc := range []struct {
			f      string
			fatals bool
		}{
			{f: "./stub/json/.project-stub-name.json"},
			{f: "./stub/toml/.project-stub-name.toml"},
			{f: "./stub/yaml/.project-stub-name.yaml"},
			{f: "./stub/yml/.project-stub-name.yml"},
			{f: "./stub/does-not-exist/foo.yml", fatals: true},
		} {
			t.Run(fmt.Sprintf("case=%d/path=%s", k, tc.f), func(t *testing.T) {
				viper.Reset()

				os.Setenv("AUTHENTICATORS_OAUTH2_INTROSPECTION_CONFIG_INTROSPECTION_URL", "http://envurl")

				cfgFile = tc.f
				if tc.fatals {
					assert.Panics(t, func() {
						InitializeConfig("project-stub-name", "")
					})
				} else {
					InitializeConfig("project-stub-name", "")
					config := GetStringMapConfig("authenticators", "oauth2_introspection", "config")

					assert.Equal(t, "http://envurl", config["introspection_url"])
				}

				cfgFile = ""
			})
		}
	})
}
