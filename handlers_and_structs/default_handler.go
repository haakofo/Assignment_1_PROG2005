package handlers_and_structs

import "net/http"

func Default_Handler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "There is nothing to do here! Try these paths instead: "+UNIVERSAL_LINEBREAK+DIAG_PATH+UNIVERSAL_LINEBREAK+UNIINFO_PATH+UNIVERSAL_LINEBREAK+NEIGHBOURUNIS_TOTAL_PATH, http.StatusOK)
}
