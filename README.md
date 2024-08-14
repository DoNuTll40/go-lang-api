## Golang API + ChatGPT
คำสั่งรันโปรเจค
```bash
$ air
```

PORT ที่ใช้งาน คือ 8080

| path | headers | body | response |
| ------------- | ------------- | ------------- | ------------- |
| /login | - | username: string, password: string | message: string, token: string |
| /register | - | username: string, password: string, role: string | message: string, result: array |
| /me | Authorization Bearer | - | message: string, result: array |