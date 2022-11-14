# Karlota Server

### Create Mock with [Mockery](https://github.com/vektra/mockery)

- Run Specified Dir & File

```bash
  mockery 
  --dir=internal/api/repository/mysql 
  --name=AccountRepository 
  --filename=account_repository_impl.go 
  --output=internal/api/domain/mocks --outpkg=mocks
```