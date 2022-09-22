# Karlota Server

### Create Mock with [Mockery](https://github.com/vektra/mockery)

- Run Specified Dir & File

```bash
  mockery 
  --dir=src/repository/mysql 
  --name=AccountRepository 
  --filename=account_repository_impl.go 
  --output=src/domain/mocks --outpkg=mocks
```