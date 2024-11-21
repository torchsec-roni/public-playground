data "hcl_schema" "app" {
  paths = fileset("schema/**.hcl")
}

env "local" {
  src = data.hcl_schema.app.url
  url = "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"
  migration {
    // URL where the migration directory resides.
    dir = "file://migrations"
  }
  // See: https://atlasgo.io/concepts/dev-database
  dev = "docker://postgres/15/dev?search_path=public"
}
