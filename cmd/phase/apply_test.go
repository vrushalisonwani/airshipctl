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

package phase_test

import (
	"testing"

	"opendev.org/airship/airshipctl/cmd/phase"
	"opendev.org/airship/airshipctl/pkg/environment"
	"opendev.org/airship/airshipctl/pkg/k8s/client"
	"opendev.org/airship/airshipctl/pkg/k8s/client/fake"
	"opendev.org/airship/airshipctl/testutil"
)

func TestNewApplyCommand(t *testing.T) {
	fakeRootSettings := &environment.AirshipCTLSettings{
		AirshipConfigPath: "../../testdata/k8s/config.yaml",
		KubeConfigPath:    "../../testdata/k8s/kubeconfig.yaml",
	}
	fakeRootSettings.InitConfig()
	testClientFactory := func(_ *environment.AirshipCTLSettings) (client.Interface, error) {
		return fake.NewClient(), nil
	}

	tests := []*testutil.CmdTest{
		{
			Name:    "phase-apply-cmd-with-help",
			CmdLine: "--help",
			Cmd:     phase.NewApplyCommand(fakeRootSettings, testClientFactory),
		},
	}
	for _, testcase := range tests {
		testutil.RunTest(t, testcase)
	}
}
