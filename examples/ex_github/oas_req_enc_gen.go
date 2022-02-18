// Code generated by ogen, DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"math"
	"math/big"
	"math/bits"
	"net"
	"net/http"
	"net/url"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
	"github.com/google/uuid"
	"github.com/ogen-go/ogen/conv"
	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/json"
	"github.com/ogen-go/ogen/otelogen"
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
)

// No-op definition for keeping imports.
var (
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = sort.Ints
	_ = http.MethodGet
	_ = io.Copy
	_ = json.Marshal
	_ = bytes.NewReader
	_ = strconv.ParseInt
	_ = time.Time{}
	_ = conv.ToInt32
	_ = uuid.UUID{}
	_ = uri.PathEncoder{}
	_ = url.URL{}
	_ = math.Mod
	_ = bits.LeadingZeros64
	_ = big.Rat{}
	_ = validate.Int{}
	_ = ht.NewRequest
	_ = net.IP{}
	_ = otelogen.Version
	_ = trace.TraceIDFromHex
	_ = otel.GetTracerProvider
	_ = metric.NewNoopMeterProvider
	_ = regexp.MustCompile
	_ = jx.Null
	_ = sync.Pool{}
	_ = codes.Unset
)

func encodeActionsCreateOrUpdateEnvironmentSecretRequestJSON(req ActionsCreateOrUpdateEnvironmentSecretReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()

	req.Encode(e)

	return e, nil
}

func encodeActionsCreateOrUpdateOrgSecretRequestJSON(req ActionsCreateOrUpdateOrgSecretReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()

	req.Encode(e)

	return e, nil
}

func encodeActionsCreateOrUpdateRepoSecretRequestJSON(req ActionsCreateOrUpdateRepoSecretReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()

	req.Encode(e)

	return e, nil
}

func encodeActionsCreateSelfHostedRunnerGroupForOrgRequestJSON(req ActionsCreateSelfHostedRunnerGroupForOrgReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()

	req.Encode(e)

	return e, nil
}

func encodeActionsReviewPendingDeploymentsForRunRequestJSON(req ActionsReviewPendingDeploymentsForRunReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()

	req.Encode(e)

	return e, nil
}

func encodeActionsSetAllowedActionsOrganizationRequestJSON(req OptSelectedActions, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()
	if req.Set {
		req.Encode(e)
	}

	return e, nil
}

func encodeActionsSetAllowedActionsRepositoryRequestJSON(req OptSelectedActions, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()
	if req.Set {
		req.Encode(e)
	}

	return e, nil
}

func encodeActionsSetGithubActionsPermissionsOrganizationRequestJSON(req ActionsSetGithubActionsPermissionsOrganizationReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()

	req.Encode(e)

	return e, nil
}

func encodeActionsSetGithubActionsPermissionsRepositoryRequestJSON(req ActionsSetGithubActionsPermissionsRepositoryReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()

	req.Encode(e)

	return e, nil
}

func encodeActionsSetRepoAccessToSelfHostedRunnerGroupInOrgRequestJSON(req ActionsSetRepoAccessToSelfHostedRunnerGroupInOrgReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()

	req.Encode(e)

	return e, nil
}

func encodeActionsSetSelectedReposForOrgSecretRequestJSON(req ActionsSetSelectedReposForOrgSecretReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()

	req.Encode(e)

	return e, nil
}

func encodeActionsSetSelectedRepositoriesEnabledGithubActionsOrganizationRequestJSON(req ActionsSetSelectedRepositoriesEnabledGithubActionsOrganizationReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()

	req.Encode(e)

	return e, nil
}

func encodeActionsSetSelfHostedRunnersInGroupForOrgRequestJSON(req ActionsSetSelfHostedRunnersInGroupForOrgReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()

	req.Encode(e)

	return e, nil
}

func encodeActionsUpdateSelfHostedRunnerGroupForOrgRequestJSON(req ActionsUpdateSelfHostedRunnerGroupForOrgReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()

	req.Encode(e)

	return e, nil
}

func encodeActivityMarkNotificationsAsReadRequestJSON(req OptActivityMarkNotificationsAsReadReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()
	if req.Set {
		req.Encode(e)
	}

	return e, nil
}

func encodeActivityMarkRepoNotificationsAsReadRequestJSON(req OptActivityMarkRepoNotificationsAsReadReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()
	if req.Set {
		req.Encode(e)
	}

	return e, nil
}

func encodeActivitySetRepoSubscriptionRequestJSON(req OptActivitySetRepoSubscriptionReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()
	if req.Set {
		req.Encode(e)
	}

	return e, nil
}

func encodeActivitySetThreadSubscriptionRequestJSON(req OptActivitySetThreadSubscriptionReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()
	if req.Set {
		req.Encode(e)
	}

	return e, nil
}

func encodeAppsCheckTokenRequestJSON(req AppsCheckTokenReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()

	req.Encode(e)

	return e, nil
}

func encodeAppsCreateContentAttachmentRequestJSON(req AppsCreateContentAttachmentReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()

	req.Encode(e)

	return e, nil
}

func encodeAppsCreateInstallationAccessTokenRequestJSON(req OptAppsCreateInstallationAccessTokenReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()
	if req.Set {
		req.Encode(e)
	}

	return e, nil
}

func encodeAppsDeleteAuthorizationRequestJSON(req AppsDeleteAuthorizationReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()

	req.Encode(e)

	return e, nil
}

func encodeAppsDeleteTokenRequestJSON(req AppsDeleteTokenReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()

	req.Encode(e)

	return e, nil
}

func encodeAppsResetTokenRequestJSON(req AppsResetTokenReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()

	req.Encode(e)

	return e, nil
}

func encodeAppsScopeTokenRequestJSON(req AppsScopeTokenReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()

	req.Encode(e)

	return e, nil
}

func encodeAppsUpdateWebhookConfigForAppRequestJSON(req OptAppsUpdateWebhookConfigForAppReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()
	if req.Set {
		req.Encode(e)
	}

	return e, nil
}

func encodeChecksCreateSuiteRequestJSON(req ChecksCreateSuiteReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()

	req.Encode(e)

	return e, nil
}

func encodeChecksSetSuitesPreferencesRequestJSON(req ChecksSetSuitesPreferencesReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()

	req.Encode(e)

	return e, nil
}

func encodeCodeScanningUpdateAlertRequestJSON(req CodeScanningUpdateAlertReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()

	req.Encode(e)

	return e, nil
}

func encodeCodeScanningUploadSarifRequestJSON(req CodeScanningUploadSarifReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()

	req.Encode(e)

	return e, nil
}

func encodeEnterpriseAdminCreateSelfHostedRunnerGroupForEnterpriseRequestJSON(req EnterpriseAdminCreateSelfHostedRunnerGroupForEnterpriseReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()

	req.Encode(e)

	return e, nil
}

func encodeEnterpriseAdminProvisionAndInviteEnterpriseGroupRequestJSON(req EnterpriseAdminProvisionAndInviteEnterpriseGroupReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()

	req.Encode(e)

	return e, nil
}

func encodeEnterpriseAdminProvisionAndInviteEnterpriseUserRequestJSON(req EnterpriseAdminProvisionAndInviteEnterpriseUserReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()

	req.Encode(e)

	return e, nil
}

func encodeEnterpriseAdminSetAllowedActionsEnterpriseRequestJSON(req SelectedActions, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()

	req.Encode(e)

	return e, nil
}

func encodeEnterpriseAdminSetGithubActionsPermissionsEnterpriseRequestJSON(req EnterpriseAdminSetGithubActionsPermissionsEnterpriseReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()

	req.Encode(e)

	return e, nil
}

func encodeEnterpriseAdminSetInformationForProvisionedEnterpriseGroupRequestJSON(req EnterpriseAdminSetInformationForProvisionedEnterpriseGroupReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()

	req.Encode(e)

	return e, nil
}

func encodeEnterpriseAdminSetInformationForProvisionedEnterpriseUserRequestJSON(req EnterpriseAdminSetInformationForProvisionedEnterpriseUserReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()

	req.Encode(e)

	return e, nil
}

func encodeEnterpriseAdminSetOrgAccessToSelfHostedRunnerGroupInEnterpriseRequestJSON(req EnterpriseAdminSetOrgAccessToSelfHostedRunnerGroupInEnterpriseReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()

	req.Encode(e)

	return e, nil
}

func encodeEnterpriseAdminSetSelectedOrganizationsEnabledGithubActionsEnterpriseRequestJSON(req EnterpriseAdminSetSelectedOrganizationsEnabledGithubActionsEnterpriseReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()

	req.Encode(e)

	return e, nil
}

func encodeEnterpriseAdminSetSelfHostedRunnersInGroupForEnterpriseRequestJSON(req EnterpriseAdminSetSelfHostedRunnersInGroupForEnterpriseReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()

	req.Encode(e)

	return e, nil
}

func encodeEnterpriseAdminUpdateAttributeForEnterpriseGroupRequestJSON(req EnterpriseAdminUpdateAttributeForEnterpriseGroupReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()

	req.Encode(e)

	return e, nil
}

func encodeEnterpriseAdminUpdateAttributeForEnterpriseUserRequestJSON(req EnterpriseAdminUpdateAttributeForEnterpriseUserReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()

	req.Encode(e)

	return e, nil
}

func encodeEnterpriseAdminUpdateSelfHostedRunnerGroupForEnterpriseRequestJSON(req OptEnterpriseAdminUpdateSelfHostedRunnerGroupForEnterpriseReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()
	if req.Set {
		req.Encode(e)
	}

	return e, nil
}

func encodeGistsCreateRequestJSON(req GistsCreateReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()

	req.Encode(e)

	return e, nil
}

func encodeGistsCreateCommentRequestJSON(req GistsCreateCommentReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()

	req.Encode(e)

	return e, nil
}

func encodeGistsUpdateCommentRequestJSON(req GistsUpdateCommentReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()

	req.Encode(e)

	return e, nil
}

func encodeGitCreateBlobRequestJSON(req GitCreateBlobReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()

	req.Encode(e)

	return e, nil
}

func encodeGitCreateCommitRequestJSON(req GitCreateCommitReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()

	req.Encode(e)

	return e, nil
}

func encodeGitCreateRefRequestJSON(req GitCreateRefReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()

	req.Encode(e)

	return e, nil
}

func encodeGitCreateTagRequestJSON(req GitCreateTagReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()

	req.Encode(e)

	return e, nil
}

func encodeGitCreateTreeRequestJSON(req GitCreateTreeReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()

	req.Encode(e)

	return e, nil
}

func encodeGitUpdateRefRequestJSON(req GitUpdateRefReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()

	req.Encode(e)

	return e, nil
}

func encodeInteractionsSetRestrictionsForAuthenticatedUserRequestJSON(req InteractionLimit, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()

	req.Encode(e)

	return e, nil
}

func encodeInteractionsSetRestrictionsForOrgRequestJSON(req InteractionLimit, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()

	req.Encode(e)

	return e, nil
}

func encodeInteractionsSetRestrictionsForRepoRequestJSON(req InteractionLimit, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()

	req.Encode(e)

	return e, nil
}

func encodeIssuesAddAssigneesRequestJSON(req OptIssuesAddAssigneesReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()
	if req.Set {
		req.Encode(e)
	}

	return e, nil
}

func encodeIssuesCreateRequestJSON(req IssuesCreateReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()

	req.Encode(e)

	return e, nil
}

func encodeIssuesCreateCommentRequestJSON(req IssuesCreateCommentReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()

	req.Encode(e)

	return e, nil
}

func encodeIssuesCreateLabelRequestJSON(req IssuesCreateLabelReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()

	req.Encode(e)

	return e, nil
}

func encodeIssuesCreateMilestoneRequestJSON(req IssuesCreateMilestoneReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()

	req.Encode(e)

	return e, nil
}

func encodeIssuesLockRequestJSON(req OptNilIssuesLockReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()
	if req.Set {
		req.Encode(e)
	}

	return e, nil
}

func encodeIssuesRemoveAssigneesRequestJSON(req OptIssuesRemoveAssigneesReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()
	if req.Set {
		req.Encode(e)
	}

	return e, nil
}

func encodeIssuesUpdateRequestJSON(req OptIssuesUpdateReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()
	if req.Set {
		req.Encode(e)
	}

	return e, nil
}

func encodeIssuesUpdateCommentRequestJSON(req IssuesUpdateCommentReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()

	req.Encode(e)

	return e, nil
}

func encodeIssuesUpdateLabelRequestJSON(req OptIssuesUpdateLabelReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()
	if req.Set {
		req.Encode(e)
	}

	return e, nil
}

func encodeIssuesUpdateMilestoneRequestJSON(req OptIssuesUpdateMilestoneReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()
	if req.Set {
		req.Encode(e)
	}

	return e, nil
}

func encodeMigrationsMapCommitAuthorRequestJSON(req OptMigrationsMapCommitAuthorReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()
	if req.Set {
		req.Encode(e)
	}

	return e, nil
}

func encodeMigrationsSetLfsPreferenceRequestJSON(req MigrationsSetLfsPreferenceReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()

	req.Encode(e)

	return e, nil
}

func encodeMigrationsStartForAuthenticatedUserRequestJSON(req MigrationsStartForAuthenticatedUserReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()

	req.Encode(e)

	return e, nil
}

func encodeMigrationsStartForOrgRequestJSON(req MigrationsStartForOrgReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()

	req.Encode(e)

	return e, nil
}

func encodeMigrationsStartImportRequestJSON(req MigrationsStartImportReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()

	req.Encode(e)

	return e, nil
}

func encodeMigrationsUpdateImportRequestJSON(req OptNilMigrationsUpdateImportReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()
	if req.Set {
		req.Encode(e)
	}

	return e, nil
}

func encodeOAuthAuthorizationsCreateAuthorizationRequestJSON(req OptOAuthAuthorizationsCreateAuthorizationReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()
	if req.Set {
		req.Encode(e)
	}

	return e, nil
}

func encodeOAuthAuthorizationsGetOrCreateAuthorizationForAppRequestJSON(req OAuthAuthorizationsGetOrCreateAuthorizationForAppReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()

	req.Encode(e)

	return e, nil
}

func encodeOAuthAuthorizationsGetOrCreateAuthorizationForAppAndFingerprintRequestJSON(req OAuthAuthorizationsGetOrCreateAuthorizationForAppAndFingerprintReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()

	req.Encode(e)

	return e, nil
}

func encodeOAuthAuthorizationsUpdateAuthorizationRequestJSON(req OptOAuthAuthorizationsUpdateAuthorizationReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()
	if req.Set {
		req.Encode(e)
	}

	return e, nil
}

func encodeOrgsCreateInvitationRequestJSON(req OptOrgsCreateInvitationReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()
	if req.Set {
		req.Encode(e)
	}

	return e, nil
}

func encodeOrgsCreateWebhookRequestJSON(req OrgsCreateWebhookReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()

	req.Encode(e)

	return e, nil
}

func encodeOrgsSetMembershipForUserRequestJSON(req OptOrgsSetMembershipForUserReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()
	if req.Set {
		req.Encode(e)
	}

	return e, nil
}

func encodeOrgsUpdateMembershipForAuthenticatedUserRequestJSON(req OrgsUpdateMembershipForAuthenticatedUserReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()

	req.Encode(e)

	return e, nil
}

func encodeOrgsUpdateWebhookRequestJSON(req OptOrgsUpdateWebhookReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()
	if req.Set {
		req.Encode(e)
	}

	return e, nil
}

func encodeOrgsUpdateWebhookConfigForOrgRequestJSON(req OptOrgsUpdateWebhookConfigForOrgReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()
	if req.Set {
		req.Encode(e)
	}

	return e, nil
}

func encodeProjectsAddCollaboratorRequestJSON(req OptNilProjectsAddCollaboratorReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()
	if req.Set {
		req.Encode(e)
	}

	return e, nil
}

func encodeProjectsCreateColumnRequestJSON(req ProjectsCreateColumnReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()

	req.Encode(e)

	return e, nil
}

func encodeProjectsCreateForAuthenticatedUserRequestJSON(req ProjectsCreateForAuthenticatedUserReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()

	req.Encode(e)

	return e, nil
}

func encodeProjectsCreateForOrgRequestJSON(req ProjectsCreateForOrgReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()

	req.Encode(e)

	return e, nil
}

func encodeProjectsCreateForRepoRequestJSON(req ProjectsCreateForRepoReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()

	req.Encode(e)

	return e, nil
}

func encodeProjectsMoveCardRequestJSON(req ProjectsMoveCardReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()

	req.Encode(e)

	return e, nil
}

func encodeProjectsMoveColumnRequestJSON(req ProjectsMoveColumnReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()

	req.Encode(e)

	return e, nil
}

func encodeProjectsUpdateRequestJSON(req OptProjectsUpdateReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()
	if req.Set {
		req.Encode(e)
	}

	return e, nil
}

func encodeProjectsUpdateCardRequestJSON(req OptProjectsUpdateCardReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()
	if req.Set {
		req.Encode(e)
	}

	return e, nil
}

func encodeProjectsUpdateColumnRequestJSON(req ProjectsUpdateColumnReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()

	req.Encode(e)

	return e, nil
}

func encodePullsCreateRequestJSON(req PullsCreateReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()

	req.Encode(e)

	return e, nil
}

func encodePullsCreateReplyForReviewCommentRequestJSON(req PullsCreateReplyForReviewCommentReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()

	req.Encode(e)

	return e, nil
}

func encodePullsCreateReviewRequestJSON(req OptPullsCreateReviewReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()
	if req.Set {
		req.Encode(e)
	}

	return e, nil
}

func encodePullsCreateReviewCommentRequestJSON(req PullsCreateReviewCommentReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()

	req.Encode(e)

	return e, nil
}

func encodePullsDismissReviewRequestJSON(req PullsDismissReviewReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()

	req.Encode(e)

	return e, nil
}

func encodePullsMergeRequestJSON(req OptNilPullsMergeReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()
	if req.Set {
		req.Encode(e)
	}

	return e, nil
}

func encodePullsRemoveRequestedReviewersRequestJSON(req PullsRemoveRequestedReviewersReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()

	req.Encode(e)

	return e, nil
}

func encodePullsSubmitReviewRequestJSON(req PullsSubmitReviewReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()

	req.Encode(e)

	return e, nil
}

func encodePullsUpdateRequestJSON(req OptPullsUpdateReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()
	if req.Set {
		req.Encode(e)
	}

	return e, nil
}

func encodePullsUpdateBranchRequestJSON(req OptNilPullsUpdateBranchReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()
	if req.Set {
		req.Encode(e)
	}

	return e, nil
}

func encodePullsUpdateReviewRequestJSON(req PullsUpdateReviewReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()

	req.Encode(e)

	return e, nil
}

func encodePullsUpdateReviewCommentRequestJSON(req PullsUpdateReviewCommentReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()

	req.Encode(e)

	return e, nil
}

func encodeReactionsCreateForCommitCommentRequestJSON(req ReactionsCreateForCommitCommentReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()

	req.Encode(e)

	return e, nil
}

func encodeReactionsCreateForIssueRequestJSON(req ReactionsCreateForIssueReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()

	req.Encode(e)

	return e, nil
}

func encodeReactionsCreateForIssueCommentRequestJSON(req ReactionsCreateForIssueCommentReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()

	req.Encode(e)

	return e, nil
}

func encodeReactionsCreateForPullRequestReviewCommentRequestJSON(req ReactionsCreateForPullRequestReviewCommentReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()

	req.Encode(e)

	return e, nil
}

func encodeReactionsCreateForReleaseRequestJSON(req ReactionsCreateForReleaseReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()

	req.Encode(e)

	return e, nil
}

func encodeReactionsCreateForTeamDiscussionCommentInOrgRequestJSON(req ReactionsCreateForTeamDiscussionCommentInOrgReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()

	req.Encode(e)

	return e, nil
}

func encodeReactionsCreateForTeamDiscussionCommentLegacyRequestJSON(req ReactionsCreateForTeamDiscussionCommentLegacyReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()

	req.Encode(e)

	return e, nil
}

func encodeReactionsCreateForTeamDiscussionInOrgRequestJSON(req ReactionsCreateForTeamDiscussionInOrgReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()

	req.Encode(e)

	return e, nil
}

func encodeReactionsCreateForTeamDiscussionLegacyRequestJSON(req ReactionsCreateForTeamDiscussionLegacyReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()

	req.Encode(e)

	return e, nil
}

func encodeReposAddAppAccessRestrictionsRequestJSON(req OptReposAddAppAccessRestrictionsReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()
	if req.Set {
		req.Encode(e)
	}

	return e, nil
}

func encodeReposAddCollaboratorRequestJSON(req OptReposAddCollaboratorReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()
	if req.Set {
		req.Encode(e)
	}

	return e, nil
}

func encodeReposAddStatusCheckContextsRequestJSON(req OptReposAddStatusCheckContextsReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()
	if req.Set {
		req.Encode(e)
	}

	return e, nil
}

func encodeReposAddTeamAccessRestrictionsRequestJSON(req OptReposAddTeamAccessRestrictionsReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()
	if req.Set {
		req.Encode(e)
	}

	return e, nil
}

func encodeReposAddUserAccessRestrictionsRequestJSON(req OptReposAddUserAccessRestrictionsReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()
	if req.Set {
		req.Encode(e)
	}

	return e, nil
}

func encodeReposCreateAutolinkRequestJSON(req ReposCreateAutolinkReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()

	req.Encode(e)

	return e, nil
}

func encodeReposCreateCommitCommentRequestJSON(req ReposCreateCommitCommentReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()

	req.Encode(e)

	return e, nil
}

func encodeReposCreateCommitStatusRequestJSON(req ReposCreateCommitStatusReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()

	req.Encode(e)

	return e, nil
}

func encodeReposCreateDeployKeyRequestJSON(req ReposCreateDeployKeyReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()

	req.Encode(e)

	return e, nil
}

func encodeReposCreateDeploymentRequestJSON(req ReposCreateDeploymentReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()

	req.Encode(e)

	return e, nil
}

func encodeReposCreateDeploymentStatusRequestJSON(req ReposCreateDeploymentStatusReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()

	req.Encode(e)

	return e, nil
}

func encodeReposCreateDispatchEventRequestJSON(req ReposCreateDispatchEventReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()

	req.Encode(e)

	return e, nil
}

func encodeReposCreateForAuthenticatedUserRequestJSON(req ReposCreateForAuthenticatedUserReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()

	req.Encode(e)

	return e, nil
}

func encodeReposCreateForkRequestJSON(req OptNilReposCreateForkReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()
	if req.Set {
		req.Encode(e)
	}

	return e, nil
}

func encodeReposCreateInOrgRequestJSON(req ReposCreateInOrgReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()

	req.Encode(e)

	return e, nil
}

func encodeReposCreateOrUpdateFileContentsRequestJSON(req ReposCreateOrUpdateFileContentsReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()

	req.Encode(e)

	return e, nil
}

func encodeReposCreatePagesSiteRequestJSON(req NilReposCreatePagesSiteReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()

	req.Encode(e)

	return e, nil
}

func encodeReposCreateReleaseRequestJSON(req ReposCreateReleaseReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()

	req.Encode(e)

	return e, nil
}

func encodeReposCreateUsingTemplateRequestJSON(req ReposCreateUsingTemplateReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()

	req.Encode(e)

	return e, nil
}

func encodeReposCreateWebhookRequestJSON(req OptNilReposCreateWebhookReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()
	if req.Set {
		req.Encode(e)
	}

	return e, nil
}

func encodeReposDeleteFileRequestJSON(req ReposDeleteFileReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()

	req.Encode(e)

	return e, nil
}

func encodeReposMergeRequestJSON(req ReposMergeReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()

	req.Encode(e)

	return e, nil
}

func encodeReposMergeUpstreamRequestJSON(req ReposMergeUpstreamReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()

	req.Encode(e)

	return e, nil
}

func encodeReposRemoveAppAccessRestrictionsRequestJSON(req OptReposRemoveAppAccessRestrictionsReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()
	if req.Set {
		req.Encode(e)
	}

	return e, nil
}

func encodeReposRemoveStatusCheckContextsRequestJSON(req OptReposRemoveStatusCheckContextsReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()
	if req.Set {
		req.Encode(e)
	}

	return e, nil
}

func encodeReposRemoveTeamAccessRestrictionsRequestJSON(req OptReposRemoveTeamAccessRestrictionsReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()
	if req.Set {
		req.Encode(e)
	}

	return e, nil
}

func encodeReposRemoveUserAccessRestrictionsRequestJSON(req OptReposRemoveUserAccessRestrictionsReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()
	if req.Set {
		req.Encode(e)
	}

	return e, nil
}

func encodeReposRenameBranchRequestJSON(req OptReposRenameBranchReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()
	if req.Set {
		req.Encode(e)
	}

	return e, nil
}

func encodeReposReplaceAllTopicsRequestJSON(req ReposReplaceAllTopicsReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()

	req.Encode(e)

	return e, nil
}

func encodeReposSetAppAccessRestrictionsRequestJSON(req OptReposSetAppAccessRestrictionsReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()
	if req.Set {
		req.Encode(e)
	}

	return e, nil
}

func encodeReposSetStatusCheckContextsRequestJSON(req OptReposSetStatusCheckContextsReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()
	if req.Set {
		req.Encode(e)
	}

	return e, nil
}

func encodeReposSetTeamAccessRestrictionsRequestJSON(req OptReposSetTeamAccessRestrictionsReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()
	if req.Set {
		req.Encode(e)
	}

	return e, nil
}

func encodeReposSetUserAccessRestrictionsRequestJSON(req OptReposSetUserAccessRestrictionsReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()
	if req.Set {
		req.Encode(e)
	}

	return e, nil
}

func encodeReposTransferRequestJSON(req ReposTransferReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()

	req.Encode(e)

	return e, nil
}

func encodeReposUpdateRequestJSON(req OptReposUpdateReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()
	if req.Set {
		req.Encode(e)
	}

	return e, nil
}

func encodeReposUpdateBranchProtectionRequestJSON(req ReposUpdateBranchProtectionReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()

	req.Encode(e)

	return e, nil
}

func encodeReposUpdateCommitCommentRequestJSON(req ReposUpdateCommitCommentReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()

	req.Encode(e)

	return e, nil
}

func encodeReposUpdateInvitationRequestJSON(req OptReposUpdateInvitationReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()
	if req.Set {
		req.Encode(e)
	}

	return e, nil
}

func encodeReposUpdatePullRequestReviewProtectionRequestJSON(req OptReposUpdatePullRequestReviewProtectionReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()
	if req.Set {
		req.Encode(e)
	}

	return e, nil
}

func encodeReposUpdateReleaseRequestJSON(req OptReposUpdateReleaseReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()
	if req.Set {
		req.Encode(e)
	}

	return e, nil
}

func encodeReposUpdateReleaseAssetRequestJSON(req OptReposUpdateReleaseAssetReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()
	if req.Set {
		req.Encode(e)
	}

	return e, nil
}

func encodeReposUpdateStatusCheckProtectionRequestJSON(req OptReposUpdateStatusCheckProtectionReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()
	if req.Set {
		req.Encode(e)
	}

	return e, nil
}

func encodeReposUpdateWebhookRequestJSON(req OptReposUpdateWebhookReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()
	if req.Set {
		req.Encode(e)
	}

	return e, nil
}

func encodeReposUpdateWebhookConfigForRepoRequestJSON(req OptReposUpdateWebhookConfigForRepoReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()
	if req.Set {
		req.Encode(e)
	}

	return e, nil
}

func encodeSecretScanningUpdateAlertRequestJSON(req SecretScanningUpdateAlertReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()

	req.Encode(e)

	return e, nil
}

func encodeTeamsAddOrUpdateMembershipForUserInOrgRequestJSON(req OptTeamsAddOrUpdateMembershipForUserInOrgReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()
	if req.Set {
		req.Encode(e)
	}

	return e, nil
}

func encodeTeamsAddOrUpdateMembershipForUserLegacyRequestJSON(req OptTeamsAddOrUpdateMembershipForUserLegacyReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()
	if req.Set {
		req.Encode(e)
	}

	return e, nil
}

func encodeTeamsAddOrUpdateProjectPermissionsInOrgRequestJSON(req OptNilTeamsAddOrUpdateProjectPermissionsInOrgReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()
	if req.Set {
		req.Encode(e)
	}

	return e, nil
}

func encodeTeamsAddOrUpdateProjectPermissionsLegacyRequestJSON(req OptTeamsAddOrUpdateProjectPermissionsLegacyReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()
	if req.Set {
		req.Encode(e)
	}

	return e, nil
}

func encodeTeamsAddOrUpdateRepoPermissionsInOrgRequestJSON(req OptTeamsAddOrUpdateRepoPermissionsInOrgReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()
	if req.Set {
		req.Encode(e)
	}

	return e, nil
}

func encodeTeamsAddOrUpdateRepoPermissionsLegacyRequestJSON(req OptTeamsAddOrUpdateRepoPermissionsLegacyReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()
	if req.Set {
		req.Encode(e)
	}

	return e, nil
}

func encodeTeamsCreateRequestJSON(req TeamsCreateReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()

	req.Encode(e)

	return e, nil
}

func encodeTeamsCreateDiscussionCommentInOrgRequestJSON(req TeamsCreateDiscussionCommentInOrgReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()

	req.Encode(e)

	return e, nil
}

func encodeTeamsCreateDiscussionCommentLegacyRequestJSON(req TeamsCreateDiscussionCommentLegacyReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()

	req.Encode(e)

	return e, nil
}

func encodeTeamsCreateDiscussionInOrgRequestJSON(req TeamsCreateDiscussionInOrgReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()

	req.Encode(e)

	return e, nil
}

func encodeTeamsCreateDiscussionLegacyRequestJSON(req TeamsCreateDiscussionLegacyReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()

	req.Encode(e)

	return e, nil
}

func encodeTeamsCreateOrUpdateIdpGroupConnectionsInOrgRequestJSON(req TeamsCreateOrUpdateIdpGroupConnectionsInOrgReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()

	req.Encode(e)

	return e, nil
}

func encodeTeamsCreateOrUpdateIdpGroupConnectionsLegacyRequestJSON(req TeamsCreateOrUpdateIdpGroupConnectionsLegacyReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()

	req.Encode(e)

	return e, nil
}

func encodeTeamsUpdateDiscussionCommentInOrgRequestJSON(req TeamsUpdateDiscussionCommentInOrgReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()

	req.Encode(e)

	return e, nil
}

func encodeTeamsUpdateDiscussionCommentLegacyRequestJSON(req TeamsUpdateDiscussionCommentLegacyReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()

	req.Encode(e)

	return e, nil
}

func encodeTeamsUpdateDiscussionInOrgRequestJSON(req OptTeamsUpdateDiscussionInOrgReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()
	if req.Set {
		req.Encode(e)
	}

	return e, nil
}

func encodeTeamsUpdateDiscussionLegacyRequestJSON(req OptTeamsUpdateDiscussionLegacyReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()
	if req.Set {
		req.Encode(e)
	}

	return e, nil
}

func encodeTeamsUpdateInOrgRequestJSON(req OptTeamsUpdateInOrgReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()
	if req.Set {
		req.Encode(e)
	}

	return e, nil
}

func encodeTeamsUpdateLegacyRequestJSON(req TeamsUpdateLegacyReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()

	req.Encode(e)

	return e, nil
}

func encodeUsersAddEmailForAuthenticatedRequestJSON(req OptUsersAddEmailForAuthenticatedReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()
	if req.Set {
		req.Encode(e)
	}

	return e, nil
}

func encodeUsersCreateGpgKeyForAuthenticatedRequestJSON(req UsersCreateGpgKeyForAuthenticatedReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()

	req.Encode(e)

	return e, nil
}

func encodeUsersCreatePublicSSHKeyForAuthenticatedRequestJSON(req UsersCreatePublicSSHKeyForAuthenticatedReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()

	req.Encode(e)

	return e, nil
}

func encodeUsersDeleteEmailForAuthenticatedRequestJSON(req OptUsersDeleteEmailForAuthenticatedReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()
	if req.Set {
		req.Encode(e)
	}

	return e, nil
}

func encodeUsersSetPrimaryEmailVisibilityForAuthenticatedRequestJSON(req UsersSetPrimaryEmailVisibilityForAuthenticatedReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()

	req.Encode(e)

	return e, nil
}

func encodeUsersUpdateAuthenticatedRequestJSON(req OptUsersUpdateAuthenticatedReq, span trace.Span) (data *jx.Writer, err error) {
	e := jx.GetWriter()
	if req.Set {
		req.Encode(e)
	}

	return e, nil
}
