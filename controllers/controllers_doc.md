## The following directories must not exceed the depth of more than 1 level as they are directly bound to routes and models
# handlers: called immediately after access to a route
# repos: have access to databases and anything in the models directory
## WARNING: do not let handlers access repos directly, intermediate steps are mandatory !!!


## The following directories do not directly respond to any routes but may be called by any other shallow level handlers
# services: intermediary steps between routes, repos, and responses. 
# responses: a special type of helper that returns structs that implement gin.Context's own responses.
## WARNING: services and responses are not complete handlers by themselves and each may be accessed by multiple handlers. Keep the variables within method's own scopes, do not initialize any global variables in those files !!!

## NOTE: anything that requires access to database->repo or other common resources is considered a service, anything that doesn't should be placed inside utils directory. 