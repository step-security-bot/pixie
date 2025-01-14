/*
 * Copyright 2018- The Pixie Authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package manifest_test

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/gogo/protobuf/types"
	"github.com/stretchr/testify/require"

	"px.dev/pixie/src/shared/artifacts/manifest"
	"px.dev/pixie/src/shared/artifacts/versionspb"
)

func TestManifest_JSONDecode(t *testing.T) {
	time1 := time.Date(2023, time.March, 23, 0, 24, 33, 10, time.UTC)
	time1Proto, err := types.TimestampProto(time1)
	require.NoError(t, err)
	time1Str, err := (&jsonpb.Marshaler{}).MarshalToString(time1Proto)
	require.NoError(t, err)
	time1Str = strings.Trim(time1Str, `"`)

	testCases := []struct {
		name                 string
		json                 string
		expectedArtifactSets []*versionspb.ArtifactSet
	}{
		{
			name: "single artifact set",
			json: fmt.Sprintf(`
[
  {"name": "vizier", "artifact": [{
    "timestamp": "%s",
    "commitHash": "1234",
    "versionStr": "0.12.17",
    "availableArtifacts": [
      "AT_CONTAINER_SET_LINUX_AMD64",
      "AT_CONTAINER_SET_YAMLS",
      "AT_CONTAINER_SET_TEMPLATE_YAMLS"
    ],
    "changelog": "changelog1"
  }]}
]
`, time1Str),
			expectedArtifactSets: []*versionspb.ArtifactSet{
				{
					Name: "vizier",
					Artifact: []*versionspb.Artifact{
						{
							Timestamp:  time1Proto,
							CommitHash: "1234",
							VersionStr: "0.12.17",
							AvailableArtifacts: []versionspb.ArtifactType{
								versionspb.AT_CONTAINER_SET_LINUX_AMD64,
								versionspb.AT_CONTAINER_SET_YAMLS,
								versionspb.AT_CONTAINER_SET_TEMPLATE_YAMLS,
							},
							Changelog: "changelog1",
						},
					},
				},
			},
		},
		{
			name: "multiple artifact sets",
			json: fmt.Sprintf(`
[
  {"name": "vizier", "artifact": [{
    "timestamp": "%s",
    "commitHash": "1234",
    "versionStr": "0.12.17",
    "availableArtifacts": [
      "AT_CONTAINER_SET_LINUX_AMD64",
      "AT_CONTAINER_SET_YAMLS",
      "AT_CONTAINER_SET_TEMPLATE_YAMLS"
    ],
    "changelog": "changelog1"
  }]},
  {"name": "cli", "artifact": [{
    "timestamp": "%s",
    "commitHash": "abcd",
    "versionStr": "0.1.1",
    "availableArtifacts": [
      "AT_LINUX_AMD64",
      "AT_DARWIN_AMD64"
    ],
    "changelog": "changelog2"
  }]}
]
`, time1Str, time1Str),
			expectedArtifactSets: []*versionspb.ArtifactSet{
				{
					Name: "vizier",
					Artifact: []*versionspb.Artifact{
						{
							Timestamp:  time1Proto,
							CommitHash: "1234",
							VersionStr: "0.12.17",
							AvailableArtifacts: []versionspb.ArtifactType{
								versionspb.AT_CONTAINER_SET_LINUX_AMD64,
								versionspb.AT_CONTAINER_SET_YAMLS,
								versionspb.AT_CONTAINER_SET_TEMPLATE_YAMLS,
							},
							Changelog: "changelog1",
						},
					},
				},
				{
					Name: "cli",
					Artifact: []*versionspb.Artifact{
						{
							Timestamp:  time1Proto,
							CommitHash: "abcd",
							VersionStr: "0.1.1",
							AvailableArtifacts: []versionspb.ArtifactType{
								versionspb.AT_LINUX_AMD64,
								versionspb.AT_DARWIN_AMD64,
							},
							Changelog: "changelog2",
						},
					},
				},
			},
		},
		{
			name: "multiple artifacts sorted",
			json: fmt.Sprintf(`
[
  {"name": "vizier", "artifact": [
    {
      "timestamp": "%s",
      "commitHash": "1234",
      "versionStr": "0.12.9",
      "availableArtifacts": [
        "AT_CONTAINER_SET_LINUX_AMD64",
        "AT_CONTAINER_SET_YAMLS",
        "AT_CONTAINER_SET_TEMPLATE_YAMLS"
      ],
      "changelog": "changelog1"
    },
    {
      "timestamp": "%s",
      "commitHash": "abcd",
      "versionStr": "0.12.10",
      "availableArtifacts": [
        "AT_CONTAINER_SET_LINUX_AMD64",
        "AT_CONTAINER_SET_YAMLS",
        "AT_CONTAINER_SET_TEMPLATE_YAMLS"
      ],
      "changelog": "changelog2"
    }
  ]}
]
`, time1Str, time1Str),
			expectedArtifactSets: []*versionspb.ArtifactSet{
				{
					Name: "vizier",
					Artifact: []*versionspb.Artifact{
						{
							Timestamp:  time1Proto,
							CommitHash: "abcd",
							VersionStr: "0.12.10",
							AvailableArtifacts: []versionspb.ArtifactType{
								versionspb.AT_CONTAINER_SET_LINUX_AMD64,
								versionspb.AT_CONTAINER_SET_YAMLS,
								versionspb.AT_CONTAINER_SET_TEMPLATE_YAMLS,
							},
							Changelog: "changelog2",
						},
						{
							Timestamp:  time1Proto,
							CommitHash: "1234",
							VersionStr: "0.12.9",
							AvailableArtifacts: []versionspb.ArtifactType{
								versionspb.AT_CONTAINER_SET_LINUX_AMD64,
								versionspb.AT_CONTAINER_SET_YAMLS,
								versionspb.AT_CONTAINER_SET_TEMPLATE_YAMLS,
							},
							Changelog: "changelog1",
						},
					},
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			r := strings.NewReader(tc.json)
			m, err := manifest.ReadArtifactManifest(r)
			require.NoError(t, err)
			require.ElementsMatch(t, tc.expectedArtifactSets, m.ArtifactSets())
		})
	}
}

func TestManifest_JSONEncode(t *testing.T) {
	time1 := time.Date(2023, time.March, 23, 0, 24, 33, 10, time.UTC)
	time1Proto, err := types.TimestampProto(time1)
	require.NoError(t, err)
	time1Str, err := (&jsonpb.Marshaler{}).MarshalToString(time1Proto)
	require.NoError(t, err)
	time1Str = strings.Trim(time1Str, `"`)

	testCases := []struct {
		name         string
		expectedJSON string
		artifactSets []*versionspb.ArtifactSet
	}{
		{
			name: "multiple artifact sets",
			expectedJSON: fmt.Sprintf(`
[
  {"name": "cli", "artifact": [{
    "timestamp": "%s",
    "commitHash": "abcd",
    "versionStr": "0.1.1",
    "availableArtifacts": [
      "AT_LINUX_AMD64",
      "AT_DARWIN_AMD64"
    ],
    "changelog": "changelog2"
  }]},
  {"name": "vizier", "artifact": [{
    "timestamp": "%s",
    "commitHash": "1234",
    "versionStr": "0.12.17",
    "availableArtifacts": [
      "AT_CONTAINER_SET_LINUX_AMD64",
      "AT_CONTAINER_SET_YAMLS",
      "AT_CONTAINER_SET_TEMPLATE_YAMLS"
    ],
    "changelog": "changelog1"
  }]}
]
`, time1Str, time1Str),
			artifactSets: []*versionspb.ArtifactSet{
				{
					Name: "vizier",
					Artifact: []*versionspb.Artifact{
						{
							Timestamp:  time1Proto,
							CommitHash: "1234",
							VersionStr: "0.12.17",
							AvailableArtifacts: []versionspb.ArtifactType{
								versionspb.AT_CONTAINER_SET_LINUX_AMD64,
								versionspb.AT_CONTAINER_SET_YAMLS,
								versionspb.AT_CONTAINER_SET_TEMPLATE_YAMLS,
							},
							Changelog: "changelog1",
						},
					},
				},
				{
					Name: "cli",
					Artifact: []*versionspb.Artifact{
						{
							Timestamp:  time1Proto,
							CommitHash: "abcd",
							VersionStr: "0.1.1",
							AvailableArtifacts: []versionspb.ArtifactType{
								versionspb.AT_LINUX_AMD64,
								versionspb.AT_DARWIN_AMD64,
							},
							Changelog: "changelog2",
						},
					},
				},
			},
		},
		{
			name: "multiple artifacts sorted",
			expectedJSON: fmt.Sprintf(`
[
  {"name": "vizier", "artifact": [
    {
      "timestamp": "%s",
      "commitHash": "abcd",
      "versionStr": "0.12.10",
      "availableArtifacts": [
        "AT_CONTAINER_SET_LINUX_AMD64",
        "AT_CONTAINER_SET_YAMLS",
        "AT_CONTAINER_SET_TEMPLATE_YAMLS"
      ],
      "changelog": "changelog2"
    },
    {
      "timestamp": "%s",
      "commitHash": "1234",
      "versionStr": "0.12.9",
      "availableArtifacts": [
        "AT_CONTAINER_SET_LINUX_AMD64",
        "AT_CONTAINER_SET_YAMLS",
        "AT_CONTAINER_SET_TEMPLATE_YAMLS"
      ],
      "changelog": "changelog1"
    }
  ]}
]
`, time1Str, time1Str),
			artifactSets: []*versionspb.ArtifactSet{
				{
					Name: "vizier",
					Artifact: []*versionspb.Artifact{
						{
							Timestamp:  time1Proto,
							CommitHash: "1234",
							VersionStr: "0.12.9",
							AvailableArtifacts: []versionspb.ArtifactType{
								versionspb.AT_CONTAINER_SET_LINUX_AMD64,
								versionspb.AT_CONTAINER_SET_YAMLS,
								versionspb.AT_CONTAINER_SET_TEMPLATE_YAMLS,
							},
							Changelog: "changelog1",
						},
						{
							Timestamp:  time1Proto,
							CommitHash: "abcd",
							VersionStr: "0.12.10",
							AvailableArtifacts: []versionspb.ArtifactType{
								versionspb.AT_CONTAINER_SET_LINUX_AMD64,
								versionspb.AT_CONTAINER_SET_YAMLS,
								versionspb.AT_CONTAINER_SET_TEMPLATE_YAMLS,
							},
							Changelog: "changelog2",
						},
					},
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			raw := &json.RawMessage{}
			err := json.Unmarshal([]byte(tc.expectedJSON), raw)
			require.NoError(t, err)
			bytes, err := json.Marshal(raw)
			require.NoError(t, err)
			expectedJSON := string(bytes)

			m := manifest.NewArtifactManifestFromProto(tc.artifactSets)

			w := &strings.Builder{}
			err = m.Write(w)
			require.NoError(t, err)
			require.Equal(t, expectedJSON, strings.Trim(w.String(), "\n"))
		})
	}
}

func TestManifest_Merge(t *testing.T) {
	time1 := time.Date(2023, time.March, 23, 0, 24, 33, 10, time.UTC)
	time1Proto, err := types.TimestampProto(time1)
	require.NoError(t, err)

	testCases := []struct {
		name                 string
		base                 []*versionspb.ArtifactSet
		updates              []*versionspb.ArtifactSet
		expectedArtifactSets []*versionspb.ArtifactSet
	}{
		{
			name: "simple append only merge",
			base: []*versionspb.ArtifactSet{
				{
					Name: "cli",
					Artifact: []*versionspb.Artifact{
						{
							Timestamp:  time1Proto,
							CommitHash: "abcd",
							VersionStr: "0.1.1",
							AvailableArtifacts: []versionspb.ArtifactType{
								versionspb.AT_LINUX_AMD64,
								versionspb.AT_DARWIN_AMD64,
							},
							Changelog: "changelog2",
						},
					},
				},
			},
			updates: []*versionspb.ArtifactSet{
				{
					Name: "vizier",
					Artifact: []*versionspb.Artifact{
						{
							Timestamp:  time1Proto,
							CommitHash: "1234",
							VersionStr: "0.12.17",
							AvailableArtifacts: []versionspb.ArtifactType{
								versionspb.AT_CONTAINER_SET_LINUX_AMD64,
								versionspb.AT_CONTAINER_SET_YAMLS,
								versionspb.AT_CONTAINER_SET_TEMPLATE_YAMLS,
							},
							Changelog: "changelog1",
						},
					},
				},
				{
					Name: "cli",
					Artifact: []*versionspb.Artifact{
						{
							Timestamp:  time1Proto,
							CommitHash: "efgh",
							VersionStr: "0.1.2",
							AvailableArtifacts: []versionspb.ArtifactType{
								versionspb.AT_LINUX_AMD64,
								versionspb.AT_DARWIN_AMD64,
							},
							Changelog: "changelog3",
						},
					},
				},
			},
			expectedArtifactSets: []*versionspb.ArtifactSet{
				{
					Name: "vizier",
					Artifact: []*versionspb.Artifact{
						{
							Timestamp:  time1Proto,
							CommitHash: "1234",
							VersionStr: "0.12.17",
							AvailableArtifacts: []versionspb.ArtifactType{
								versionspb.AT_CONTAINER_SET_LINUX_AMD64,
								versionspb.AT_CONTAINER_SET_YAMLS,
								versionspb.AT_CONTAINER_SET_TEMPLATE_YAMLS,
							},
							Changelog: "changelog1",
						},
					},
				},
				{
					Name: "cli",
					Artifact: []*versionspb.Artifact{
						{
							Timestamp:  time1Proto,
							CommitHash: "efgh",
							VersionStr: "0.1.2",
							AvailableArtifacts: []versionspb.ArtifactType{
								versionspb.AT_LINUX_AMD64,
								versionspb.AT_DARWIN_AMD64,
							},
							Changelog: "changelog3",
						},
						{
							Timestamp:  time1Proto,
							CommitHash: "abcd",
							VersionStr: "0.1.1",
							AvailableArtifacts: []versionspb.ArtifactType{
								versionspb.AT_LINUX_AMD64,
								versionspb.AT_DARWIN_AMD64,
							},
							Changelog: "changelog2",
						},
					},
				},
			},
		},
		{
			name: "update artifacts with the same version",
			base: []*versionspb.ArtifactSet{
				{
					Name: "cli",
					Artifact: []*versionspb.Artifact{
						{
							Timestamp:  time1Proto,
							CommitHash: "abcd",
							VersionStr: "0.1.1",
							AvailableArtifacts: []versionspb.ArtifactType{
								versionspb.AT_LINUX_AMD64,
								versionspb.AT_DARWIN_AMD64,
							},
							Changelog: "changelog2",
						},
					},
				},
			},
			updates: []*versionspb.ArtifactSet{
				{
					Name: "cli",
					Artifact: []*versionspb.Artifact{
						{
							Timestamp:  time1Proto,
							CommitHash: "efgh",
							VersionStr: "0.1.1",
							AvailableArtifacts: []versionspb.ArtifactType{
								versionspb.AT_CONTAINER_SET_YAMLS,
								versionspb.AT_LINUX_AMD64,
								versionspb.AT_DARWIN_AMD64,
							},
							Changelog: "changelog3",
						},
					},
				},
			},
			expectedArtifactSets: []*versionspb.ArtifactSet{
				{
					Name: "cli",
					Artifact: []*versionspb.Artifact{
						{
							Timestamp:  time1Proto,
							CommitHash: "efgh",
							VersionStr: "0.1.1",
							AvailableArtifacts: []versionspb.ArtifactType{
								versionspb.AT_CONTAINER_SET_YAMLS,
								versionspb.AT_LINUX_AMD64,
								versionspb.AT_DARWIN_AMD64,
							},
							Changelog: "changelog3",
						},
					},
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			base := manifest.NewArtifactManifestFromProto(tc.base)
			updates := manifest.NewArtifactManifestFromProto(tc.updates)
			merged := base.Merge(updates)

			require.ElementsMatch(t, tc.expectedArtifactSets, merged.ArtifactSets())
		})
	}
}

func TestManifest_GetArtifact(t *testing.T) {
	testCases := []struct {
		name             string
		artifactSets     []*versionspb.ArtifactSet
		artifactName     string
		artifactVersion  string
		expectedErr      error
		expectedArtifact *versionspb.Artifact
	}{
		{
			name: "basic",
			artifactSets: []*versionspb.ArtifactSet{
				{
					Name: "cli",
					Artifact: []*versionspb.Artifact{
						{
							CommitHash: "abcd",
							VersionStr: "1.0.0",
						},
						{
							CommitHash: "efgh",
							VersionStr: "0.1.2",
						},
						{
							CommitHash: "1234",
							VersionStr: "0.1.1",
						},
					},
				},
			},
			artifactName:    "cli",
			artifactVersion: "0.1.2",
			expectedErr:     nil,
			expectedArtifact: &versionspb.Artifact{
				CommitHash: "efgh",
				VersionStr: "0.1.2",
			},
		},
		{
			name: "prerelease versions",
			artifactSets: []*versionspb.ArtifactSet{
				{
					Name: "cli",
					Artifact: []*versionspb.Artifact{
						{
							CommitHash: "abcd",
							VersionStr: "1.0.0",
						},
						{
							CommitHash: "1234",
							VersionStr: "0.1.1-prerelease-1234",
						},
						{
							CommitHash: "efgh",
							VersionStr: "0.1.1-prerelease-efgh",
						},
					},
				},
			},
			artifactName:    "cli",
			artifactVersion: "0.1.1-prerelease-efgh",
			expectedErr:     nil,
			expectedArtifact: &versionspb.Artifact{
				CommitHash: "efgh",
				VersionStr: "0.1.1-prerelease-efgh",
			},
		},
		{
			name: "error cant find artifact set",
			artifactSets: []*versionspb.ArtifactSet{
				{
					Name:     "cli",
					Artifact: []*versionspb.Artifact{},
				},
			},
			artifactName:    "vizier",
			artifactVersion: "0.1.1",
			expectedErr:     manifest.ErrArtifactSetNotFound,
		},
		{
			name: "error cant find artifact version",
			artifactSets: []*versionspb.ArtifactSet{
				{
					Name: "cli",
					Artifact: []*versionspb.Artifact{
						{
							CommitHash: "abcd",
							VersionStr: "1.0.0",
						},
						{
							CommitHash: "1234",
							VersionStr: "0.1.1-prerelease.2",
						},
						{
							CommitHash: "efgh",
							VersionStr: "0.1.1-prerelease.1",
						},
					},
				},
			},
			artifactName:    "cli",
			artifactVersion: "0.1.1-prerelease.3",
			expectedErr:     manifest.ErrArtifactNotFound,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			m := manifest.NewArtifactManifestFromProto(tc.artifactSets)

			a, err := m.GetArtifact(tc.artifactName, tc.artifactVersion)
			if tc.expectedErr != nil {
				require.Equal(t, tc.expectedErr, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expectedArtifact, a)
			}
		})
	}
}

func TestManifest_ListArtifacts(t *testing.T) {
	testCases := []struct {
		name              string
		artifactSets      []*versionspb.ArtifactSet
		artifactName      string
		limit             int64
		filters           []manifest.ArtifactFilter
		expectedErr       error
		expectedArtifacts []*versionspb.Artifact
	}{
		{
			name: "basic",
			artifactSets: []*versionspb.ArtifactSet{
				{
					Name: "cli",
					Artifact: []*versionspb.Artifact{
						{
							CommitHash: "abcd",
							VersionStr: "1.0.0",
						},
						{
							CommitHash: "efgh",
							VersionStr: "0.1.2",
						},
						{
							CommitHash: "1234",
							VersionStr: "0.1.1",
						},
					},
				},
			},
			artifactName: "cli",
			expectedErr:  nil,
			expectedArtifacts: []*versionspb.Artifact{
				{
					CommitHash: "abcd",
					VersionStr: "1.0.0",
				},
				{
					CommitHash: "efgh",
					VersionStr: "0.1.2",
				},
				{
					CommitHash: "1234",
					VersionStr: "0.1.1",
				},
			},
		},
		{
			name: "filter out prerelease versions",
			artifactSets: []*versionspb.ArtifactSet{
				{
					Name: "cli",
					Artifact: []*versionspb.Artifact{
						{
							CommitHash: "abcd",
							VersionStr: "1.0.0",
						},
						{
							CommitHash: "1111",
							VersionStr: "0.1.1",
						},
						{
							CommitHash: "1234",
							VersionStr: "0.1.1-prerelease.1234",
						},
						{
							CommitHash: "efgh",
							VersionStr: "0.1.1-pre-main-efgh",
						},
					},
				},
			},
			artifactName: "cli",
			filters: []manifest.ArtifactFilter{
				manifest.RemovePrereleasesFilter(),
			},
			expectedErr: nil,
			expectedArtifacts: []*versionspb.Artifact{
				{
					CommitHash: "abcd",
					VersionStr: "1.0.0",
				},
				{
					CommitHash: "1111",
					VersionStr: "0.1.1",
				},
			},
		},
		{
			name: "filter out prerelease versions with limit",
			artifactSets: []*versionspb.ArtifactSet{
				{
					Name: "cli",
					Artifact: []*versionspb.Artifact{
						{
							CommitHash: "abcd",
							VersionStr: "1.0.0",
						},
						{
							CommitHash: "1111",
							VersionStr: "0.1.1",
						},
						{
							CommitHash: "1234",
							VersionStr: "0.1.1-prerelease.1234",
						},
						{
							CommitHash: "efgh",
							VersionStr: "0.1.1-pre-main-efgh",
						},
						{
							CommitHash: "0000",
							VersionStr: "0.1.0",
						},
					},
				},
			},
			artifactName: "cli",
			filters: []manifest.ArtifactFilter{
				manifest.RemovePrereleasesFilter(),
			},
			limit:       2,
			expectedErr: nil,
			expectedArtifacts: []*versionspb.Artifact{
				{
					CommitHash: "abcd",
					VersionStr: "1.0.0",
				},
				{
					CommitHash: "1111",
					VersionStr: "0.1.1",
				},
			},
		},
		{
			name: "error cant find artifact set",
			artifactSets: []*versionspb.ArtifactSet{
				{
					Name:     "cli",
					Artifact: []*versionspb.Artifact{},
				},
			},
			artifactName: "vizier",
			expectedErr:  manifest.ErrArtifactSetNotFound,
		},
		{
			name: "filter for available artifacts",
			artifactSets: []*versionspb.ArtifactSet{
				{
					Name: "vizier",
					Artifact: []*versionspb.Artifact{
						{
							CommitHash: "abcd",
							VersionStr: "1.0.0",
							AvailableArtifacts: []versionspb.ArtifactType{
								versionspb.AT_CONTAINER_SET_LINUX_AMD64,
							},
						},
						{
							CommitHash: "1111",
							VersionStr: "0.1.1",
							AvailableArtifacts: []versionspb.ArtifactType{
								versionspb.AT_CONTAINER_SET_YAMLS,
							},
						},
						{
							CommitHash: "1234",
							VersionStr: "0.1.1-prerelease.2",
							AvailableArtifacts: []versionspb.ArtifactType{
								versionspb.AT_CONTAINER_SET_YAMLS,
							},
						},
						{
							CommitHash: "efgh",
							VersionStr: "0.1.1-prerelease.1",
							AvailableArtifacts: []versionspb.ArtifactType{
								versionspb.AT_CONTAINER_SET_YAMLS,
							},
						},
					},
				},
			},
			artifactName: "vizier",
			filters: []manifest.ArtifactFilter{
				manifest.ArtifactTypeFilter(versionspb.AT_CONTAINER_SET_YAMLS),
			},
			expectedErr: nil,
			expectedArtifacts: []*versionspb.Artifact{
				{
					CommitHash: "1111",
					VersionStr: "0.1.1",
					AvailableArtifacts: []versionspb.ArtifactType{
						versionspb.AT_CONTAINER_SET_YAMLS,
					},
				},
				{
					CommitHash: "1234",
					VersionStr: "0.1.1-prerelease.2",
					AvailableArtifacts: []versionspb.ArtifactType{
						versionspb.AT_CONTAINER_SET_YAMLS,
					},
				},
				{
					CommitHash: "efgh",
					VersionStr: "0.1.1-prerelease.1",
					AvailableArtifacts: []versionspb.ArtifactType{
						versionspb.AT_CONTAINER_SET_YAMLS,
					},
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			m := manifest.NewArtifactManifestFromProto(tc.artifactSets)

			l, err := m.ListArtifacts(tc.artifactName, tc.limit, tc.filters...)
			if tc.expectedErr != nil {
				require.Equal(t, tc.expectedErr, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expectedArtifacts, l)
			}
		})
	}
}
