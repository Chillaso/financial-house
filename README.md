# Financial house

Financial house is an API REST developed with Go using Gin gonic framework and MongoDB.
Deployed in docker

The main purpose of this project is control house finances as an excel book but adding
filters, budgets and exposing this API to Internet in order to access everywhere.

DIY, avoid companies control your finances.

### TODO:

- [x] Finish CRUD
- [ ] Add calculation logic
- [ ] Add validations
- [ ] Add budgets CRUD 
- [ ] Add filters
- [ ] Hand errors in a clean way, view [this](https://github.com/henrmota/errors-handling-example) for more details
- [ ] Add a base controller and base service to reuse CRUD common logic
- [ ] Add result pagination
- [ ] Infrastructure (Going to deploy in Raspberry pi 4 with MicroK8s)
