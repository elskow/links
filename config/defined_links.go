package config

type url struct {
	Logo  string `json:"logo,omitempty"`
	Title string `json:"title"`
	Url   string `json:"url"`
}

type urls []url

const (
	LinkedIn = "https://www.linkedin.com/in/helmyluqman/"
	Twitter  = "https://twitter.com/helmy_lh"
	GitHub   = "https://www.github.com/elskow"
	Kaggle   = "https://www.kaggle.com/elskowaski"
	Wandb    = "https://wandb.ai/elskow"
)

func NewDefinedUrls() urls {
	return urls{
		{
			Title: "helmyl.com",
			Url:   "https://helmyl.com",
		},
		{
			Logo:  "https://linkedin.com/favicon.ico",
			Title: "LinkedIn",
			Url:   LinkedIn,
		},
		{
			Logo:  "https://twitter.com/favicon.ico",
			Title: "Twitter",
			Url:   Twitter,
		},
		{
			Logo:  "https://github.githubassets.com/favicons/favicon.png",
			Title: "GitHub",
			Url:   GitHub,
		},
		{
			Logo:  "https://www.kaggle.com/static/images/favicon.ico",
			Title: "Kaggle",
			Url:   Kaggle,
		},
		{
			Logo:  "https://cdn.wandb.ai/production/3df823730/favicon.png",
			Title: "Wandb",
			Url:   Wandb,
		},
	}
}
