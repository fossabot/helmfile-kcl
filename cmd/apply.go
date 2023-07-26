package cmd

import (
	"github.com/spf13/cobra"
	"kcl-lang.io/helmfile-kcl/pkg/app"
	"kcl-lang.io/helmfile-kcl/pkg/config"
)

// NewApplyCmd returns the template command.
func NewApplyCmd() *cobra.Command {
	templateOptions := config.NewTemplateOptions()

	cmd := &cobra.Command{
		Use:   "apply",
		Short: "Apply all resources from state file only when there are changes",
		RunE: func(*cobra.Command, []string) error {
			err := app.New().Template(config.NewTemplateImpl(templateOptions))
			if err != nil {
				return err
			}
			return nil
		},
		SilenceUsage: true,
	}

	f := cmd.Flags()
	f.StringVar(&templateOptions.File, "file", "", "input kcl file to pass to helm kcl template")
	f.StringArrayVar(&templateOptions.Set, "set", nil, "additional values to be merged into the helm command --set flag")
	f.StringArrayVar(&templateOptions.Values, "values", nil, "additional value files to be merged into the helm command --values flag")
	f.StringVar(&templateOptions.OutputDir, "output-dir", "", "output directory to pass to helm template (helm template --output-dir)")
	f.StringVar(&templateOptions.OutputDirTemplate, "output-dir-template", "", "go text template for generating the output directory. Default: {{ .OutputDir }}/{{ .State.BaseName }}-{{ .State.AbsPathSHA1 }}-{{ .Release.Name}}")
	f.IntVar(&templateOptions.Concurrency, "concurrency", 0, "maximum number of concurrent helm processes to run, 0 is unlimited")
	f.BoolVar(&templateOptions.Validate, "validate", false, "validate your manifests against the Kubernetes cluster you are currently pointing at. Note that this requires access to a Kubernetes cluster to obtain information necessary for validating, like the template of available API versions")
	f.BoolVar(&templateOptions.IncludeCRDs, "include-crds", false, "include CRDs in the templated output")
	f.BoolVar(&templateOptions.SkipTests, "skip-tests", false, "skip tests from templated output")
	f.BoolVar(&templateOptions.SkipNeeds, "skip-needs", true, `do not automatically include releases from the target release's "needs" when --selector/-l flag is provided. Does nothing when --selector/-l flag is not provided. Defaults to true when --include-needs or --include-transitive-needs is not provided`)
	f.BoolVar(&templateOptions.IncludeNeeds, "include-needs", false, `automatically include releases from the target release's "needs" when --selector/-l flag is provided. Does nothing when --selector/-l flag is not provided`)
	f.BoolVar(&templateOptions.IncludeTransitiveNeeds, "include-transitive-needs", false, `like --include-needs, but also includes transitive needs (needs of needs). Does nothing when --selector/-l flag is not provided. Overrides exclusions of other selectors and conditions.`)
	f.BoolVar(&templateOptions.SkipDeps, "skip-deps", false, `skip running "helm repo update" and "helm dependency build"`)
	f.StringVar(&templateOptions.PostRenderer, "post-renderer", "", `pass --post-renderer to "helm template" or "helm upgrade --install"`)

	return cmd
}
