package stacks

import (
	"fmt"
	"path/filepath"
)

type KotlinStack struct{}

func (k KotlinStack) Name() string { return "Kotlin (Android)" }

func (k KotlinStack) Generate(cwd, featureName string, opts Options) error {
	pascal, _, snake := ParseFeatureName(featureName)
	base := filepath.Join(cwd, "features", snake)

	files := map[string]string{
		filepath.Join(base, "view", pascal+"Fragment.kt"):      kotlinFragmentTemplate(pascal),
		filepath.Join(base, "viewmodel", pascal+"ViewModel.kt"): kotlinViewModelTemplate(pascal),
	}
	if opts.IncludeModels {
		files[filepath.Join(base, "model", pascal+"Model.kt")] = kotlinModelTemplate(pascal)
	}
	if opts.IncludeWidgets {
		files[filepath.Join(base, "widget", pascal+"Widget.kt")] = kotlinWidgetTemplate(pascal)
	}

	for path, content := range files {
		if err := WriteFile(path, content); err != nil {
			return fmt.Errorf("kotlin: failed to write %s: %w", path, err)
		}
	}
	return nil
}

func kotlinFragmentTemplate(pascal string) string {
	return fmt.Sprintf(`import android.os.Bundle
import android.view.View
import androidx.fragment.app.Fragment

class %sFragment : Fragment() {

    override fun onViewCreated(view: View, savedInstanceState: Bundle?) {
        super.onViewCreated(view, savedInstanceState)
    }
}
`, pascal)
}

func kotlinViewModelTemplate(pascal string) string {
	return fmt.Sprintf(`import androidx.lifecycle.ViewModel

class %sViewModel : ViewModel() {
}
`, pascal)
}

func kotlinModelTemplate(pascal string) string {
	return fmt.Sprintf(`data class %sModel(
    val id: String = ""
)
`, pascal)
}

func kotlinWidgetTemplate(pascal string) string {
	return fmt.Sprintf(`import android.content.Context
import android.util.AttributeSet
import android.view.View

class %sWidget @JvmOverloads constructor(
    context: Context,
    attrs: AttributeSet? = null,
    defStyleAttr: Int = 0
) : View(context, attrs, defStyleAttr) {
}
`, pascal)
}
