// Binary ogen generates go source code from OAS.
package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/go-faster/errors"
	"go.uber.org/zap"

	"github.com/ogen-go/ogen"
	"github.com/ogen-go/ogen/gen"
	"github.com/ogen-go/ogen/gen/genfs"
	"github.com/ogen-go/ogen/internal/location"
	"github.com/ogen-go/ogen/internal/ogenzap"
)

func cleanDir(targetDir string, files []os.DirEntry) error {
	for _, f := range files {
		if f.IsDir() {
			continue
		}
		name := f.Name()
		if !(strings.HasSuffix(name, "_gen.go") || strings.HasSuffix(name, "_gen_test.go")) {
			continue
		}
		if !(strings.HasPrefix(name, "openapi") || strings.HasPrefix(name, "oas")) {
			continue
		}
		if err := os.Remove(filepath.Join(targetDir, name)); err != nil {
			return err
		}
	}
	return nil
}

func generate(data []byte, packageName string, fs gen.FileSystem, opts gen.Options) error {
	log := opts.Logger

	spec, err := ogen.Parse(data)
	if err != nil {
		return errors.Wrap(err, "parse spec")
	}

	start := time.Now()
	g, err := gen.NewGenerator(spec, opts)
	if err != nil {
		return errors.Wrap(err, "build IR")
	}
	log.Debug("Build IR", zap.Duration("took", time.Since(start)))

	start = time.Now()
	if err := g.WriteSource(fs, packageName); err != nil {
		return errors.Wrap(err, "write")
	}
	log.Debug("Write", zap.Duration("took", time.Since(start)))

	return nil
}

func run() error {
	set := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	set.Usage = func() {
		_, _ = fmt.Fprintf(set.Output(), "Usage: %s [options] <spec>\n", os.Args[0])
		set.PrintDefaults()
	}
	var (
		targetDir         = set.String("target", "api", "Path to target dir")
		packageName       = set.String("package", "api", "Target package name")
		inferTypes        = set.Bool("infer-types", false, "Infer schema types, if type is not defined explicitly")
		performFormat     = set.Bool("format", true, "Perform code formatting")
		clean             = set.Bool("clean", false, "Clean generated files before generation")
		generateTests     = set.Bool("generate-tests", false, "Generate tests encode-decode/based on schema examples")
		allowRemote       = set.Bool("allow-remote", false, "Enables remote references resolving")
		skipTestsRegex    = set.String("skip-tests", "", "Skip tests matched by regex")
		skipUnimplemented = set.Bool("skip-unimplemented", false, "Disables generation of UnimplementedHandler")
		noClient          = set.Bool("no-client", false, "Disables client generation")
		noServer          = set.Bool("no-server", false, "Disables server generation")

		debugIgnoreNotImplemented = set.String("debug.ignoreNotImplemented", "",
			"Ignore methods having functionality which is not implemented ")
		debugNoerr = set.Bool("debug.noerr", false, "Ignore errors")

		logOptions ogenzap.Options
	)
	logOptions.RegisterFlags(set)

	var (
		ctAliases gen.ContentTypeAliases

		filterPath    *regexp.Regexp
		filterMethods []string
	)
	set.Var(&ctAliases, "ct-alias", "Content type alias, e.g. text/x-markdown=text/plain")
	set.Func("filter-path", "Filter operations by path regex", func(s string) (err error) {
		filterPath, err = regexp.Compile(s)
		return err
	})
	set.Func("filter-methods", "Filter operations by HTTP methods (comma-separated)", func(s string) error {
		for _, m := range strings.Split(s, ",") {
			m = strings.TrimSpace(m)
			if m == "" {
				continue
			}
			filterMethods = append(filterMethods, m)
		}
		return nil
	})

	if err := set.Parse(os.Args[1:]); err != nil {
		return err
	}

	specPath := set.Arg(0)
	if set.NArg() == 0 || specPath == "" {
		set.Usage()
		return errors.New("no spec provided")
	}
	specPath = filepath.Clean(specPath)

	switch files, err := os.ReadDir(*targetDir); {
	case os.IsNotExist(err):
		if err := os.MkdirAll(*targetDir, 0o750); err != nil {
			return err
		}
	case err != nil:
		return err
	default:
		if *clean {
			if err := cleanDir(*targetDir, files); err != nil {
				return errors.Wrap(err, "clean")
			}
		}
	}

	logger, err := ogenzap.Create(logOptions)
	if err != nil {
		return err
	}
	defer func() {
		_ = logger.Sync()
	}()

	_, fileName := filepath.Split(specPath)
	opts := gen.Options{
		NoClient:             *noClient,
		NoServer:             *noServer,
		GenerateExampleTests: *generateTests,
		SkipTestRegex:        nil, // Set below.
		SkipUnimplemented:    *skipUnimplemented,
		InferSchemaType:      *inferTypes,
		AllowRemote:          *allowRemote,
		Filters: gen.Filters{
			PathRegex: filterPath,
			Methods:   filterMethods,
		},
		IgnoreNotImplemented: strings.Split(*debugIgnoreNotImplemented, ","),
		ContentTypeAliases:   ctAliases,
		Filename:             fileName,
		Logger:               logger,
	}
	if expr := *skipTestsRegex; expr != "" {
		r, err := regexp.Compile(expr)
		if err != nil {
			return errors.Wrap(err, "skipTestsRegex")
		}
		opts.SkipTestRegex = r
	}
	if *debugNoerr {
		opts.IgnoreNotImplemented = []string{"all"}
	}

	data, err := os.ReadFile(specPath)
	if err != nil {
		return err
	}

	fs := genfs.FormattedSource{
		Root:   *targetDir,
		Format: *performFormat,
	}
	if err := generate(data, *packageName, fs, opts); err != nil {
		if location.PrintPrettyError(os.Stderr, specPath, data, err) {
			return errors.New("generation failed")
		}
		return errors.Wrap(err, "generate")
	}

	return nil
}

func main() {
	if err := run(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%+v\n", err)
		os.Exit(1)
	}
}
