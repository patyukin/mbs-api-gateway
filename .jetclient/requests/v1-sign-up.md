```toml
name = 'v1-sign-up'
method = 'POST'
url = 'http://localhost:5001/v1/sign-up'
sortWeight = 2000000
id = 'a5b1fb0f-9e96-4f25-bac0-2712c40c0d00'

[body]
type = 'JSON'
raw = '''
{
  "email": "john.doe@example.com",
  "password": "securepassword123",
  "telegram_login": "johndoe_telegram",
  "first_name": "John",
  "last_name": "Doe",
  "patronymic": "Jonathan",
  "date_of_birth": "1990-01-01",
  "phone": "79000000000",
  "address": "Moscow, Russia, 1, 1, 1"
}'''
```
