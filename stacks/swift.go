package stacks

import (
	"fmt"
	"path/filepath"
)

type SwiftStack struct{}

func (s SwiftStack) Name() string { return "Swift (iOS)" }

func (s SwiftStack) Generate(cwd, featureName string, opts Options) error {
	pascal, _, _ := ParseFeatureName(featureName)
	base := filepath.Join(cwd, "Features", pascal)

	files := map[string]string{
		filepath.Join(base, "Views", pascal+"View.swift"):           swiftViewTemplate(pascal),
		filepath.Join(base, "ViewModels", pascal+"ViewModel.swift"): swiftViewModelTemplate(pascal),
	}
	if opts.IncludeModels {
		files[filepath.Join(base, "Models", pascal+"Model.swift")] = swiftModelTemplate(pascal)
	}
	if opts.IncludeWidgets {
		files[filepath.Join(base, "Widgets", pascal+"Widget.swift")] = swiftWidgetTemplate(pascal)
	}

	for path, content := range files {
		if err := WriteFile(path, content); err != nil {
			return fmt.Errorf("swift: failed to write %s: %w", path, err)
		}
	}
	return nil
}

func swiftViewTemplate(pascal string) string {
	return fmt.Sprintf(`import SwiftUI

struct %sView: View {
    @StateObject private var viewModel = %sViewModel()

    var body: some View {
        Text("%s")
    }
}

#Preview {
    %sView()
}
`, pascal, pascal, pascal, pascal)
}

func swiftViewModelTemplate(pascal string) string {
	return fmt.Sprintf(`import Foundation
import Combine

final class %sViewModel: ObservableObject {
    init() {}
}
`, pascal)
}

func swiftModelTemplate(pascal string) string {
	return fmt.Sprintf(`import Foundation

struct %sModel: Codable, Equatable {
}
`, pascal)
}

func swiftWidgetTemplate(pascal string) string {
	return fmt.Sprintf(`import SwiftUI

struct %sWidget: View {
    var body: some View {
        EmptyView()
    }
}
`, pascal)
}
