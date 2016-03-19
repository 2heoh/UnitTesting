Run service:

```bash
./script/gazetiere daemon
```

this runs service on port 3000

Run tests:
```bash
./script/gazetiere test
``` 

Generate coverage report:
```bash
cover -delete
PERL5OPT=-MDevel::Cover ./script/gazetiere test
cover
```

Install dependencies: 
```bash
cpanm DBIx::Connector Mojolicious CGI Devel::Cover 
```