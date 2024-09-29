schema "public" {
  comment = "standard public schema"
}
table "users" {
  schema = schema.public
  column "name" {
    null = false
    type = string
  }
}
