package autocomplete_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"

	"pixielabs.ai/pixielabs/src/cloud/autocomplete"
	"pixielabs.ai/pixielabs/src/cloud/autocomplete/mock"
	"pixielabs.ai/pixielabs/src/cloud/cloudapipb"
)

var orgID uuid.UUID = uuid.NewV4()

func TestParseIntoCommand(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		requests    [][]*autocomplete.SuggestionRequest
		responses   [][]*autocomplete.SuggestionResult
		expectedCmd *autocomplete.Command
	}{
		{
			name:  "valid",
			input: "script:px/svc_info svc:pl/test",
			requests: [][]*autocomplete.SuggestionRequest{
				[]*autocomplete.SuggestionRequest{
					&autocomplete.SuggestionRequest{
						OrgID:        orgID,
						Input:        "px/svc_info",
						AllowedKinds: []cloudapipb.AutocompleteEntityKind{cloudapipb.AEK_SCRIPT},
						AllowedArgs:  []cloudapipb.AutocompleteEntityKind{},
					},
				},
				[]*autocomplete.SuggestionRequest{
					&autocomplete.SuggestionRequest{
						OrgID:        orgID,
						Input:        "pl/test",
						AllowedKinds: []cloudapipb.AutocompleteEntityKind{cloudapipb.AEK_SVC},
						AllowedArgs:  []cloudapipb.AutocompleteEntityKind{cloudapipb.AEK_SVC},
					},
				},
			},
			responses: [][]*autocomplete.SuggestionResult{
				[]*autocomplete.SuggestionResult{
					&autocomplete.SuggestionResult{
						Suggestions: []*autocomplete.Suggestion{
							&autocomplete.Suggestion{
								Name:  "px/svc_info",
								Score: 1,
								Args:  []cloudapipb.AutocompleteEntityKind{cloudapipb.AEK_SVC},
							},
						},
						ExactMatch: true,
					},
				},
				[]*autocomplete.SuggestionResult{
					&autocomplete.SuggestionResult{
						Suggestions: []*autocomplete.Suggestion{
							&autocomplete.Suggestion{
								Name:  "px/test",
								Score: 1,
							},
						},
						ExactMatch: true,
					},
				},
			},
			expectedCmd: &autocomplete.Command{
				TabStops: []*autocomplete.TabStop{
					&autocomplete.TabStop{
						Value: "run",
						Kind:  cloudapipb.AEK_UNKNOWN,
						Valid: true,
					},
					&autocomplete.TabStop{
						Value: "px/svc_info",
						Kind:  cloudapipb.AEK_SCRIPT,
						Valid: true,
					},
					&autocomplete.TabStop{
						Value: "pl/test",
						Kind:  cloudapipb.AEK_SVC,
						Valid: true,
					},
				},
				Executable: true,
			},
		},
		{
			name:  "invalid",
			input: "script:px/svc_info pod:pl/test",
			requests: [][]*autocomplete.SuggestionRequest{
				[]*autocomplete.SuggestionRequest{
					&autocomplete.SuggestionRequest{
						OrgID:        orgID,
						Input:        "px/svc_info",
						AllowedKinds: []cloudapipb.AutocompleteEntityKind{cloudapipb.AEK_SCRIPT},
						AllowedArgs:  []cloudapipb.AutocompleteEntityKind{},
					},
				},
				[]*autocomplete.SuggestionRequest{
					&autocomplete.SuggestionRequest{
						OrgID:        orgID,
						Input:        "pod:pl/test",
						AllowedKinds: []cloudapipb.AutocompleteEntityKind{cloudapipb.AEK_SVC},
						AllowedArgs:  []cloudapipb.AutocompleteEntityKind{},
					},
				},
			},
			responses: [][]*autocomplete.SuggestionResult{
				[]*autocomplete.SuggestionResult{
					&autocomplete.SuggestionResult{
						Suggestions: []*autocomplete.Suggestion{
							&autocomplete.Suggestion{
								Name:  "px/svc_info",
								Score: 1,
								Args:  []cloudapipb.AutocompleteEntityKind{cloudapipb.AEK_SVC},
							},
						},
						ExactMatch: true,
					},
				},
				[]*autocomplete.SuggestionResult{
					&autocomplete.SuggestionResult{
						Suggestions: []*autocomplete.Suggestion{
							&autocomplete.Suggestion{
								Name:  "pl/test",
								Score: 1,
							},
						},
						ExactMatch: true,
					},
				},
			},
			expectedCmd: &autocomplete.Command{
				TabStops: []*autocomplete.TabStop{
					&autocomplete.TabStop{
						Value: "run",
						Kind:  cloudapipb.AEK_UNKNOWN,
						Valid: true,
					},
					&autocomplete.TabStop{
						Value: "px/svc_info",
						Kind:  cloudapipb.AEK_SCRIPT,
						Valid: true,
					},
					&autocomplete.TabStop{
						Value: "pod:pl/test",
						Kind:  cloudapipb.AEK_UNKNOWN,
						Valid: false,
					},
				},
				Executable: false,
			},
		},
		{
			name:  "no script, defined entity",
			input: "px/svc_info pod:pl/test",
			requests: [][]*autocomplete.SuggestionRequest{
				[]*autocomplete.SuggestionRequest{
					&autocomplete.SuggestionRequest{
						OrgID:        orgID,
						Input:        "px/svc_info",
						AllowedKinds: []cloudapipb.AutocompleteEntityKind{cloudapipb.AEK_POD, cloudapipb.AEK_SVC, cloudapipb.AEK_NAMESPACE, cloudapipb.AEK_SCRIPT},
						AllowedArgs:  []cloudapipb.AutocompleteEntityKind{cloudapipb.AEK_POD},
					},
					&autocomplete.SuggestionRequest{
						OrgID:        orgID,
						Input:        "pl/test",
						AllowedKinds: []cloudapipb.AutocompleteEntityKind{cloudapipb.AEK_POD},
						AllowedArgs:  []cloudapipb.AutocompleteEntityKind{cloudapipb.AEK_POD},
					},
				},
			},
			responses: [][]*autocomplete.SuggestionResult{
				[]*autocomplete.SuggestionResult{
					&autocomplete.SuggestionResult{
						Suggestions: []*autocomplete.Suggestion{
							&autocomplete.Suggestion{
								Name:  "px/svc_info",
								Score: 1,
								Args:  []cloudapipb.AutocompleteEntityKind{cloudapipb.AEK_SVC},
							},
						},
						ExactMatch: false,
					},
					&autocomplete.SuggestionResult{
						Suggestions: []*autocomplete.Suggestion{
							&autocomplete.Suggestion{
								Name:  "pl/test",
								Score: 1,
							},
						},
						ExactMatch: true,
					},
				},
			},
			expectedCmd: &autocomplete.Command{
				TabStops: []*autocomplete.TabStop{
					&autocomplete.TabStop{
						Value: "run",
						Kind:  cloudapipb.AEK_UNKNOWN,
						Valid: true,
					},
					&autocomplete.TabStop{
						Value: "px/svc_info",
						Kind:  cloudapipb.AEK_UNKNOWN,
						Valid: false,
					},
					&autocomplete.TabStop{
						Value: "pl/test",
						Kind:  cloudapipb.AEK_POD,
						Valid: true,
					},
				},
				Executable: false,
			},
		},
		{
			name:  "invalid script",
			input: "script:px/svc_info pl/$0test",
			requests: [][]*autocomplete.SuggestionRequest{
				[]*autocomplete.SuggestionRequest{
					&autocomplete.SuggestionRequest{
						OrgID:        orgID,
						Input:        "px/svc_info",
						AllowedKinds: []cloudapipb.AutocompleteEntityKind{cloudapipb.AEK_SCRIPT},
						AllowedArgs:  []cloudapipb.AutocompleteEntityKind{},
					},
				},
				[]*autocomplete.SuggestionRequest{
					&autocomplete.SuggestionRequest{
						OrgID:        orgID,
						Input:        "pl/test",
						AllowedKinds: []cloudapipb.AutocompleteEntityKind{cloudapipb.AEK_POD, cloudapipb.AEK_SVC, cloudapipb.AEK_NAMESPACE},
						AllowedArgs:  []cloudapipb.AutocompleteEntityKind{},
					},
				},
			},
			responses: [][]*autocomplete.SuggestionResult{
				[]*autocomplete.SuggestionResult{
					&autocomplete.SuggestionResult{
						Suggestions: []*autocomplete.Suggestion{
							&autocomplete.Suggestion{
								Name:  "px/svc_infos",
								Score: 1,
								Args:  []cloudapipb.AutocompleteEntityKind{cloudapipb.AEK_SVC},
							},
						},
						ExactMatch: false,
					},
				},
				[]*autocomplete.SuggestionResult{
					&autocomplete.SuggestionResult{
						Suggestions: []*autocomplete.Suggestion{
							&autocomplete.Suggestion{
								Name:  "pl/test",
								Score: 1,
							},
						},
						ExactMatch: true,
					},
				},
			},
			expectedCmd: &autocomplete.Command{
				TabStops: []*autocomplete.TabStop{
					&autocomplete.TabStop{
						Value: "run",
						Kind:  cloudapipb.AEK_UNKNOWN,
						Valid: true,
					},
					&autocomplete.TabStop{
						Value: "px/svc_info",
						Kind:  cloudapipb.AEK_SCRIPT,
						Valid: false,
					},
					&autocomplete.TabStop{
						Value:          "pl/$0test",
						Kind:           cloudapipb.AEK_UNKNOWN,
						Valid:          false,
						ContainsCursor: true,
					},
				},
				Executable: false,
			},
		},
		{
			name:  "script with two args",
			input: "script:$0px/svc_info svc:pl/test test",
			requests: [][]*autocomplete.SuggestionRequest{
				[]*autocomplete.SuggestionRequest{
					&autocomplete.SuggestionRequest{
						OrgID:        orgID,
						Input:        "px/svc_info",
						AllowedKinds: []cloudapipb.AutocompleteEntityKind{cloudapipb.AEK_SCRIPT},
						AllowedArgs:  []cloudapipb.AutocompleteEntityKind{},
					},
				},
				[]*autocomplete.SuggestionRequest{
					&autocomplete.SuggestionRequest{
						OrgID:        orgID,
						Input:        "pl/test",
						AllowedKinds: []cloudapipb.AutocompleteEntityKind{cloudapipb.AEK_SVC},
						AllowedArgs:  []cloudapipb.AutocompleteEntityKind{cloudapipb.AEK_SVC},
					},
					&autocomplete.SuggestionRequest{
						OrgID:        orgID,
						Input:        "test",
						AllowedKinds: []cloudapipb.AutocompleteEntityKind{cloudapipb.AEK_SVC},
						AllowedArgs:  []cloudapipb.AutocompleteEntityKind{cloudapipb.AEK_SVC},
					},
				},
			},
			responses: [][]*autocomplete.SuggestionResult{
				[]*autocomplete.SuggestionResult{
					&autocomplete.SuggestionResult{
						Suggestions: []*autocomplete.Suggestion{
							&autocomplete.Suggestion{
								Name:  "px/svc_info",
								Score: 1,
								Args:  []cloudapipb.AutocompleteEntityKind{cloudapipb.AEK_SVC, cloudapipb.AEK_SVC},
							},
						},
						ExactMatch: true,
					},
				},
				[]*autocomplete.SuggestionResult{
					&autocomplete.SuggestionResult{
						Suggestions: []*autocomplete.Suggestion{
							&autocomplete.Suggestion{
								Name:  "pl/test",
								Score: 1,
							},
						},
						ExactMatch: true,
					},
					&autocomplete.SuggestionResult{
						Suggestions: []*autocomplete.Suggestion{
							&autocomplete.Suggestion{
								Name:  "pl/test",
								Score: 1,
							},
						},
						ExactMatch: false,
					},
				},
			},
			expectedCmd: &autocomplete.Command{
				TabStops: []*autocomplete.TabStop{
					&autocomplete.TabStop{
						Value: "run",
						Kind:  cloudapipb.AEK_UNKNOWN,
						Valid: true,
					},
					&autocomplete.TabStop{
						Value:          "$0px/svc_info",
						Kind:           cloudapipb.AEK_SCRIPT,
						Valid:          true,
						ContainsCursor: true,
					},
					&autocomplete.TabStop{
						Value: "pl/test",
						Kind:  cloudapipb.AEK_SVC,
						Valid: true,
					},
					&autocomplete.TabStop{
						Value: "test",
						Kind:  cloudapipb.AEK_UNKNOWN,
						Valid: false,
					},
				},
				Executable: false,
			},
		},
		{
			name:  "invalid label",
			input: "script:px/svc_info no:pl/test",
			requests: [][]*autocomplete.SuggestionRequest{
				[]*autocomplete.SuggestionRequest{
					&autocomplete.SuggestionRequest{
						OrgID:        orgID,
						Input:        "px/svc_info",
						AllowedKinds: []cloudapipb.AutocompleteEntityKind{cloudapipb.AEK_SCRIPT},
						AllowedArgs:  []cloudapipb.AutocompleteEntityKind{},
					},
				},
				[]*autocomplete.SuggestionRequest{
					&autocomplete.SuggestionRequest{
						OrgID:        orgID,
						Input:        "no:pl/test",
						AllowedKinds: []cloudapipb.AutocompleteEntityKind{cloudapipb.AEK_SVC},
						AllowedArgs:  []cloudapipb.AutocompleteEntityKind{},
					},
				},
			},
			responses: [][]*autocomplete.SuggestionResult{
				[]*autocomplete.SuggestionResult{
					&autocomplete.SuggestionResult{
						Suggestions: []*autocomplete.Suggestion{
							&autocomplete.Suggestion{
								Name:  "px/svc_info",
								Score: 1,
								Args:  []cloudapipb.AutocompleteEntityKind{cloudapipb.AEK_SVC},
							},
						},
						ExactMatch: true,
					},
				},
				[]*autocomplete.SuggestionResult{
					&autocomplete.SuggestionResult{
						Suggestions: []*autocomplete.Suggestion{
							&autocomplete.Suggestion{
								Name:  "pl/test",
								Score: 1,
							},
						},
						ExactMatch: false,
					},
				},
			},
			expectedCmd: &autocomplete.Command{
				TabStops: []*autocomplete.TabStop{
					&autocomplete.TabStop{
						Value: "run",
						Kind:  cloudapipb.AEK_UNKNOWN,
						Valid: true,
					},
					&autocomplete.TabStop{
						Value: "px/svc_info",
						Kind:  cloudapipb.AEK_SCRIPT,
						Valid: true,
					},
					&autocomplete.TabStop{
						Value: "no:pl/test",
						Kind:  cloudapipb.AEK_UNKNOWN,
						Valid: false,
					},
				},
				Executable: false,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			s := mock_autocomplete.NewMockSuggester(ctrl)
			suggestionCalls := 0

			s.EXPECT().
				GetSuggestions(gomock.Any()).
				DoAndReturn(func(req []*autocomplete.SuggestionRequest) ([]*autocomplete.SuggestionResult, error) {
					assert.ElementsMatch(t, test.requests[suggestionCalls], req)
					resp := test.responses[suggestionCalls]
					suggestionCalls++
					return resp, nil
				}).
				Times(len(test.requests))

			cmd, err := autocomplete.ParseIntoCommand(test.input, s, orgID)
			assert.Nil(t, err)
			assert.NotNil(t, cmd)

			assert.Equal(t, test.expectedCmd.Executable, cmd.Executable)
			assert.Equal(t, len(test.expectedCmd.TabStops), len(cmd.TabStops))
			for i, a := range test.expectedCmd.TabStops {
				assert.Equal(t, a.Value, cmd.TabStops[i].Value)
				assert.Equal(t, a.Valid, cmd.TabStops[i].Valid)
				assert.Equal(t, a.Kind, cmd.TabStops[i].Kind)
			}
		})
	}
}

func TestToFormatString(t *testing.T) {
	tests := []struct {
		name                string
		cmd                 *autocomplete.Command
		action              cloudapipb.AutocompleteActionType
		expectedStr         string
		expectedSuggestions []*cloudapipb.TabSuggestion
	}{
		{
			name: "edit",
			cmd: &autocomplete.Command{
				TabStops: []*autocomplete.TabStop{
					&autocomplete.TabStop{
						Value: "run",
						Kind:  cloudapipb.AEK_UNKNOWN,
						Valid: true,
					},
					&autocomplete.TabStop{
						Value:          "px/$0svc_info",
						Kind:           cloudapipb.AEK_SCRIPT,
						Valid:          true,
						ContainsCursor: true,
					},
					&autocomplete.TabStop{
						Value: "no:pl/test",
						Kind:  cloudapipb.AEK_UNKNOWN,
						Valid: false,
						Suggestions: []*autocomplete.Suggestion{
							&autocomplete.Suggestion{
								Name: "pl/test",
								Desc: "a svc",
								Kind: cloudapipb.AEK_SVC,
							},
						},
					},
				},
				Executable: false,
			},
			action:      cloudapipb.AAT_EDIT,
			expectedStr: "${2:run} ${3:script:px/$0svc_info} ${1:no:pl/test}",
			expectedSuggestions: []*cloudapipb.TabSuggestion{
				&cloudapipb.TabSuggestion{
					TabIndex:              2,
					ExecutableAfterSelect: false,
					Suggestions:           []*cloudapipb.AutocompleteSuggestion{},
				},
				&cloudapipb.TabSuggestion{
					TabIndex:              3,
					ExecutableAfterSelect: false,
					Suggestions:           []*cloudapipb.AutocompleteSuggestion{},
				},
				&cloudapipb.TabSuggestion{
					TabIndex:              1,
					ExecutableAfterSelect: true,
					Suggestions: []*cloudapipb.AutocompleteSuggestion{
						&cloudapipb.AutocompleteSuggestion{
							Kind:        cloudapipb.AEK_SVC,
							Name:        "pl/test",
							Description: "a svc",
						},
					},
				},
			},
		},
		{
			name: "edit",
			cmd: &autocomplete.Command{
				TabStops: []*autocomplete.TabStop{
					&autocomplete.TabStop{
						Value: "run",
						Kind:  cloudapipb.AEK_UNKNOWN,
						Valid: true,
					},
					&autocomplete.TabStop{
						Value:          "px/$0svc_info",
						Kind:           cloudapipb.AEK_SCRIPT,
						Valid:          true,
						ContainsCursor: true,
					},
					&autocomplete.TabStop{
						Value: "no:pl/test",
						Kind:  cloudapipb.AEK_UNKNOWN,
						Valid: false,
						Suggestions: []*autocomplete.Suggestion{
							&autocomplete.Suggestion{
								Name: "pl/test",
								Desc: "a svc",
								Kind: cloudapipb.AEK_SVC,
							},
						},
					},
				},
				Executable: false,
			},
			action:      cloudapipb.AAT_SELECT,
			expectedStr: "${1:run} ${2:script:px/svc_info} ${3:no:pl/test$0}",
			expectedSuggestions: []*cloudapipb.TabSuggestion{
				&cloudapipb.TabSuggestion{
					TabIndex:              1,
					ExecutableAfterSelect: false,
					Suggestions:           []*cloudapipb.AutocompleteSuggestion{},
				},
				&cloudapipb.TabSuggestion{
					TabIndex:              2,
					ExecutableAfterSelect: false,
					Suggestions:           []*cloudapipb.AutocompleteSuggestion{},
				},
				&cloudapipb.TabSuggestion{
					TabIndex:              3,
					ExecutableAfterSelect: true,
					Suggestions: []*cloudapipb.AutocompleteSuggestion{
						&cloudapipb.AutocompleteSuggestion{
							Kind:        cloudapipb.AEK_SVC,
							Name:        "pl/test",
							Description: "a svc",
						},
					},
				},
			},
		},
		{
			name: "empty value",
			cmd: &autocomplete.Command{
				TabStops: []*autocomplete.TabStop{
					&autocomplete.TabStop{
						Value: "run",
						Kind:  cloudapipb.AEK_UNKNOWN,
						Valid: true,
					},
					&autocomplete.TabStop{
						Value:          "px/svc_info$0",
						Kind:           cloudapipb.AEK_SCRIPT,
						Valid:          true,
						ContainsCursor: true,
					},
					&autocomplete.TabStop{
						Value: "",
						Kind:  cloudapipb.AEK_UNKNOWN,
						Valid: false,
						Suggestions: []*autocomplete.Suggestion{
							&autocomplete.Suggestion{
								Name: "pl/test",
								Desc: "a svc",
								Kind: cloudapipb.AEK_SVC,
							},
						},
					},
				},
				Executable: false,
			},
			action:      cloudapipb.AAT_EDIT,
			expectedStr: "${2:run} ${3:script:px/svc_info$0} $1",
			expectedSuggestions: []*cloudapipb.TabSuggestion{
				&cloudapipb.TabSuggestion{
					TabIndex:              2,
					ExecutableAfterSelect: false,
					Suggestions:           []*cloudapipb.AutocompleteSuggestion{},
				},
				&cloudapipb.TabSuggestion{
					TabIndex:              3,
					ExecutableAfterSelect: false,
					Suggestions:           []*cloudapipb.AutocompleteSuggestion{},
				},
				&cloudapipb.TabSuggestion{
					TabIndex:              1,
					ExecutableAfterSelect: true,
					Suggestions: []*cloudapipb.AutocompleteSuggestion{
						&cloudapipb.AutocompleteSuggestion{
							Kind:        cloudapipb.AEK_SVC,
							Name:        "pl/test",
							Description: "a svc",
						},
					},
				},
			},
		},
		{
			name: "invalid before current cursor",
			cmd: &autocomplete.Command{
				TabStops: []*autocomplete.TabStop{
					&autocomplete.TabStop{
						Value: "run",
						Kind:  cloudapipb.AEK_UNKNOWN,
						Valid: true,
					},
					&autocomplete.TabStop{
						Value: "blah",
						Kind:  cloudapipb.AEK_UNKNOWN,
						Valid: false,
						Suggestions: []*autocomplete.Suggestion{
							&autocomplete.Suggestion{
								Name: "pl/blah",
								Desc: "a svc",
								Kind: cloudapipb.AEK_SVC,
							},
						},
					},
					&autocomplete.TabStop{
						Value:          "pl/frontend$0",
						Kind:           cloudapipb.AEK_SVC,
						Valid:          true,
						ContainsCursor: true,
					},
				},
				Executable: false,
			},
			action:      cloudapipb.AAT_SELECT,
			expectedStr: "${2:run} ${3:blah$0} ${1:svc:pl/frontend}",
			expectedSuggestions: []*cloudapipb.TabSuggestion{
				&cloudapipb.TabSuggestion{
					TabIndex:              2,
					ExecutableAfterSelect: false,
					Suggestions:           []*cloudapipb.AutocompleteSuggestion{},
				},
				&cloudapipb.TabSuggestion{
					TabIndex:              3,
					ExecutableAfterSelect: true,
					Suggestions: []*cloudapipb.AutocompleteSuggestion{
						&cloudapipb.AutocompleteSuggestion{
							Kind:        cloudapipb.AEK_SVC,
							Name:        "pl/blah",
							Description: "a svc",
						},
					},
				},
				&cloudapipb.TabSuggestion{
					TabIndex:              1,
					ExecutableAfterSelect: false,
					Suggestions:           []*cloudapipb.AutocompleteSuggestion{},
				},
			},
		},
		{
			name: "all valid",
			cmd: &autocomplete.Command{
				TabStops: []*autocomplete.TabStop{
					&autocomplete.TabStop{
						Value: "run",
						Kind:  cloudapipb.AEK_UNKNOWN,
						Valid: true,
					},
					&autocomplete.TabStop{
						Value:          "px/svc_info$0",
						Kind:           cloudapipb.AEK_SCRIPT,
						Valid:          true,
						ContainsCursor: true,
					},
					&autocomplete.TabStop{
						Value: "pl/frontend",
						Kind:  cloudapipb.AEK_SVC,
						Valid: true,
					},
				},
				Executable: false,
			},
			action:      cloudapipb.AAT_SELECT,
			expectedStr: "${1:run} ${2:script:px/svc_info} ${3:svc:pl/frontend$0}",
			expectedSuggestions: []*cloudapipb.TabSuggestion{
				&cloudapipb.TabSuggestion{
					TabIndex:              1,
					ExecutableAfterSelect: false,
					Suggestions:           []*cloudapipb.AutocompleteSuggestion{},
				},
				&cloudapipb.TabSuggestion{
					TabIndex:              2,
					ExecutableAfterSelect: false,
					Suggestions:           []*cloudapipb.AutocompleteSuggestion{},
				},
				&cloudapipb.TabSuggestion{
					TabIndex:              3,
					ExecutableAfterSelect: false,
					Suggestions:           []*cloudapipb.AutocompleteSuggestion{},
				},
			},
		},
		{
			name: "invalid args",
			cmd: &autocomplete.Command{
				TabStops: []*autocomplete.TabStop{
					&autocomplete.TabStop{
						Value: "run",
						Kind:  cloudapipb.AEK_UNKNOWN,
						Valid: true,
					},
					&autocomplete.TabStop{
						Value: "px/svc_info",
						Kind:  cloudapipb.AEK_SCRIPT,
						Valid: false,
						Suggestions: []*autocomplete.Suggestion{
							&autocomplete.Suggestion{
								Name: "pl/svc_info_abc",
								Desc: "a svc",
								Kind: cloudapipb.AEK_SCRIPT,
							},
						},
					},
					&autocomplete.TabStop{
						Value:          "pl/frontend$0",
						Kind:           cloudapipb.AEK_POD,
						Valid:          false,
						ContainsCursor: true,
						Suggestions: []*autocomplete.Suggestion{
							&autocomplete.Suggestion{
								Name: "pl/frontend-test",
								Desc: "a pod",
								Kind: cloudapipb.AEK_POD,
							},
						},
					},
				},
				Executable: false,
			},
			action:      cloudapipb.AAT_EDIT,
			expectedStr: "${1:run} ${2:script:px/svc_info} ${3:pod:pl/frontend$0}",
			expectedSuggestions: []*cloudapipb.TabSuggestion{
				&cloudapipb.TabSuggestion{
					TabIndex:              1,
					ExecutableAfterSelect: false,
					Suggestions:           []*cloudapipb.AutocompleteSuggestion{},
				},
				&cloudapipb.TabSuggestion{
					TabIndex:              2,
					ExecutableAfterSelect: false,
					Suggestions: []*cloudapipb.AutocompleteSuggestion{
						&cloudapipb.AutocompleteSuggestion{
							Kind:        cloudapipb.AEK_SCRIPT,
							Name:        "pl/svc_info_abc",
							Description: "a svc",
						},
					},
				},
				&cloudapipb.TabSuggestion{
					TabIndex:              3,
					ExecutableAfterSelect: false,
					Suggestions: []*cloudapipb.AutocompleteSuggestion{
						&cloudapipb.AutocompleteSuggestion{
							Kind:        cloudapipb.AEK_POD,
							Name:        "pl/frontend-test",
							Description: "a pod",
						},
					},
				},
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output, suggestions := test.cmd.ToFormatString(test.action)
			assert.Equal(t, test.expectedStr, output)
			assert.ElementsMatch(t, test.expectedSuggestions, suggestions)
		})
	}
}
