## The following directories must not exceed the depth of more than 1 level as they are directly bound to routes and models
# handlers: 
Called immediately after access to a route
# repos: 
Have access to databases and anything in the models directory
# WARNING: do not let handlers access repos directly, intermediate steps are mandatory !!!


## The following directories do not directly respond to any routes but may be called by any other shallow level handlers
# services: 
Intermediary steps between routes, repos, and responses. 
# responses: 
A special type of helper that returns structs that implement gin.Context's own responses. Each file is independent and therefore contains its own struct objs.
# WARNING: services and responses are not complete handlers by themselves and each may be accessed by multiple handlers. Keep the variables within method's own scopes, do not initialize any global variables in those files !!!

## NOTE: anything that requires access to database->repo or other common resources is considered a service, anything that doesn't should be placed inside utils directory. 


## Responsibilities:
Handlers: checks request format -> cycles through: (calls validators -> calls services) -> generate responses
Services: main controllers and sanitizers before accessing repository
Validators: check whether the responses received from services can be used to proceed with further actions 
Repos: data access layer, the only layer that has direct access to databases