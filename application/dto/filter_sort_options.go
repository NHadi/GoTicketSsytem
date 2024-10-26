package dto

// FilterOptions defines the available filter criteria
type FilterOptions struct {
	FilterName  string `json:"filter_name"`
	FilterType  string `json:"filter_type"`  // "before", "after", "between"
	FilterValue string `json:"filter_value"` // single date string for "before" and "after", or a date range for "between"
}

// SortOptions defines the available sort criteria
type SortOptions struct {
	SortName string `json:"sort_name"` // "created_at" or "user_id"
	SortDir  string `json:"sort_dir"`  // "asc" or "desc"
}
