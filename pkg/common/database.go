package common

type Database struct {
    Type string `name:"Type" question:"What's the type?"`
    Version string `name:"Version" question:"What's the version?"`
    Storage string `name:"Storage in GB" question:"What's the storage in GB?"`
    OriginalHost string `name:"Original Host" question:"What's the host of the original database?"`
    OriginalDatabase string `name:"Original Database" question:"What's the name of the original database?"`
    OriginalUsername string `name:"Original Username" question:"What's the username of the original database?"`
    OriginalPassword string `name:"Original Password" question:"What's the password of the original database?"`
    EnvVarHost string `name:"The Environment Variable for Host" question:"What's the environment variable for the host?"`
    EnvVarDatabase string `name:"The Environment Variable for Database" question:"What's the environment variable for the database name?"`
    EnvVarUsername string `name:"The Environment Variable for Username" question:"What's the environment variable for the username?"`
    EnvVarPassword string `name:"The Environment Variable for Password" question:"What's the environment variable for the password?"`
}
