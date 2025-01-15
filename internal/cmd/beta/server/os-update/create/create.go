package create

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/goccy/go-yaml"
	"github.com/stackitcloud/stackit-cli/internal/pkg/args"
	cliErr "github.com/stackitcloud/stackit-cli/internal/pkg/errors"
	"github.com/stackitcloud/stackit-cli/internal/pkg/examples"
	"github.com/stackitcloud/stackit-cli/internal/pkg/flags"
	"github.com/stackitcloud/stackit-cli/internal/pkg/globalflags"
	"github.com/stackitcloud/stackit-cli/internal/pkg/print"
	"github.com/stackitcloud/stackit-cli/internal/pkg/services/serverosupdate/client"

	"github.com/spf13/cobra"
	"github.com/stackitcloud/stackit-sdk-go/services/serverupdate"
)

const (
	serverIdFlag             = "server-id"
	maintenanceWindowFlag    = "maintenance-window"
	defaultMaintenanceWindow = 1
)

type inputModel struct {
	*globalflags.GlobalFlagModel

	ServerId          string
	MaintenanceWindow int64
}

func NewCmd(p *print.Printer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Creates a Server os-update.",
		Long:  "Creates a Server os-update. Operation always is async.",
		Args:  args.NoArgs,
		Example: examples.Build(
			examples.NewExample(
				`Create a Server os-update with name "myupdate"`,
				`$ stackit beta server os-update create --server-id xxx --name=myupdate`),
			examples.NewExample(
				`Create a Server os-update with name "myupdate" and maintenance window for 13 o'clock.`,
				`$ stackit beta server os-update create --server-id xxx --name=myupdate --maintenance-window=13`),
		),
		RunE: func(cmd *cobra.Command, _ []string) error {
			ctx := context.Background()

			model, err := parseInput(p, cmd)
			if err != nil {
				return err
			}

			// Configure API client
			apiClient, err := client.ConfigureClient(p)
			if err != nil {
				return err
			}

			if !model.AssumeYes {
				prompt := fmt.Sprintf("Are you sure you want to create a os-update for server %s?", model.ServerId)
				err = p.PromptForConfirmation(prompt)
				if err != nil {
					return err
				}
			}

			// Call API
			req, err := buildRequest(ctx, model, apiClient)
			if err != nil {
				return err
			}
			resp, err := req.Execute()
			if err != nil {
				return fmt.Errorf("create Server os-update: %w", err)
			}

			return outputResult(p, model, resp)
		},
	}
	configureFlags(cmd)
	return cmd
}

func configureFlags(cmd *cobra.Command) {
	cmd.Flags().VarP(flags.UUIDFlag(), serverIdFlag, "s", "Server ID")
	cmd.Flags().Int64P(maintenanceWindowFlag, "m", defaultMaintenanceWindow, "Maintenance window (in hours, 1-24)")

	err := flags.MarkFlagsRequired(cmd, serverIdFlag)
	cobra.CheckErr(err)
}

func parseInput(p *print.Printer, cmd *cobra.Command) (*inputModel, error) {
	globalFlags := globalflags.Parse(p, cmd)
	if globalFlags.ProjectId == "" {
		return nil, &cliErr.ProjectIdError{}
	}

	model := inputModel{
		GlobalFlagModel:   globalFlags,
		ServerId:          flags.FlagToStringValue(p, cmd, serverIdFlag),
		MaintenanceWindow: flags.FlagWithDefaultToInt64Value(p, cmd, maintenanceWindowFlag),
	}

	if p.IsVerbosityDebug() {
		modelStr, err := print.BuildDebugStrFromInputModel(model)
		if err != nil {
			p.Debug(print.ErrorLevel, "convert model to string for debugging: %v", err)
		} else {
			p.Debug(print.DebugLevel, "parsed input values: %s", modelStr)
		}
	}

	return &model, nil
}

func buildRequest(ctx context.Context, model *inputModel, apiClient *serverupdate.APIClient) (serverupdate.ApiCreateUpdateRequest, error) {
	req := apiClient.CreateUpdate(ctx, model.ProjectId, model.ServerId)
	payload := serverupdate.CreateUpdatePayload{
		MaintenanceWindow: &model.MaintenanceWindow,
	}
	req = req.CreateUpdatePayload(payload)
	return req, nil
}

func outputResult(p *print.Printer, model *inputModel, resp *serverupdate.Update) error {
	switch model.OutputFormat {
	case print.JSONOutputFormat:
		details, err := json.MarshalIndent(resp, "", "  ")
		if err != nil {
			return fmt.Errorf("marshal server os-update: %w", err)
		}
		p.Outputln(string(details))

		return nil
	case print.YAMLOutputFormat:
		details, err := yaml.MarshalWithOptions(resp, yaml.IndentSequence(true))
		if err != nil {
			return fmt.Errorf("marshal server os-update: %w", err)
		}
		p.Outputln(string(details))

		return nil
	default:
		p.Outputf("Triggered creation of server os-update for server %s. Update ID: %d\n", model.ServerId, *resp.Id)
		return nil
	}
}