package starr

/* This file contains shared structs or constants for all the *arr apps. */

// StatusMessage represents the status of the item. All apps use this.
type StatusMessage struct {
	Title    string   `json:"title"`
	Messages []string `json:"messages"`
}

// Quality is a download quality attached to a movie, book, track or series.
type Quality struct {
	Quality struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
	} `json:"quality"`
	Revision struct {
		Version  int64 `json:"version"`
		Real     int64 `json:"real"`
		IsRepack bool  `json:"isRepack"`
	} `json:"revision"`
}
