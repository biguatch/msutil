package msutil

import "net/http"

func RequestToFilter(r *http.Request) map[string][]string {
	params := r.Form
	filters := map[string][]string{}

	for k := range params {
		if len(k) > 7 && k[:7] == "filter_" {
			filters[k[7:]] = params[k]
		}
	}

	return filters
}
