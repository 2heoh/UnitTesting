```bash
virtualenv --no-site-packages .env
source .env/bin/activate
pip install -r requirements.txt
./manage.py migrate
./manage.py loaddata address.json
./manage.py runserver  # localhost:8000
REUSE_DB=1 ./manage.py test
```
