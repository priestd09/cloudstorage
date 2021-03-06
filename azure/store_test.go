package azure_test

import (
	"flag"
	"os"
	"testing"

	"github.com/araddon/gou"
	"github.com/bmizerany/assert"

	"github.com/lytics/cloudstorage"
	"github.com/lytics/cloudstorage/azure"
	"github.com/lytics/cloudstorage/testutils"
)

/*

# to use azure tests ensure you have exported

export AZURE_KEY="aaa"
export AZURE_PROJECT="bbb"
export AZURE_BUCKET="cloudstorageunittests"

*/
func init() {
	flag.Parse()
	if testing.Verbose() {
		gou.SetupLogging("debug")
		gou.SetColorOutput()
	}
}

var config = &cloudstorage.Config{
	Type:       azure.StoreType,
	AuthMethod: azure.AuthKey,
	Bucket:     os.Getenv("AZURE_BUCKET"),
	TmpDir:     "/tmp/localcache/azure",
	Settings:   make(gou.JsonHelper),
}

func TestConfig(t *testing.T) {
	conf := &cloudstorage.Config{
		Type:     azure.StoreType,
		Settings: make(gou.JsonHelper),
	}
	// Should error with empty config
	_, err := cloudstorage.NewStore(conf)
	assert.NotEqual(t, nil, err)

	conf.AuthMethod = azure.AuthKey
	conf.Settings[azure.ConfKeyAuthKey] = ""
	_, err = cloudstorage.NewStore(conf)
	assert.NotEqual(t, nil, err)
}

func TestAll(t *testing.T) {
	config.Project = os.Getenv("AZURE_PROJECT")
	if config.Project == "" {
		t.Logf("must provide AZURE_PROJECT")
		t.Skip()
		return
	}
	config.Settings[azure.ConfKeyAuthKey] = os.Getenv("AZURE_KEY")
	store, err := cloudstorage.NewStore(config)
	if err != nil {
		t.Logf("No valid auth provided, skipping azure testing %v", err)
		t.Skip()
		return
	}
	testutils.RunTests(t, store)
}
