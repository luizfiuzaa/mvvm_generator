package stacks

import (
	"fmt"
	"path/filepath"
)

type FlutterStack struct{}

func (f FlutterStack) Name() string { return "Flutter (Dart)" }

func (f FlutterStack) Generate(cwd, featureName string, opts Options) error {
	pascal, _, snake := ParseFeatureName(featureName)
	base := filepath.Join(cwd, "lib", "features", snake)

	files := map[string]string{
		filepath.Join(base, "views", snake+"_view.dart"):           flutterViewTemplate(pascal, snake),
		filepath.Join(base, "viewmodels", snake+"_viewmodel.dart"): flutterViewModelTemplate(pascal),
	}
	if opts.IncludeModels {
		files[filepath.Join(base, "models", snake+"_model.dart")] = flutterModelTemplate(pascal)
	}
	if opts.IncludeWidgets {
		files[filepath.Join(base, "widgets", snake+"_widget.dart")] = flutterWidgetTemplate(pascal)
	}

	for path, content := range files {
		if err := WriteFile(path, content); err != nil {
			return fmt.Errorf("flutter: failed to write %s: %w", path, err)
		}
	}
	return nil
}

func flutterViewTemplate(pascal, snake string) string {
	return fmt.Sprintf(`import 'package:flutter/material.dart';

class %sView extends StatelessWidget {
  const %sView({super.key});

  @override
  Widget build(BuildContext context) {
    return const Scaffold(
      body: Center(
        child: Text('%s view'),
      ),
    );
  }
}
`, pascal, pascal, snake)
}

func flutterViewModelTemplate(pascal string) string {
	return fmt.Sprintf(`import 'package:flutter/foundation.dart';

class %sViewModel extends ChangeNotifier {
  %sViewModel();
}
`, pascal, pascal)
}

func flutterModelTemplate(pascal string) string {
	return fmt.Sprintf(`class %sModel {
  const %sModel();
}
`, pascal, pascal)
}

func flutterWidgetTemplate(pascal string) string {
	return fmt.Sprintf(`import 'package:flutter/material.dart';

class %sWidget extends StatelessWidget {
  const %sWidget({super.key});

  @override
  Widget build(BuildContext context) {
    return const SizedBox.shrink();
  }
}
`, pascal, pascal)
}
