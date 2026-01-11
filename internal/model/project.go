package model

type ProjectMetadata struct {
    Name           string
    RootPath       string
    Languages      []string
    PackageManager string

    InstallCommands []string
    RunCommands     []string

    HasDocker        bool
    HasDockerCompose bool
    HasEnvFile       bool
    EnvExampleFile   string
    License string
}
