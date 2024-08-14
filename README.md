## Golang API + ChatGPT
คำสั่งรันโปรเจค
```bash
$ air
```

PORT ที่ใช้งาน คือ 8080

| method | path | headers | body | response |
| ------------- | ------------- | ------------- | ------------- | ------------- |
| POST | /login | - | username: string, password: string | message: string, token: string |
| POST | /register | - | username: string, password: string, role: string | message: string, result: array |
| GET | /me | Authorization Bearer | - | message: string, role: string, user: string, userid: int |