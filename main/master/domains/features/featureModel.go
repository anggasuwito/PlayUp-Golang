package features

type Feature struct {
	ID           		string `json:"feature_id"`
	FeatureName 		string `json:"feature_name"`
	FeatureDescription  string `json:"feature_description"`
	FeatureImage	  	string `json:"feature_image"`
	FeatureStatus  		string `json:"feature_status"`
	Created      		string `json:"created"`
	Updated      		string `json:"updated"`
}

