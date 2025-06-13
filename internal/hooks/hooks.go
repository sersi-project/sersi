package hooks

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"

	"github.com/charmbracelet/lipgloss"
)

var hookStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#00008B"))

type Hooks struct {
	preHook     bool
	postHook    bool
	projectName string
	httpClient  *http.Client
}

func InitHooks(projectName string, preHook bool, postHook bool) *Hooks {
	return &Hooks{projectName: projectName, preHook: preHook, postHook: postHook, httpClient: &http.Client{}}
}

func (h *Hooks) RunPreHook() error {
	if h.preHook {
		fmt.Printf("%s\n", hookStyle.Render("Running pre-hook ..."))
		err := h.GetFileFromServer("pre")
		if err != nil {
			return fmt.Errorf("error getting pre-hook file: %s", err)
		}
		out, err := exec.Command("python", "install.py").CombinedOutput()
		if err != nil {
			return fmt.Errorf("error executing pre-hook: %s", err)
		}
		fmt.Printf("\nOutput: %s\n\n", hookStyle.Render(string(out)))
		fmt.Printf("Pre-hook executed successfully!\n\n")
		err = os.Remove("install.py")
		if err != nil {
			return fmt.Errorf("error removing pre-hook file: %s", err)
		}
		return nil
	}
	return nil
}

func (h *Hooks) RunPostHook() error {
	if h.postHook {
		fmt.Printf("%s\n", hookStyle.Render("Running post-hook ..."))
		err := h.GetFileFromServer("post")
		if err != nil {
			return fmt.Errorf("error getting post-hook file: %s", err)
		}
		err = exec.Command("python", "install.py", h.projectName).Run()
		if err != nil {
			return fmt.Errorf("error executing post-hook: %s", err)
		}
		fmt.Printf("%s\n", hookStyle.Render("Post-hook executed successfully!"))
		err = os.Remove("install.py")
		if err != nil {
			return fmt.Errorf("error removing post-hook file: %s", err)
		}
		return nil
	}
	return nil
}

func (h *Hooks) GetFileFromServer(endpoint string) error {
	baseUrl := os.Getenv("SERSI_HOOKS_BASE_URL")
	fullUrl := fmt.Sprintf("%s/%s", baseUrl, endpoint)
	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return fmt.Errorf("error creating request: %s", err)
	}

	req.Header.Set("Accept", "text/x-python")

	res, err := h.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("error sending request: %s", err)
	}

	defer res.Body.Close() //nolint

	if res.StatusCode != 200 {
		return fmt.Errorf("pre-hook failed with status code: %d", res.StatusCode)
	}

	out, err := os.Create("install.py")
	if err != nil {
		return fmt.Errorf("error creating file: %s", err)
	}

	defer out.Close() //nolint

	_, err = io.Copy(out, res.Body)
	if err != nil {
		return fmt.Errorf("error copying file: %s", err)
	}

	return nil
}
