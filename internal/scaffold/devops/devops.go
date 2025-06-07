package devops

type Devops struct {
	ProjectName string
	Docker      bool
	CI          string
	Monitoring  string
}

type DevopsBuilder struct {
	devops *Devops
}

func NewDevopsBuilder() *DevopsBuilder {
	return &DevopsBuilder{
		devops: &Devops{},
	}
}

func (b *DevopsBuilder) ProjectName(projectName string) *DevopsBuilder {
	b.devops.ProjectName = projectName
	return b
}

func (b *DevopsBuilder) Docker(docker bool) *DevopsBuilder {
	b.devops.Docker = docker
	return b
}

func (b *DevopsBuilder) CI(ci string) *DevopsBuilder {
	b.devops.CI = ci
	return b
}

func (b *DevopsBuilder) Monitoring(monitoring string) *DevopsBuilder {
	b.devops.Monitoring = monitoring
	return b
}

func (b *DevopsBuilder) Build() *Devops {
	return b.devops
}

func (b *Devops) Generate() error {
	template := NewDTemplateBuilder().ProjectName(b.ProjectName).CI(b.CI).Docker(b.Docker).Monitoring(b.Monitoring).Build()
	err := template.Execute()
	if err != nil {
		return b.ProcessError(err)
	}
	return nil
}

func (b *Devops) ProcessError(err error) error {
	return err
}
