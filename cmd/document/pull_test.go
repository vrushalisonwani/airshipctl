/*
 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at

     https://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
*/

package document

import (
	"testing"

	"github.com/stretchr/testify/require"
	fixtures "gopkg.in/src-d/go-git-fixtures.v3"

	"opendev.org/airship/airshipctl/pkg/config"
	"opendev.org/airship/airshipctl/pkg/environment"

	"opendev.org/airship/airshipctl/testutil"
)

func getDummyAirshipSettings(t *testing.T) *environment.AirshipCTLSettings {
	settings := new(environment.AirshipCTLSettings)
	conf := testutil.DummyConfig()
	mfst := conf.Manifests["dummy_manifest"]

	err := fixtures.Init()
	require.NoError(t, err)

	fx := fixtures.Basic().One()

	mfst.Repositories = map[string]*config.Repository{"primary": {
		URLString: fx.DotGit().Root(),
		CheckoutOptions: &config.RepoCheckout{
			Branch:        "master",
			ForceCheckout: false,
		},
		Auth: &config.RepoAuth{
			Type: "http-basic",
		},
	},
	}
	settings.SetConfig(conf)
	return settings
}

func TestPull(t *testing.T) {
	cmdTests := []*testutil.CmdTest{
		{
			Name:    "document-pull-cmd-with-defaults",
			CmdLine: "",
			Cmd:     NewDocumentPullCommand(getDummyAirshipSettings(t)),
		},
		{
			Name:    "document-pull-cmd-with-help",
			CmdLine: "--help",
			Cmd:     NewDocumentPullCommand(nil),
		},
	}

	for _, tt := range cmdTests {
		testutil.RunTest(t, tt)
	}

	testutil.CleanUpGitFixtures(t)
}
