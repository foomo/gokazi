package main_test

import (
	"encoding/json"
	"errors"
	"os"
	"path"
	"reflect"
	"strings"
	"testing"

	testingx "github.com/foomo/go/testing"
	tagx "github.com/foomo/go/testing/tag"
	"github.com/foomo/gokazi/pkg/config"
	"github.com/invopop/jsonschema"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestConfig(t *testing.T) {
	testingx.Tags(t, tagx.Short)
	t.Parallel()

	cwd, err := os.Getwd()
	require.NoError(t, err)

	reflector := new(jsonschema.Reflector)
	reflector.RequiredFromJSONSchemaTags = true
	reflector.Namer = func(t reflect.Type) string {
		if t.Name() == "" {
			return t.String()
		}
		return strings.ReplaceAll(t.PkgPath(), "/", ".") + "." + t.Name()
	}
	require.NoError(t, reflector.AddGoComments("github.com/foomo/gokazi", "./"))
	schema := reflector.Reflect(&config.Config{})
	actual, err := json.MarshalIndent(schema, "", "  ")
	require.NoError(t, err)

	filename := path.Join(cwd, "gokazi.schema.json")
	expected, err := os.ReadFile(filename)
	if !errors.Is(err, os.ErrNotExist) {
		require.NoError(t, err)
	}

	if !assert.Equal(t, string(expected), string(actual)) {
		require.NoError(t, os.WriteFile(filename, actual, 0600))
	}
}
