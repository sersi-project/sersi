package pkg

type Preset struct {
	Frontend FrontendConfig
	Backend  BackendConfig
	Devops   DevopsConfig
}

var MernPreset Preset = Preset{
	Frontend: FrontendConfig{
		Framework: "react",
		CSS:       "tailwind",
		Language:  "js",
	},
	Backend: BackendConfig{
		Language:  "js",
		Framework: "express",
		Database:  "mongodb",
	},
	Devops: DevopsConfig{
		CI:         "github",
		Docker:     false,
		Monitoring: "none",
	},
}

var FastAPIPreset Preset = Preset{
	Frontend: FrontendConfig{
		Framework: "react",
		CSS:       "tailwind",
		Language:  "js",
	},
	Backend: BackendConfig{
		Language:  "python",
		Framework: "fastapi",
		Database:  "postgresql",
	},
	Devops: DevopsConfig{
		CI:         "github",
		Docker:     false,
		Monitoring: "none",
	},
}

var SveltePreset Preset = Preset{
	Frontend: FrontendConfig{
		Framework: "svelte",
		CSS:       "tailwind",
		Language:  "js",
	},
	Backend: BackendConfig{
		Language:  "go",
		Framework: "gin",
		Database:  "postgresql",
	},
	Devops: DevopsConfig{
		CI:         "github",
		Docker:     false,
		Monitoring: "none",
	},
}
