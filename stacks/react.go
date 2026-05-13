package stacks

import (
	"fmt"
	"path/filepath"
)

type ReactStack struct{}

func (r ReactStack) Name() string { return "React (TypeScript)" }

func (r ReactStack) Generate(cwd, featureName string, opts Options) error {
	pascal, camel, _ := ParseFeatureName(featureName)
	base := filepath.Join(cwd, "src", "features", camel)

	files := map[string]string{
		filepath.Join(base, "views", pascal+"View.tsx"):                    reactViewTemplate(pascal),
		filepath.Join(base, "viewmodels", "use"+pascal+"ViewModel.ts"):     reactViewModelTemplate(pascal),
	}
	if opts.IncludeModels {
		files[filepath.Join(base, "models", pascal+"Model.ts")] = reactModelTemplate(pascal)
	}
	if opts.IncludeWidgets {
		files[filepath.Join(base, "widgets", pascal+"Widget.tsx")] = reactWidgetTemplate(pascal)
	}

	for path, content := range files {
		if err := WriteFile(path, content); err != nil {
			return fmt.Errorf("react: failed to write %s: %w", path, err)
		}
	}
	return nil
}

func reactViewTemplate(pascal string) string {
	return fmt.Sprintf(`import React from 'react';
import { use%sViewModel } from '../viewmodels/use%sViewModel';

const %sView: React.FC = () => {
  const viewModel = use%sViewModel();

  return (
    <div>
      <h1>%s</h1>
    </div>
  );
};

export default %sView;
`, pascal, pascal, pascal, pascal, pascal, pascal)
}

func reactViewModelTemplate(pascal string) string {
	return fmt.Sprintf(`export function use%sViewModel() {
  return {};
}
`, pascal)
}

func reactModelTemplate(pascal string) string {
	return fmt.Sprintf(`export interface %sModel {
  id: string;
}
`, pascal)
}

func reactWidgetTemplate(pascal string) string {
	return fmt.Sprintf(`import React from 'react';

const %sWidget: React.FC = () => {
  return <div />;
};

export default %sWidget;
`, pascal, pascal)
}
