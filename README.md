# Karlota Server

### Create Mock with [Mockery](https://github.com/vektra/mockery)

- Run Specified Dir & File

```bash
  mockery 
  --dir=internal/app/repository/mysql 
  --name=AccountRepository 
  --filename=account_repository_impl.go 
  --output=internal/app/domain/mocks --outpkg=mocks
```