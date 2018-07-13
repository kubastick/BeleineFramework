package beleine

import "fmt"

func PostRequestJS(URL string, jsData string, onSuccesJS string) string {
	return fmt.Sprintf(`$.post("%s",%s,function(data){%s});`, URL, jsData, onSuccesJS)
}
